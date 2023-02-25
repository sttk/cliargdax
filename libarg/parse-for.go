// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package libarg

import (
	"github.com/sttk-go/sabi"
	"reflect"
	"strconv"
	"strings"
)

type /* error reason */ (
	// OptionStoreIsNotChangeable is an error reason which indicates that
	// the second parameter of ParseFor function, which is set options produced
	// by parsing command line arguments, is not a pointer.
	OptionStoreIsNotChangeable struct{}

	// FailToParseInt is an error reaason which indicates that an option
	// parameter in command line arguments should be an integer but is invalid.
	FailToParseInt struct {
		Field   string
		Input   string
		BitSize int
	}

	// FailToParseUint is an error reason which indicates that an option
	// parameter in command line arguments should be an unsigned integer but is
	// invalid.
	FailToParseUint struct {
		Field   string
		Input   string
		BitSize int
	}

	// FailToParseFloat is an error reason which indicates that an option
	// parameter in command line arguments should be a floating point number but
	// is invalid.
	FailToParseFloat struct {
		Field   string
		Input   string
		BitSize int
	}

	// IllegalOptionType is an error reason which indicates that a type of a
	// field of the option store is neither a boolean, a number, a string, nor
	// an array of numbers or strings.
	IllegalOptionType struct {
		Field string
		Type  reflect.Type
	}
)

// ParseFor is a function to parse command line arguments and set their values
// to the option store which is the second parameter of this function.
// This function divides command line arguments to command parameters and
// options, then stores the options to the option store, and returns the
// command parameters.
//
// The configurations of options are determined by types and struct tags of
// fields of the option store.
// If the type is bool, the option takes no parameter.
// If the type is integer, floating point number or string, the option can
// takes one  option parameter, therefore it can appear once in command line
// arguments.
// If the type is an array, the option can takes multiple option parameters,
// therefore it can appear multiple times in command line arguments.
//
// A struct tag can specify an option name, aliases, and a default value.
// It has a special format, like `opt:foo-bar,f=123`.
// This opt: is the struct tag key for the option configuration.
// The string following this key and rounded by double quotes is the content
// of the option configuration.
// The first part of the option configuration is an option name and aliases,
// which are separated by commas, and ends with "=" mark or end of string.
// If the option name is empty or no struct tag, the option's name becomes same
// with the field name of the option store.
//
// The string after the "=" mark is default value(s).
// If the type of the option is a boolean, the string after "=" mark is ignored
// because a boolean option takes no option parameter.
// If the type of the option is a number or a string, the whole string after
// "=" mark is a default value.
// If the type of the option is an array, the string after "=" mark have to be
// rounded by square brackets and separate the elements with commas, like
// [elem1,elem2,elem3].
// The element separator can be used other than a comma by put the separator
// before the open square bracket, like :[elem1:elem2:elem3].
// It's useful when some array elements include commas.
//
// NOTE: A default value of a string array option in a struct tag is [], like
// `opt:"name=[]"`, it doesn't represent an array which contains only an empty
// string but an empty array.
// If you want to specify an array which contains only an empty string, write
// nothing after "=" mark, like `opt:"name="`.
//
// Usage example:
//
//	type MyOptions struct {
//	  FooBar bool     `opt:"foo-bar,f"`
//	  Baz    int      `opt:"baz,b=99"`
//	  Qux    string   `opt:"=XXX"`
//	  Quux   []string `opt:"quux=[A,B,C]"`
//	  Corge  []int
//	}
//	options := MyOptions{}
//
//	osArgs := []string{
//	  "--foo-bar", "c1", "-b", "12", "--Qux", "ABC", "c2",
//	  "--Corge", "20", "--Corge=21",
//	}
//
//	cmdParams, err := ParseFor(osArgs, &options)
//	cmdParams      // [c1 c2]
//	options.FooBar // true
//	options.Baz    // 12
//	options.Qux    // ABC
//	options.Quux   // [A B C]
//	options.Corge  // [20 21]
func ParseFor(args []string, options any) ([]string, sabi.Err) {
	v := reflect.ValueOf(options)
	if v.Kind() != reflect.Ptr {
		return empty, sabi.NewErr(OptionStoreIsNotChangeable{})
	}
	v = v.Elem()

	t := v.Type()
	n := t.NumField()

	optCfgs := make([]OptCfg, n)
	var err sabi.Err

	for i := 0; i < n; i++ {
		optCfgs[i], err = newOptCfg(t.Field(i))
		if !err.IsOk() {
			return empty, err
		}
		var setter func([]string) sabi.Err
		setter, err = newValueSetter(optCfgs[i].Name, v.Field(i))
		if !err.IsOk() {
			return empty, err
		}
		optCfgs[i].OnParsed = &setter
	}

	a, err := ParseWith(args, optCfgs)
	if !err.IsOk() {
		return empty, err
	}

	return a.cmdParams, sabi.Ok()
}

func newOptCfg(fld reflect.StructField) (OptCfg, sabi.Err) {
	opt := fld.Tag.Get("opt")
	arr := strings.SplitN(opt, "=", 2)

	names := strings.Split(arr[0], ",")

	var name string
	var aliases []string
	if len(names) == 0 || len(names[0]) == 0 {
		name = fld.Name
		aliases = nil
	} else {
		name = names[0]
		aliases = names[1:]
	}

	isArray := false
	hasParam := true
	switch fld.Type.Kind() {
	case reflect.Slice | reflect.Array:
		isArray = true
	case reflect.Bool:
		hasParam = false
	}

	var defaults []string
	if len(arr) > 1 && hasParam {
		def := arr[1]
		n := len(def)
		if !isArray {
			defaults = []string{def}
		} else if n > 1 && def[0] == '[' && def[n-1] == ']' {
			defs := def[1 : n-1]
			if len(defs) > 0 {
				defaults = strings.Split(defs, ",")
			} else {
				defaults = empty
			}
		} else if n > 2 && def[1] == '[' && def[n-1] == ']' {
			defs := def[2 : n-1]
			if len(defs) > 0 {
				defaults = strings.Split(defs, def[0:1])
			} else {
				defaults = empty
			}
		} else {
			defaults = []string{def}
		}
	}

	return OptCfg{
		Name:     name,
		Aliases:  aliases,
		HasParam: hasParam,
		IsArray:  isArray,
		Default:  defaults,
	}, sabi.Ok()
}

func newValueSetter(
	name string,
	fld reflect.Value,
) (func([]string) sabi.Err, sabi.Err) {
	t := fld.Type()
	switch t.Kind() {
	case reflect.Bool:
		return newBoolSetter(name, fld)
	case reflect.Int:
		return newIntSetter(name, fld, strconv.IntSize)
	case reflect.Int8:
		return newIntSetter(name, fld, 8)
	case reflect.Int16:
		return newIntSetter(name, fld, 16)
	case reflect.Int32:
		return newIntSetter(name, fld, 32)
	case reflect.Int64:
		return newIntSetter(name, fld, 64)
	case reflect.Uint:
		return newUintSetter(name, fld, strconv.IntSize)
	case reflect.Uint8:
		return newUintSetter(name, fld, 8)
	case reflect.Uint16:
		return newUintSetter(name, fld, 16)
	case reflect.Uint32:
		return newUintSetter(name, fld, 32)
	case reflect.Uint64:
		return newUintSetter(name, fld, 64)
	case reflect.Float32:
		return newFloatSetter(name, fld, 32)
	case reflect.Float64:
		return newFloatSetter(name, fld, 64)
	case reflect.Array | reflect.Slice:
		elm := t.Elem()
		switch elm.Kind() {
		case reflect.Int:
			return newIntArraySetter(name, fld, strconv.IntSize)
		case reflect.Int8:
			return newIntArraySetter(name, fld, 8)
		case reflect.Int16:
			return newIntArraySetter(name, fld, 16)
		case reflect.Int32:
			return newIntArraySetter(name, fld, 32)
		case reflect.Int64:
			return newIntArraySetter(name, fld, 64)
		case reflect.Uint:
			return newUintArraySetter(name, fld, strconv.IntSize)
		case reflect.Uint8:
			return newUintArraySetter(name, fld, 8)
		case reflect.Uint16:
			return newUintArraySetter(name, fld, 16)
		case reflect.Uint32:
			return newUintArraySetter(name, fld, 32)
		case reflect.Uint64:
			return newUintArraySetter(name, fld, 64)
		case reflect.Float32:
			return newFloatArraySetter(name, fld, 32)
		case reflect.Float64:
			return newFloatArraySetter(name, fld, 64)
		case reflect.String:
			return newStringArraySetter(name, fld)
		default:
			return newIllegalOptionTypeErr(name, t)
		}
	case reflect.String:
		return newStringSetter(name, fld)
	default:
		return newIllegalOptionTypeErr(name, t)
	}
}

func newIllegalOptionTypeErr(
	name string, t reflect.Type,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func([]string) sabi.Err {
		return sabi.Ok()
	}
	r := IllegalOptionType{Field: name, Type: t}
	return fn, sabi.NewErr(r)
}

func newBoolSetter(
	name string, fld reflect.Value,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if s != nil {
			fld.SetBool(true)
		}
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

func newIntSetter(
	name string, fld reflect.Value, bitSize int,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if len(s) == 0 {
			return sabi.Ok()
		}
		n, e := strconv.ParseInt(s[0], 0, bitSize)
		if e != nil {
			r := FailToParseInt{Field: name, Input: s[0], BitSize: bitSize}
			return sabi.NewErr(r, e)
		}
		fld.SetInt(n)
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

func newUintSetter(
	name string, fld reflect.Value, bitSize int,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if len(s) == 0 {
			return sabi.Ok()
		}
		n, e := strconv.ParseUint(s[0], 0, bitSize)
		if e != nil {
			r := FailToParseUint{Field: name, Input: s[0], BitSize: bitSize}
			return sabi.NewErr(r, e)
		}
		fld.SetUint(n)
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

func newFloatSetter(
	name string, fld reflect.Value, bitSize int,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if len(s) == 0 {
			return sabi.Ok()
		}
		n, e := strconv.ParseFloat(s[0], bitSize)
		if e != nil {
			r := FailToParseFloat{Field: name, Input: s[0], BitSize: bitSize}
			return sabi.NewErr(r, e)
		}
		fld.SetFloat(n)
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

func newStringSetter(
	name string, fld reflect.Value,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if len(s) == 0 {
			return sabi.Ok()
		}
		fld.SetString(s[0])
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

func newIntArraySetter(
	name string, fld reflect.Value, bitSize int,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if s == nil {
			return sabi.Ok()
		}
		emp := reflect.MakeSlice(fld.Type(), 0, 0)
		n := len(s)
		if n == 0 {
			fld.Set(emp)
			return sabi.Ok()
		}
		t := fld.Type().Elem()
		a := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			v, e := strconv.ParseInt(s[i], 0, bitSize)
			if e != nil {
				r := FailToParseInt{Field: name, Input: s[i], BitSize: bitSize}
				return sabi.NewErr(r, e)
			}
			a[i] = reflect.ValueOf(v).Convert(t)
		}
		fld.Set(reflect.Append(emp, a...))
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

func newUintArraySetter(
	name string, fld reflect.Value, bitSize int,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if s == nil {
			return sabi.Ok()
		}
		emp := reflect.MakeSlice(fld.Type(), 0, 0)
		n := len(s)
		if n == 0 { // If "=[]" then n==0, else if "=" then n==1 and s[0]=""
			fld.Set(emp)
			return sabi.Ok()
		}
		t := fld.Type().Elem()
		a := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			v, e := strconv.ParseUint(s[i], 0, bitSize)
			if e != nil {
				r := FailToParseUint{Field: name, Input: s[i], BitSize: bitSize}
				return sabi.NewErr(r, e)
			}
			a[i] = reflect.ValueOf(v).Convert(t)
		}
		fld.Set(reflect.Append(emp, a...))
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

func newFloatArraySetter(
	name string, fld reflect.Value, bitSize int,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if s == nil {
			return sabi.Ok()
		}
		emp := reflect.MakeSlice(fld.Type(), 0, 0)
		n := len(s)
		if n == 0 { // If "=[]" then n==0, else if "=" then n==1 and s[0]=""
			fld.Set(emp)
			return sabi.Ok()
		}
		t := fld.Type().Elem()
		a := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			v, e := strconv.ParseFloat(s[i], bitSize)
			if e != nil {
				r := FailToParseFloat{Field: name, Input: s[i], BitSize: bitSize}
				return sabi.NewErr(r, e)
			}
			a[i] = reflect.ValueOf(v).Convert(t)
		}
		fld.Set(reflect.Append(emp, a...))
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

func newStringArraySetter(
	name string, fld reflect.Value,
) (func([]string) sabi.Err, sabi.Err) {
	fn := func(s []string) sabi.Err {
		if s == nil {
			return sabi.Ok()
		}
		emp := reflect.MakeSlice(fld.Type(), 0, 0)
		n := len(s)
		if n == 0 { // If "=[]" then n==0, else if "=" then n==1 and s[0]=""
			fld.Set(emp)
			return sabi.Ok()
		}
		a := make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			a[i] = reflect.ValueOf(s[i])
		}
		fld.Set(reflect.Append(emp, a...))
		return sabi.Ok()
	}
	return fn, sabi.Ok()
}

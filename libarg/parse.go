// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package libarg

import (
	"github.com/sttk-go/sabi"
	"os"
	"strings"
	"unicode"
)

type /* error reason */ (
	// OptionHasInvalidChar is an error reason which indicates that an invalid
	// character is found in the option.
	OptionHasInvalidChar struct{ Option string }
)

var (
	//noentry []string = nil
	empty            = make([]string, 0)
	rangeOfAlphabets = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x0041, 0x005a, 1}, // A-Z
			{0x0061, 0x007a, 1}, // a-z
		},
	}
	rangeOfAlNumMarks = &unicode.RangeTable{
		R16: []unicode.Range16{
			{0x002d, 0x002d, 1}, // -
			{0x0030, 0x0039, 1}, // 0-9
			{0x0041, 0x005a, 1}, // A-Z
			{0x0061, 0x007a, 1}, // a-z
		},
	}
)

// Args is a structure which contains command parameters and option parameters
// that are parsed from command line arguments without configurations.
// And this provides methods to check if they are specified or to obtain them.
type Args struct {
	optParams map[string][]string
	cmdParams []string
}

// HasOpt is a method which checks if the option is specified in command line
// arguments.
func (args Args) HasOpt(opt string) bool {
	_, exists := args.optParams[opt]
	return exists
}

// OptParam is a method to get a option parameter which is firstly specified
// with opt in command line arguments.
func (args Args) OptParam(opt string) string {
	arr := args.optParams[opt]
	// If no entry, map returns a nil slice.
	// If a value of a found entry is an empty slice.
	// Both returned values are zero length in common.
	if len(arr) == 0 {
		return ""
	} else {
		return arr[0]
	}
}

// OptParams is a method to get option parameters which are all specified with
// opt in command line arguments.
func (args Args) OptParams(opt string) []string {
	return args.optParams[opt]
}

// CmdParams is a method to get command parameters which are specified in
// command line parameters and are not associated with any options.
func (args Args) CmdParams() []string {
	return args.cmdParams
}

// Parse is a function to parse command line arguments without configurations.
// This function divides command line arguments to command parameters, which
// are not associated with any options, and options, of which each has a name
// and option parameters.
// If an option appears multiple times in command line arguments, the option
// has multiple option parameters.
// Options are divided to long format options and short format options.
//
// A long format option starts with "--" and follows multiple characters which
// consists of alphabets, numbers, and '-'.
// (A character immediately after the heading "--" allows only an alphabet.)
// A long format option can be followed by "=" and its option parameter.
//
// A short format option starts with "-" and follows single character which is
// an alphabet.
// Multiple short options can be combined into one argument.
// (For example -a -b -c can be combined into -abc.)
// Moreover, a short option can be followed by "=" and its option parameter.
// In case of combined short options, only the last short option can take an
// option parameter.
// (For example, -abc=3 is equal to -a -b -c=3.)
//
// Usage example:
//
//	// os.Args[1:]  ==>  [--foo-bar=A -a --baz -bc=3 qux -c=4 quux]
//	args, _ := Parse()
//	args.HasOpt("a")          // true
//	args.HasOpt("b")          // true
//	args.HasOpt("c")          // true
//	args.HasOpt("foo-bar")    // true
//	args.HasOpt("baz")        // true
//	args.OptParam("c")        // 3
//	args.OptParams("c")       // [3 4]
//	args.CmdParams()          // [qux quux]
func Parse() (Args, sabi.Err) {
	var cmdParams = make([]string, 0)
	var optParams = make(map[string][]string)

	var collCmdParams = func(params ...string) sabi.Err {
		cmdParams = append(cmdParams, params...)
		return sabi.Ok()
	}
	var collOptParams = func(opt string, params ...string) sabi.Err {
		arr, exists := optParams[opt]
		if !exists {
			arr = empty
		}
		optParams[opt] = append(arr, params...)
		return sabi.Ok()
	}

	err := parseArgs(os.Args[1:], collCmdParams, collOptParams, _false)
	if !err.IsOk() {
		return Args{cmdParams: empty}, err
	}

	return Args{cmdParams: cmdParams, optParams: optParams}, err
}

func _false(_ string) bool {
	return false
}

func parseArgs(
	args []string,
	collectCmdParams func(...string) sabi.Err,
	collectOptParams func(string, ...string) sabi.Err,
	takeParams func(string) bool,
) sabi.Err {

	isNonOpt := false
	prevOptTakingParams := ""

	for iArg, arg := range args {
		if isNonOpt {
			err := collectCmdParams(arg)
			if !err.IsOk() {
				return err
			}

		} else if len(prevOptTakingParams) > 0 {
			err := collectOptParams(prevOptTakingParams, arg)
			if !err.IsOk() {
				return err
			}
			prevOptTakingParams = ""

		} else if strings.HasPrefix(arg, "--") {
			if len(arg) == 2 {
				isNonOpt = true
				continue
			}

			arg = arg[2:]
			i := 0
			for _, r := range arg {
				if i > 0 {
					if r == '=' {
						err := collectOptParams(arg[0:i], arg[i+1:])
						if !err.IsOk() {
							return err
						}
						break
					}
					if !unicode.Is(rangeOfAlNumMarks, r) {
						return sabi.NewErr(OptionHasInvalidChar{Option: arg})
					}
				} else {
					if !unicode.Is(rangeOfAlphabets, r) {
						return sabi.NewErr(OptionHasInvalidChar{Option: arg})
					}
				}
				i++
			}

			if i == len(arg) {
				if takeParams(arg) && iArg < len(args)-1 {
					prevOptTakingParams = arg
					continue
				}
				err := collectOptParams(arg)
				if !err.IsOk() {
					return err
				}
			}

		} else if strings.HasPrefix(arg, "-") {
			if len(arg) == 1 {
				err := collectCmdParams(arg)
				if !err.IsOk() {
					return err
				}
				continue
			}

			arg := arg[1:]
			var opt string
			i := 0
			for _, r := range arg {
				if i > 0 {
					if r == '=' {
						err := collectOptParams(opt, arg[i+1:])
						if !err.IsOk() {
							return err
						}
						break
					}
					err := collectOptParams(opt)
					if !err.IsOk() {
						return err
					}
				}
				opt = string(r)
				if !unicode.Is(rangeOfAlphabets, r) {
					return sabi.NewErr(OptionHasInvalidChar{Option: opt})
				}
				i++
			}

			if i == len(arg) {
				if takeParams(opt) && iArg < len(args)-1 {
					prevOptTakingParams = opt
				} else {
					err := collectOptParams(opt)
					if !err.IsOk() {
						return err
					}
				}
			}

		} else {
			err := collectCmdParams(arg)
			if !err.IsOk() {
				return err
			}
		}
	}

	return sabi.Ok()
}

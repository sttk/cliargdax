package libarg_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/clidax/libarg"
	"reflect"
	"testing"
)

func TestParseFor_emptyOptionStoreAndNoArgs(t *testing.T) {
	type MyOptions struct{}
	args := []string{}
	options := MyOptions{}
	cmdParams, err := libarg.ParseFor(args, &options)
	assert.True(t, err.IsOk())
	assert.Equal(t, cmdParams, []string{})
}

func TestParseFor_nonEmptyOptionStoreAndNoArgs(t *testing.T) {
	type MyOptions struct {
		BoolVal    bool
		IntVal     int
		Int8Val    int8
		Int16Val   int16
		Int32Val   int32
		Int64Val   int64
		UintVal    uint
		Uint8Val   uint8
		Uint16Val  uint16
		Uint32Val  uint32
		Uint64Val  uint64
		Float32Val float32
		Float64Val float64
		StringVal  string
		IntArr     []int
		Int8Arr    []int8
		Int16Arr   []int16
		Int32Arr   []int32
		Int64Arr   []int64
		UintArr    []uint
		Uint8Arr   []uint8
		Uint16Arr  []uint16
		Uint32Arr  []uint32
		Uint64Arr  []uint64
		Float32Arr []float32
		Float64Arr []float64
		StringArr  []string
	}
	options := MyOptions{}

	args := []string{}
	cmdParams, err := libarg.ParseFor(args, &options)
	assert.True(t, err.IsOk())
	assert.Equal(t, cmdParams, []string{})
	assert.False(t, options.BoolVal)
	assert.Equal(t, options.IntVal, 0)
	assert.Equal(t, options.Int8Val, int8(0))
	assert.Equal(t, options.Int16Val, int16(0))
	assert.Equal(t, options.Int32Val, int32(0))
	assert.Equal(t, options.Int64Val, int64(0))
	assert.Equal(t, options.UintVal, uint(0))
	assert.Equal(t, options.Uint8Val, uint8(0))
	assert.Equal(t, options.Uint16Val, uint16(0))
	assert.Equal(t, options.Uint32Val, uint32(0))
	assert.Equal(t, options.Uint64Val, uint64(0))
	assert.Equal(t, options.Float32Val, float32(0.0))
	assert.Equal(t, options.Float64Val, 0.0)
	assert.Equal(t, options.StringVal, "")
	assert.Equal(t, options.IntArr, []int(nil))
	assert.Equal(t, options.Int8Arr, []int8(nil))
	assert.Equal(t, options.Int16Arr, []int16(nil))
	assert.Equal(t, options.Int32Arr, []int32(nil))
	assert.Equal(t, options.Int64Arr, []int64(nil))
	assert.Equal(t, options.UintArr, []uint(nil))
	assert.Equal(t, options.Uint8Arr, []uint8(nil))
	assert.Equal(t, options.Uint16Arr, []uint16(nil))
	assert.Equal(t, options.Uint32Arr, []uint32(nil))
	assert.Equal(t, options.Uint64Arr, []uint64(nil))
	assert.Equal(t, options.Float32Arr, []float32(nil))
	assert.Equal(t, options.Float64Arr, []float64(nil))
	assert.Equal(t, options.StringArr, []string(nil))
}

func TestParseFor_dontOverwriteOptionsIfNoArgs(t *testing.T) {
	type MyOptions struct {
		BoolVal    bool
		IntVal     int
		Int8Val    int8
		Int16Val   int16
		Int32Val   int32
		Int64Val   int64
		UintVal    uint
		Uint8Val   uint8
		Uint16Val  uint16
		Uint32Val  uint32
		Uint64Val  uint64
		Float32Val float32
		Float64Val float64
		StringVal  string
		IntArr     []int
		Int8Arr    []int8
		Int16Arr   []int16
		Int32Arr   []int32
		Int64Arr   []int64
		UintArr    []uint
		Uint8Arr   []uint8
		Uint16Arr  []uint16
		Uint32Arr  []uint32
		Uint64Arr  []uint64
		Float32Arr []float32
		Float64Arr []float64
		StringArr  []string
	}
	options := MyOptions{
		BoolVal:    true,
		IntVal:     111,
		Int8Val:    22,
		Int16Val:   333,
		Int32Val:   444,
		Int64Val:   555,
		UintVal:    666,
		Uint8Val:   77,
		Uint16Val:  888,
		Uint32Val:  999,
		Uint64Val:  1111,
		Float32Val: 0.123,
		Float64Val: 0.456789,
		StringVal:  "abcdefg",
		IntArr:     []int{1, 1, 1},
		Int8Arr:    []int8{2, 2},
		Int16Arr:   []int16{3, 3, 3},
		Int32Arr:   []int32{4, 4, 4},
		Int64Arr:   []int64{5, 5, 5},
		UintArr:    []uint{6, 6, 6},
		Uint8Arr:   []uint8{7, 7},
		Uint16Arr:  []uint16{8, 8, 8},
		Uint32Arr:  []uint32{9, 9, 9},
		Uint64Arr:  []uint64{1, 1, 1, 1},
		Float32Arr: []float32{0.1, 2.3},
		Float64Arr: []float64{0.45, 6.789},
		StringArr:  []string{"ab", "cd", "efg"},
	}

	args := []string{}
	cmdParams, err := libarg.ParseFor(args, &options)
	assert.True(t, err.IsOk())
	assert.Equal(t, cmdParams, []string{})
	assert.True(t, options.BoolVal)
	assert.Equal(t, options.IntVal, 111)
	assert.Equal(t, options.Int8Val, int8(22))
	assert.Equal(t, options.Int16Val, int16(333))
	assert.Equal(t, options.Int32Val, int32(444))
	assert.Equal(t, options.Int64Val, int64(555))
	assert.Equal(t, options.UintVal, uint(666))
	assert.Equal(t, options.Uint8Val, uint8(77))
	assert.Equal(t, options.Uint16Val, uint16(888))
	assert.Equal(t, options.Uint32Val, uint32(999))
	assert.Equal(t, options.Uint64Val, uint64(1111))
	assert.Equal(t, options.Float32Val, float32(0.123))
	assert.Equal(t, options.Float64Val, 0.456789)
	assert.Equal(t, options.StringVal, "abcdefg")
	assert.Equal(t, options.IntArr, []int{1, 1, 1})
	assert.Equal(t, options.Int8Arr, []int8{2, 2})
	assert.Equal(t, options.Int16Arr, []int16{3, 3, 3})
	assert.Equal(t, options.Int32Arr, []int32{4, 4, 4})
	assert.Equal(t, options.Int64Arr, []int64{5, 5, 5})
	assert.Equal(t, options.UintArr, []uint{6, 6, 6})
	assert.Equal(t, options.Uint8Arr, []uint8{7, 7})
	assert.Equal(t, options.Uint16Arr, []uint16{8, 8, 8})
	assert.Equal(t, options.Uint32Arr, []uint32{9, 9, 9})
	assert.Equal(t, options.Uint64Arr, []uint64{1, 1, 1, 1})
	assert.Equal(t, options.Float32Arr, []float32{0.1, 2.3})
	assert.Equal(t, options.Float64Arr, []float64{0.45, 6.789})
	assert.Equal(t, options.StringArr, []string{"ab", "cd", "efg"})
}

func TestParseFor_optionIsBoolAndArgIsName(t *testing.T) {
	type MyOptions struct {
		Flag bool `opt:"flag,f"`
	}
	options := MyOptions{}

	args := []string{"--flag", "abc"}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.True(t, options.Flag)
}

func TestParseFor_optionIsBoolAndArgIsAlias(t *testing.T) {
	type MyOptions struct {
		Flag bool `opt:"flag,f"`
	}
	options := MyOptions{}

	args := []string{"-f", "abc"}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.True(t, options.Flag)
}

func TestParseFor_optionsAreIntAndArgsAreNames(t *testing.T) {
	type MyOptions struct {
		IntVal   int   `opt:"int-val,i"`
		Int8Val  int8  `opt:"int8-val,j"`
		Int16Val int16 `opt:"int16-val,k"`
		Int32Val int32 `opt:"int32-val,m"`
		Int64Val int64 `opt:"int64-val,n"`
	}
	options := MyOptions{}

	args := []string{
		"--int-val", "1",
		"--int8-val", "2",
		"--int16-val", "3",
		"--int32-val", "4",
		"--int64-val", "5",
		"abc",
	}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.Equal(t, options.IntVal, 1)
	assert.Equal(t, options.Int8Val, int8(2))
	assert.Equal(t, options.Int16Val, int16(3))
	assert.Equal(t, options.Int32Val, int32(4))
	assert.Equal(t, options.Int64Val, int64(5))
}

func TestParseFor_optionsAreIntAndArgsAreAliases(t *testing.T) {
	type MyOptions struct {
		IntVal   int   `opt:"int-val,i"`
		Int8Val  int8  `opt:"int8-val,j"`
		Int16Val int16 `opt:"int16-val,k"`
		Int32Val int32 `opt:"int32-val,m"`
		Int64Val int64 `opt:"int64-val,n"`
	}
	options := MyOptions{}

	args := []string{
		"-i", "1",
		"-j", "2",
		"-k", "3",
		"-m", "4",
		"-n", "5",
		"abc",
	}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.Equal(t, options.IntVal, 1)
	assert.Equal(t, options.Int8Val, int8(2))
	assert.Equal(t, options.Int16Val, int16(3))
	assert.Equal(t, options.Int32Val, int32(4))
	assert.Equal(t, options.Int64Val, int64(5))
}

func TestParseFor_optionsAreUintAndArgsAreNames(t *testing.T) {
	type MyOptions struct {
		UintVal   uint   `opt:"uint-val,i"`
		Uint8Val  uint8  `opt:"uint8-val,j"`
		Uint16Val uint16 `opt:"uint16-val,k"`
		Uint32Val uint32 `opt:"uint32-val,m"`
		Uint64Val uint64 `opt:"uint64-val,n"`
	}
	options := MyOptions{}

	args := []string{
		"--uint-val", "1",
		"--uint8-val", "2",
		"--uint16-val", "3",
		"--uint32-val", "4",
		"--uint64-val", "5",
		"abc",
	}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.Equal(t, options.UintVal, uint(1))
	assert.Equal(t, options.Uint8Val, uint8(2))
	assert.Equal(t, options.Uint16Val, uint16(3))
	assert.Equal(t, options.Uint32Val, uint32(4))
	assert.Equal(t, options.Uint64Val, uint64(5))
}

func TestParseFor_optionsAreUintAndArgsAreAliases(t *testing.T) {
	type MyOptions struct {
		UintVal   uint   `opt:"uint-val,i"`
		Uint8Val  uint8  `opt:"uint8-val,j"`
		Uint16Val uint16 `opt:"uint16-val,k"`
		Uint32Val uint32 `opt:"uint32-val,m"`
		Uint64Val uint64 `opt:"uint64-val,n"`
	}
	options := MyOptions{}

	args := []string{
		"-i", "1",
		"-j", "2",
		"-k", "3",
		"-m", "4",
		"-n", "5",
		"abc",
	}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.Equal(t, options.UintVal, uint(1))
	assert.Equal(t, options.Uint8Val, uint8(2))
	assert.Equal(t, options.Uint16Val, uint16(3))
	assert.Equal(t, options.Uint32Val, uint32(4))
	assert.Equal(t, options.Uint64Val, uint64(5))
}

func TestParseFor_optionsAreFloatAndArgsAreNames(t *testing.T) {
	type MyOptions struct {
		Float32Val float32 `opt:"float32-val,m"`
		Float64Val float64 `opt:"float64-val,n"`
	}
	options := MyOptions{}

	args := []string{
		"--float32-val", "0.1234",
		"--float64-val", "0.5678",
		"abc",
	}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.Equal(t, options.Float32Val, float32(0.1234))
	assert.Equal(t, options.Float64Val, 0.5678)
}

func TestParseFor_optionsAreFloatAndArgsAreAliases(t *testing.T) {
	type MyOptions struct {
		Float32Val float32 `opt:"float32-val,m"`
		Float64Val float64 `opt:"float64-val,n"`
	}
	options := MyOptions{}

	args := []string{
		"-m", "0.1234",
		"-n", "0.5678",
		"abc",
	}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.Equal(t, options.Float32Val, float32(0.1234))
	assert.Equal(t, options.Float64Val, 0.5678)
}

func TestParseFor_optionsAreStringAndArgsAreNames(t *testing.T) {
	type MyOptions struct {
		StringVal string `opt:"string-val,s"`
	}
	options := MyOptions{}

	args := []string{
		"--string-val", "def",
		"abc",
	}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.Equal(t, options.StringVal, "def")
}

func TestParseFor_optionsAreStringAndArgsAreAliases(t *testing.T) {
	type MyOptions struct {
		StringVal string `opt:"string-val,s"`
	}
	options := MyOptions{}

	args := []string{
		"-s", "def",
		"abc",
	}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{"abc"})
	assert.Equal(t, options.StringVal, "def")
}

func TestParseFor_defaultValueIsInt(t *testing.T) {
	type MyOptions struct {
		IntVal   int   `opt:"=11"`
		Int8Val  int8  `opt:"=22"`
		Int16Val int16 `opt:"=33"`
		Int32Val int32 `opt:"=44"`
		Int64Val int64 `opt:"=55"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntVal, 11)
	assert.Equal(t, options.Int8Val, int8(22))
	assert.Equal(t, options.Int16Val, int16(33))
	assert.Equal(t, options.Int32Val, int32(44))
	assert.Equal(t, options.Int64Val, int64(55))
}

func TestParseFor_defaultValueIsNegativeInt(t *testing.T) {
	type MyOptions struct {
		IntVal   int   `opt:"=-11"`
		Int8Val  int8  `opt:"=-22"`
		Int16Val int16 `opt:"=-33"`
		Int32Val int32 `opt:"=-44"`
		Int64Val int64 `opt:"=-55"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntVal, -11)
	assert.Equal(t, options.Int8Val, int8(-22))
	assert.Equal(t, options.Int16Val, int16(-33))
	assert.Equal(t, options.Int32Val, int32(-44))
	assert.Equal(t, options.Int64Val, int64(-55))
}

func TestParseFor_defaultValueIsUint(t *testing.T) {
	type MyOptions struct {
		UintVal   uint   `opt:"=11"`
		Uint8Val  uint8  `opt:"=22"`
		Uint16Val uint16 `opt:"=33"`
		Uint32Val uint32 `opt:"=44"`
		Uint64Val uint64 `opt:"=55"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.UintVal, uint(11))
	assert.Equal(t, options.Uint8Val, uint8(22))
	assert.Equal(t, options.Uint16Val, uint16(33))
	assert.Equal(t, options.Uint32Val, uint32(44))
	assert.Equal(t, options.Uint64Val, uint64(55))
}

func TestParseFor_defaultValueIsFloat(t *testing.T) {
	type MyOptions struct {
		Float32Val float32 `opt:"=0.123"`
		Float64Val float64 `opt:"=0.456789"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Val, float32(0.123))
	assert.Equal(t, options.Float64Val, float64(0.456789))
}

func TestParseFor_defaultValueIsNegativeFloat(t *testing.T) {
	type MyOptions struct {
		Float32Val float32 `opt:"=-0.123"`
		Float64Val float64 `opt:"=-0.456789"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Val, float32(-0.123))
	assert.Equal(t, options.Float64Val, float64(-0.456789))
}

func TestParseFor_defaultValueIsIntArrayAndSize0(t *testing.T) {
	type MyOptions struct {
		IntArr   []int   `opt:"=[]"`
		Int8Arr  []int8  `opt:"=[]"`
		Int16Arr []int16 `opt:"=[]"`
		Int32Arr []int32 `opt:"=[]"`
		Int64Arr []int64 `opt:"=[]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntArr, []int{})
	assert.Equal(t, options.Int8Arr, []int8{})
	assert.Equal(t, options.Int16Arr, []int16{})
	assert.Equal(t, options.Int32Arr, []int32{})
	assert.Equal(t, options.Int64Arr, []int64{})
}

func TestParseFor_overwriteIntArrayWithDefaultValueIfSize0(t *testing.T) {
	type MyOptions struct {
		IntArr   []int   `opt:"=[]"`
		Int8Arr  []int8  `opt:"=[]"`
		Int16Arr []int16 `opt:"=[]"`
		Int32Arr []int32 `opt:"=[]"`
		Int64Arr []int64 `opt:"=[]"`
	}
	options := MyOptions{
		IntArr:   []int{1},
		Int8Arr:  []int8{2},
		Int16Arr: []int16{3},
		Int32Arr: []int32{4},
		Int64Arr: []int64{5},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntArr, []int{})
	assert.Equal(t, options.Int8Arr, []int8{})
	assert.Equal(t, options.Int16Arr, []int16{})
	assert.Equal(t, options.Int32Arr, []int32{})
	assert.Equal(t, options.Int64Arr, []int64{})
}

func TestParseFor_defaultValueIsIntArrayAndSize1(t *testing.T) {
	type MyOptions struct {
		IntArr   []int   `opt:"=[1]"`
		Int8Arr  []int8  `opt:"=[2]"`
		Int16Arr []int16 `opt:"=[3]"`
		Int32Arr []int32 `opt:"=[4]"`
		Int64Arr []int64 `opt:"=[5]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntArr, []int{1})
	assert.Equal(t, options.Int8Arr, []int8{2})
	assert.Equal(t, options.Int16Arr, []int16{3})
	assert.Equal(t, options.Int32Arr, []int32{4})
	assert.Equal(t, options.Int64Arr, []int64{5})
}

func TestParseFor_overwriteIntArrayWithDefaultValueIfSize1(t *testing.T) {
	type MyOptions struct {
		IntArr   []int   `opt:"=[1]"`
		Int8Arr  []int8  `opt:"=[2]"`
		Int16Arr []int16 `opt:"=[3]"`
		Int32Arr []int32 `opt:"=[4]"`
		Int64Arr []int64 `opt:"=[5]"`
	}
	options := MyOptions{
		IntArr:   []int{11},
		Int8Arr:  []int8{22},
		Int16Arr: []int16{33},
		Int32Arr: []int32{44},
		Int64Arr: []int64{55},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntArr, []int{1})
	assert.Equal(t, options.Int8Arr, []int8{2})
	assert.Equal(t, options.Int16Arr, []int16{3})
	assert.Equal(t, options.Int32Arr, []int32{4})
	assert.Equal(t, options.Int64Arr, []int64{5})
}

func TestParseFor_defaultValueIsIntArrayAndSize2(t *testing.T) {
	type MyOptions struct {
		IntArr   []int   `opt:"=[1,2]"`
		Int8Arr  []int8  `opt:"=[2,3]"`
		Int16Arr []int16 `opt:"=[3,4]"`
		Int32Arr []int32 `opt:"=[4,5]"`
		Int64Arr []int64 `opt:"=[5,6]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntArr, []int{1, 2})
	assert.Equal(t, options.Int8Arr, []int8{2, 3})
	assert.Equal(t, options.Int16Arr, []int16{3, 4})
	assert.Equal(t, options.Int32Arr, []int32{4, 5})
	assert.Equal(t, options.Int64Arr, []int64{5, 6})
}

func TestParseFor_overwriteIntArrayWithDefaultValueIfSize2(t *testing.T) {
	type MyOptions struct {
		IntArr   []int   `opt:"=[1,2]"`
		Int8Arr  []int8  `opt:"=[2,3]"`
		Int16Arr []int16 `opt:"=[3,4]"`
		Int32Arr []int32 `opt:"=[4,5]"`
		Int64Arr []int64 `opt:"=[5,6]"`
	}
	options := MyOptions{
		IntArr:   []int{11},
		Int8Arr:  []int8{22},
		Int16Arr: []int16{33},
		Int32Arr: []int32{44},
		Int64Arr: []int64{55},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntArr, []int{1, 2})
	assert.Equal(t, options.Int8Arr, []int8{2, 3})
	assert.Equal(t, options.Int16Arr, []int16{3, 4})
	assert.Equal(t, options.Int32Arr, []int32{4, 5})
	assert.Equal(t, options.Int64Arr, []int64{5, 6})
}

func TestParseFor_defaultValueIsNegativeIntArray(t *testing.T) {
	type MyOptions struct {
		IntArr   []int   `opt:"=[-1,-2]"`
		Int8Arr  []int8  `opt:"=[-2,-3]"`
		Int16Arr []int16 `opt:"=[-3,-4]"`
		Int32Arr []int32 `opt:"=[-4,-5]"`
		Int64Arr []int64 `opt:"=[-5,-6]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntArr, []int{-1, -2})
	assert.Equal(t, options.Int8Arr, []int8{-2, -3})
	assert.Equal(t, options.Int16Arr, []int16{-3, -4})
	assert.Equal(t, options.Int32Arr, []int32{-4, -5})
	assert.Equal(t, options.Int64Arr, []int64{-5, -6})
}

func TestParseFor_defaultValueIsIntArraySeparatedByColons(t *testing.T) {
	type MyOptions struct {
		IntArr   []int   `opt:"=:[-1:-2]"`
		Int8Arr  []int8  `opt:"=/[-2/-3]"`
		Int16Arr []int16 `opt:"=![-3!-4]"`
		Int32Arr []int32 `opt:"=|[-4|-5]"`
		Int64Arr []int64 `opt:"='[-5'-6]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.IntArr, []int{-1, -2})
	assert.Equal(t, options.Int8Arr, []int8{-2, -3})
	assert.Equal(t, options.Int16Arr, []int16{-3, -4})
	assert.Equal(t, options.Int32Arr, []int32{-4, -5})
	assert.Equal(t, options.Int64Arr, []int64{-5, -6})
}

func TestParseFor_defaultValueIsUintArrayAndSize0(t *testing.T) {
	type MyOptions struct {
		UintArr   []uint   `opt:"=[]"`
		Uint8Arr  []uint8  `opt:"=[]"`
		Uint16Arr []uint16 `opt:"=[]"`
		Uint32Arr []uint32 `opt:"=[]"`
		Uint64Arr []uint64 `opt:"=[]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.UintArr, []uint{})
	assert.Equal(t, options.Uint8Arr, []uint8{})
	assert.Equal(t, options.Uint16Arr, []uint16{})
	assert.Equal(t, options.Uint32Arr, []uint32{})
	assert.Equal(t, options.Uint64Arr, []uint64{})
}

func TestParseFor_overwriteUintArrayWithDefaultValueIfSize0(t *testing.T) {
	type MyOptions struct {
		UintArr   []uint   `opt:"=[]"`
		Uint8Arr  []uint8  `opt:"=[]"`
		Uint16Arr []uint16 `opt:"=[]"`
		Uint32Arr []uint32 `opt:"=[]"`
		Uint64Arr []uint64 `opt:"=[]"`
	}
	options := MyOptions{
		UintArr:   []uint{1},
		Uint8Arr:  []uint8{2},
		Uint16Arr: []uint16{3},
		Uint32Arr: []uint32{4},
		Uint64Arr: []uint64{5},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.UintArr, []uint{})
	assert.Equal(t, options.Uint8Arr, []uint8{})
	assert.Equal(t, options.Uint16Arr, []uint16{})
	assert.Equal(t, options.Uint32Arr, []uint32{})
	assert.Equal(t, options.Uint64Arr, []uint64{})
}

func TestParseFor_defaultValueIsUintArrayAndSize1(t *testing.T) {
	type MyOptions struct {
		UintArr   []uint   `opt:"=[1]"`
		Uint8Arr  []uint8  `opt:"=[2]"`
		Uint16Arr []uint16 `opt:"=[3]"`
		Uint32Arr []uint32 `opt:"=[4]"`
		Uint64Arr []uint64 `opt:"=[5]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.UintArr, []uint{1})
	assert.Equal(t, options.Uint8Arr, []uint8{2})
	assert.Equal(t, options.Uint16Arr, []uint16{3})
	assert.Equal(t, options.Uint32Arr, []uint32{4})
	assert.Equal(t, options.Uint64Arr, []uint64{5})
}

func TestParseFor_overwriteUintArrayWithDefaultValueIfSize1(t *testing.T) {
	type MyOptions struct {
		UintArr   []uint   `opt:"=[1]"`
		Uint8Arr  []uint8  `opt:"=[2]"`
		Uint16Arr []uint16 `opt:"=[3]"`
		Uint32Arr []uint32 `opt:"=[4]"`
		Uint64Arr []uint64 `opt:"=[5]"`
	}
	options := MyOptions{
		UintArr:   []uint{11},
		Uint8Arr:  []uint8{22},
		Uint16Arr: []uint16{33},
		Uint32Arr: []uint32{44},
		Uint64Arr: []uint64{55},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.UintArr, []uint{1})
	assert.Equal(t, options.Uint8Arr, []uint8{2})
	assert.Equal(t, options.Uint16Arr, []uint16{3})
	assert.Equal(t, options.Uint32Arr, []uint32{4})
	assert.Equal(t, options.Uint64Arr, []uint64{5})
}

func TestParseFor_defaultValueIsUintArrayAndSize2(t *testing.T) {
	type MyOptions struct {
		UintArr   []uint   `opt:"=[1,2]"`
		Uint8Arr  []uint8  `opt:"=[2,3]"`
		Uint16Arr []uint16 `opt:"=[3,4]"`
		Uint32Arr []uint32 `opt:"=[4,5]"`
		Uint64Arr []uint64 `opt:"=[5,6]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.UintArr, []uint{1, 2})
	assert.Equal(t, options.Uint8Arr, []uint8{2, 3})
	assert.Equal(t, options.Uint16Arr, []uint16{3, 4})
	assert.Equal(t, options.Uint32Arr, []uint32{4, 5})
	assert.Equal(t, options.Uint64Arr, []uint64{5, 6})
}

func TestParseFor_overwriteUintArrayWithDefaultValueIfSize2(t *testing.T) {
	type MyOptions struct {
		UintArr   []uint   `opt:"=[1,2]"`
		Uint8Arr  []uint8  `opt:"=[2,3]"`
		Uint16Arr []uint16 `opt:"=[3,4]"`
		Uint32Arr []uint32 `opt:"=[4,5]"`
		Uint64Arr []uint64 `opt:"=[5,6]"`
	}
	options := MyOptions{
		UintArr:   []uint{11},
		Uint8Arr:  []uint8{22},
		Uint16Arr: []uint16{33},
		Uint32Arr: []uint32{44},
		Uint64Arr: []uint64{55},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.UintArr, []uint{1, 2})
	assert.Equal(t, options.Uint8Arr, []uint8{2, 3})
	assert.Equal(t, options.Uint16Arr, []uint16{3, 4})
	assert.Equal(t, options.Uint32Arr, []uint32{4, 5})
	assert.Equal(t, options.Uint64Arr, []uint64{5, 6})
}

func TestParseFor_defaultValueIsUintArraySeparatedByColons(t *testing.T) {
	type MyOptions struct {
		UintArr   []uint   `opt:"=:[1:2]"`
		Uint8Arr  []uint8  `opt:"=/[2/3]"`
		Uint16Arr []uint16 `opt:"=![3!4]"`
		Uint32Arr []uint32 `opt:"=|[4|5]"`
		Uint64Arr []uint64 `opt:"='[5'6]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.UintArr, []uint{1, 2})
	assert.Equal(t, options.Uint8Arr, []uint8{2, 3})
	assert.Equal(t, options.Uint16Arr, []uint16{3, 4})
	assert.Equal(t, options.Uint32Arr, []uint32{4, 5})
	assert.Equal(t, options.Uint64Arr, []uint64{5, 6})
}

func TestParseFor_defaultValueIsFloatArrayAndSize0(t *testing.T) {
	type MyOptions struct {
		Float32Arr []float32 `opt:"=[]"`
		Float64Arr []float64 `opt:"=[]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Arr, []float32{})
	assert.Equal(t, options.Float64Arr, []float64{})
}

func TestParseFor_overwriteFloatArrayWithDefaultValueIfSize0(t *testing.T) {
	type MyOptions struct {
		Float32Arr []float32 `opt:"=[]"`
		Float64Arr []float64 `opt:"=[]"`
	}
	options := MyOptions{
		Float32Arr: []float32{0.999},
		Float64Arr: []float64{0.888},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Arr, []float32{})
	assert.Equal(t, options.Float64Arr, []float64{})
}

func TestParseFor_defaultValueIsFloatArrayAndSize1(t *testing.T) {
	type MyOptions struct {
		Float32Arr []float32 `opt:"=[0.1]"`
		Float64Arr []float64 `opt:"=[0.2]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Arr, []float32{0.1})
	assert.Equal(t, options.Float64Arr, []float64{0.2})
}

func TestParseFor_overwriteFloatArrayWithDefaultValueIfSize1(t *testing.T) {
	type MyOptions struct {
		Float32Arr []float32 `opt:"=[0.1]"`
		Float64Arr []float64 `opt:"=[0.2]"`
	}
	options := MyOptions{
		Float32Arr: []float32{0.99},
		Float64Arr: []float64{0.88},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Arr, []float32{0.1})
	assert.Equal(t, options.Float64Arr, []float64{0.2})
}

func TestParseFor_defaultValueIsFloatArrayAndSize2(t *testing.T) {
	type MyOptions struct {
		Float32Arr []float32 `opt:"=[0.1,0.2]"`
		Float64Arr []float64 `opt:"=[0.3,0.4]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Arr, []float32{0.1, 0.2})
	assert.Equal(t, options.Float64Arr, []float64{0.3, 0.4})
}

func TestParseFor_overwriteFloatArrayWithDefaultValueIfSize2(t *testing.T) {
	type MyOptions struct {
		Float32Arr []float32 `opt:"=[0.1,0.2]"`
		Float64Arr []float64 `opt:"=[0.3,0.4]"`
	}
	options := MyOptions{
		Float32Arr: []float32{0.99},
		Float64Arr: []float64{0.88},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Arr, []float32{0.1, 0.2})
	assert.Equal(t, options.Float64Arr, []float64{0.3, 0.4})
}

func TestParseFor_defaultValueIsNegativeFloatArray(t *testing.T) {
	type MyOptions struct {
		Float32Arr []float32 `opt:"=[-0.1,-0.2]"`
		Float64Arr []float64 `opt:"=[-0.3,-0.4]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Arr, []float32{-0.1, -0.2})
	assert.Equal(t, options.Float64Arr, []float64{-0.3, -0.4})
}

func TestParseFor_defaultValueIsFloatArraySeparatedByColons(t *testing.T) {
	type MyOptions struct {
		Float32Arr []float32 `opt:"=|[-0.1|-0.2]"`
		Float64Arr []float64 `opt:"='[-0.3'-0.4]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.Float32Arr, []float32{-0.1, -0.2})
	assert.Equal(t, options.Float64Arr, []float64{-0.3, -0.4})
}

func TestParseFor_defaultValueIsStringAndSize0(t *testing.T) {
	type MyOptions struct {
		StringArr []string `opt:"=[]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringArr, []string{})
}

func TestParseFor_overwriteStringArrayWithDefaultValueIfSize0(t *testing.T) {
	type MyOptions struct {
		StringArr []string `opt:"=[]"`
	}
	options := MyOptions{
		StringArr: []string{"ZZZ"},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringArr, []string{})
}

func TestParseFor_defaultValueIsStringArrayAndSize1(t *testing.T) {
	type MyOptions struct {
		StringArr []string `opt:"=[ABC]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringArr, []string{"ABC"})
}

func TestParseFor_overwriteStringArrayWithDefaultValueIfSize1(t *testing.T) {
	type MyOptions struct {
		StringArr []string `opt:"=[ABC]"`
	}
	options := MyOptions{
		StringArr: []string{"ZZZ"},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringArr, []string{"ABC"})
}

func TestParseFor_defaultValueIsStringArrayAndSize2(t *testing.T) {
	type MyOptions struct {
		StringArr []string `opt:"=[ABC,DEF]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringArr, []string{"ABC", "DEF"})
}

func TestParseFor_overwriteStringArrayWithDefaultValueIfSize2(t *testing.T) {
	type MyOptions struct {
		StringArr []string `opt:"=[ABC,DEF]"`
	}
	options := MyOptions{
		StringArr: []string{"ZZZ"},
	}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringArr, []string{"ABC", "DEF"})
}

func TestParseFor_defaultValueIsStringArraySeparatedByColons(t *testing.T) {
	type MyOptions struct {
		StringArr []string `opt:"=|[ABC|DEF]"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringArr, []string{"ABC", "DEF"})
}

func TestParseFor_ignoreEmptyDefaultValueIfOptionIsBool(t *testing.T) {
	type MyOptions struct {
		BoolVar bool `opt:"="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.False(t, options.BoolVar)
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsInt(t *testing.T) {
	type MyOptions struct {
		IntVar int `opt:"="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.Equal(t, params, []string{})
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.FailToParseInt:
		assert.Equal(t, err.Get("Field"), "IntVar")
		assert.Equal(t, err.Get("Input"), "")
		assert.Equal(t, err.Get("BitSize"), 64)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsUint(t *testing.T) {
	type MyOptions struct {
		UintVar uint `opt:"="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.Equal(t, params, []string{})
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.FailToParseUint:
		assert.Equal(t, err.Get("Field"), "UintVar")
		assert.Equal(t, err.Get("Input"), "")
		assert.Equal(t, err.Get("BitSize"), 64)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsFloat(t *testing.T) {
	type MyOptions struct {
		Float64Var float64 `opt:"="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.Equal(t, params, []string{})
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.FailToParseFloat:
		assert.Equal(t, err.Get("Field"), "Float64Var")
		assert.Equal(t, err.Get("Input"), "")
		assert.Equal(t, err.Get("BitSize"), 64)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsString(t *testing.T) {
	type MyOptions struct {
		StringVar string `opt:"="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringVar, "")
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsIntArray(t *testing.T) {
	type MyOptions struct {
		IntArr []int `opt:"="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.Equal(t, params, []string{})
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.FailToParseInt:
		assert.Equal(t, err.Get("Field"), "IntArr")
		assert.Equal(t, err.Get("Input"), "")
		assert.Equal(t, err.Get("BitSize"), 64)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsUintArray(t *testing.T) {
	type MyOptions struct {
		UintArr []uint `opt:"="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.Equal(t, params, []string{})
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.FailToParseUint:
		assert.Equal(t, err.Get("Field"), "UintArr")
		assert.Equal(t, err.Get("Input"), "")
		assert.Equal(t, err.Get("BitSize"), 64)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_errorEmptyDefaultValueIfOptionIsFloatArray(t *testing.T) {
	type MyOptions struct {
		Float64Arr []float64 `opt:"="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.Equal(t, params, []string{})
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.FailToParseFloat:
		assert.Equal(t, err.Get("Field"), "Float64Arr")
		assert.Equal(t, err.Get("Input"), "")
		assert.Equal(t, err.Get("BitSize"), 64)
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_optionIsStringArrayAndSetOneEmptyStringByDefaultArray(t *testing.T) {
	type MyOptions struct {
		StringArr []string `opt:"str-arr="`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.Equal(t, options.StringArr, []string{""})
}

func TestParseFor_defaultValueIsIgnoreWhenTypeIsBool(t *testing.T) {
	type MyOptions struct {
		BoolVar bool `opt:"bool-var=true"`
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, params, []string{})
	assert.False(t, options.BoolVar)
}

func TestParseFor_errorIfDefaultValueIsInvalidType(t *testing.T) {
	type MyOptions struct {
		BoolArr []bool
	}
	options := MyOptions{}

	args := []string{}
	params, err := libarg.ParseFor(args, &options)

	assert.Equal(t, params, []string{})
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.IllegalOptionType:
		assert.Equal(t, err.Get("Field"), "BoolArr")
		assert.Equal(t, err.Get("Type"), reflect.TypeOf(options.BoolArr))
	default:
		assert.Fail(t, err.Error())
	}
}

func TestParseFor_multipleOptsAndMultipleArgs(t *testing.T) {
	type MyOptions struct {
		FooBar bool     `opt:"foo-bar,f"`
		Baz    int      `opt:"baz,b=99"`
		Qux    string   `opt:"=XXX"`
		Quux   []string `opt:"quux=/[A/B/C]"`
		Corge  []int
	}
	options := MyOptions{}

	args := []string{
		"--foo-bar", "c1", "-b", "12", "--Qux", "ABC", "c2",
		"--Corge", "20", "--Corge=21",
	}

	cmdParams, err := libarg.ParseFor(args, &options)

	assert.True(t, err.IsOk())
	assert.Equal(t, cmdParams, []string{"c1", "c2"})
	assert.True(t, options.FooBar)
	assert.Equal(t, options.Baz, 12)
	assert.Equal(t, options.Qux, "ABC")
	assert.Equal(t, options.Quux, []string{"A", "B", "C"})
	assert.Equal(t, options.Corge, []int{20, 21})
}

func TestMakeOptCfgsFor_multipleOptsAndMultipleArgs(t *testing.T) {
	type MyOptions struct {
		FooBar bool     `opt:"foo-bar,f"`
		Baz    int      `opt:"baz,b=99"`
		Qux    string   `opt:"=XXX"`
		Quux   []string `opt:"quux=/[A/B/C]"`
		Corge  []int
	}
	options := MyOptions{}

	args := []string{
		"--foo-bar", "c1", "-b", "12", "--Qux", "ABC", "c2",
		"--Corge", "20", "--Corge=21",
	}

	optCfgs, err0 := libarg.MakeOptCfgsFor(&options)
	assert.True(t, err0.IsOk())
	assert.Equal(t, optCfgs[0].Name, "foo-bar")
	assert.Equal(t, optCfgs[0].Aliases, []string{"f"})
	assert.False(t, optCfgs[0].HasParam)
	assert.False(t, optCfgs[0].IsArray)
	assert.Nil(t, optCfgs[0].Default)
	assert.NotNil(t, optCfgs[0].OnParsed)
	assert.Equal(t, optCfgs[1].Name, "baz")
	assert.Equal(t, optCfgs[1].Aliases, []string{"b"})
	assert.True(t, optCfgs[1].HasParam)
	assert.False(t, optCfgs[1].IsArray)
	assert.Equal(t, optCfgs[1].Default, []string{"99"})
	assert.NotNil(t, optCfgs[1].OnParsed)
	assert.Equal(t, optCfgs[2].Name, "Qux")
	assert.Equal(t, optCfgs[2].Aliases, []string(nil))
	assert.True(t, optCfgs[2].HasParam)
	assert.False(t, optCfgs[2].IsArray)
	assert.Equal(t, optCfgs[2].Default, []string{"XXX"})
	assert.NotNil(t, optCfgs[2].OnParsed)
	assert.Equal(t, optCfgs[3].Name, "quux")
	assert.Equal(t, optCfgs[3].Aliases, []string{})
	assert.True(t, optCfgs[3].HasParam)
	assert.True(t, optCfgs[3].IsArray)
	assert.Equal(t, optCfgs[3].Default, []string{"A", "B", "C"})
	assert.NotNil(t, optCfgs[3].OnParsed)
	assert.Equal(t, optCfgs[4].Name, "Corge")
	assert.Equal(t, optCfgs[4].Aliases, []string(nil))
	assert.True(t, optCfgs[4].HasParam)
	assert.True(t, optCfgs[4].IsArray)
	assert.Equal(t, optCfgs[4].Default, []string(nil))
	assert.NotNil(t, optCfgs[4].OnParsed)

	a, err1 := libarg.ParseWith(args, optCfgs)
	assert.True(t, err1.IsOk())
	assert.Equal(t, a.CmdParams(), []string{"c1", "c2"})
	assert.True(t, a.HasOpt("foo-bar"))
	assert.True(t, a.HasOpt("baz"))
	assert.True(t, a.HasOpt("Qux"))
	assert.True(t, a.HasOpt("quux"))
	assert.True(t, a.HasOpt("Corge"))
	assert.Equal(t, a.OptParam("foo-bar"), "")
	assert.Equal(t, a.OptParam("baz"), "12")
	assert.Equal(t, a.OptParam("Qux"), "ABC")
	assert.Equal(t, a.OptParam("quux"), "A")
	assert.Equal(t, a.OptParam("Corge"), "20")
	assert.Equal(t, a.OptParams("foo-bar"), []string{})
	assert.Equal(t, a.OptParams("baz"), []string{"12"})
	assert.Equal(t, a.OptParams("Qux"), []string{"ABC"})
	assert.Equal(t, a.OptParams("quux"), []string{"A", "B", "C"})
	assert.Equal(t, a.OptParams("Corge"), []string{"20", "21"})
	assert.True(t, options.FooBar)
	assert.Equal(t, options.Baz, 12)
	assert.Equal(t, options.Qux, "ABC")
	assert.Equal(t, options.Quux, []string{"A", "B", "C"})
	assert.Equal(t, options.Corge, []int{20, 21})
}

package libarg

import (
	"github.com/sttk-go/sabi"
	"os"
	"strings"
	"unicode"
)

type /* error reason */ (
	// InvalidOption is an error reason which indicates that an invalid option is
	// found in command line arguments.
	InvalidOption struct{ Option string }
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
func (a Args) HasOpt(opt string) bool {
	_, exists := a.optParams[opt]
	return exists
}

// OptParam is a method to get a option parameter which is firstly specified
// with opt in command line arguments.
func (a Args) OptParam(opt string) string {
	arr := a.optParams[opt]
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
func (a Args) OptParams(opt string) []string {
	return a.optParams[opt]
}

// CmdParams is a method to get command parameters which are specified in
// command line parameters and are not associated with any options.
func (a Args) CmdParams() []string {
	return a.cmdParams
}

// Parse is a function to parse command line arguments without configurations.
// This function divides command line arguments to command parameters, which
// are not associated with any options, and options, of which each has a name
// and option parametes.
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
//	// os.Args[1:]  ==>  [--foo-bar=A -a --baz -bc=3 qux]
//	a, _ := Parse()
//	a.HasOpt("a")          // true
//	a.HasOpt("b")          // true
//	a.HasOpt("c")          // true
//	a.HasOpt("foo-bar")    // true
//	a.HasOpt("baz")        // true
//	a.OptParam("foo-bar")  // A
//	a.OptParams("foo-bar") // [A]
//	a.OptParam("c")        // 3
//	a.OptParams("c")       // [3]
//	a.CmdParams()          // [qux]
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
		return Args{}, err
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

	for _, arg := range args {
		if isNonOpt {
			err := collectCmdParams(arg)
			if !err.IsOk() {
				return err
			}
		} else if arg == "--" {
			isNonOpt = true

		} else if len(prevOptTakingParams) > 0 {
			err := collectOptParams(prevOptTakingParams, arg)
			if !err.IsOk() {
				return err
			}
			prevOptTakingParams = ""

		} else if strings.HasPrefix(arg, "--") {
			arg = arg[2:]
			i := 0
			for _, r := range arg {
				if i > 0 {
					if r == '=' {
						break
					}
					if !unicode.Is(rangeOfAlNumMarks, r) {
						return sabi.NewErr(InvalidOption{Option: arg})
					}
				} else {
					if !unicode.Is(rangeOfAlphabets, r) {
						return sabi.NewErr(InvalidOption{Option: arg})
					}
				}
				i++
			}
			if i == len(arg) {
				err := collectOptParams(arg)
				if !err.IsOk() {
					return err
				}
				if takeParams(arg) {
					prevOptTakingParams = arg
				}
			} else {
				err := collectOptParams(arg[0:i], arg[i+1:])
				if !err.IsOk() {
					return err
				}
			}
		} else if strings.HasPrefix(arg, "-") {
			arg := arg[1:]
			var opt string
			i := 0
			for _, r := range arg {
				if i > 0 && r == '=' {
					err := collectOptParams(opt, arg[i+1:])
					if !err.IsOk() {
						return err
					}
					break
				}
				opt = string(r)
				if !unicode.Is(rangeOfAlphabets, r) {
					return sabi.NewErr(InvalidOption{Option: opt})
				}
				err := collectOptParams(opt)
				if !err.IsOk() {
					return err
				}
				i++
			}
			if i == len(arg) && takeParams(opt) {
				prevOptTakingParams = opt
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

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
	return a.optParams[opt] != nil
}

// OptParam is a method to get a option parameter which is firstly specified
// with opt in command line arguments.
func (a Args) OptParam(opt string) string {
	arr := a.optParams[opt]
	if len(arr) > 0 {
		return arr[0]
	} else {
		return ""
	}
}

// OptParams is a method to get option parameters which are all specified with
// opt in command line arguments.
func (a Args) OptParams(opt string) []string {
	arr := a.optParams[opt]
	if len(arr) > 0 {
		return arr
	} else {
		return empty
	}
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
	return parseArgs(os.Args[1:], _false)
}

func _false(_ string) bool {
	return false
}

func parseArgs(args []string, takeParams func(string) bool) (Args, sabi.Err) {
	var a Args

	var cmdParams = make([]string, 0)
	var optParams = make(map[string][]string)

	isNonOpt := false
	prevOptTakingParams := ""

	for _, arg := range args {
		if isNonOpt {
			cmdParams = append(cmdParams, arg)

		} else if arg == "--" {
			isNonOpt = true

		} else if len(prevOptTakingParams) > 0 {
			addKeyValueToMap(optParams, prevOptTakingParams, arg)
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
						return a, sabi.NewErr(InvalidOption{Option: arg})
					}
				} else {
					if !unicode.Is(rangeOfAlphabets, r) {
						return a, sabi.NewErr(InvalidOption{Option: arg})
					}
				}
				i++
			}
			if i == len(arg) {
				addKeyToMap(optParams, arg)
				if takeParams(arg) {
					prevOptTakingParams = arg
				}
			} else {
				addKeyValueToMap(optParams, arg[0:i], arg[i+1:])
			}

		} else if strings.HasPrefix(arg, "-") {
			arg := arg[1:]
			var opt string
			i := 0
			for _, r := range arg {
				if i > 0 && r == '=' {
					addKeyValueToMap(optParams, opt, arg[i+1:])
					break
				}
				opt = string(r)
				if !unicode.Is(rangeOfAlphabets, r) {
					return a, sabi.NewErr(InvalidOption{Option: opt})
				}
				addKeyToMap(optParams, opt)
				i++
			}
			if i == len(arg) && takeParams(opt) {
				prevOptTakingParams = opt
			}

		} else {
			cmdParams = append(cmdParams, arg)
		}
	}

	a.cmdParams = cmdParams
	a.optParams = optParams

	return a, sabi.Ok()
}

func addKeyToMap(m map[string][]string, key string) {
	arr := m[key]
	if arr == nil {
		m[key] = empty
	}
}

func addKeyValueToMap(m map[string][]string, key, val string) {
	arr := m[key]
	if arr == nil {
		m[key] = []string{val}
	} else {
		m[key] = append(arr, val)
	}
}

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
func (args Args) HasOpt(opt string) bool {
	return args.optParams[opt] != nil
}

// OptParam is a method to get a option parameter which is firstly specified
// with opt in command line arguments.
func (args Args) OptParam(opt string) string {
	arr := args.optParams[opt]
	if len(arr) > 0 {
		return arr[0]
	} else {
		return ""
	}
}

// OptParams is a method to get option parameters which are all specified with
// opt in command line arguments.
func (args Args) OptParams(opt string) []string {
	arr := args.optParams[opt]
	if len(arr) > 0 {
		return arr
	} else {
		return empty
	}
}

// CmdParams is a method to get command parameters which are specified in
// command line parameters and are not associated with any options.
func (args Args) CmdParams() []string {
	return args.cmdParams
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
//	os.Args[1:]               // [--foo-bar=A -a --baz -bc=3 qux]
//	args, _ := Parse()
//	args.HasOpt("a")          // true
//	args.HasOpt("b")          // true
//	args.HasOpt("c")          // true
//	args.HasOpt("foo-bar")    // true
//	args.HasOpt("baz")        // true
//	args.OptParam("foo-bar")  // A
//	args.OptParams("foo-bar") // [A]
//	args.OptParam("c")        // 3
//	args.OptParams("c")       // [3]
//	args.CmdParams()          // [qux]
func Parse() (Args, sabi.Err) {
	var args Args

	var cmdParams = make([]string, 0)
	var optParams = make(map[string][]string)

	isNonOpt := false

	for _, arg := range os.Args[1:] {
		if isNonOpt {
			cmdParams = append(cmdParams, arg)

		} else if arg == "--" {
			isNonOpt = true

		} else if strings.HasPrefix(arg, "--") {
			arg = arg[2:]
			i := 0
			for _, r := range arg {
				if i > 0 {
					if r == '=' {
						break
					}
					if !unicode.Is(rangeOfAlNumMarks, r) {
						return args, sabi.NewErr(InvalidOption{Option: arg})
					}
				} else {
					if !unicode.Is(rangeOfAlphabets, r) {
						return args, sabi.NewErr(InvalidOption{Option: arg})
					}
				}
				i++
			}
			if i == len(arg) {
				addKeyToMap(optParams, arg)
			} else {
				addKeyValueToMap(optParams, arg[0:i], arg[i+1:])
			}

		} else if strings.HasPrefix(arg, "-") {
			arg := arg[1:]
			var last string
			for i, r := range arg {
				if i > 0 && r == '=' {
					addKeyValueToMap(optParams, last, arg[i+1:])
					break
				}
				last = string(r)
				if !unicode.Is(rangeOfAlphabets, r) {
					return args, sabi.NewErr(InvalidOption{Option: last})
				}
				addKeyToMap(optParams, last)
			}

		} else {
			cmdParams = append(cmdParams, arg)
		}
	}

	args.cmdParams = cmdParams
	args.optParams = optParams

	return args, sabi.Ok()
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

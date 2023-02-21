// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

package libarg

import (
	"github.com/sttk-go/sabi"
)

type /* error reason */ (
	// ConfigIsArrayButHasNoParam is an error reason which indicates that
	// an option configuration contradicts that the option must be an array
	// (.IsArray = true) but must have no option parameter (.HasParam = false).
	ConfigIsArrayButHasNoParam struct{ Opt string }

	// ConfigHasDefaultButHasNoParam is an error reason which indicates that
	// an option configuration contradicts that the option has default value
	// (.Default != nil) but must have no option parameter (.HasParam = false).
	ConfigHasDefaultButHasNoParam struct{ Opt string }

	// UnconfiguredOption is an error reason which indicates that there is no
	// configuration about the input option.
	UnconfiguredOption struct{ Opt string }

	// OptionNeedsParam is an error reason which indicates that an option is
	// input with no option parameter though its option configuration requires
	// option parameters (.HasParam = true).
	OptionNeedsParam struct{ Opt string }

	// OptionTakesNoParam is an error reason which indicates that an option is
	// input with an option parameter though its option configuration does not
	// accept option parameters (.HasParam = false).
	OptionTakesNoParam struct{ Opt string }

	// OptionIsNotArray is an error reason which indicates that an option is
	// input with an option parameter multiple times though its option
	// configuration specifies the option is not an array (.IsArray = false).
	OptionIsNotArray struct{ Opt string }
)

const anyOpt = "*"

// OptCfg is a structure that represents an option configuration.
// An option configuration consists of fields: Name, Aliases, HasParam,
// IsArray.

// Name is the option name and Aliases are the another names.
// Options given by those names in command line arguments are all registered to
// Args with the Name.
//
// HasParam and IsArray are flags which allows the option to take option
// parameters.
// If both HasParam and IsArray are true, the option can take one or multiple
// option parameters.
// If HasParam is true and IsArray is false, the option can take only one
// option parameter.
// If both HasParam and IsArray are false, the option can take no option
// parameter.
//
// Default is the field to specify the default value for when the option is not
// given in command line arguments.
type OptCfg struct {
	Name     string
	Aliases  []string
	HasParam bool
	IsArray  bool
	Default  []string
}

// ParseWith is a function which parses command line arguments with option
// configurations.
// This function divides command line arguments to command parameters and
// options, and an option consists of a name and option parameters.
// Options are divided to long format options and short format options.
// About long/short format options, since they are same with Parse function,
// see the comment of the function.
//
// This function allows only options declared in option configurations.
// A option configuration has fields: Name, Aliases, HasParam, IsArray, and
// Default.
// When an option matches Name or includes in Aliases in an option
// configuration, the option is registered in Args with the Name.
// If both HasParam and IsArray are true, the option can has one or multiple
// option parameters, and if HasParam is true and IsArray is false, the option
// can has only one option parameter, otherwise the option cannot have option
// parameter.
// If Default is specified and the option is not given in command line
// arguments, the value of Default is set to the option parameter.
//
// If options not declared in option configurations are given in command line
// arguments, this function basically returns UnconfiguredOption error.
// If you want to allow other options, add an option configuration of which
// Name is "*" (but HasParam and IsArray of this configuration is ignored).
//
// Usage example:
//
//	osArgs := []string{"--foo-bar", "quz", "--baz", "1", "-z=2", "-X", "quux"}
//	optCfgs := []OptCfg{
//		OptCfg{Name:"foo-bar"},
//		OptCfg{Name:"baz", Aliases:[]string{"z"}, HasParam:true, IsArray:true},
//		OptCfg{Name:"corge", Default:[]string{"99"}},
//		OptCfg{Name:"*"},
//	}
//
//	args, err := ParseWith(osArgs, optCfgs)
//	args.HasOpt("foo-bar")  // true
//	args.HasOpt("baz")      // true
//	args.HasOpt("X")        // true, due to "*" config
//	args.HasOpt("corge")    // true, due to .Default.
//	args.OptParam("baz")    // 1
//	args.OptParams("baz")   // [1 2]
//	args.OptParam("corge")  // 99
//	args.OptParams("corge") // [99]
//	args.CmdParams()        // [qux quux]
func ParseWith(args []string, optCfgs []OptCfg) (Args, sabi.Err) {
	hasAnyOpt := false
	cfgMap := make(map[string]int)
	defMap := make(map[string]int)
	for i, cfg := range optCfgs {
		if cfg.IsArray && !cfg.HasParam {
			err := sabi.NewErr(ConfigIsArrayButHasNoParam{Opt: cfg.Name})
			return Args{cmdParams: empty}, err
		}
		if cfg.Name == anyOpt {
			hasAnyOpt = true
			continue
		}
		if cfg.Default != nil {
			if !cfg.HasParam {
				err := sabi.NewErr(ConfigHasDefaultButHasNoParam{Opt: cfg.Name})
				return Args{cmdParams: empty}, err
			}
			defMap[cfg.Name] = i
		}
		cfgMap[cfg.Name] = i
		for _, a := range cfg.Aliases {
			cfgMap[a] = i
		}
	}

	var takeParam = func(opt string) bool {
		i, exists := cfgMap[opt]
		if exists {
			return optCfgs[i].HasParam
		}
		return false
	}

	var cmdParams = make([]string, 0)
	var optParams = make(map[string][]string)

	var collCmdParams = func(params ...string) sabi.Err {
		cmdParams = append(cmdParams, params...)
		return sabi.Ok()
	}
	var collOptParams = func(opt string, params ...string) sabi.Err {
		i, exists := cfgMap[opt]
		if !exists {
			if !hasAnyOpt {
				return sabi.NewErr(UnconfiguredOption{Opt: opt})
			}

			arr := optParams[opt]
			if arr == nil {
				arr = empty
			}
			optParams[opt] = append(arr, params...)
			return sabi.Ok()
		}

		cfg := optCfgs[i]
		if !cfg.HasParam {
			if len(params) > 0 {
				return sabi.NewErr(OptionTakesNoParam{Opt: cfg.Name})
			}
		} else {
			if len(params) == 0 {
				return sabi.NewErr(OptionNeedsParam{Opt: cfg.Name})
			}
		}

		arr := optParams[cfg.Name]
		if arr == nil {
			arr = empty
		}
		arr = append(arr, params...)

		if !cfg.IsArray {
			if len(arr) > 1 {
				return sabi.NewErr(OptionIsNotArray{Opt: cfg.Name})
			}
		}

		optParams[cfg.Name] = arr
		return sabi.Ok()
	}

	err := parseArgs(args, collCmdParams, collOptParams, takeParam)
	if !err.IsOk() {
		return Args{cmdParams: empty}, err
	}

	for name, i := range defMap {
		_, exists := optParams[name]
		if !exists {
			optParams[name] = optCfgs[i].Default
		}
	}

	return Args{cmdParams: cmdParams, optParams: optParams}, err
}

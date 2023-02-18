package libarg_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sttk-go/clidax/libarg"
	"testing"
)

func TestParseWith_zeroCfgAndZeroArg(t *testing.T) {
	optCfgs := []libarg.OptCfg{}

	osArgs := []string{}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_zeroCfgAndOneCommandParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{}

	osArgs := []string{"foo-bar"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{"foo-bar"})
}

func testParseWith_zeroCfgAndOneLongOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{}

	osArgs := []string{"--foo-bar"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.UnconfiguredOption:
		assert.Equal(t, err.Get("Opt"), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_zeroCfgAndOneShortOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{}

	osArgs := []string{"-f"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.UnconfiguredOption:
		assert.Equal(t, err.Get("Opt"), "f")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgAndZeroOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgAndOneCmdParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"foo-bar"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{"foo-bar"})
}

func TestParseWith_oneCfgAndOneLongOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"--foo-bar"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string{})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgAndOneShortOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "f"},
	}

	osArgs := []string{"-f"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string{})
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgAndOneDifferentLongOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"--boo-far"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.UnconfiguredOption:
		assert.Equal(t, err.Get("Opt"), "boo-far")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgAndOneDifferentShortOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "f"},
	}

	osArgs := []string{"-b"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.UnconfiguredOption:
		assert.Equal(t, err.Get("Opt"), "b")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_anyOptCfgAndOneDifferentLongOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar"},
		libarg.OptCfg{Name: "*"},
	}

	osArgs := []string{"--boo-far"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	assert.True(t, args.HasOpt("boo-far"))
	assert.Equal(t, args.OptParam("boo-far"), "")
	assert.Equal(t, args.OptParams("boo-far"), []string{})
}

func TestParseWith_anyOptCfgAndOneDifferentShortOpt(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "f"},
		libarg.OptCfg{Name: "*"},
	}

	osArgs := []string{"-b"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	assert.True(t, args.HasOpt("b"))
	assert.Equal(t, args.OptParam("b"), "")
	assert.Equal(t, args.OptParams("b"), []string{})
}

func TestParseWith_oneCfgHasParamAndOneLongOptHasParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar", HasParam: true},
	}

	osArgs := []string{"--foo-bar", "ABC"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"--foo-bar=ABC"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"--foo-bar", ""}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string{""})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"--foo-bar="}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string{""})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgHasParamAndOneShortOptHasParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "f", HasParam: true},
	}

	osArgs := []string{"-f", "ABC"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "ABC")
	assert.Equal(t, args.OptParams("f"), []string{"ABC"})
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"-f=ABC"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "ABC")
	assert.Equal(t, args.OptParams("f"), []string{"ABC"})
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"-f", ""}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string{""})
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"-f="}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string{""})
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgHasParamButOneLongOptHasNoParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar", HasParam: true},
	}

	osArgs := []string{"--foo-bar"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.OptionNeedsParam:
		assert.Equal(t, err.Get("Opt"), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgHasParamAndOneShortOptHasNoParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "f", HasParam: true},
	}

	osArgs := []string{"-f"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.OptionNeedsParam:
		assert.Equal(t, err.Get("Opt"), "f")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgHasNoParamAndOneLongOptHasParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar"},
	}

	osArgs := []string{"--foo-bar", "ABC"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string{})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{"ABC"})

	osArgs = []string{"--foo-bar=ABC"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.OptionTakesNoParam:
		assert.Equal(t, err.Get("Opt"), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"--foo-bar", ""}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string{})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{""})

	osArgs = []string{"--foo-bar="}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.OptionTakesNoParam:
		assert.Equal(t, err.Get("Opt"), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgHasNoParamAndOneShortOptHasParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "f"},
	}

	osArgs := []string{"-f", "ABC"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string{})
	assert.Equal(t, args.CmdParams(), []string{"ABC"})

	osArgs = []string{"-f=ABC"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.OptionTakesNoParam:
		assert.Equal(t, err.Get("Opt"), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"-f", ""}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string{})
	assert.Equal(t, args.CmdParams(), []string{""})

	osArgs = []string{"-f="}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.OptionTakesNoParam:
		assert.Equal(t, err.Get("Opt"), "f")
	default:
		assert.Fail(t, err.Error())
	}

	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgHasNoParamButIsArray(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar", HasParam: false, IsArray: true},
	}

	osArgs := []string{"--foo-bar", "ABC"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.ConfigIsArrayButHasNoParam:
		assert.Equal(t, err.Get("Opt"), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgIsArrayAndOptHasOneParam(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar", HasParam: true, IsArray: true},
		libarg.OptCfg{Name: "f", HasParam: true, IsArray: true},
	}

	osArgs := []string{"--foo-bar", "ABC"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"--foo-bar", "ABC", "--foo-bar=DEF"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"-f", "ABC"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "ABC")
	assert.Equal(t, args.OptParams("f"), []string{"ABC"})
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"-f", "ABC", "-f=DEF"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.True(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "ABC")
	assert.Equal(t, args.OptParams("f"), []string{"ABC", "DEF"})
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgHasAliasesAndArgMatchesName(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{
			Name:     "foo-bar",
			Aliases:  []string{"f", "b"},
			HasParam: true,
		},
	}

	osArgs := []string{"--foo-bar", "ABC"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"--foo-bar=ABC"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgHasAliasesAndArgMatchesAliases(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{
			Name:     "foo-bar",
			Aliases:  []string{"f"},
			HasParam: true,
		},
	}

	osArgs := []string{"-f", "ABC"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"-f=ABC"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_combineOptsByNameAndAliases(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{
			Name:     "foo-bar",
			Aliases:  []string{"f"},
			HasParam: true,
			IsArray:  true,
		},
	}

	osArgs := []string{"-f", "ABC", "--foo-bar=DEF"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})

	osArgs = []string{"-f=ABC", "--foo-bar", "DEF"}

	args, err = libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "ABC")
	assert.Equal(t, args.OptParams("foo-bar"), []string{"ABC", "DEF"})
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_oneCfgIsNotArrayButOptsAreMultiple(t *testing.T) {
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{
			Name:     "foo-bar",
			Aliases:  []string{"f"},
			HasParam: true,
			IsArray:  false,
		},
	}

	osArgs := []string{"-f", "ABC", "--foo-bar=DEF"}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.False(t, err.IsOk())
	switch err.Reason().(type) {
	case libarg.OptionIsNotArray:
		assert.Equal(t, err.Get("Opt"), "foo-bar")
	default:
		assert.Fail(t, err.Error())
	}
	assert.False(t, args.HasOpt("foo-bar"))
	assert.Equal(t, args.OptParam("foo-bar"), "")
	assert.Equal(t, args.OptParams("foo-bar"), []string(nil))
	assert.False(t, args.HasOpt("f"))
	assert.Equal(t, args.OptParam("f"), "")
	assert.Equal(t, args.OptParams("f"), []string(nil))
	assert.Equal(t, args.CmdParams(), []string{})
}

func TestParseWith_multipleArgs(t *testing.T) {
	osArgs := []string{"--foo-bar", "qux", "--baz", "1", "-z=2", "-X", "quux"}
	optCfgs := []libarg.OptCfg{
		libarg.OptCfg{Name: "foo-bar"},
		libarg.OptCfg{
			Name:     "baz",
			Aliases:  []string{"z"},
			HasParam: true,
			IsArray:  true,
		},
		libarg.OptCfg{Name: "*"},
	}

	args, err := libarg.ParseWith(osArgs, optCfgs)
	assert.True(t, err.IsOk())
	assert.True(t, args.HasOpt("foo-bar"))
	assert.True(t, args.HasOpt("baz"))
	assert.True(t, args.HasOpt("X"))
	assert.Equal(t, args.OptParam("baz"), "1")
	assert.Equal(t, args.OptParams("baz"), []string{"1", "2"})
	assert.Equal(t, args.CmdParams(), []string{"qux", "quux"})
}

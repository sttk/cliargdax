package cliargdax_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sttk/cliargdax"
	"github.com/sttk/cliargs"
	"github.com/sttk/sabi"
	"github.com/sttk/sabi/errs"
)

var origOsArgs = os.Args

func resetOsArgs() {
	os.Args = origOsArgs
}

type noopAsyncGroup struct{}

func (ag *noopAsyncGroup) Add(fn func() errs.Err) {}

func TestCliArgDax_NewDaxSrc_ok(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{"/path/to/app", "--foo", "bar", "--baz=123"}

	ds := cliargdax.NewDaxSrc()

	ag := &noopAsyncGroup{}
	err := ds.Setup(ag)
	defer ds.Close()

	dc, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())

	conn, ok := dc.(cliargdax.DaxConn)
	assert.True(t, ok)

	cmd := conn.Cmd()
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{"bar"})
	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.True(t, cmd.HasOpt("baz"))
	assert.Equal(t, cmd.OptArg("baz"), "123")
	assert.Equal(t, cmd.OptArgs("baz"), []string{"123"})

	optCfgs := conn.OptCfgs()
	assert.Equal(t, len(optCfgs), 0)

	options := conn.Options()
	assert.Nil(t, options)
}

func TestCliArgDax_NewDaxSrc_error(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{"/path/to/app", "--foo", "bar", "--123"}

	ds := cliargdax.NewDaxSrc()

	ag := &noopAsyncGroup{}
	err := ds.Setup(ag)
	defer ds.Close()

	switch r := err.Reason().(type) {
	case cliargs.OptionHasInvalidChar:
		assert.Equal(t, r.Option, "123")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestCliArgDax_NewDaxSrcWithOptCfgs_ok(t *testing.T) {
	defer resetOsArgs()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "foo",
			Desc: "foo description",
		},
		cliargs.OptCfg{
			Name:   "baz",
			Desc:   "baz description",
			HasArg: true,
		},
	}

	os.Args = []string{"/path/to/app", "--foo", "bar", "--baz=123"}

	ds := cliargdax.NewDaxSrcWithOptCfgs(optCfgs)

	ag := &noopAsyncGroup{}
	err := ds.Setup(ag)
	defer ds.Close()

	dc, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())

	conn, ok := dc.(cliargdax.DaxConn)
	assert.True(t, ok)

	cmd := conn.Cmd()
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{"bar"})
	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.True(t, cmd.HasOpt("baz"))
	assert.Equal(t, cmd.OptArg("baz"), "123")
	assert.Equal(t, cmd.OptArgs("baz"), []string{"123"})

	cfgs := conn.OptCfgs()
	assert.Equal(t, len(optCfgs), 2)
	assert.Equal(t, cfgs[0].Name, "foo")
	assert.Equal(t, cfgs[0].Desc, "foo description")
	assert.False(t, cfgs[0].HasArg)
	assert.Equal(t, cfgs[1].Name, "baz")
	assert.Equal(t, cfgs[1].Desc, "baz description")
	assert.True(t, cfgs[1].HasArg)

	options := conn.Options()
	assert.Nil(t, options)
}

func TestCliArgDax_NewDaxSrcWithOptCfgs_error(t *testing.T) {
	defer resetOsArgs()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "foo",
			Desc: "foo description",
		},
		cliargs.OptCfg{
			Name:   "baz",
			Desc:   "baz description",
			HasArg: true,
		},
	}

	os.Args = []string{"/path/to/app", "--foo", "bar", "--qux", "--baz=123"}

	ds := cliargdax.NewDaxSrcWithOptCfgs(optCfgs)

	ag := &noopAsyncGroup{}
	err := ds.Setup(ag)
	defer ds.Close()

	switch r := err.Reason().(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, r.Option, "qux")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestCliArgDax_NewDaxSrcForOptions_ok(t *testing.T) {
	defer resetOsArgs()

	type Options struct {
		Foo bool `optcfg:"foo" optdesc:"foo description"`
		Baz int  `optcfg:"baz" optdesc:"baz description"`
	}

	options := Options{}

	os.Args = []string{"/path/to/app", "--foo", "bar", "--baz=123"}

	ds := cliargdax.NewDaxSrcForOptions(&options)

	ag := &noopAsyncGroup{}
	err := ds.Setup(ag)
	defer ds.Close()

	dc, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())

	conn, ok := dc.(cliargdax.DaxConn)
	assert.True(t, ok)

	cmd := conn.Cmd()
	assert.Equal(t, cmd.Name, "app")
	assert.Equal(t, cmd.Args(), []string{"bar"})
	assert.True(t, cmd.HasOpt("foo"))
	assert.Equal(t, cmd.OptArg("foo"), "")
	assert.Equal(t, cmd.OptArgs("foo"), []string{})
	assert.True(t, cmd.HasOpt("baz"))
	assert.Equal(t, cmd.OptArg("baz"), "123")
	assert.Equal(t, cmd.OptArgs("baz"), []string{"123"})

	cfgs := conn.OptCfgs()
	assert.Equal(t, len(cfgs), 2)
	assert.Equal(t, cfgs[0].Name, "foo")
	assert.Equal(t, cfgs[0].Desc, "foo description")
	assert.False(t, cfgs[0].HasArg)
	assert.Equal(t, cfgs[1].Name, "baz")
	assert.Equal(t, cfgs[1].Desc, "baz description")
	assert.True(t, cfgs[1].HasArg)

	opts, ok := conn.Options().(*Options)
	assert.True(t, ok)
	assert.True(t, opts.Foo)
	assert.Equal(t, opts.Baz, 123)
}

func TestCliArgDax_NewDaxSrcForOptions_error(t *testing.T) {
	defer resetOsArgs()

	type Options struct {
		Foo bool `optcfg:"foo" optdesc:"foo description"`
		Baz int  `optcfg:"baz" optdesc:"baz description"`
	}

	options := Options{}

	os.Args = []string{"/path/to/app", "--foo", "bar", "--qux", "--baz=123"}

	ds := cliargdax.NewDaxSrcForOptions(&options)

	ag := &noopAsyncGroup{}
	err := ds.Setup(ag)
	defer ds.Close()

	switch r := err.Reason().(type) {
	case cliargs.UnconfiguredOption:
		assert.Equal(t, r.Option, "qux")
	default:
		assert.Fail(t, err.Error())
	}
}

func TestCliArgDax_txn_commit(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{"/path/to/app", "--foo", "bar", "--baz=123"}

	base := sabi.NewDaxBase()
	defer base.Close()

	base.Uses("cliarg", cliargdax.NewDaxSrc())

	err := sabi.Txn(base, func(dax sabi.Dax) errs.Err {
		conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
		assert.True(t, err.IsOk())

		cmd := conn.Cmd()
		assert.Equal(t, cmd.Name, "app")
		assert.Equal(t, cmd.Args(), []string{"bar"})
		assert.True(t, cmd.HasOpt("foo"))
		assert.Equal(t, cmd.OptArg("foo"), "")
		assert.Equal(t, cmd.OptArgs("foo"), []string{})
		assert.True(t, cmd.HasOpt("baz"))
		assert.Equal(t, cmd.OptArg("baz"), "123")
		assert.Equal(t, cmd.OptArgs("baz"), []string{"123"})

		return errs.Ok()
	})
	assert.True(t, err.IsOk())
}

func TestCliArgDax_txn_rollback(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{"/path/to/app", "--foo", "bar", "--baz=123"}

	base := sabi.NewDaxBase()
	defer base.Close()

	base.Uses("cliarg", cliargdax.NewDaxSrc())

	type FailToDoSomething struct{}

	err := sabi.Txn(base, func(dax sabi.Dax) errs.Err {
		_, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
		assert.True(t, err.IsOk())
		return errs.New(FailToDoSomething{})
	})
	switch err.Reason().(type) {
	case FailToDoSomething:
	default:
		assert.Fail(t, err.Error())
	}
}

func TestCliArgDax_DaxConn_SetOption(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{"/path/to/app"}

	base := sabi.NewDaxBase()
	defer base.Close()

	base.Uses("cliarg", cliargdax.NewDaxSrc())

	type MyOption struct {
		Flag int
	}

	err := sabi.Txn(base, func(dax sabi.Dax) errs.Err {
		conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
		assert.True(t, err.IsOk())
		conn.SetOptions(MyOption{Flag: 111})
		return errs.Ok()
	})
	assert.True(t, err.IsOk())

	err = sabi.Txn(base, func(dax sabi.Dax) errs.Err {
		conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
		assert.True(t, err.IsOk())
		assert.Equal(t, conn.Options().(MyOption).Flag, 111)
		return errs.Ok()
	})
	assert.True(t, err.IsOk())
}

func TestCliArgDax_forCoverage(t *testing.T) {
	defer resetOsArgs()

	os.Args = []string{"/path/to/app"}

	ds := cliargdax.NewDaxSrc()

	ag := &noopAsyncGroup{}
	err := ds.Setup(ag)
	defer ds.Close()

	dc, err := ds.CreateDaxConn()
	assert.True(t, err.IsOk())

	conn, ok := dc.(cliargdax.DaxConn)
	assert.True(t, ok)

	conn.Rollback(ag)
}

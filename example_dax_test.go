package cliargdax_test

import (
	"fmt"
	"os"

	"github.com/sttk/cliargdax"
	"github.com/sttk/cliargs"
	"github.com/sttk/sabi"
)

func ExampleDaxConn_Cmd() {
	os.Args = []string{"path/to/app", "--foo", "bar"}

	base := sabi.NewDaxBase()
	defer base.Close()

	base.Uses("cliarg", cliargdax.NewDaxSrc())

	conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
	fmt.Printf("err.IsOk = %t\n", err.IsOk())

	cmd := conn.Cmd()
	fmt.Printf("cmd.Name = %s\n", cmd.Name)
	fmt.Printf("cmd.Args = %v\n", cmd.Args())
	fmt.Printf("cmd.HasOpts: foo = %t\n", cmd.HasOpt("foo"))

	// Output:
	// err.IsOk = true
	// cmd.Name = app
	// cmd.Args = [bar]
	// cmd.HasOpts: foo = true

	resetOsArgs()
}

func ExampleDaxConn_OptCfgs() {
	os.Args = []string{"path/to/app", "--foo", "bar"}

	base := sabi.NewDaxBase()
	defer base.Close()

	opts := struct {
		Foo bool `optcfg:"foo"`
	}{}
	base.Uses("cliarg", cliargdax.NewDaxSrcForOptions(&opts))

	conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
	fmt.Printf("err.IsOk = %t\n", err.IsOk())

	optCfgs := conn.OptCfgs()
	fmt.Printf("len(optCfgs) = %d\n", len(optCfgs))
	fmt.Printf("optCfgs[0].Name = %v\n", optCfgs[0].Name)

	// Output:
	// err.IsOk = true
	// len(optCfgs) = 1
	// optCfgs[0].Name = foo

	resetOsArgs()
}

func ExampleDaxConn_Options() {
	os.Args = []string{"path/to/app", "--foo", "bar"}

	base := sabi.NewDaxBase()
	defer base.Close()

	type MyOptions struct {
		Foo bool `optcfg:"foo"`
	}
	base.Uses("cliarg", cliargdax.NewDaxSrcForOptions(&MyOptions{}))

	conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
	fmt.Printf("err.IsOk = %t\n", err.IsOk())

	options := conn.Options().(*MyOptions)
	fmt.Printf("options.Foo = %t\n", options.Foo)

	// Output:
	// err.IsOk = true
	// options.Foo = true

	resetOsArgs()
}

func ExampleDaxConn_SetOptions() {
	os.Args = []string{"path/to/app", "--foo", "bar"}

	base := sabi.NewDaxBase()
	defer base.Close()

	base.Uses("cliarg", cliargdax.NewDaxSrc())

	conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
	fmt.Printf("err.IsOk = %t\n", err.IsOk())

	fmt.Printf("options = %v\n", conn.Options())

	type MyOptions struct {
		Foo bool
	}
	conn.SetOptions(&MyOptions{Foo: true})

	options := conn.Options().(*MyOptions)
	fmt.Printf("options.Foo = %t\n", options.Foo)

	// Output:
	// err.IsOk = true
	// options = <nil>
	// options.Foo = true
}

func ExampleNewDaxSrc() {
	os.Args = []string{"path/to/app", "--foo", "bar"}

	base := sabi.NewDaxBase()
	defer base.Close()

	base.Uses("cliarg", cliargdax.NewDaxSrc())

	conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
	fmt.Printf("err.IsOk = %t\n", err.IsOk())

	cmd := conn.Cmd()
	fmt.Printf("cmd.Name = %s\n", cmd.Name)
	fmt.Printf("cmd.Args = %v\n", cmd.Args())
	fmt.Printf("cmd.HasOpts: foo = %t\n", cmd.HasOpt("foo"))
	fmt.Printf("optCfgs = %v\n", conn.OptCfgs())
	fmt.Printf("options = %v\n", conn.Options())

	// Output:
	// err.IsOk = true
	// cmd.Name = app
	// cmd.Args = [bar]
	// cmd.HasOpts: foo = true
	// optCfgs = []
	// options = <nil>

	resetOsArgs()
}

func ExampleNewDaxSrcForOptions() {
	os.Args = []string{"path/to/app", "--foo", "bar"}

	base := sabi.NewDaxBase()
	defer base.Close()

	optCfgs := []cliargs.OptCfg{
		cliargs.OptCfg{
			Name: "foo",
		},
	}
	base.Uses("cliarg", cliargdax.NewDaxSrcWithOptCfgs(optCfgs))

	conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
	fmt.Printf("err.IsOk = %t\n", err.IsOk())

	cmd := conn.Cmd()
	fmt.Printf("cmd.Name = %s\n", cmd.Name)
	fmt.Printf("cmd.Args = %v\n", cmd.Args())
	fmt.Printf("cmd.HasOpts: foo = %t\n", cmd.HasOpt("foo"))
	fmt.Printf("optCfgs[0].Name = %s\n", conn.OptCfgs()[0].Name)
	fmt.Printf("options = %v\n", conn.Options())

	// Output:
	// err.IsOk = true
	// cmd.Name = app
	// cmd.Args = [bar]
	// cmd.HasOpts: foo = true
	// optCfgs[0].Name = foo
	// options = <nil>

	resetOsArgs()
}

func ExampleNewDaxSrcWithOptCfgs() {
	os.Args = []string{"path/to/app", "--foo", "bar"}

	base := sabi.NewDaxBase()
	defer base.Close()

	type MyOptions struct {
		Foo bool `optcfg:"foo"`
	}
	base.Uses("cliarg", cliargdax.NewDaxSrcForOptions(&MyOptions{}))

	conn, err := sabi.GetDaxConn[cliargdax.DaxConn](base, "cliarg")
	fmt.Printf("err.IsOk = %t\n", err.IsOk())

	cmd := conn.Cmd()
	fmt.Printf("cmd.Name = %s\n", cmd.Name)
	fmt.Printf("cmd.Args = %v\n", cmd.Args())
	fmt.Printf("cmd.HasOpts: foo = %t\n", cmd.HasOpt("foo"))

	fmt.Printf("optCfgs[0].Name = %s\n", conn.OptCfgs()[0].Name)

	options := conn.Options().(*MyOptions)
	fmt.Printf("options.Foo = %v\n", options.Foo)

	// Output:
	// err.IsOk = true
	// cmd.Name = app
	// cmd.Args = [bar]
	// cmd.HasOpts: foo = true
	// optCfgs[0].Name = foo
	// options.Foo = true

	resetOsArgs()
}

// Copyright (C) 2023 Takayuki Sato. All Rights Reserved.
// This program is free software under MIT License.
// See the file LICENSE in this distribution for more details.

/*
Package github.com/sttk/cliargdax is a dax in sabi framework for data access.

This package uses github.com/sttk/cliargs package to parse command-line arguments and store the results of parsing.

# Usage of dax source

This package provides a dax source named DaxSrc.

This dax source can be used both as a global and local dax source.
However, since command-line arguments can be obtained globally, it is typically
used as a global dax source.
And it would be appropriate to declare its usage within init function.

	import "github.com/sttk/cliargdax"

	func init() {
	    sabi.Uses("cliopts", cliargdax.NewDaxSrc())
	}

A DaxSrc instance can be instantiated by the functions: NewDaxSrc,
NewDaxSrcWithOptCfgs, NewDaxSrcForOptions.

NewDaxSrc function creates a DaxSrc instance without configuration (see
above).
And it's Setup method separates command-line arguments to normal arguments and options only by their formats.

NewDaxSrcWithOptCfgs function creates a DaxSrc instance with an array of
cliargs.OptCfg.

	optCfgs := []cliargs.OptCfg{
	  cliargs.OptCfg{ ... },
	  ... ,
	}
	sabi.Uses("cliopts", cliargdax.NewDaxSrcWithOptCfgs(optCfgs))

And it's Setup method parses command-line arguments with the array.
This configuration array can be retrieve by using DaxConn#OptCfgs method.

NewDaxSrcForOptions function creates a DaxSrc instance with a pointer of any
type struct instance that stores results of command-line argument parsing.

	type MyOptions struct { ... }
	opts := MyOptions{ ... }
	sabi.Uses("cliopts", cliargdax.NewDaxSrcForOptions(&opts)

And it's Setup method parses command-line arguments with an array of
cliargs.OptCfg that is created from the option store instance.
These configuration array and store instance can be retrieve by using
DaxConn#OptCfgs and DaxConn#Options methods.

# Usage of dax connection

This package provides a dax connection named DaxConn.
This dax connection is created by DaxSrc#CreateDaxConn method.

Within dax implementations using command-line arguments, this dax connection is
obtained by sabi.GetDaxConn method.
And by this dac connection, the results of command-line argument parsing can be
obtained.

	func (dax MyDax) doSomeDataAccess() errs.Err {
	    conn, err := sabi.GetDaxConn[cliargdax.DaxConn](dax, "cliopts")
	    if err.IsNotOk() {
	        return err
	    }

	    var cmd cliargs.Cmd = conn.Cmd()
	    var optCfgs []cliargs.OptCfg = conn.OptCfgs()
	    var options *MyOptions = conn.Options().(*MyOptions)

	    return errs.Ok()
	}
*/
package cliargdax

import (
	"os"

	"github.com/sttk/cliargs"
	"github.com/sttk/sabi"
	"github.com/sttk/sabi/errs"
)

// DaxConn is the dax connection struct for command line argument operations.
// In addition to methods for transactions: Commit, IsCommitted, Rollback,
// ForceBack, and Close, this structure provides methods to retrieve the
// cliargs.Cmd struct that stores the results of command line argument parsing,
// an array of cliargs.OptCfg struct for storing command line argument
// configurations, and methods to set and retrieve any type struct instance
// generated from the results of command line argument parsing.
type DaxConn struct {
	ds *DaxSrc
}

// Cmd is the method to retrieve a cliargs.Cmd struct instance that stores the
// results of command line argument parsing.
func (conn DaxConn) Cmd() cliargs.Cmd {
	return conn.ds.cmd
}

// OptCfgs is the method to retrieve an array of cliargs.OptCfg struct
// instances.
// This array is either passed as an argument to NewDaxSrcWithOptCfgs function
// or parsed from the struct instance passed as an argument to
// NewDaxSrcForOptions function.
func (conn DaxConn) OptCfgs() []cliargs.OptCfg {
	return conn.ds.optCfgs
}

// Options is the method to retrieve a struct instance of any type, which
// is either passed as an argument to NewDaxSrcForOptions or set by
// DaxConn#SetOptions method.
func (conn DaxConn) Options() any {
	return conn.ds.options
}

// SetOptions is the method to set a struct instance of any type to a DaxSrc
// instance through this DaxConn instance..
// Because this argument is set to a DaxSrc instance, it is persists even
// after the transaction has ended.
// If the DaxSrc instance is global, the argument instance will persist until
// the application is terminated (until the sabi.Close function is called).
func (conn DaxConn) SetOptions(opts any) {
	conn.ds.options = opts
}

// Commit is the one of the required methods for a struct that inherits
// sabi.DaxConn.
// It is called by sabi.Txn function.
// This method is empty and only returns a result of errs.Ok().
func (conn DaxConn) Commit(ag sabi.AsyncGroup) errs.Err {
	return errs.Ok()
}

// IsCommitted is the one of the required methods for a struct that inherits
// sabi.DaxConn.
// It is called by sabi.Txn function.
// This method always returns true.
func (conn DaxConn) IsCommitted() bool {
	return true
}

// Rollback is the one of the required methods for a struct that inherits
// sabi.DaxConn.
// This method never be called because IsCommitted always returns true.
func (conn DaxConn) Rollback(ag sabi.AsyncGroup) {
	// never be run because IsCommitted always returns true.
}

// ForceBack is the one of the required methods for a struct that inherits
// sabi.DaxConn.
// This method is empty and does nothing.
func (conn DaxConn) ForceBack(ag sabi.AsyncGroup) {
}

// Close is the one of the required methods for a struct that inherits
// sabi.DaxConn.
// This method is empty and does nothing.
func (conn DaxConn) Close() {
}

// DaxSrc is the dax source struct for command line argument operations.
// This struct stores the results of command line argument parsing, and
// provides them via a DaxConn instance.
type DaxSrc struct {
	cmd     cliargs.Cmd
	optCfgs []cliargs.OptCfg
	options any
}

// Setup is the one of the required methods for a struct that inherits
// sabi.DaxSrc.
// This method parses command line arguments and sets the results of the
// parsing to this DaxSrc instance.
// If failing to parse, this method returns errs.Err instnace that holds an
// error instance from cliargs.Parse/ParseWith/ParseFor function as the error
// reason.
func (ds *DaxSrc) Setup(ag sabi.AsyncGroup) errs.Err {
	if ds.options != nil {
		cmd, optCfgs, e := cliargs.ParseFor(os.Args, ds.options)
		if e != nil {
			return errs.New(e)
		}
		ds.cmd = cmd
		ds.optCfgs = optCfgs
	} else if len(ds.optCfgs) > 0 {
		cmd, e := cliargs.ParseWith(os.Args, ds.optCfgs)
		if e != nil {
			return errs.New(e)
		}
		ds.cmd = cmd
	} else {
		cmd, e := cliargs.Parse()
		if e != nil {
			return errs.New(e)
		}
		ds.cmd = cmd
	}

	return errs.Ok()
}

// Close is the one of the required methods for a struct that inherits
// sabi.DaxSrc.
// This method is empty and does nothing.
func (ds *DaxSrc) Close() {
}

// CreateDaxConn is the one of the required methods for a struct that inherits
// sabi.DaxSrc.
// This method creates a new instance of cliargdax.DaxConn struct.
func (ds *DaxSrc) CreateDaxConn() (sabi.DaxConn, errs.Err) {
	return DaxConn{ds: ds}, errs.Ok()
}

// NewDaxSrc is the constructor function of cliargdax.DaxSrc struct.
func NewDaxSrc() *DaxSrc {
	return &DaxSrc{}
}

// NewDaxSrcWithOptCfgs is the constructor function for cliargdax.DaxSrc struct
// that takes an array of instances of the cliargs.OptCfg struct.
func NewDaxSrcWithOptCfgs(cfgs []cliargs.OptCfg) *DaxSrc {
	return &DaxSrc{optCfgs: cfgs}
}

// NewDaxSrcForOptions is the constructor function for cliargdax.DaxSrc struct
// that takes an instnace of a struct of any type, which stores the results of
// command line argument parsing.
func NewDaxSrcForOptions(opts any) *DaxSrc {
	return &DaxSrc{options: opts}
}

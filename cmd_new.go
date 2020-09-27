////////////////////////////////////////////////////////////////////////////
// Program: fiber
// Purpose: Fiber application generator
// Authors: fiber https://github.com/gofiber/fiber (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// new

func newCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*newT)
	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.Self, Opts.Verbose =
		rootArgv.Self, rootArgv.Verbose.Value()
	// argv.OptHotReload, argv.OptNoView, argv.OptView, argv.OptGit, argv.OptForce,
	//return nil
	return DoNew()
}

//
// DoNew implements the business logic of command `new`
func DoNew() error {
	fmt.Fprintf(os.Stderr, "%s v%s. new - Fiber application generator\n", progname, version)
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
	return nil
}

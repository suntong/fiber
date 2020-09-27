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

////////////////////////////////////////////////////////////////////////////
// help

func helpCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*helpT)
	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.Self, Opts.Verbose =
		rootArgv.Self, rootArgv.Verbose.Value()
	//
	//return nil
	return DoHelp()
}

//
// DoHelp implements the business logic of command `help`
func DoHelp() error {
	fmt.Fprintf(os.Stderr, "%s v%s. help - Output usage information\n", progname, version)
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
	return nil
}

////////////////////////////////////////////////////////////////////////////
// version

func versionCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*versionT)
	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.Self, Opts.Verbose =
		rootArgv.Self, rootArgv.Verbose.Value()
	//
	//return nil
	return DoVersion()
}

//
// DoVersion implements the business logic of command `version`
func DoVersion() error {
	fmt.Fprintf(os.Stderr, "%s v%s. version - Output the fiber version number\n", progname, version)
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
	return nil
}

////////////////////////////////////////////////////////////////////////////
// run

func runCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*runT)
	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.Self, Opts.Verbose =
		rootArgv.Self, rootArgv.Verbose.Value()
	//
	//return nil
	return DoRun()
}

//
// DoRun implements the business logic of command `run`
func DoRun() error {
	fmt.Fprintf(os.Stderr, "%s v%s. run - Starts your app with hot reload support\n", progname, version)
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
	return nil
}

////////////////////////////////////////////////////////////////////////////
// pack

func packCLI(ctx *cli.Context) error {
	rootArgv = ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*packT)
	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
	Opts.Self, Opts.Verbose =
		rootArgv.Self, rootArgv.Verbose.Value()
	//
	//return nil
	return DoPack()
}

//
// DoPack implements the business logic of command `pack`
func DoPack() error {
	fmt.Fprintf(os.Stderr, "%s v%s. pack - Compress a Fiber application into a single binary\n", progname, version)
	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
	return nil
}

////////////////////////////////////////////////////////////////////////////
// Program: fiber
// Purpose: Fiber application framework
// Authors: fiber https://github.com/gofiber/fiber (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	//  	"fmt"
	//  	"os"

	"github.com/mkideal/cli"
	//  	"github.com/mkideal/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// fiber

type rootT struct {
	cli.Helper
	Self    *rootT      `cli:"c,config" usage:"config file\n" json:"-" parser:"jsoncfg" dft:"fiber_cfg.json"`
	Verbose cli.Counter `cli:"V,verbose" usage:"Verbose mode (Multiple -V options increase the verbosity)\n"`
}

var root = &cli.Command{
	Name: "fiber",
	Desc: "Fiber application framework\nVersion " + version + " built on " + date +
		"\nCopyright (C) 2020, fiber https://github.com/gofiber/fiber",
	Text:   "Fiber -- Express inspired web framework in Go",
	Global: true,
	Argv:   func() interface{} { t := new(rootT); t.Self = t; return t },
	Fn:     fiber,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Self	*rootT
//  	Verbose	cli.Counter
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "fiber"
//          version   = "0.1.0"
//          date = "2020-09-27"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,
//  		cli.Tree(helpDef),
//  		cli.Tree(versionDef),
//  		cli.Tree(newDef),
//  		cli.Tree(runDef),
//  		cli.Tree(packDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  		os.Exit(1)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

//  func fiber(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// help

//  func helpCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*helpT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Self, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Self, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	//
//  	//return nil
//  	return DoHelp()
//  }
//
// DoHelp implements the business logic of command `help`
//  func DoHelp() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. help - Output usage information\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
//  	return nil
//  }

type helpT struct {
}

var helpDef = &cli.Command{
	Name: "help",
	Desc: "Output usage information",

	Argv: func() interface{} { return new(helpT) },
	Fn:   helpCLI,
}

////////////////////////////////////////////////////////////////////////////
// version

//  func versionCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*versionT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Self, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Self, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	//
//  	//return nil
//  	return DoVersion()
//  }
//
// DoVersion implements the business logic of command `version`
//  func DoVersion() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. version - Output the fiber version number\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
//  	return nil
//  }

type versionT struct {
}

var versionDef = &cli.Command{
	Name: "version",
	Desc: "Output the fiber version number",

	Argv: func() interface{} { return new(versionT) },
	Fn:   versionCLI,
}

////////////////////////////////////////////////////////////////////////////
// new

//  func newCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*newT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Self, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Self, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	// argv.OptHotReload, argv.OptNoView, argv.OptView, argv.OptGit, argv.OptForce,
//  	//return nil
//  	return DoNew()
//  }
//
// DoNew implements the business logic of command `new`
//  func DoNew() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. new - Fiber application generator\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
//  	return nil
//  }

type newT struct {
	OptHotReload bool   `cli:"hot-reload" usage:"generate hot-reload program logic"`
	OptNoView    bool   `cli:"no-view" usage:"generate without view engine"`
	OptView      string `cli:"v,view" usage:"add view <engine> support (html|ace|amber|django|handlebars|jet|mustache|pug). defaults to" dft:"html"`
	OptGit       bool   `cli:"git" usage:"add .gitignore"`
	OptForce     string `cli:"f,force" usage:"force on non-empty directory"`
}

var newDef = &cli.Command{
	Name: "new",
	Desc: "Fiber application generator",
	Text: "Usage:\n  fiber new [options] [dir]",
	Argv: func() interface{} { return new(newT) },
	Fn:   newCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}

////////////////////////////////////////////////////////////////////////////
// run

//  func runCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*runT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Self, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Self, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	//
//  	//return nil
//  	return DoRun()
//  }
//
// DoRun implements the business logic of command `run`
//  func DoRun() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. run - Starts your app with hot reload support\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
//  	return nil
//  }

type runT struct {
}

var runDef = &cli.Command{
	Name: "run",
	Desc: "Starts your app with hot reload support",

	Argv: func() interface{} { return new(runT) },
	Fn:   runCLI,
}

////////////////////////////////////////////////////////////////////////////
// pack

//  func packCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*packT)
//  	clis.Setup(fmt.Sprintf("%s::%s", progname, ctx.Path()), rootArgv.Verbose.Value())
//  	clis.Verbose(2, "<%s> -\n  %+v\n  %+v\n  %v\n", ctx.Path(), rootArgv, argv, ctx.Args())
//  	Opts.Self, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Self, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	//
//  	//return nil
//  	return DoPack()
//  }
//
// DoPack implements the business logic of command `pack`
//  func DoPack() error {
//  	fmt.Fprintf(os.Stderr, "%s v%s. pack - Compress a Fiber application into a single binary\n", progname, version)
//  	// fmt.Fprintf(os.Stderr, "Copyright (C) 2020, fiber https://github.com/gofiber/fiber\n\n")
//  	return nil
//  }

type packT struct {
}

var packDef = &cli.Command{
	Name: "pack",
	Desc: "Compress a Fiber application into a single binary",

	Argv: func() interface{} { return new(packT) },
	Fn:   packCLI,
}

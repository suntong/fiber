////////////////////////////////////////////////////////////////////////////
// Program: fiber
// Purpose: Fiber application generator
// Authors: fiber https://github.com/gofiber/fiber (c) 2020, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v fiber_cliGen.sh

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
type OptsT struct {
	Self    *rootT
	Verbose int
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "fiber"
	version  = "0.1.0"
	date     = "2020-09-27"

	rootArgv *rootT
	// Opts store all the configurable options
	Opts OptsT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root,
		cli.Tree(helpDef),
		cli.Tree(versionDef),
		cli.Tree(newDef),
		cli.Tree(runDef),
		cli.Tree(packDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("")
}

//==========================================================================
// Dumb root handler

func fiber(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}

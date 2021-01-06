package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	name = "gitu"

	debugFlag = "debug"
)

func main() {
	a := &action{}

	app := &cli.App{
		Name:  name,
		Usage: "switch git user",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    debugFlag,
				Aliases: []string{"d"},
				Usage:   "show debugging info",
			},
		},
		Action: a.Run,
	}

	if err := app.Run(os.Args); err != nil {
		// Highlight error
		fmtErr := color.New(color.FgRed)
		fmtErr.Printf("[%s error]: ", name)
		fmt.Printf("%s\n", err.Error())
	}
}

type action struct {
	debug bool
}

func (a *action) Run(c *cli.Context) error {
	// Show help if there is nothing
	if c.NArg() == 0 && c.NumFlags() == 0 {
		return cli.ShowAppHelp(c)
	}

	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.debug = c.Bool(debugFlag)
}

func (a *action) logDebug(format string, v ...interface{}) {
	if a.debug {
		log.Printf(format, v...)
	}
}

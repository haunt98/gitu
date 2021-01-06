package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

const (
	name = "gitu"

	debugFlag = "debug"
	nameFlag  = "name"
	emailFlag = "email"

	addCommand = "add"
)

func main() {
	a := &action{
		flags: make(map[string]string),
	}

	app := &cli.App{
		Name:  name,
		Usage: "switch git user",
		Commands: []*cli.Command{
			{
				Name:    addCommand,
				Aliases: []string{"a"},
				Usage:   "add git user",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  nameFlag,
						Usage: "gitconfig user.name",
					},
					&cli.StringFlag{
						Name:  emailFlag,
						Usage: "gitconfig user.email",
					},
				},
				Action: a.RunAdd,
			},
		},
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
	flags map[string]string
}

// Show help by default
func (a *action) Run(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) RunAdd(c *cli.Context) error {
	a.getFlags(c)

	if a.flags[nameFlag] == "" {
		fmt.Println("What is your name?")
		a.flags[nameFlag] = readStdin()
		fmt.Printf("Hello %s\n", a.flags[nameFlag])
	}

	if a.flags[emailFlag] == "" {
		fmt.Println("What is your email?")
		a.flags[emailFlag] = readStdin()
		fmt.Printf("Copy that %s\n", a.flags[emailFlag])
	}

	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.debug = c.Bool(debugFlag)
	a.flags[nameFlag] = c.String(nameFlag)
	a.flags[emailFlag] = c.String(emailFlag)
}

func (a *action) logDebug(format string, v ...interface{}) {
	if a.debug {
		log.Printf(format, v...)
	}
}

func readStdin() string {
	bs := bufio.NewScanner(os.Stdin)
	for bs.Scan() {
		line := bs.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		return line
		break
	}

	return ""
}

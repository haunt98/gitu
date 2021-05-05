package cli

import (
	"os"

	"github.com/haunt98/color"
	"github.com/urfave/cli/v2"
)

const (
	appName  = "gitu"
	appUsage = "switch git user fastly"

	// flags
	nameFlag     = "name"
	emailFlag    = "email"
	nicknameFlag = "nickname"
	allFlag      = "all"

	// commands
	addCommand    = "add"
	switchCommand = "switch"
	statusCommand = "status"
	listCommand   = "list"
	deleteCommand = "delete"

	// flag usage
	nameUsage     = "gitconfig user.name"
	emailUsage    = "gitconfig user.email"
	nicknameUsage = "nickname to choose"
	allUsage      = "select all nicknames"

	// command usage
	addUsage    = "add git user"
	switchUsage = "switch git user"
	statusUsage = "show current name and email"
	listUsage   = "list all saved names and emails"
	deleteUsage = "delete saved name and email"
)

var (
	// flag aliases
	allAliases = []string{"a"}

	// command aliases
	addAliases    = []string{"a"}
	switchAliases = []string{"sw"}
	statusAliases = []string{"st"}
	listAliases   = []string{"l"}
	deleteAliases = []string{"d"}
)

type App struct {
	cliApp *cli.App
}

func NewApp() *App {
	a := &action{}

	cliApp := &cli.App{
		Name:  appName,
		Usage: appUsage,
		Commands: []*cli.Command{
			{
				Name:    addCommand,
				Aliases: addAliases,
				Usage:   addUsage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  nameFlag,
						Usage: nameUsage,
					},
					&cli.StringFlag{
						Name:  emailFlag,
						Usage: emailUsage,
					},
					&cli.StringFlag{
						Name:  nicknameFlag,
						Usage: nicknameUsage,
					},
				},
				Action: a.runAdd,
			},
			{
				Name:    switchCommand,
				Aliases: switchAliases,
				Usage:   switchUsage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  nicknameFlag,
						Usage: nicknameUsage,
					},
				},
				Action: a.runSwitch,
			},
			{
				Name:    statusCommand,
				Aliases: statusAliases,
				Usage:   statusUsage,
				Action:  a.runStatus,
			},
			{
				Name:    listCommand,
				Aliases: listAliases,
				Usage:   listUsage,
				Action:  a.runList,
			},
			{
				Name:    deleteCommand,
				Aliases: deleteAliases,
				Usage:   deleteUsage,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  nicknameFlag,
						Usage: nicknameUsage,
					},
					&cli.BoolFlag{
						Name:    allFlag,
						Aliases: allAliases,
						Usage:   allUsage,
					},
				},
				Action: a.runDelete,
			},
		},
		Action: a.runHelp,
	}

	return &App{
		cliApp: cliApp,
	}
}

func (a *App) Run() {
	if err := a.cliApp.Run(os.Args); err != nil {
		color.PrintAppError(appName, err.Error())
	}
}

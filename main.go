package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
	"github.com/urfave/cli/v2"
)

const (
	appName = "gitu"

	nameFlag     = "name"
	emailFlag    = "email"
	nicknameFlag = "nickname"
	allFlag      = "all"

	addCommand    = "add"
	switchCommand = "switch"
	statusCommand = "status"
	listCommand   = "list"
	deleteCommand = "delete"
)

func main() {
	a := &action{}

	app := &cli.App{
		Name:  appName,
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
					&cli.StringFlag{
						Name:  nicknameFlag,
						Usage: "nickname for quick access",
					},
				},
				Action: a.RunAdd,
			},
			{
				Name:    switchCommand,
				Aliases: []string{"sw"},
				Usage:   "switch git user",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  nicknameFlag,
						Usage: "nickname to switch",
					},
				},
				Action: a.RunSwitch,
			},
			{
				Name:    statusCommand,
				Aliases: []string{"st"},
				Usage:   "show current name and email",
				Action:  a.RunStatus,
			},
			{
				Name:    listCommand,
				Aliases: []string{"l"},
				Usage:   "list all saved name and email in",
				Action:  a.RunList,
			},
			{
				Name:    deleteCommand,
				Aliases: []string{"d"},
				Usage:   "delete saved name and email",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  nicknameFlag,
						Usage: "nickname to delete",
					},
					&cli.BoolFlag{
						Name:    allFlag,
						Aliases: []string{"-a"},
						Usage:   "delete all, be careful",
					},
				},
				Action: a.RunDelete,
			},
		},
		Action: a.Run,
	}

	if err := app.Run(os.Args); err != nil {
		// Highlight error
		fmtErr := color.New(color.FgRed)
		fmtErr.Printf("[%s error]: ", appName)
		fmt.Printf("%s\n", err.Error())
	}
}

type action struct {
	flags struct {
		name     string
		email    string
		nickname string
		all      bool
	}
}

// Show help by default
func (a *action) Run(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) RunAdd(c *cli.Context) error {
	a.getFlags(c)

	if a.flags.name == "" {
		fmt.Println("What is your name?")
		a.flags.name = readStdin()
		fmt.Printf("Hello %s\n", a.flags.name)
	}

	if a.flags.email == "" {
		fmt.Println("What is your email?")
		a.flags.email = readStdin()
		fmt.Printf("Copy that %s\n", a.flags.email)
	}

	if a.flags.nickname == "" {
		fmt.Println("What should I call you?")
		a.flags.nickname = readStdin()
		fmt.Printf("Nice nickname %s\n", a.flags.nickname)
	}

	cfg, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.CheckExist(a.flags.nickname) {
		fmt.Printf("Nickname %s already exist, replace it with new user (y/n)? ", a.flags.nickname)
		answer := readStdin()
		if !strings.EqualFold(answer, "y") {
			fmt.Println("Nothing changed :D")
			return nil
		}
	}

	cfg.Update(a.flags.nickname, User{
		Name:  a.flags.name,
		Email: a.flags.email,
	})

	if err := SaveConfig(&cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

func (a *action) RunSwitch(c *cli.Context) error {
	a.getFlags(c)

	if a.flags.nickname == "" {
		fmt.Println("Which nickname you choose?")
		a.flags.nickname = readStdin()
		fmt.Printf("Switching to nickname %s...\n", a.flags.nickname)
	}

	cfg, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	user, ok := cfg.Get(a.flags.nickname)
	if !ok {
		fmt.Printf("Nickname %s is not exist :(\n", a.flags.nickname)
		return nil
	}

	repo, err := git.PlainOpen(".")
	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			fmt.Println("This is not git repository mate :(")
			return nil
		}

		return fmt.Errorf("failed to open repository: %w", err)
	}

	repoCfg, err := repo.Config()
	if err != nil {
		return fmt.Errorf("failed to get repository config: %w", err)
	}

	// Update name and email
	repoCfg.User.Name = user.Name
	repoCfg.User.Email = user.Email
	if err := repo.SetConfig(repoCfg); err != nil {
		return fmt.Errorf("failed to set repository config: %w", err)
	}

	return nil
}

func (a *action) RunStatus(c *cli.Context) error {
	repo, err := git.PlainOpen(".")
	if err != nil {
		if errors.Is(err, git.ErrRepositoryNotExists) {
			fmt.Println("This is not git repository mate :(")
			return nil
		}

		return fmt.Errorf("failed to open repository: %w", err)
	}

	repoCfg, err := repo.Config()
	if err != nil {
		return fmt.Errorf("failed to get repository config: %w", err)
	}

	fmt.Printf("Name: %s\n", repoCfg.User.Name)
	fmt.Printf("Email: %s\n", repoCfg.User.Email)

	return nil
}

func (a *action) RunList(c *cli.Context) error {
	cfg, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	users := cfg.GetAll()
	for nickname, user := range users {
		fmt.Printf("Nickname: %s\n", nickname)
		fmt.Printf("Name: %s\n", user.Name)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Println()
	}

	return nil
}

func (a *action) RunDelete(c *cli.Context) error {
	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.name = c.String(nameFlag)
	a.flags.email = c.String(emailFlag)
	a.flags.nickname = c.String(nicknameFlag)
	a.flags.all = c.Bool(allFlag)
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
	}

	return ""
}

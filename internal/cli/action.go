package cli

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/haunt98/gitu/internal/config"
	"github.com/haunt98/ioe-go"
	"github.com/urfave/cli/v2"
)

type action struct {
	flags struct {
		name     string
		email    string
		nickname string
		all      bool
	}
}

func (a *action) runHelp(c *cli.Context) error {
	return cli.ShowAppHelp(c)
}

func (a *action) runAdd(c *cli.Context) error {
	a.getFlags(c)

	if a.flags.name == "" {
		fmt.Println("What is your name?")
		a.flags.name = ioe.ReadInput()
		fmt.Printf("Hello %s\n", a.flags.name)
	}

	if a.flags.email == "" {
		fmt.Println("What is your email?")
		a.flags.email = ioe.ReadInput()
		fmt.Printf("Copy that %s\n", a.flags.email)
	}

	if a.flags.nickname == "" {
		fmt.Println("What should I call you?")
		a.flags.nickname = ioe.ReadInput()
		fmt.Printf("Nice nickname %s\n", a.flags.nickname)
	}

	cfg, err := config.LoadConfig(appName)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if cfg.CheckExist(a.flags.nickname) {
		fmt.Printf("Nickname %s already exist, replace it with new user (y/n)? ", a.flags.nickname)
		answer := ioe.ReadInput()
		if !strings.EqualFold(answer, "y") {
			fmt.Println("Nothing changed :D")
			return nil
		}
	}

	cfg.Update(a.flags.nickname, config.User{
		Name:  a.flags.name,
		Email: a.flags.email,
	})

	if err := config.SaveConfig(appName, &cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

func (a *action) runSwitch(c *cli.Context) error {
	a.getFlags(c)

	if a.flags.nickname == "" {
		fmt.Println("Which nickname you choose?")
		a.flags.nickname = ioe.ReadInput()
		fmt.Printf("Switching to nickname %s\n", a.flags.nickname)
	}

	cfg, err := config.LoadConfig(appName)
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

func (a *action) runStatus(c *cli.Context) error {
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

func (a *action) runList(c *cli.Context) error {
	cfg, err := config.LoadConfig(appName)
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

func (a *action) runDelete(c *cli.Context) error {
	a.getFlags(c)

	cfg, err := config.LoadConfig(appName)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if a.flags.all {
		fmt.Print("Do you sure you want to wipe out all saved user (y/n) ")
		answer := ioe.ReadInput()
		if strings.EqualFold(answer, "y") {
			cfg.DeleteAll()
			fmt.Println("Eveything is deleted")

			if err := config.SaveConfig(appName, &cfg); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			return nil
		}
	}

	if a.flags.nickname == "" {
		fmt.Println("Which nickname you want to delete?")
		a.flags.nickname = ioe.ReadInput()
		fmt.Printf("Deleting nickname %s\n", a.flags.nickname)
	}

	cfg.Delete(a.flags.nickname)
	if err := config.SaveConfig(appName, &cfg); err != nil {
		return fmt.Errorf("failed to save config: %w", err)
	}

	return nil
}

func (a *action) getFlags(c *cli.Context) {
	a.flags.name = c.String(nameFlag)
	a.flags.email = c.String(emailFlag)
	a.flags.nickname = c.String(nicknameFlag)
	a.flags.all = c.Bool(allFlag)
}

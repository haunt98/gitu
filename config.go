// Config stores at $XDG_CONFIG_HOME/gitu/config.json

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/haunt98/xdg"
)

const (
	configFile = "config.json"

	indent = "  "
)

type Config struct {
	Users map[string]User `json:"users"`
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Load config from file, return empty if file not found
func Load() (Config, error) {
	path := getPath()
	f, err := os.Open(path)
	if err != nil {
		// https://github.com/golang/go/blob/3e1e13ce6d1271f49f3d8ee359689145a6995bad/src/os/error.go#L90-L91
		if errors.Is(err, os.ErrNotExist) {
			return Config{
				Users: make(map[string]User),
			}, nil
		}

		return Config{}, fmt.Errorf("failed to open %s: %w", path, err)
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read %s: %w", path, err)
	}

	var result Config
	if err := json.Unmarshal(bytes, &result); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal: %w,", err)
	}

	if result.Users == nil {
		result.Users = make(map[string]User)
	}

	return result, nil
}

// Save config to file
func Save(c *Config) error {
	bytes, err := json.MarshalIndent(c, "", indent)
	if err != nil {
		return fmt.Errorf("failed to marshall: %w", err)
	}

	path := getPath()
	if err := ioutil.WriteFile(path, bytes, 0644); err != nil {
		return fmt.Errorf("failed to write %s: %w", path, err)
	}

	return nil
}

func (c *Config) CheckExist(nickname string) bool {
	_, ok := c.Users[nickname]
	return ok
}

func (c *Config) Update(nickname string, user User) {
	c.Users[nickname] = user
}

func (c *Config) Delete(nickname string) {
	delete(c.Users, nickname)
}

func getPath() string {
	return filepath.Join(xdg.GetConfigHome(), name, configFile)
}

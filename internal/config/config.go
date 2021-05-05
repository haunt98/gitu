// Config location
// Unix: $XDG_CONFIG_HOME/gitu/config.json
// Windows: %LocalAppData%\gitu\config.json

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
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

// LoadConfig config from file, return empty if file not found
func LoadConfig(appName string) (Config, error) {
	_, filePath, err := getConfigPath(appName)
	if err != nil {
		return Config{}, err
	}

	f, err := os.Open(filePath)
	if err != nil {
		// https://github.com/golang/go/blob/3e1e13ce6d1271f49f3d8ee359689145a6995bad/src/os/error.go#L90-L91
		if errors.Is(err, os.ErrNotExist) {
			return Config{
				Users: make(map[string]User),
			}, nil
		}

		return Config{}, fmt.Errorf("failed to open %s: %w", filePath, err)
	}
	defer f.Close()

	bytes, err := io.ReadAll(f)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read %s: %w", filePath, err)
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

// SaveConfig config to file
func SaveConfig(appName string, c *Config) error {
	dirPath, filePath, err := getConfigPath(appName)
	if err != nil {
		return fmt.Errorf("failed to get config path: %w", err)
	}

	// Make sure dir is exist
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to mkdir %s: %w", dirPath, err)
	}

	bytes, err := json.MarshalIndent(c, "", indent)
	if err != nil {
		return fmt.Errorf("failed to marshall: %w", err)
	}

	if err := os.WriteFile(filePath, bytes, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write file: %s: %w", filePath, err)
	}

	return nil
}

func (c *Config) CheckExist(nickname string) bool {
	_, ok := c.Users[nickname]
	return ok
}

func (c *Config) Get(nickname string) (User, bool) {
	user, ok := c.Users[nickname]
	if !ok {
		return User{}, false
	}

	return user, true
}

func (c *Config) GetAll() map[string]User {
	return c.Users
}

func (c *Config) Update(nickname string, user User) {
	c.Users[nickname] = user
}

func (c *Config) Delete(nickname string) {
	delete(c.Users, nickname)
}

func (c *Config) DeleteAll() {
	c.Users = make(map[string]User)
}

func getConfigPath(appName string) (dirPath, filePath string, err error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", "", fmt.Errorf("failed to get user config dir: %w", err)
	}

	dirPath = filepath.Join(cfgDir, appName)
	filePath = filepath.Join(cfgDir, appName, configFile)
	return
}

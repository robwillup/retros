package config

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"

	"github.com/robwillup/retros/src/clientos"
	"github.com/robwillup/retros/src/sshutils"
)

func Create(config sshutils.SSHConfig) (*os.File, error) {
	if config.Host == "" {
		return nil, errors.New("Host is required")
	}

	if config.Username == "" {
		return nil, errors.New("Username is required")
	}

	if config.KeyPath == "" {
		return nil, errors.New("KeyPath is required")
	}

	home := clientos.GetHomeDir()
	f, err := os.Create(filepath.Clean(home + "/.retros"))

	if err != nil {
		return nil, err
	}

	defer f.Close()

	_, err = f.Write([]byte(config.Host + "\n"))
	if err != nil {
		return nil, err
	}

	_, err = f.Write([]byte(config.Username + "\n"))
	if err != nil {
		return nil, err
	}

	_, err = f.Write([]byte(config.KeyPath))
	if err != nil {
		return nil, err
	}

	return f, nil
}

func CheckIfExists() bool {
	home := clientos.GetHomeDir()
	if _, err := os.Stat(home + "/.retros"); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func Read() (sshutils.SSHConfig, error) {
	configValues := []string{}
	config := sshutils.SSHConfig{}
	home := clientos.GetHomeDir()
	file, err := os.Open(filepath.Clean(home + "/.retros"))
	if err != nil {
		return config, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		configValues = append(configValues, scanner.Text())
	}

	if len(configValues) < 3 {
		return config, errors.New("Configuration file is incomplete")
	}

	config.Host = configValues[0]
	config.Username = configValues[1]
	config.KeyPath = configValues[2]
	config.Port = 22

	return config, nil
}

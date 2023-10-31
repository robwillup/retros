package config

import (
	"bufio"
	"errors"
	"os"

	"github.com/robwillup/rosy/src/sshutils"
)

func Create(config sshutils.SSHConfig) error {
	if config.Host == "" {
		return errors.New("Host is required")
	}

	if config.Username == "" {
		return errors.New("Username is required")
	}

	if config.KeyPath == "" {
		return errors.New("KeyPath is required")
	}

	home := os.Getenv("HOME")
	f, err := os.Create(home + "/.rosy")
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write([]byte(config.Host + "\n"))
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(config.Username + "\n"))
	if err != nil {
		return err
	}

	_, err = f.Write([]byte(config.KeyPath))
	if err != nil {
		return err
	}

	return nil
}

func CheckIfExists() bool {
	home := os.Getenv("HOME")
	if _, err := os.Stat(home + "/.rosy"); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func Read() (sshutils.SSHConfig, error) {
	configValues := []string{}
	config := sshutils.SSHConfig{}
	home := os.Getenv("HOME")
	file, err := os.Open(home + "/.rosy")
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

package config

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/robwillup/retros/internal/clientos"
	"github.com/robwillup/retros/internal/sshutils"
	"gopkg.in/yaml.v3"
)

const ConfigFileName = ".retros.yml"

func Create(config sshutils.SSHConfig) error {
	configPath := filepath.Join(clientos.GetHomeDir(), ".retros")

	err := os.MkdirAll(configPath, 0750)
	if err != nil {
		return err
	}

	f, err := os.Create(filepath.Clean(filepath.Join(configPath, ConfigFileName)))

	if err != nil {
		return err
	}

	defer f.Close()

	yaml, err := yaml.Marshal(config)

	if err != nil {
		return err
	}

	err = os.WriteFile(filepath.Join(configPath, ConfigFileName), yaml, fs.FileMode(os.O_CREATE|os.O_WRONLY))

	if err != nil {
		return err
	}

	return nil
}

func Read() (sshutils.SSHConfig, error) {
	config := sshutils.SSHConfig{}
	bytes, err := os.ReadFile(filepath.Join(clientos.GetHomeDir(), ".retros", ConfigFileName))
	if err != nil {
		return config, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return config, err
	}

	return config, nil
}

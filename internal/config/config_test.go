package config

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/robwillup/retros/internal/clientos"
	"github.com/robwillup/retros/internal/filesystem"
	"github.com/robwillup/retros/internal/sshutils"
)

var configPath string = filepath.Join(clientos.GetHomeDir(), ".retros/", CONFIG_FILE_NAME)

// Calls config.Create checking if the config file `.retros.yml` is created.
func TestCreate(t *testing.T) {
	// Arrange
	conf := sshutils.SSHConfig{
		Host:     "test",
		Username: "test",
		KeyPath:  "test",
	}

	if filesystem.CheckIfExists(configPath) {
		existingConf, err := Read()

		if err != nil {
			t.Fatalf("Failed to read config file. Error: %v\n", err)
		}

		conf = existingConf
	}

	// Act
	err := Create(conf)

	// Assert
	if err != nil {
		t.Fatalf("Failed to create config file with error: %v", err)
	}

	f, err := os.Open(configPath)
	if err != nil {
		t.Fatalf("Failed to create config file with error: %v", err)
	}

	defer f.Close()

	// Clean up
	if conf.Host == "test" {
		os.Remove(configPath)
	}
}

// Calls config.Read checking if the configuration returned is correct.
func TestRead(t *testing.T) {
	// Arrange
	if filesystem.CheckIfExists(configPath) {
		src, err := os.Open(configPath)

		if err != nil {
			t.Fatalf("Failed to backup original config file. Error: %v", err)
		}

		defer src.Close()

		dest, err := os.Create(fmt.Sprintf("%s_bak", configPath))

		if err != nil {
			t.Fatalf("Failed to backup original config file. Error: %v", err)
		}

		defer dest.Close()
		_, err = io.Copy(dest, src)
	}

	expected := sshutils.SSHConfig{
		Host:     "test",
		Username: "test",
		KeyPath:  "test",
		Port:     0,
	}

	err := Create(expected)

	if err != nil {
		t.Fatalf("Failed to create test config file. Error: %v", err)
	}

	// Act
	actual, err := Read()

	// Assert
	if err != nil {
		t.Fatalf("Failed to read config file. Error: %v", err)
	}

	if expected != actual {
		t.Fatalf("Expected: %v.\nActual: %v.\n", expected, actual)
	}

	// Clean up
	if _, err := os.Stat(fmt.Sprintf("%s_bak", configPath)); errors.Is(err, os.ErrNotExist) {
		return
	}

	src, err := os.Open(fmt.Sprintf("%s_bak", configPath))

	if err != nil {
		t.Fatalf("Failed to restore original config file. Error: %v", err)
	}

	defer src.Close()

	dest, err := os.Create(configPath)

	if err != nil {
		t.Fatalf("Failed to backup original config file. Error: %v", err)
	}

	defer dest.Close()
	_, err = io.Copy(dest, src)
}

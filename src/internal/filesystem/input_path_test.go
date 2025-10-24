package filesystem

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/robwillup/retros/src/internal/clientos"
	"github.com/robwillup/retros/src/internal/config"
	"github.com/robwillup/retros/src/internal/sshutils"
)

var configPath string = filepath.Join(clientos.GetHomeDir(), ".retros/", config.CONFIG_FILE_NAME)

func TestCheckDir(t *testing.T) {
	expected := false

	actual, err := IsDir("input_path.go")
	if err != nil {
		t.Fatalf("Failed to CheckDir(). Error: %t", err)
		return
	}

	if expected != actual {
		t.Fatal("CheckDir() failed. Expected file, returned dir.")
		return
	}
}

func TestGetFiles(t *testing.T) {
	expected := []string{"input_path.go", "input_path_test.go"}

	actual, err := GetFiles(".")
	if err != nil {
		t.Fatalf("Failed to GetFiles(). Error: %t", err)
		return
	}

	for i, actualFile := range actual {
		if expected[i] != actualFile {
			t.Fatalf("Failed to GetFiles().\nExpected: %v\nActual: %v", expected[i], actualFile)
			return
		}
	}
}

// Calls config.CheckIfExists to test if it's capable of detecting the config file.
func TestCheckIfExists(t *testing.T) {
	// Arrange
	expected := true

	if _, err := os.Stat(configPath); errors.Is(err, os.ErrNotExist) {
		conf := sshutils.SSHConfig{
			Host:     "test",
			Username: "test",
			KeyPath:  "test",
		}
		err = config.Create(conf)
		if err != nil {
			t.Fatalf("Failed to arrange for test with error: %v", err)
		}
	}

	// Act
	actual := CheckIfExists(configPath)

	// Assert
	if expected != actual {
		t.Fatal("Failed to check if config file exists")
	}

	// Clean up
	conf, err := config.Read()
	if err != nil {
		log.Println("Failed to complete test clean up")
		return
	}

	if conf.Host == "test" {
		os.Remove(configPath)
	}
}

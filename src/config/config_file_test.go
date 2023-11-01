package config

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/robwillup/rosy/src/sshutils"
)

// Calls config.Create checking if the config file `.rosy` is created.
func TestCreate(t *testing.T) {
	// Arrange
	home := os.Getenv("HOME")
	conf := sshutils.SSHConfig{
		Host:     "test",
		Username: "test",
		KeyPath:  "test",
	}

	if CheckIfExists() {
		conf, _ = Read()
	}

	// Act
	err := Create(conf)

	// Assert
	if err != nil {
		t.Fatalf("Failed to create config file with error: %v", err)
		return
	}

	_, err = os.Open(home + "/.rosy")
	if err != nil {
		t.Fatalf("Failed to create config file with error: %v", err)
		return
	}

	// Clean up
	if conf.Host == "test" {
		os.Remove(home + "/.rosy")
	}
}

// Calls config.CheckIfExists to test if it's capable of detecting the config file.
func TestCheckIfExists(t *testing.T) {
	// Arrange
	expected := true
	home := os.Getenv("HOME")

	if _, err := os.Stat(home + "/.rosy"); errors.Is(err, os.ErrNotExist) {
		conf := sshutils.SSHConfig{
			Host:     "test",
			Username: "test",
			KeyPath:  "test",
		}
		err = Create(conf)
		if err != nil {
			t.Fatalf("Failed to arrange for test with error: %v", err)
			return
		}
	}

	// Act
	actual := CheckIfExists()

	// Assert
	if expected != actual {
		t.Fatal("Failed to check if config file exists")
	}

	// Clean up
	conf, err := Read()
	if err != nil {
		log.Println("Failed to complete test clean up")
		return
	}

	if conf.Host == "test" {
		os.Remove(home + "/.rosy")
	}
}

// Calls config.Read checking if the configuration returned is correct.
func TestRead(t *testing.T) {
	// Arrange
	home := os.Getenv("HOME")

	if CheckIfExists() {
		err := os.Rename(home+"/.rosy", home+"/.rosy_bak")

		if err != nil {
			t.Fatalf("Failed to backup original config file. Error: %v", err)
			return
		}
	}

	expected := sshutils.SSHConfig{
		Host:     "test",
		Username: "test",
		KeyPath:  "test",
	}

	err := Create(expected)

	if err != nil {
		t.Fatalf("Failed to create test config file. Error: %v", err)
		return
	}

	// Act
	actual, err := Read()

	// Assert
	if err != nil {
		t.Fatalf("Failed to read config file. Error: %v", err)
		return
	}

	if expected.Host != actual.Host {
		t.Fatalf("Actual host (%s) does not match expected host (%s).", actual.Host, expected.Host)
	}

	// Clean up
	err = os.Remove(home + "/.rosy")

	if err != nil {
		t.Fatalf("Failed to remove test file. Error: %v", err)
	}

	if _, err := os.Stat(home + "/.rosy_bak"); errors.Is(err, os.ErrNotExist) {
		return
	}

	err = os.Rename(home+"/.rosy_bak", home+"/.rosy")

	if err != nil {
		t.Fatalf("Failed to restore config backup. Error: %v", err)
	}
}

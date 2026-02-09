package sshutils

import (
	"errors"
	"os"
	"testing"

	"github.com/robwillup/retros/internal/clientos"
)

func TestReadKnownHosts(t *testing.T) {
	home := clientos.GetHomeDir()
	if _, err := os.Stat(home + "/.ssh/known_hosts"); errors.Is(err, os.ErrNotExist) {
		return
	}

	actual, err := readKnownHosts()

	if err != nil {
		t.Fatalf("Failed to readKnownHosts(). Error: %t", err)
		return
	}

	if len(actual) < 1 {
		t.Fatalf("Failed to readKnownHosts(). No hosts found")
	}
}

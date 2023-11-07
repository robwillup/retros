package sshutils

import "testing"

func TestReadKnownHosts(t *testing.T) {
	actual, err := readKnownHosts()

	if err != nil {
		t.Fatalf("Failed to readKnownHosts(). Error: %t", err)
		return
	}

	if len(actual) < 1 {
		t.Fatalf("Failed to readKnownHosts(). No hosts found")
	}
}
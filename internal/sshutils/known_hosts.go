package sshutils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/robwillup/retros/internal/clientos"
)

// TODO: This function simply returns every line in `.ssh/known_hosts`.
// Improve it so it only returns hosts.
func readKnownHosts() ([]string, error) {
	known_hosts := []string{}
	home := clientos.GetHomeDir()
	file, err := os.Open(filepath.Clean(home + "/.ssh/known_hosts"))

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		known_hosts = append(known_hosts, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return known_hosts, nil
}

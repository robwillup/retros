package sshutils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readKnownHosts() ([]string, error) {
	known_hosts := []string{}
	home := os.Getenv("HOME")
	file, err := os.Open(home + "/.ssh/known_hosts")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "|1|") {
			host := scanner.Text()[61:len(scanner.Text())]
			known_hosts = append(known_hosts, host)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	return known_hosts, nil
}

package sshutils

import (
	"encoding/base64"
	"fmt"
	"net"
	"os"
	"strconv"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func EstablishSSHConnection(config SSHConfig) (*ssh.Client, error) {
	sshConfig := &ssh.ClientConfig{
		User:            config.Username,
		Auth:            []ssh.AuthMethod{},
		HostKeyCallback: trustedHostKeyCallback(config.HostKey),
	}

	if config.Password != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.Password(config.Password))
	} else if config.KeyPath != "" {
		key, err := os.ReadFile(config.KeyPath)
		if err != nil {
			return nil, err
		}
		signer, err := ssh.ParsePrivateKey(key)
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeys(signer))
	} else {
		socket := os.Getenv("SSH_AUTH_SOCKET")
		conn, err := net.Dial("unix", socket)
		if err != nil {
			return nil, err
		}

		agentClient := agent.NewClient(conn)
		if err != nil {
			return nil, err
		}
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agentClient.Signers))
	}

	fmt.Println("Got here")
	client, err := ssh.Dial("tcp", config.Host+":"+strconv.Itoa(config.Port), sshConfig)
	if err != nil {
		return nil, err
	}
	fmt.Println("Got here too")

	return client, nil
}

func trustedHostKeyCallback(trustedKey string) ssh.HostKeyCallback {

	if trustedKey == "" {
		return func(_ string, _ net.Addr, k ssh.PublicKey) error {
			var a []any = []any{"WARNING: SSH-key verification is *NOT* in effect: to fix, add this trustedKey: %q", keyString(k)}
			fmt.Fprintln(os.Stdout, a...)
			return nil
		}
	}

	return func(_ string, _ net.Addr, k ssh.PublicKey) error {
		ks := keyString(k)
		if trustedKey != ks {
			return fmt.Errorf("SSH-key verification: expected %q but got %q", trustedKey, ks)
		}

		return nil
	}
}

func keyString(k ssh.PublicKey) string {
	return k.Type() + " " + base64.StdEncoding.EncodeToString(k.Marshal())
}

func ExecuteRemoteCommand(client *ssh.Client, command string) (string, error) {
    // Create a new SSH session
    session, err := client.NewSession()
    if err != nil {
        return "", err
    }
    defer session.Close()

    // Run the remote command
    output, err := session.CombinedOutput(command)
    if err != nil {
        return "", err
    }

    return string(output), nil
}

package sshutils

type SSHConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	KeyPath  string
	HostKey  string
}
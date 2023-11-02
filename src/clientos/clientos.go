package clientos

import (
	"os"
	"runtime"
)

func GetHomeDir() string {
	plat := runtime.GOOS
	var home string

	if plat == "windows" {
		home = os.Getenv("USERPROFILE")
	} else {
		home = os.Getenv("HOME")
	}

	return home
}

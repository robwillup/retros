package clientos

import (
	"os/user"
	"runtime"
	"testing"
)

func TestGetHomeDir(t *testing.T) {
	plat := runtime.GOOS
	u, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get home dir. Error: %t", err)
		return
	}

	if plat == "linux" {
		expected := "/home/" + u.Username

		actual := GetHomeDir()

		if actual != expected {
			t.Fatalf("Failed to get home dir.\nexpected: %s\nactual: %s", expected, actual)
			return
		}

	} else if plat == "windows" {
		expected := "C:\\Users\\" + u.Username

		actual := GetHomeDir()

		if actual != expected {
			t.Fatalf("Failed to get home dir.\nexpected: %s\nactual: %s", expected, actual)
			return
		}
	}
}
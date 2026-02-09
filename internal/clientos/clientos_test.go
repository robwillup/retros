package clientos

import (
	"os/user"
	"runtime"
	"strings"
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
		// In Windows, user.Current().Username returns the "<MACHINE_NAME>\USERNAME".
		// That is why we split and get second element.
		expected := "C:\\Users\\" + strings.Split(u.Username, "\\")[1]

		actual := GetHomeDir()

		if actual != expected {
			t.Fatalf("Failed to get home dir.\nexpected: %s\nactual: %s", expected, actual)
			return
		}
	}
}
package filesystem

import (
	"testing"
)

func TestCheckDir(t *testing.T) {
	expected := false

	actual, err := CheckDir("input_path.go")
	if err != nil {
		t.Fatalf("Failed to CheckDir(). Error: %t", err)
		return
	}

	if expected != actual {
		t.Fatal("CheckDir() failed. Expected file, returned dir.")
		return
	}
}

func TestGetFiles(t *testing.T) {
	expected := []string{"input_path.go", "input_path_test.go"}

	actual, err := GetFiles(".")
	if err != nil {
		t.Fatalf("Failed to GetFiles(). Error: %t", err)
		return
	}

	for i, actualFile := range actual {
		if expected[i] != actualFile {
			t.Fatalf("Failed to GetFiles().\nExpected: %v\nActual: %v", expected[i], actualFile)
			return
		}
	}
}
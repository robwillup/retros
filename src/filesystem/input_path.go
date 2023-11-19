package filesystem

import (
	"errors"
	"os"
)

// IsDir calls os.Stat for the path provided and then calls fileInfo.IsDir to check whether it is a
// directory.
// It returns a boolean and and error. True if the path is a directory, or false if the path is a file.
func IsDir(fsPath string) (bool, error) {
	fileInfo, err := os.Stat(fsPath)

	if err != nil {
		return false, err
	}

	if fileInfo.IsDir() {
		return true, nil
	}

	return false, nil
}

// GetFiles returns a slice with the names of all files found under the directory provided.
// Only the file names are returned, not the full path.
func GetFiles(dir string) ([]string, error) {
	files := []string{}
	items, err := os.ReadDir(dir)

	if err != nil {
		return []string{}, nil
	}

	for _, item := range items {
		if !item.IsDir() {
			files = append(files, item.Name())
		}
	}

	return files, nil
}

func CheckIfExists(fsPath string) bool {
	if _, err := os.Stat(fsPath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

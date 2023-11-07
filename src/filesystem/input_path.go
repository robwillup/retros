package filesystem

import (
	"os"
)

func CheckDir(fsPath string) (bool, error) {
	fileInfo, err := os.Stat(fsPath)
	if err != nil {
		return false, err
	}
	if fileInfo.IsDir() {
		return true, nil
	}

	return false, nil
}

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
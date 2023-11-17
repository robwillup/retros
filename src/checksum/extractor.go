package checksum

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func GetChecksums(emulator, fsPath string) (map[string]ROM, error) {
	dataFilePath := "src/checksum/data/"

	if fsPath != "" {
		dataFilePath = fsPath
	}

	checksumsPath := filepath.Clean(filepath.Join(dataFilePath, fmt.Sprintf("%s.yml", emulator)))
	f, err := os.ReadFile(checksumsPath)
	if err != nil {
		return nil, err
	}

	out := make(map[string]ROM)
	if err := yaml.Unmarshal(f, &out); err != nil {
		return nil, err
	}

	return out, err
}
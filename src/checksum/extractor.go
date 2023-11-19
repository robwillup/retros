package checksum

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/robwillup/retros/src/clientos"
	"github.com/robwillup/retros/src/filesystem"
	"gopkg.in/yaml.v3"
)

func GetChecksums(emulator, fsPath string) (map[string]ROM, error) {
	var checksumPath string = filepath.Join(clientos.GetHomeDir(), ".retros")

	if fsPath != "" {
		checksumPath = fsPath
	}

	fsPath = filepath.Join(checksumPath, fmt.Sprintf("%s.yml", emulator))

	if !filesystem.CheckIfExists(fsPath) {
		err := download(fsPath, emulator)

		if err != nil {
			return nil, err
		}
	}

	f, err := os.ReadFile(fsPath)
	if err != nil {
		return nil, err
	}

	out := make(map[string]ROM)
	if err := yaml.Unmarshal(f, &out); err != nil {
		return nil, err
	}

	return out, err
}

func download(fsPath, emulator string) error {
	out, err := os.Create(fsPath)
	if err != nil {
		return err
	}

	defer out.Close()

	url := fmt.Sprintf("https://raw.githubusercontent.com/robwillup/retros/main/src/checksum/data/%s.yml", emulator)

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		return err
	}

	return nil
}
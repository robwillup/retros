package checksum

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ROMChecksum struct {
	Name  string `yaml:"name"`
	SHA256 string `yaml:"sha256"`
}

type Data struct {
	Roms map[string]string `yaml:"roms"`
}

func GetChecksums() (map[string]string, error) {
	f, err := os.ReadFile("src/checksum/snes.yml")
	if err != nil {
		return nil, err
	}

	out := Data{}
	if err := yaml.Unmarshal(f, &out); err != nil {
		return nil, err
	}

	fmt.Printf("HERE\n")

	return out.Roms, err
}
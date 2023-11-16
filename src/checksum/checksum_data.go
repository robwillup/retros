package checksum

import (
	"os"

	"gopkg.in/yaml.v3"
)

func GetChecksums() (map[string]ROM, error) {
	f, err := os.ReadFile("src/checksum/data/atari2600.yml")
	if err != nil {
		return nil, err
	}

	out := make(map[string]ROM)
	if err := yaml.Unmarshal(f, &out); err != nil {
		return nil, err
	}

	return out, err
}
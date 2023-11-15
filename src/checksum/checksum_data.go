package checksum

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Rom map[string]string

type Data struct {
	Roms []Rom `yaml:"roms"`
}

func GetChecksums() ([]Rom, error) {
	f, err := os.ReadFile("src/checksum/snes.yml")
	if err != nil {
		return nil, err
	}

	out := Data{}
	if err := yaml.Unmarshal(f, &out); err != nil {
		return nil, err
	}

	return out.Roms, err
}
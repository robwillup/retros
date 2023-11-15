package checksum

import (
	"encoding/json"
	"fmt"
	"os"
)

type ROM struct {
	Name   string `json:"_name"`
	SHA256 string `json:"_sha256"`
}

type Game struct {
	ROM interface{} `json:"rom"`
}

type Datafile struct {
	Game []Game `json:"game"`
}

type Checksums struct {
	Datafile Datafile `json:"datafile"`
}

func GetChecksumsFromJSON() ([]Game, error) {
	f, err := os.ReadFile("src/checksum/snes.json")
	if err != nil {
		return nil, err
	}

	var checksums Checksums

	if err := json.Unmarshal(f, &checksums); err != nil {
		fmt.Printf("%v\n", err)
		return nil, err
	}

	return checksums.Datafile.Game, nil
}

func WriteChecksumsToYaml(games []Game) error {
	f, err := os.OpenFile("src/checksum/snes.yml", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, game := range games {
		var yamlLine string
		switch rom := game.ROM.(type) {
		case map[string]interface{}:
			name := rom["_name"].(string)
			if sha256, ok := rom["_sha256"].(string); ok {
				yamlLine = fmt.Sprintf(" \"%s\": %s\n", name, sha256)
			} else {
				yamlLine = fmt.Sprintf(" \"%s\": %s\n", name, "No SHA256")
			}
		case []interface{}:
			for _, rom := range rom {
				romMap := rom.(map[string]interface{})
				name := romMap["_name"].(string)
				if sha256, ok := romMap["_sha256"].(string); ok {
					yamlLine = fmt.Sprintf(" \"%s\": %s\n", name, sha256)
				} else {
					yamlLine = fmt.Sprintf(" \"%s\": %s\n", name, "No SHA256")
				}
			}
		}

		if _, err := f.WriteString(yamlLine); err != nil {
			return err
		}
	}

	return nil
}
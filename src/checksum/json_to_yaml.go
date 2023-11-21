package checksum

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v3"
)

const JSONFile = "src/checksum/data/atari2600.json"
const YAMLFile = "src/checksum/data/atari2600.yml"

type ROM struct {
	Name   string `json:"_name"`
	MD5    string `json:"_md5"`
	SHA1   string `json:"_sha1"`
	SHA256 string `json:"_sha256"`
	Size   int    `json:"_size"`
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

func getChecksumsFromJSON() ([]Game, error) {
	f, err := os.ReadFile(JSONFile)

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

func WriteChecksumsToYaml() error {
	games, err := getChecksumsFromJSON()

	if err != nil {
		return err
	}

	f, err := os.OpenFile(YAMLFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()
	yamlLine := make(map[string]ROM)

	for _, game := range games {
		switch rom := game.ROM.(type) {
		case map[string]interface{}:
			md5Key := rom["_md5"].(string)

			size, err := strconv.Atoi(rom["_size"].(string))

			if err != nil {
				log.Printf("Failed to convert size from string to int. Error: %v\n", err)
			}

			romData := ROM{
				Name: rom["_name"].(string),
				MD5:  md5Key,
				Size: size,
			}

			if sha1, ok := rom["_sha1"].(string); ok {
				romData.SHA1 = sha1
			}

			if sha256, ok := rom["_sha256"].(string); ok {
				romData.SHA256 = sha256
			}

			yamlLine[md5Key] = romData

		case []interface{}:
			for _, rom := range rom {
				romMap := rom.(map[string]interface{})
				md5Key := romMap["_md5"].(string)
				size, err := strconv.Atoi(romMap["_size"].(string))

				if err != nil {
					log.Printf("Failed to convert size from string to int. Error: %v\n", err)
				}

				romData := ROM{
					Name: romMap["_name"].(string),
					MD5:  md5Key,
					Size: size,
				}

				if sha1, ok := romMap["_sha1"].(string); ok {
					romData.SHA1 = sha1
				}

				if sha256, ok := romMap["_sha256"].(string); ok {
					romData.SHA256 = sha256
				}

				yamlLine[md5Key] = romData
			}
		}
	}

	yaml, err := yaml.Marshal(yamlLine)

	if err != nil {
		return err
	}

	err = os.WriteFile(YAMLFile, yaml, fs.FileMode(os.O_CREATE|os.O_WRONLY))

	if err != nil {
		return err
	}

	return nil
}

/*
Copyright © 2023 Robson William

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/hex"
	"fmt"
	"log"
	"path/filepath"

	"github.com/robwillup/retros/src/checksum"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks ROM file integrity",
	Long: `Checks ROM file integrity by comparing its checksum with original's.
For example:

retros check GameFile.snes     Lists all ROM files
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking ROM file integrity")
		fmt.Println()

		res, err := checkROM(args[0])

		if err != nil {
			log.Fatalf("Failed to check ROM. Error: %v\n", err)
			return
		}

		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func checkROM(romPath string) (string, error) {
	// games, _ := checksum.GetChecksumsFromJSON()

	// err := checksum.WriteChecksumsToYaml(games)

	// if err != nil {
	// 	return "", err
	// }

	list, err := checksum.GetChecksums()

	if err != nil {
		return "", err
	}

	localChecksum, err := checksum.CalcSha256(romPath)

	if err != nil {
		return "", err
	}

	romName := filepath.Base(romPath)
	encoded := hex.EncodeToString(localChecksum)

	fmt.Printf("Local file name: %s\n", romName)
	fmt.Printf("Local file hash: %s\n", encoded)
	fmt.Printf("Original hash:   %s\n", list[romName])

	_, ok := list[romName]

	if ok && list[romName] == encoded {
		return "Good ROM file", nil
	}

	return "Bad ROM file", nil
}

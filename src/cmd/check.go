/*
Copyright © 2023 Robson William

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"errors"
	"fmt"
	"log"

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

		err := checkROM(args[0])

		if err != nil {
			log.Fatalf("Failed to check ROM. Error: %v\n", err)
			return
		}

		fmt.Println("ROM file is OK.")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func checkROM(romPath string) (error) {
	// err := checksum.WriteChecksumsToYaml()

	// if err != nil {
	// 	return err
	// }

	list, err := checksum.GetChecksums()

	if err != nil {
		return err
	}

	localROM, err := checksum.CalcChecksum(romPath)

	if err != nil {
		return err
	}

	fmt.Printf("Local file name: %s\n", localROM.Name)
	fmt.Printf("Local file MD5: %s\n", localROM.MD5)
	fmt.Printf("Local file SHA1: %s\n", localROM.SHA1)
	fmt.Printf("Local file SHA256: %s\n", localROM.SHA256)
	fmt.Printf("Local file Size: %d\n", localROM.Size)

	s, ok := list[localROM.MD5]

	if ok {
		fmt.Printf("Original file name: %s\n", s.Name)
		fmt.Printf("Original file MD5: %s\n", s.MD5)
		fmt.Printf("Original file SHA1: %s\n", s.SHA1)
		fmt.Printf("Original file SHA256: %s\n", s.SHA256)
		fmt.Printf("Original file Size: %d\n", s.Size)

		return nil
	}

	return errors.New(fmt.Sprintf("Invalid ROM hash for game %s", s.Name))
}

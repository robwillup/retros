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

// Package cmd handles CLI commands
package cmd

import (
	"errors"
	"fmt"
	"log"
	"path"
	"path/filepath"

	checksum2 "github.com/robwillup/retros/src/internal/checksum"
	"github.com/robwillup/retros/src/internal/emulators"
	"github.com/robwillup/retros/src/internal/filesystem"
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

		check(args[0])
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func check(fsPath string) {
	isDir, err := filesystem.IsDir(fsPath)

	if err != nil {
		log.Fatalf("Failed to check path for ROM verification.\nError: %v\n", err)
	}

	if isDir {
		files, err := filesystem.GetFiles(fsPath)

		if err != nil {
			log.Fatalf("Failed to get files for ROM verification.\nError: %v\n", err)
		}

		if len(files) < 1 {
			log.Println("No files in the path provided.")
			return
		}

		for _, file := range files {
			err := verifyFileIntegrity(path.Join(fsPath, file))

			if err != nil {
				fmt.Printf("Failed to verify ROM.\nError: %v\n", err)
			}
		}

		return
	}

	err = verifyFileIntegrity(fsPath)

	if err != nil {
		fmt.Printf("Failed to verify ROM.\nError: %v\n", err)
	}

	return
}

func verifyFileIntegrity(fsPath string) error {
	emulator := emulators.FindEmulatorFromExtension(fsPath)

	if emulator == "" {
		return errors.New(fmt.Sprintf("The file extension of '%s' is not yet supported.", filepath.Base(fsPath)))
	}

	originalChecksums, err := checksum2.GetChecksums(emulator, "")

	if err != nil {
		return err
	}

	localFile, err := checksum2.CalcChecksum(fsPath)

	if err != nil {
		return err
	}

	originalFile, ok := originalChecksums[localFile.MD5]

	if ok &&
		(originalFile.SHA1 == localFile.SHA1 || originalFile.SHA1 == "") &&
		(originalFile.SHA256 == localFile.SHA256 || originalFile.SHA256 == "") {

		fmt.Printf("Local file name:      %s\n", localFile.Name)
		fmt.Printf("Original file name:   %s\n", originalFile.Name)
		fmt.Printf("Local file MD5:       %s\n", localFile.MD5)
		fmt.Printf("Original file MD5:    %s\n", originalFile.MD5)
		fmt.Printf("Local file SHA1:      %s\n", localFile.SHA1)
		fmt.Printf("Original file SHA1:   %s\n", originalFile.SHA1)
		fmt.Printf("Local file SHA256:    %s\n", localFile.SHA256)
		fmt.Printf("Original file SHA256: %s\n", originalFile.SHA256)
		fmt.Printf("Local file Size:      %d\n", localFile.Size)
		fmt.Printf("Original file Size:   %d\n", originalFile.Size)
		fmt.Println()
		fmt.Println("ROM is authentic.")
		fmt.Println("--------------------------------------------")
		return nil
	}

	return errors.New(fmt.Sprintf("Invalid ROM: %s", originalFile.Name))
}

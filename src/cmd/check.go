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

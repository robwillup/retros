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
	"fmt"
	"log"
	"os"

	"github.com/robwillup/retros/src/config"
	"github.com/robwillup/retros/src/sshutils"
	"github.com/spf13/cobra"
)

const screenshotSourcePath = "/opt/retropie/configs/all/retroarch/screenshots"

// screenshotCmd represents the screenshot command
var screenshotCmd = &cobra.Command{
	Use:   "screenshot",
	Short: "Copies screenshots from the RetroPie setup.",
	Long: `Copies screenshots from the RetroPie setup.
If the destination already contains a file with the same name, that file can be skipped or overwritten.
It's possible to specify whether files in the origin should be deleted.
For example:

retros cp screenshot . --overwrite --clear

The above command will copy the screenshots from RetroPie to the current directory, overwrite duplicates and clear the origin.`,

	Run: func(cmd *cobra.Command, args []string) {
		dest, err := cmd.Flags().GetString("destination")

		if err != nil {
			log.Fatalln(err)
		}

		overwrite, err := cmd.Flags().GetBool("overwrite")

		if err != nil {
			log.Fatalln(err)
		}

		clear, err := cmd.Flags().GetBool("clear")

		if err != nil {
			log.Fatalln(err)
		}

		if dest == "" {
			dest, err = os.Getwd()
			if err != nil {
				log.Fatalln(err)
			}
		}

		copyScreenshots(dest, overwrite, clear)
	},
}

func init() {
	cpCmd.AddCommand(screenshotCmd)
	var destination string
	var overwrite   bool
	var clear       bool

	screenshotCmd.PersistentFlags().StringVarP(&destination, "destination", "d", "", "The destination where the screenshot files will be copied to. When empty, the current directory is used")
	screenshotCmd.PersistentFlags().BoolVarP(&overwrite, "overwrite", "o", false, "Whether same file should be overwritten. When false, the file is skipped")
	screenshotCmd.PersistentFlags().BoolVarP(&clear, "clear", "c", false, "Whether to delete files from source once they're copied. Default is false")
}

func copyScreenshots(dest string, overwrite, clear bool) error {
	fmt.Printf("Copy screenshots to %s\n", dest)
	fmt.Printf("Overwrite files: %t\n", overwrite)
	fmt.Printf("Clear files: %t\n", clear)

	config, err := config.Read()

	if err != nil {
		return err
	}

	client, err := sshutils.EstablishSSHConnection(config)

	if err != nil {
		return err
	}

	lsCmd := "ls " + screenshotSourcePath

	output, err := sshutils.ExecuteRemoteCommand(client, lsCmd)

	if err != nil {
		return err
	}

	println(output)

	return nil
}
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
	"path"
	"strings"

	"github.com/robwillup/retros/src/config"
	"github.com/robwillup/retros/src/sshutils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists ROM files",
	Long: `Lists ROM files in the remote machine where RetroPie is installed.
For example:

retros ls             Lists all ROM files
retros ls -p=snes     Lists all ROM files under snes/
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ROM files found: ")
		fmt.Println()

		emulator, err := cmd.Flags().GetString("emulator")

		if err != nil {
			log.Fatalf("Failed to get cp flags: Error %t\n", err)
		}

		output, err := listROMFiles(emulator)

		if err != nil {
			log.Fatalf("Failed to list ROM files. Error: %v\n", err)
			return
		}

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	var emulator string
	lsCmd.PersistentFlags().StringVarP(&emulator, "emulator", "e", "", "The emulator for which to list ROM files.")
}

func listROMFiles(emulator string) (string, error) {
	romsPath := "/home/pi/RetroPie/roms/"

	config, err := config.Read()

	if err != nil {
		return "", err
	}

	client, err := sshutils.EstablishSSHConnection(config)
	if err != nil {
		return "", err
	}

	if emulator != "" {
		output, err := runLs(path.Join(romsPath, emulator), client)

		if err != nil {
			return "", err
		}

		return output, nil
	}

	output, err := runLs(romsPath, client)

	if err != nil {
		return "", err
	}

	emulators := strings.Split(output, "\n")
	output = ""

	var sb strings.Builder

	for _, emu := range emulators {
		if emu != "" {
			output, err = runLs(path.Join(romsPath, emu), client)

			if err != nil {
				return "", err
			}

			if output != "" {
				sb.WriteString(fmt.Sprintf("%s\n\n", strings.ToUpper(emu)))
				sb.WriteString(fmt.Sprintf("%s\n", output))
				sb.WriteString("====================================================\n")
			}
		}
	}

	return sb.String(), nil
}

func runLs(dirPath string, client *ssh.Client) (string, error) {
	cmd := "ls " + dirPath
	output, err := sshutils.ExecuteRemoteCommand(client, cmd)

	if err != nil {
		log.Printf("Sorry. Failed to list ROM files under: %s\n\n", dirPath)
	}

	return output, nil
}

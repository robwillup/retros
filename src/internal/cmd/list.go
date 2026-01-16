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
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/robwillup/retros/src/internal/clientos"
	"github.com/robwillup/retros/src/internal/config"
	"github.com/robwillup/retros/src/internal/sshutils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

const LOCAL_MACHINE = ""

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls [OPTIONS]",
	Short: "Lists ROM files",
	Long: `Lists ROM files in the remote machine where RetroPie is installed.
For example:

retros ls                          Lists all ROM files
retros ls -e=snes                  Lists all ROM files under snes/
retros ls --emulator=mastersystem  Lists all ROM files under mastersystem/
`,
	Run: func(cmd *cobra.Command, args []string) {
		emulator, err := cmd.Flags().GetString("emulator")

		if err != nil {
			log.Fatalf("Failed to get cp flags: Error %t\n", err)
		}

		output, err := listROMFiles(emulator)

		if err != nil {
			log.Fatalf("Failed to list ROM files. Error: %v\n", err)
		}

		fmt.Printf("%s games found: \n", strings.ToUpper(emulator))
		fmt.Println()

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	var emulator string
	lsCmd.PersistentFlags().StringVarP(&emulator, "emulator", "e", "", "The emulator to list ROM files.")
}

func listROMFiles(emulator string) (string, error) {
	romsPath := "/home/pi/RetroPie/roms/"

	config, err := config.Read()

	if config.Host == LOCAL_MACHINE {
		romsPath = filepath.Join(clientos.GetHomeDir(), "RetroPie", "roms")
	}

	if err != nil {
		return "", err
	}

	var client *ssh.Client = nil

	if config.Host != LOCAL_MACHINE {
		client, err = sshutils.EstablishSSHConnection(config)
		if err != nil {
			return "", err
		}
	}

	if emulator != "" {
		output, err := runFind(path.Join(romsPath, emulator), client)

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
			output, err = runFind(path.Join(romsPath, emu), client)

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
	lsCmd := "ls " + dirPath

	// Target is a remote machine
	if client != nil {
		output, err := sshutils.ExecuteRemoteCommand(client, lsCmd)

		if err != nil {
			log.Printf("Failed to list ROM files under: %s\n\n", dirPath)
		}

		return output, nil
	}

	// Target is the local machine
	var subDirs []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && info.IsDir() {
			subDirs = append(subDirs, info.Name())
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	sort.Strings(subDirs)

	var sb strings.Builder

	for _, dir := range subDirs {
		_, _ = sb.WriteString(fmt.Sprintf("%s\n", dir))
	}

	return string(sb.String()), nil
}

func runFind(dirPath string, client *ssh.Client) (string, error) {
	findCmd := "find " + dirPath + " -type f ! -name '*.state*' ! -name '*.srm' ! -name '*.xml' ! -name '*.txt' -exec basename {} \\; | sort"

	// Target is a remote machine
	if client != nil {
		output, err := sshutils.ExecuteRemoteCommand(client, findCmd)

		if err != nil {
			log.Printf("Failed to list ROM files under: %s\n\n", dirPath)
		}

		return output, nil
	}

	// Target is the local machine
	var fileNames []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() && !strings.HasSuffix(info.Name(), ".state") && !strings.HasSuffix(info.Name(), ".srm") {
			fileNames = append(fileNames, info.Name())
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	sort.Strings(fileNames)

	var sb strings.Builder

	for _, fileName := range fileNames {
		_, _ = sb.WriteString(fmt.Sprintf("%s\n", fileName))
	}

	return string(sb.String()), nil
}

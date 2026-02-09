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
package cli

import (
	"fmt"
	"log"
	"path"
	"path/filepath"

	"github.com/robwillup/retros/internal/config"
	"github.com/robwillup/retros/internal/emulators"
	"github.com/robwillup/retros/internal/filesystem"
	"github.com/robwillup/retros/internal/sshutils"
	"github.com/spf13/cobra"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp [OPTIONS]",
	Short: "Copies a ROM file",
	Long: `Copies a ROM file from the current machine to a remote machine where RetroPie is running.
RetroS will copy the ROM file to the correct directory in $HOME/RetroPie/roms/.
For example:

retros cp Game.md

copies Game.md to $HOME/RetroPie/roms/genesis.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("Path to local ROM file is required.")
		}

		fmt.Println("Copying ROM files")

		emulator, err := cmd.Flags().GetString("emulator")

		if err != nil {
			log.Fatalln(err)
		}

		err = copy(args[0], emulator)

		if err != nil {
			log.Fatalf("Failed to copy ROM file. Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)
	var emulator string
	cpCmd.PersistentFlags().StringVarP(&emulator, "emulator", "e", "", "The emulator where the ROM file(s) will be copied to")
}

func copy(fsPath, emulator string) error {
	isDir, err := filesystem.IsDir(fsPath)

	if err != nil {
		return err
	}

	if isDir {
		files, err := filesystem.GetFiles(fsPath)

		if err != nil {
			return err
		}

		if len(files) > 0 {
			for _, file := range files {
				err := copyROMFile(path.Join(fsPath, file), emulator)

				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	err = copyROMFile(fsPath, emulator)

	if err != nil {
		return err
	}

	return nil
}

func copyROMFile(romFile, emulator string) error {
	romsPath := "/home/pi/RetroPie/roms/"

	if emulator == "" {
		emulator = emulators.FindEmulatorFromExtension(romFile)
	}

	romsPath = path.Join(romsPath, emulator, filepath.Base(romFile))

	config, err := config.Read()

	if err != nil {
		return err
	}

	client, err := sshutils.EstablishSSHConnection(config)

	if err != nil {
		return err
	}

	err = sshutils.CopyROMToRemote(client, romFile, romsPath)

	if err != nil {
		return err
	}

	return nil
}

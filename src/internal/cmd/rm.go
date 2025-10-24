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
	"log"
	"path"

	"github.com/robwillup/retros/src/internal/config"
	"github.com/robwillup/retros/src/internal/emulators"
	"github.com/robwillup/retros/src/internal/sshutils"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm [OPTIONS]",
	Short: "Removes a ROM file",
	Long: `Removes a ROM file from the RetroPie setup. For example:

retros rm Game.gba

Removes Game.gba from $HOME/RetroPie/roms/gba.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Path of the ROM file to be removed is required.")
		}

		emulator, err := cmd.Flags().GetString("emulator")

		if err != nil {
			log.Fatalln(err)
		}

		err = remove(args[0], emulator)

		if err != nil {
			log.Fatalf("Failed to delete ROM file. Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	var emulator string
	rmCmd.PersistentFlags().StringVarP(&emulator, "emulator", "e", "", "The emulator from where the ROM file will be removed.")
}

func remove(romFile, emulator string) error {
	romsPath := "/home/pi/RetroPie/roms/"

	if emulator == "" {
		emulator = emulators.FindEmulatorFromExtension(romFile)
	}

	romPath := path.Join(romsPath, emulator, romFile)

	config, err := config.Read()

	if err != nil {
		return err
	}

	client, err := sshutils.EstablishSSHConnection(config)

	if err != nil {
		return err
	}

	err = sshutils.DeleteROMFromRemote(client, romPath)

	if err != nil {
		return err
	}

	return nil
}

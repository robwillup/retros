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

		// Remove
	},
}

func init() {
	rootCmd.AddCommand(rmCmd)
	var emulator string
	rmCmd.PersistentFlags().StringVarP(&emulator, "emulator", "e", "", "The emulator from where the ROM file will be removed.")
}

func remove(fsPath, emulator string) error {
	// TODO: Add code to remove the file
}
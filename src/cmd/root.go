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
	"os"
	"path/filepath"

	"github.com/robwillup/retros/src/clientos"
	"github.com/robwillup/retros/src/config"
	"github.com/robwillup/retros/src/filesystem"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "retros",
	Short: "RetroS is a file manager for retro gamers",
	Long: `An easy to use tool that helps you maintain
your retro gaming setup organized and clean.`,

	Run: func(cmd *cobra.Command, args []string) {
		checkConfig()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func checkConfig() {
	configFile := filepath.Join(clientos.GetHomeDir(), ".retros", config.CONFIG_FILE_NAME)
	if !filesystem.CheckIfExists(configFile) {
		fmt.Println("RetroS must be configured. Run `retros cf` and follow the prompts")
		return
	}

	fmt.Println("RetroS: File manager for retro gamers")
}

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

// Package cmd handles CLI commands.
package cmd

import (
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/robwillup/retros/internal/clientos"
	"github.com/robwillup/retros/internal/config"
	"github.com/robwillup/retros/internal/sshutils"
	"github.com/spf13/cobra"
)

const SSH_PORT = 22

// cfCmd represents the cf command
var cfCmd = &cobra.Command{
	Use:   "cf",
	Short: "Used for configuring RetroS",
	Long: `This commands allows you to change different RetroS settings.
For example:

retros cf ssh         Configure SSH username, host, key path, etc.
retros cf retropie    Configure RetroPie path`,

	Run: func(cmd *cobra.Command, args []string) {
		home := clientos.GetHomeDir()
		conf := sshutils.SSHConfig{}
		fmt.Println("Host IP address (leave empty and press enter for local machine):")
		_, err := fmt.Scanln(&conf.Host)

		if conf.Host == "" {
			fmt.Println("Host is empty. RetroS will manage game files on this machine.")

			err = config.Create(conf)

			if err != nil {
				log.Fatal("Failed to create config file")
			}

			return
		}

		if err != nil {
			log.Fatal("Failed to read host")
		}

		fmt.Println("Username:")
		_, err = fmt.Scanln(&conf.Username)

		if err != nil {
			log.Fatal("Failed to read username")
		}

		fmt.Println("SSH key path [default '$HOME/.ssh/id_rsa']:")
		_, err = fmt.Scanln(&conf.KeyPath)

		if err != nil && !strings.Contains(err.Error(), "unexpected newline") {
			log.Fatal("Failed to read key path")
		}

		if conf.KeyPath == "" {
			conf.KeyPath = path.Join(home, "/.ssh/id_rsa")
		}

		conf.Port = SSH_PORT
		err = config.Create(conf)

		if err != nil {
			log.Fatal("Failed to create config file")
		}
	},
}

func init() {
	rootCmd.AddCommand(cfCmd)
}

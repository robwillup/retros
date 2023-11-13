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
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/robwillup/retros/src/clientos"
	"github.com/robwillup/retros/src/config"
	"github.com/robwillup/retros/src/sshutils"
	"github.com/spf13/cobra"
)

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

		fmt.Println("Host IP address:")
		_, err := fmt.Scanln(&conf.Host)
		if err != nil {
			log.Fatal("Failed to read host")
		}

		fmt.Println("Username:")
		_, err = fmt.Scanln(&conf.Username)
		if err != nil {
			log.Fatal("Failed to read username")
		}

		fmt.Println("SSH key path [default '$HOME/.ssh/id_rsa']:")
		_, err = fmt.Scanln(&conf.KeyPath,)
		if err != nil && !strings.Contains(err.Error(), "unexpected newline") {
			log.Fatal("Failed to read key path")
		}

		if conf.KeyPath == "" {
			conf.KeyPath = path.Join(home, "/.ssh/id_rsa")
		}

		_, err = config.Create(conf)

		if err != nil {
			log.Fatal("Failed to create config file")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(cfCmd)
}

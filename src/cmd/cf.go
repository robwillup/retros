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
	"os"

	"github.com/robwillup/rosy/src/config"
	"github.com/robwillup/rosy/src/sshutils"
	"github.com/spf13/cobra"
)

// cfCmd represents the cf command
var cfCmd = &cobra.Command{
	Use:   "cf",
	Short: "Used for configuring RoSy",
	Long: `This commands allows you to change different RoSy settings.
For example:

rosy cf ssh         Configure SSH username, host, key path, etc.
rosy cf retropie    Configure RetroPie path`,
	Run: func(cmd *cobra.Command, args []string) {
		home := os.Getenv("HOME")
		conf := sshutils.SSHConfig{
			KeyPath: home + "/.ssh/id_rsa",
		}

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

		if config.Create(conf) != nil {
			log.Fatal("Failed to create config file")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(cfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cfCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cfCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

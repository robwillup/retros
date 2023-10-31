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

	"github.com/robwillup/rosy/src/config"
	"github.com/robwillup/rosy/src/sshutils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "Lists ROM files",
	Long: `Lists ROM files in the remote machine where RetroPie is installed.
For example:

rosy ls				Lists all ROM files
rosy ls -p snes		Lists all ROM files under snes/
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ROM files found: ")
		fmt.Println()

		platform, _ := cmd.Flags().GetString("platform")

		config, err := config.Read()

		if err != nil {
			log.Fatal("Failed to read config file")
			return
		}

		client, err := sshutils.EstablishSSHConnection(config)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer client.Close()

		output, err := listROMFiles(client, platform)

		if err != nil {
			log.Fatalln(err)
			return
		}

		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(lsCmd)
	var platform string
	lsCmd.PersistentFlags().StringVarP(&platform, "platform", "p", "", "The platform for which to list ROM files.")
}

func listROMFiles(client *ssh.Client, platform string) (string, error) {
	romsPath := "/home/pi/RetroPie/roms/"

	if platform != "" {
		romsPath = path.Join(romsPath, platform)
	}

	remoteCommand := "ls " + romsPath

	output, err := sshutils.ExecuteRemoteCommand(client, remoteCommand)
	if err != nil {
		return "", err
	}

	return output, nil
}

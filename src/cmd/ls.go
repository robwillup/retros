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

rosy ls             Lists all ROM files
rosy ls -p=snes     Lists all ROM files under snes/
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ROM files found: ")
		fmt.Println()

		platform, _ := cmd.Flags().GetString("platform")

		output, err := listROMFiles(platform)

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

func listROMFiles(platform string) (string, error) {
	romsPath := "/home/pi/RetroPie/roms/"

	config, err := config.Read()

	if err != nil {
		return "", err
	}

	client, err := sshutils.EstablishSSHConnection(config)
	if err != nil {
		return "", err
	}

	if platform != "" {
		output, err := runLs(path.Join(romsPath, platform), client)

		if err != nil {
			return "", err
		}

		return output, nil
	}

	output, err := runLs(romsPath, client)

	if err != nil {
		return "", err
	}

	platforms := strings.Split(output, "\n")
	output = ""

	var sb strings.Builder

	for _, plat := range platforms {
		if plat != "" {
			output, err = runLs(path.Join(romsPath, plat), client)

			if err != nil {
				return "", err
			}

			if output != "" {
				sb.WriteString(fmt.Sprintf("%s\n\n", strings.ToUpper(plat)))
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
		return "", err
	}

	return output, nil
}

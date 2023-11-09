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
	"path/filepath"

	"github.com/robwillup/rosy/src/config"
	"github.com/robwillup/rosy/src/filesystem"
	"github.com/robwillup/rosy/src/emulators"
	"github.com/robwillup/rosy/src/sshutils"
	"github.com/spf13/cobra"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp",
	Short: "Copies a ROM file",
	Long: `Copies a ROM file from the current machine to a remote machine where RetroPie is running.
RoSy will copy the ROM file to the correct directory in $HOME/RetroPie/roms/.
For example:

rosy cp Game.md

copies Game.md to $HOME/RetroPie/roms/genesis.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln("Path to local ROM file is required.")
		}

		fmt.Println("Copying ROM files")
		fmt.Println()

		platform, err := cmd.Flags().GetString("platform")

		if err != nil {
			log.Fatalln(err)
		}

		err = copy(args[0], platform)

		if err != nil {
			log.Fatalf("Failed to copy ROM file. Error: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)
	var platform string
	cpCmd.PersistentFlags().StringVarP(&platform, "platform", "p", "", "The platform where the ROM file(s) will be copied to")
}

func copy(fsPath, platform string) error {
	isDir, err := filesystem.CheckDir(fsPath)

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
				err := copyROMFile(path.Join(fsPath, file), platform)

				if err != nil {
					return err
				}
			}
		}

		return nil
	}

	err = copyROMFile(fsPath, platform)

	if err != nil {
		return err
	}

	return nil
}

func copyROMFile(romFile, plat string) error {
	romsPath := "/home/pi/RetroPie/roms/"

	if plat == "" {
		plat = emulators.FindEmulatorFromExtension(romFile)
	}

	romsPath = path.Join(romsPath, plat, filepath.Base(romFile))

	config, err := config.Read()

	if err != nil {
		log.Fatal("Failed to read config file")
		return nil
	}

	client, err := sshutils.EstablishSSHConnection(config)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	err = sshutils.CopyROMToRemote(client, romFile, romsPath)

	if err != nil {
		return err
	}

	return nil
}

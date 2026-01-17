package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/robwillup/retros/src/internal/clientos"
	"github.com/robwillup/retros/src/internal/config"
	"github.com/robwillup/retros/src/internal/sshutils"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh"
)

var emulatorCmd = &cobra.Command{
	Use:   "emulator [OPTIONS]",
	Short: "List emulator systems",
	Long: `List emulator systems in the RetroPie setup.
For example:

retros ls emulator        Lists emulator systems with at least one ROM file.
retros ls emulator --all  Lists all emulator systems.`,

	Run: func(cmd *cobra.Command, args []string) {
		all, err := cmd.Flags().GetBool("all")

		if err != nil {
			log.Fatalf("Failed to get ls emulator flags: Error %t\n", err)
		}

		output, err := listEmulators(all)

		if err != nil {
			log.Fatalf("Failed to list emulators: Error %t\n", err)
		}

		fmt.Println("Emulator systems found:")
		fmt.Println()
		fmt.Println(output)
	},
}

func init() {
	lsCmd.AddCommand(emulatorCmd)
	var all bool
	emulatorCmd.Flags().BoolVarP(&all, "all", "a", false, "Include emulator systems without any ROM files.")
}

func listEmulators(all bool) (string, error) {
	emulatorsPath := "/home/pi/RetroPie/roms/"

	config, err := config.Read()

	if config.Host == LOCAL_MACHINE {
		emulatorsPath = filepath.Join(clientos.GetHomeDir(), "RetroPie", "roms")
	}

	if err != nil {
		return "", err
	}

	var client *ssh.Client = nil

	if config.Host != LOCAL_MACHINE {
		client, err = sshutils.EstablishSSHConnection(config)
		if err != nil {
			return "", err
		}
	}

	if all {
		return "TODO", err
	}

	output, err := runLs(emulatorsPath, true, client)

	return output, err
}

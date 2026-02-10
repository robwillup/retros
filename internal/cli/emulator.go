package cli

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/robwillup/retros/internal/clientos"
	"github.com/robwillup/retros/internal/config"
	emulators2 "github.com/robwillup/retros/internal/emulators"
	"github.com/robwillup/retros/internal/sshutils"
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
			log.Fatalf("Failed to get ls emulator flags: Error %t", err)
		}

		fmt.Println("Looking for emulator systems on the target machine ...")

		output, err := listEmulators(all)

		if err != nil {
			log.Fatalf("Failed to list emulators: Error %t\n", err)
		}

		fmt.Println("Emulator systems found:")
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

	if config.Host == LocalMachine {
		emulatorsPath = filepath.Join(clientos.GetHomeDir(), "RetroPie", "roms")
	}

	if err != nil {
		return "", err
	}

	var client *ssh.Client = nil

	if config.Host != LocalMachine {
		client, err = sshutils.EstablishSSHConnection(config)
		if err != nil {
			return "", err
		}
	}

	output, err := runLs(emulatorsPath, true, client)

	if err != nil {
		return "", err
	}

	emulators := strings.Split(output, "\n")

	if all {
		var emulatorsDisplayNames string
		for _, e := range emulators {
			emulatorsDisplayNames = emulatorsDisplayNames + emulators2.GetEmulatorDisplayName(e) + "\n"
		}
		return emulatorsDisplayNames, err
	}

	var emulatorsInUse string

	for _, e := range emulators {
		files, err := runFind(filepath.Join(emulatorsPath, e), client)
		if err != nil {
			return "", err
		}

		if files != "" {
			emulatorsInUse = emulatorsInUse + emulators2.GetEmulatorDisplayName(e) + "\n"
		}
	}

	return emulatorsInUse, err
}

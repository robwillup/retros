/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// screenshotCmd represents the screenshot command
var screenshotCmd = &cobra.Command{
	Use:   "screenshot",
	Short: "Copies screenshots from the RetroPie setup.",
	Long: `Copies screenshots from the RetroPie setup.
If the destination already contains a file with the same name, that file can be skipped or overwritten.
It's possible to specify whether files in the origin should be deleted.
For example:

retros cp screenshot . --overwrite --clear

The above command will copy the screenshots from RetroPie to the current directory, overwrite duplicates and clear the origin.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("screenshot called")
	},
}

func init() {
	cpCmd.AddCommand(screenshotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// screenshotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// screenshotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

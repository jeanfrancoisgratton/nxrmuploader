// nxrmuploader
// src/cmd/root.go

package cmd

import (
	"github.com/spf13/cobra"
	"nxrmuploader/helpers"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "uploadNxRM",
	Short:   "Add a short description here",
	Long:    "Add a long description here",
	Version: "1.00.00-0 (2023.12.31)",
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.ChangeLog()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableAutoGenTag = true
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(clCmd)
}

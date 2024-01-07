// nxrmuploader
// src/cmd/root.go

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"nxrmuploader/env"
	"nxrmuploader/exec"
	"nxrmuploader/helpers"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "uploadNxRM",
	Short:   "Upload a binary package to a NxRM repository",
	Long:    "The target repository will be chosen according to the extension of the file to be uploaded.",
	Version: "1.01.00-0 (2024.01.04)",
}

var clCmd = &cobra.Command{
	Use:     "changelog",
	Aliases: []string{"cl"},
	Short:   "Shows changelog",
	Run: func(cmd *cobra.Command, args []string) {
		helpers.ChangeLog()
	},
}

var upCmd = &cobra.Command{
	Use:     "upload",
	Aliases: []string{"up", "push"},
	Short:   "Uploads a binary (RPM, DEB) package to its appropriate repository",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("You need to provide at least one filename")
			os.Exit(0)
		}
		var debianPkgs, rhPkgs, apkPkgs []string
		for _, argument := range args {
			if strings.HasSuffix(strings.ToLower(argument), ".deb") {
				debianPkgs = append(debianPkgs, strings.ToLower(argument))
			}
			if strings.HasSuffix(strings.ToLower(argument), ".rpm") {
				rhPkgs = append(rhPkgs, strings.ToLower(argument))
			}
			if strings.HasSuffix(strings.ToLower(argument), ".apk") {
				apkPkgs = append(apkPkgs, strings.ToLower(argument))
			}
		}
		if len(debianPkgs) > 0 {
			if err := exec.Upload(debianPkgs); err != nil {
				fmt.Println(err)
			}
		}
		if len(rhPkgs) > 0 {
			if err := exec.Upload(rhPkgs); err != nil {
				fmt.Println(err)
			}
		}
		if len(apkPkgs) > 0 {
			if err := exec.Upload(apkPkgs); err != nil {
				fmt.Println(err)
			}
		}
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
	rootCmd.AddCommand(clCmd, upCmd)

	rootCmd.PersistentFlags().StringVarP(&env.EnvConfigFile, "env", "e", "defaultEnv.json", "Default environment configuration file; this is a per-user setting.")
	upCmd.Flags().Int8VarP(&exec.IndexNumber, "index", "i", 0, "Index of repository; this is zero-based.")
}

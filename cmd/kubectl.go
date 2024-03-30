/*
Copyright © 2024 Nik Ogura <nik.ogura@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/nikogura/nikstools/pkg/nikstools"
	"github.com/spf13/cobra"
	"log"
)

// kubectlCmd represents the kubectl command
var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Add kubectl functions to your shell profile.",
	Long: `
Add kubectl functions to your shell profile.
`,
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := homedir.Dir()
		if err != nil {
			log.Fatalf("Failed to find homedir: %s", err)
		}
		profileName, err := nikstools.SetProfile(homeDir)
		if err != nil {
			log.Fatalf("err setting profile: %s", err)
		}

		if profileName == "" {
			fmt.Printf("No bash profile available, source ~/%s in your profile of choice to use kubernetes related utility functions", nikstools.DEFAULT_PROFILE_DIR)
		} else {
			fmt.Printf("%s updated successfully to include profile\n", profileName)
		}

	},
}

func init() {
	rootCmd.AddCommand(kubectlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kubectlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kubectlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

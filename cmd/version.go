package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/constants"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "See your Runtipi CLI version",
	Long: "This command prints the current Runtipi CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Current Runtipi CLI version: ")
		color.Set(color.FgBlue)
		fmt.Print(constants.Version)
		color.Unset()
		fmt.Println()
	},
}
package cmd

import (
	"fmt"

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
		fmt.Printf("Current Runtipi version: %s\n", constants.Blue(constants.RuntipiVersion))
		fmt.Printf("Current CLI version: %s\n", constants.Blue(constants.CliVersion))
	},
}
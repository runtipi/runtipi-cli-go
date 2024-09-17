package cmd

import (
	"runtipi-cli-go/internal/commands"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use: "stop",
	Short: "Stop Runtipi",
	Long: "Use this command to stop the Runtipi docker stack",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Stop()
	},
}
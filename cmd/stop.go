package cmd

import (
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/commands"
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
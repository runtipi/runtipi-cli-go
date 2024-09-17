package cmd

import (
	"runtipi-cli-go/internal/commands"

	"github.com/spf13/cobra"
)

func init() {
	startCmd.Flags().BoolVar(&noPermissions, "no-permissions", false, "Skip setting permissions.")
	startCmd.Flags().StringVar(&envFile, "env-file", "", "Path to custom .env file, it has to be absolute")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use: "start",
	Short: "Start Runtipi",
	Long: "Use this command to start the Runtipi docker stack",
	Run: func(cmd *cobra.Command, args []string) {
		commands.Start(envFile, noPermissions)
	},
}

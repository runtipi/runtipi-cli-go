package app

import (
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/cmd/app/subcommands"
)

func AppCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "app",
		Short: "App commands",
		Long: "Control your Runtipi apps through the CLI",
	}
	cmd.AddCommand(subcommands.StartAppCmd)
	cmd.AddCommand(subcommands.StopAppCmd)
	cmd.AddCommand(subcommands.RestartAppCmd)
	cmd.AddCommand(subcommands.ResetAppCmd)
	cmd.AddCommand(subcommands.UpdateAppCmd)
	cmd.AddCommand(subcommands.UninstallAppCmd)
	cmd.AddCommand(subcommands.StartAllCmd)
	return cmd
}
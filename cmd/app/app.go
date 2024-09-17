package app

import (
	"runtipi-cli-go/cmd/app/subcommands"

	"github.com/spf13/cobra"
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
	cmd.AddCommand(subcommands.ListAppBackupsCmd)
	return cmd
}
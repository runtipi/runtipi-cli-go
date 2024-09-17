package system

import (
	"runtipi-cli-go/cmd/system/subcommands"

	"github.com/spf13/cobra"
)

func SystemCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "system",
		Short: "System commands",
		Long: "Control your Runtipi system through the CLI",
	}
	cmd.AddCommand(subcommands.HealthCheckCmd)
	cmd.AddCommand(subcommands.StatusCmd)
	return cmd
}
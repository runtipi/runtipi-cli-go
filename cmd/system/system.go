package system

import (
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/cmd/system/subcommands"
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
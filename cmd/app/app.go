package app

import (
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/cmd/app/start"
)

func AppCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "app",
		Short: "App commands",
		Long: "Control your Runtipi apps through the CLI",
	}
	cmd.AddCommand(start.StartAppCmd)
	return cmd
}
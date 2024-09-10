package subcommands

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/api"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
)

var StartAllCmd = &cobra.Command{
	Use: "start-all",
	Short: "Starts all apps using the Runtipi API",
	Long: "This command starts all apps using the Runtipi worker API",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Define Path
		path := "apps/start-all"

		// Start Spinner
		spinner.SetMessage("Starting apps")
		spinner.Start()

		// Start apps
		err := api.ApiRequest(path, "POST", 15 * time.Minute)

		if err != nil {
			spinner.Fail("Failed to start apps")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// Succeed
		spinner.Succeed("Apps succeessfully")
		spinner.Stop()
	},
}
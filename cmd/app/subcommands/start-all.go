package subcommands

import (
	"fmt"
	"os"
	"time"

	"runtipi-cli-go/internal/api"
	"runtipi-cli-go/internal/spinner"

	"github.com/spf13/cobra"
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
		response, err := api.ApiRequest(path, "POST", 15 * time.Minute)

		if err != nil {
			spinner.Fail("Failed to start apps")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		defer response.Body.Close()

		// Succeed
		spinner.Succeed("Apps succeessfully")
		spinner.Stop()
	},
}
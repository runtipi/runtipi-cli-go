package subcommands

import (
	"fmt"
	"os"
	"time"

	"runtipi-cli-go/internal/api"
	"runtipi-cli-go/internal/spinner"

	"github.com/spf13/cobra"
)

var	UpdateAppCmd = &cobra.Command{
	Use: "update [app]",
	Short: "Update an app using the Runtipi API",
	Long: "This command updates the specified app using the Runtipi worker API",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Define Path
		path := fmt.Sprintf("apps/%s/update", args[0])

		// Start Spinner
		spinner.SetMessage("Updating app")
		spinner.Start()

		// Updating app
		response, err := api.ApiRequest(path, "POST", 15 * time.Minute)

		if err != nil {
			spinner.Fail("Failed to update app")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		defer response.Body.Close()
		
		// Succeed
		spinner.Succeed("App updated succeessfully")
		spinner.Stop()
	},
}
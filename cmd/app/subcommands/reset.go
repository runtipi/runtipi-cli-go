package subcommands

import (
	"fmt"
	"os"
	"time"

	"runtipi-cli-go/internal/api"
	"runtipi-cli-go/internal/spinner"

	"github.com/spf13/cobra"
)

var	ResetAppCmd = &cobra.Command{
	Use: "reset [app]",
	Short: "Resets an app using the Runtipi API",
	Long: "This command resets the specified app using the Runtipi worker API",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Define Path
		path := fmt.Sprintf("apps/%s/reset", args[0])

		// Start Spinner
		spinner.SetMessage("Resetting app")
		spinner.Start()

		// Reset app
		response, err := api.ApiRequest(path, "POST", 5 * time.Minute)

		if err != nil {
			spinner.Fail("Failed to reset app")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		defer response.Body.Close()

		// Succeed
		spinner.Succeed("App reset succeessfully")
		spinner.Stop()
	},
}
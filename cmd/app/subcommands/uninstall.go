package subcommands

import (
	"fmt"
	"os"
	"time"

	"runtipi-cli-go/internal/api"
	"runtipi-cli-go/internal/spinner"

	"github.com/spf13/cobra"
)

var	UninstallAppCmd = &cobra.Command{
	Use: "uninstall [app]",
	Short: "Uninstalls an app using the Runtipi API",
	Long: "This command uninstalls the specified app using the Runtipi worker API",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Define Path
		path := fmt.Sprintf("apps/%s/uninstall", args[0])

		// Start Spinner
		spinner.SetMessage("Uninstalling app")
		spinner.Start()

		// Uninstall app
		response, err := api.ApiRequest(path, "POST", 5 * time.Minute)

		if err != nil {
			spinner.Fail("Failed to uninstall app")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		defer response.Body.Close()

		// Succeed
		spinner.Succeed("App uninstalled succeessfully")
		spinner.Stop()
	},
}
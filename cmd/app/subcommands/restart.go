package subcommands

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/api"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
)

var	RestartAppCmd = &cobra.Command{
	Use: "restart [app]",
	Short: "Restart an app using the Runtipi API",
	Long: "This command restarts the specified app using the Runtipi worker API",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Define Path
		path := fmt.Sprintf("apps/%s/restart", args[0])

		// Start Spinner
		spinner.SetMessage("Restarting app")
		spinner.Start()

		// Restart app
		err := api.ApiRequest(path, "POST", 5 * time.Minute)

		if err != nil {
			spinner.Fail("Failed to restart app")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// Succeed
		spinner.Succeed("App restarted succeessfully")
		spinner.Stop()
	},
}
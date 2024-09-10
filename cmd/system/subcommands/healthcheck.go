package subcommands

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/api"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
)

var	HealthCheckCmd = &cobra.Command{
	Use: "healthcheck",
	Short: "Checks if the Runtipi system is up and running use the API",
	Long: "Checks if the Runtipi system is up and running using the worker API",
	Run: func(cmd *cobra.Command, args []string) {
		// Define Path
		path := "healthcheck"

		// Start Spinner
		spinner.SetMessage("Checking system")
		spinner.Start()

		// Do Check
		response, err := api.ApiRequest(path, "GET", 1 * time.Minute)

		if err != nil {
			spinner.Fail("Failed to check system, is runtipi running?")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		defer response.Body.Close()

		// Succeed
		spinner.Succeed("Runtipi system up and running")
		spinner.Stop()
	},
}
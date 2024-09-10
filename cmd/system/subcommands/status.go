package subcommands

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/api"
	"github.com/steveiliop56/runtipi-cli-go/internal/constants"
	"github.com/steveiliop56/runtipi-cli-go/internal/schemas"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
)

var	StatusCmd = &cobra.Command{
	Use: "status",
	Short: "Gets system readings from the Runtipi API",
	Long: "Shows the system readings (e.g. cpu usage, disk usage) from the Runtipi worker API",
	Run: func(cmd *cobra.Command, args []string) {
		// Define Path
		path := "system-status"

		// Start Spinner
		spinner.SetMessage("Getting system status")
		spinner.Start()

		// Do Check
		response, err := api.ApiRequest(path, "GET", 1 * time.Minute)

		if err != nil {
			spinner.Fail("Failed to get system status, is runtipi running?")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// Parse Json

		status := new(schemas.SystemStatusApi)

		jsonErr := json.NewDecoder(response.Body).Decode(&status)

		if jsonErr != nil {
			spinner.Fail("Failed to decode system status json")
			spinner.Stop()
			fmt.Printf("Error: %s\n", jsonErr)
			os.Exit(1)
		}

		defer response.Body.Close()

		// Succeed
		spinner.Succeed("Successfully got system status")
		spinner.Stop()

		// Print status
		fmt.Printf("Your CPU usage is %s\n", constants.Blue(fmt.Sprintf("%.2f%%", status.Data.CpuLoad)))
		fmt.Printf("Your Disk size is %s, you are using %s which is %s\n", constants.Blue(fmt.Sprintf("%dGB", status.Data.DiskSize)), constants.Blue(fmt.Sprintf("%dGB", status.Data.DiskUsed)), constants.Blue(fmt.Sprintf("%0.f%%", status.Data.PercentUsed)))
		fmt.Printf("Your Memory size is %s and you are using %s\n", constants.Blue(fmt.Sprintf("%dGB", status.Data.MemoryTotal)), constants.Blue(fmt.Sprintf("%0.f%%", status.Data.PercentUsedMemory)))
	},
}
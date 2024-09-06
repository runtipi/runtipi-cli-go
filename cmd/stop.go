package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use: "stop",
	Short: "Stop Runtipi",
	Long: "Use this command to stop the Runtipi docker stack",
	Run: func(cmd *cobra.Command, args []string) {
		// Stop containers
		spinner.SetMessage("Stopping containers")
		spinner.Start()

		_, err := exec.Command("docker", "compose", "down", "--remove-orphans", "--rmi", "local").Output()

		if err != nil {
			spinner.Fail("Error in stopping containers")
			spinner.Stop()
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		containersToRm := []string{"runtipi-reverse-proxy", "runtipi-db", "runtipi-redis", "runtipi", "tipi-db", "tipi-redis", "tipi-reverse-proxy", "tipi-docker-proxy", "tipi-dashboard", "tipi-worker"}

		for _, container := range containersToRm {
			exec.Command("docker", "stop", container).Output()
			exec.Command("docker", "rm", container).Output()
		}

		spinner.Succeed("Runtipi stopped successfully")
		spinner.Stop()
	},
}
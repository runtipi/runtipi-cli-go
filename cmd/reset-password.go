package cmd

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/env"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
)

func init() {
	rootCmd.AddCommand(resetPasswordCmd)
}

var resetPasswordCmd = &cobra.Command{
	Use: "reset-password",
	Short: "Reset Runtipi's Password",
	Long: "Use this command to reset your Runtipi's password if you forget it",
	Run: func(cmd *cobra.Command, args []string) {
		spinner.SetMessage("Creating reset password request")
		spinner.Start()

		rootFolder, osErr := os.Getwd()
	
		if osErr != nil {
			spinner.Fail("Failed to get root folder")
			spinner.Stop()
			fmt.Printf("Error: %s\n", osErr)
			os.Exit(1)
		}

		time := time.Now().Unix()
		writeErr := os.WriteFile(path.Join(rootFolder, "state", "password-change-request"), []byte(strconv.Itoa(int(time))), 0644)

		if writeErr != nil {
			spinner.Fail("Failed to create password change request")
			spinner.Stop()
			fmt.Printf("Error: %s\n", writeErr)
			os.Exit(1)
		}

		internalIp, _ := env.GetEnvValue("INTERNAL_IP")
		nginxPort, _ := env.GetEnvValue("NGINX_PORT")

		message := fmt.Sprintf("Password request created. Head back to http://%s:%s/reset-password to reset your password.", internalIp, nginxPort)

		spinner.Succeed(message)
		spinner.Stop()
	},
}
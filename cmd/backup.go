package cmd

import (
	"fmt"
	"os"

	"runtipi-cli-go/internal/backups"
	"runtipi-cli-go/internal/commands"
	"runtipi-cli-go/internal/constants"
	"runtipi-cli-go/internal/spinner"
	"runtipi-cli-go/internal/system"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(backupCmd)
}

var backupCmd = &cobra.Command{
	Use: "backup",
	Short: "Backup Runtipi",
	Long: "Use this command to backup your Runtipi instance",
	Run: func(cmd *cobra.Command, args []string) {
		// Stop runtipi
		commands.Stop()

		// Start Spinner
		spinner.SetMessage("Backing up...")
		spinner.Start()

		// Ensure tar
		tarErr := system.EnsureTar()

		if tarErr != nil {
			spinner.Fail("Tar is not installed")
			spinner.Stop()
			fmt.Printf("Error: %s\n", tarErr)
			os.Exit(1)
		}
		
		// Backup
		archivePath, backupErr := backups.CreateBackup()

		if backupErr != nil {
			spinner.Fail("Failed to backup")
			spinner.Stop()
			fmt.Printf("Error: %s\n", backupErr)
			os.Exit(1)
		}

		spinner.Succeed("Backed up successfully")
		spinner.Stop()

		// Print archive name
		fmt.Printf("%s Your backup is located at %s\n", constants.Blue("🛈"), archivePath)
		fmt.Printf("%s Check https://runtipi.io/docs/guides/backup-and-restore on how to restore\n", constants.Blue("🛈"))
	},
}
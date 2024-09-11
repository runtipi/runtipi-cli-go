package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/backups"
	"github.com/steveiliop56/runtipi-cli-go/internal/commands"
	"github.com/steveiliop56/runtipi-cli-go/internal/constants"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
	"github.com/steveiliop56/runtipi-cli-go/internal/system"
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
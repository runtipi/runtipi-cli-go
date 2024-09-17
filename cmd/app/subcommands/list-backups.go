package subcommands

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"runtipi-cli-go/internal/constants"
	"runtipi-cli-go/internal/utils"

	"github.com/aquasecurity/table"
	"github.com/spf13/cobra"
)

var	ListAppBackupsCmd = &cobra.Command{
	Use: "list-backups [app]",
	Short: "List app backups",
	Long: "Lists all the backups of an app",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Get root folder
		rootFolder, osErr := os.Getwd()
	
		if osErr != nil {
			fmt.Printf("%s Failed to get root folder\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", osErr)
			os.Exit(1)
		}

		// Define paths
		backupPath := path.Join(rootFolder, "backups", args[0])

		// Check if folder exists
		_, pathCheckErr := os.Stat(backupPath)
		if pathCheckErr != nil {
			fmt.Printf("%s App backup path doesn't exist\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", pathCheckErr)
			os.Exit(1)
		}

		// Read directory
		backups, readErr := os.ReadDir(backupPath)
		if readErr != nil {
			fmt.Printf("%s Failed to read app backups folder\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", readErr)
			os.Exit(1)
		}

		// List backups
		backupsTable := table.New(os.Stdout)
		backupsTable.SetHeaders("Name", "Size", "Date Created")
		for _, backup := range backups {
			backupInfo, _ := backup.Info()
			backupSize := utils.FormatFileSize(float64(backupInfo.Size()))
			backupDateIso, _ := strconv.ParseInt(strings.Replace(strings.Replace(strings.Replace(backup.Name(), args[0], "", 1), ".tar.gz", "", 1), "-", "", 1), 10, 64)
			backupDate := time.UnixMilli(backupDateIso).String()
			backupsTable.AddRow(backup.Name(), backupSize, backupDate)
		}

		backupsTable.Render()
	},
}
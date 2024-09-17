package cmd

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"runtipi-cli-go/internal/constants"
	"runtipi-cli-go/internal/utils"

	"github.com/aquasecurity/table"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listBackupsCmd)
}

var listBackupsCmd = &cobra.Command{
	Use: "list-backups",
	Short: "List Runtipi backups",
	Long: "Use this command to list all the available Runtipi backups",
	Run: func(cmd *cobra.Command, args []string) {
		// Get root folder
		rootFolder, osErr := os.Getwd()
	
		if osErr != nil {
			fmt.Printf("%s Failed to get root folder\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", osErr)
			os.Exit(1)
		}

		// Define paths
		backupPath := path.Join(rootFolder, "backups")

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
		backupNum := 0
		for _, backup := range backups {
			if !backup.IsDir() {
				backupInfo, _ := backup.Info()
				backupSize := utils.FormatFileSize(float64(backupInfo.Size()))
				backupDateIso, _ := strconv.ParseInt(strings.Replace(strings.Replace(strings.Replace(backup.Name(), filepath.Base(rootFolder), "", 1), ".tar.gz", "", 1), "-", "", 1), 10, 64)
				backupDate := time.UnixMilli(backupDateIso).String()
				backupsTable.AddRow(backup.Name(), backupSize, backupDate)
				backupNum += 1
			}
		}

		if backupNum == 0 {
			fmt.Printf("%s No backups found\n", constants.Red("✗"))
			os.Exit(0)
		}

		backupsTable.Render()
	},
}
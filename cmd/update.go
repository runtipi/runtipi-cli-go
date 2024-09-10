package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/Delta456/box-cli-maker"
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/constants"
	"github.com/steveiliop56/runtipi-cli-go/internal/env"
	"github.com/steveiliop56/runtipi-cli-go/internal/release"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
)

func init() {
	updateCmd.Flags().BoolVar(&noPermissions, "no-permissions", false, "Skip setting permissions.")
	updateCmd.Flags().StringVar(&envFile, "env-file", "", "Path to custom .env file")
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use: "update [version]",
	Short: "Update to the latest version",
	Long: "Use this command to update your runtipi instance to the latest version",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Permission Warning
		if noPermissions {
			fmt.Printf("%s No permissions mode enabled, you may face issues with Runtipi\n", constants.Yellow("⚠"))
		}

		// Root folder
		rootFolder, osErr := os.Getwd()
	
		if osErr != nil {
			fmt.Printf("%s Faild to get root folder\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", osErr)
			os.Exit(1)
		}

		// Define paths
		cliPath := path.Join(rootFolder, "runtipi-cli-go")

		// Start spinner
		spinner.SetMessage("Updating runtipi...")
		spinner.Start()

		// Get versions
		version := args[0]
		currentVersion, currentVersionErr := env.GetEnvValue("TIPI_VERSION")
		if currentVersionErr != nil {
			fmt.Printf("%s Failed to get current environment version\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", currentVersionErr)
			os.Exit(1)
		}

		spinner.Update(fmt.Sprintf("Updating from %s to %s", constants.Blue(currentVersion), constants.Blue(version)))

		// Validate
		spinner.SetMessage("Validating version")

		isValid, validateErr := release.ValidateVersion(version)

		if validateErr != nil {
			spinner.Fail("Error in validating version")
			spinner.Stop()
			fmt.Printf("Error: %s\n", validateErr)
			os.Exit(1)
		}

		if !isValid {
			spinner.Fail("Version is not valid")
			spinner.Stop()
			os.Exit(1)
		}
		
		spinner.Succeed("Version is valid")

		// Compare versions
		spinner.SetMessage("Comparing versions...")

		versionToUpdate := ""

		if version == "latest" {
			latestVersion, latestVersionErr := release.GetLatestVersion()
			if latestVersionErr != nil {
				spinner.Fail("Failed to get latest version")
				spinner.Stop()
				fmt.Printf("Error: %s\n", latestVersionErr)
				os.Exit(1)
			}
			versionToUpdate = latestVersion
		} else if version == "nightly" {
			spinner.Fail("Nightly currently not supported")
			spinner.Stop()
			os.Exit(1)
		} else {
			if currentVersion != "nightly" {
				isMajor, isMajorErr := release.IsMajorBump(version, currentVersion)

				if isMajorErr != nil {
					spinner.Fail("Failed to compare versions")
					spinner.Stop()
					fmt.Printf("Error: %s\n", isMajorErr)
					os.Exit(1)
				}
	
				if isMajor {
					spinner.Fail("You are trying to update to a new major version. Please update manually using the update instructions on the website. https://runtipi.io/docs/reference/breaking-updates")
					spinner.Stop()
					os.Exit(1)
				}
			}

			versionToUpdate = version
		}

		spinner.Succeed("Versions compared")

		// Backup CLI
		spinner.SetMessage("Backing up current CLI")

		backupErr := release.BackupCurrentCLI()

		if backupErr != nil {
			spinner.Fail("Failed to backup current CLI, no modification were made")
			spinner.Stop()
			fmt.Printf("Error: %s\n", backupErr)
			os.Exit(1)
		}

		spinner.Succeed("CLI backed up")

		// Download latest CLI
		spinner.SetMessage("Downloading latest CLI")

		downloadErr := release.DownloadLatestCLI(versionToUpdate)

		if downloadErr != nil {
			spinner.Fail("Failed to download latest CLI, please copy the runtipi-cli-go.bak file to runtipi-cli-go and try again")
			spinner.Stop()
			fmt.Printf("Error: %s\n", downloadErr)
			os.Exit(1)
		}

		spinner.Succeed("New CLI downloaded successfully")

		// Start new CLI
		spinner.SetMessage("Starting new CLI")

		cliArgs := []string{"start"}

		if envFile != "" {
			cliArgs = append(cliArgs, "--env-file")
			cliArgs = append(cliArgs, envFile)
		}

		if noPermissions {
			cliArgs = append(cliArgs, "--no-permissions")
		}

		_, startErr := exec.Command(cliPath, cliArgs...).Output()

		if startErr != nil {
			spinner.Fail("Failed to start the new CLI, please copy the runtipi-cli-go.bak file to runtipi-cli-go and try again")
			spinner.Stop()
			fmt.Printf("Error: %s\n", downloadErr)
			os.Exit(1)
		}

		spinner.Succeed("New CLI started successfully, you are good to go")
		
		// Succeed
		spinner.Stop()

		internalIp, _ := env.GetEnvValue("INTERNAL_IP")
		nginxPort, _ := env.GetEnvValue("NGINX_PORT")

		boxMessage := fmt.Sprintf("Visit http://%s:%s to access the dashboard\n\nFind documentation and guides at: https://runtipi.io\n\nTipi is entirely written in TypeScript and we are looking for contributors!", internalIp, nginxPort)

		Box := box.New(box.Config{Py: 2, Px: 2, Type: "Double", Color: "Green", TitlePos: "Top", ContentAlign: "Center"})
		Box.Print("Runtipi updated successfully 🎉", boxMessage)
	},
}
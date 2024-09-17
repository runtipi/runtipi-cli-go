package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"runtipi-cli-go/internal/constants"
	"runtipi-cli-go/internal/env"
	"runtipi-cli-go/internal/system"
	"runtipi-cli-go/internal/utils"

	"github.com/aquasecurity/table"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/spf13/cobra"
)

func init() {
	debugCmd.Flags().BoolVar(&showLogs, "logs", false,  "Show last 15 lines of all container logs.")
	rootCmd.AddCommand(debugCmd)
}

func GetEnvSafe(key string) (string) {
	val, err := env.GetEnvValue(key)
	if err != nil {
		return "Error"
	}
	if val == "" {
		return "Not set"
	}
	return val
}

func GetEnvSafeRedact(key string) (string) {
	val, err := env.GetEnvValue(key)
	if err != nil {
		return "Error"
	}
	if val == "" {
		return "Not set"
	}
	return "<redacted>"
}

var debugCmd = &cobra.Command{
	Use: "debug",
	Short: "Debug runtipi",
	Long: "Use this command to debug your runtipi instance (useful for issues)",
	Run: func(cmd *cobra.Command, args []string) {
		// Print warning
		fmt.Printf("\n%s Make sure you have started tipi before running this command\n\n", constants.Yellow("⚠"))

		// Containers
		containers := []string{"runtipi", "runtipi-db", "runtipi-redis", "runtipi-reverse-proxy"}

		// Root folder
		rootFolder, rootFolderErr := os.Getwd()

		if rootFolderErr != nil {
			fmt.Printf("%s Failed to get root folder\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", rootFolderErr)
			os.Exit(1)
		}

		// System Info
		fmt.Printf("---- %s ----\n", constants.Blue("System Info"))
		operatingSystem := runtime.GOOS
		kernel, kernelErr := exec.Command("uname", "-r").Output()
		if kernelErr != nil {
			fmt.Printf("%s Failed to run uname command\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", kernelErr)
			os.Exit(1)
		}
		memory, memoryErr := mem.VirtualMemory()
		if memoryErr != nil {
			fmt.Printf("%s Failed to get total memory\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", memoryErr)
			os.Exit(1)
		}
		arch := system.GetArch()
		sysInfoTable := table.New(os.Stdout)
		sysInfoTable.AddRow("OS", operatingSystem)
		sysInfoTable.AddRow("OS Version", string(kernel[:]))
		sysInfoTable.AddRow("Memory (GB)", utils.FormatFileSize(float64(memory.Total)))
		sysInfoTable.AddRow("Architecture", arch)
		sysInfoTable.Render()
		fmt.Println()

		// User config
		fmt.Printf("---- %s ----\n", constants.Blue("User Config"))
		userConfigTable := table.New(os.Stdout)
		if _, err := os.Stat(path.Join(rootFolder, "user-config", "tipi-compose.yml")); errors.Is(err, os.ErrNotExist) {
			userConfigTable.AddRow("Custom tipi docker compose", "false")
		} else {
			userConfigTable.AddRow("Custom tipi docker compose", "true")
		}
		userConfigTable.Render()
		fmt.Println()

		// Settings
		fmt.Printf("---- %s ----\n", constants.Blue("Settings"))
		settings, settingsErr := os.ReadFile(path.Join(rootFolder, "state", "settings.json"))
		if settingsErr != nil {
			fmt.Printf("%s Failed to read settings file\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", settingsErr)
			os.Exit(1)
		}
		var prettyJson bytes.Buffer
		prettifyErr := json.Indent(&prettyJson, settings, "", "\t")
		if prettifyErr != nil {
			fmt.Printf("%s Failed to prettify json\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", prettifyErr)
			os.Exit(1)
		}
		fmt.Println(prettyJson.String())
		fmt.Println()

		// Env
		fmt.Printf("---- %s ----\n", constants.Blue("Environment"))
		envTable := table.New(os.Stdout)
		envTable.AddRow("POSTGRES_PASSWORD", GetEnvSafeRedact("POSTGRES_PASSWORD"))
		envTable.AddRow("REDIS_PASSWORD", GetEnvSafeRedact("REDIS_PASSWORD"))
		envTable.AddRow("APPS_REPO_ID", GetEnvSafe("APPS_REPO_ID"))
		envTable.AddRow("APPS_REPO_URL", GetEnvSafe("APPS_REPO_URL"))
		envTable.AddRow("TIPI_VERSION", GetEnvSafe("TIPI_VERSION"))
		envTable.AddRow("INTERNAL_IP", GetEnvSafe("INTERNAL_IP"))
		envTable.AddRow("ARCHITECTURE", GetEnvSafe("ARCHITECTURE"))
		envTable.AddRow("JWT_SECRET", GetEnvSafeRedact("JWT_SECRET"))
		envTable.AddRow("ROOT_FOLDER_HOST", GetEnvSafe("ROOT_FOLDER_HOST"))
		envTable.AddRow("RUNTIPI_APP_DATA_PATH", GetEnvSafe("RUNTIPI_APP_DATA_PATH"))
		envTable.AddRow("NGINX_PORT", GetEnvSafe("NGINX_PORT"))
		envTable.AddRow("NGINX_PORT_SSL", GetEnvSafe("NGINX_PORT_SSL"))
		envTable.AddRow("DOMAIN", GetEnvSafeRedact("DOMAIN"))
		envTable.AddRow("POSTGRES_HOST", GetEnvSafe("POSTGRES_HOST"))
		envTable.AddRow("POSTGRES_DBNAME", GetEnvSafe("POSTGRES_DBNAME"))
		envTable.AddRow("POSTGRES_USERNAME", GetEnvSafe("POSTGRES_USERNAME"))
		envTable.AddRow("POSTGRES_PORT", GetEnvSafe("POSTGRES_PORT"))
		envTable.AddRow("REDIS_HOST", GetEnvSafe("REDIS_HOST"))
		envTable.AddRow("DEMO_MODE", GetEnvSafe("DEMO_MODE"))
		envTable.AddRow("LOCAL_DOMAIN", GetEnvSafe("LOCAL_DOMAIN"))
		envTable.Render()
		fmt.Println()

		// Containers
		fmt.Printf("---- %s ----\n", constants.Blue("Containers"))
		containerTable := table.New(os.Stdout)
		for _, container := range containers {
			status, err := exec.Command("docker", "ps", "-a", "--filter", "name=" + container, "--format", "{{.Status}}").Output()
			if err != nil {
				containerTable.AddRow(container, constants.Red("down"))
			} else {
				if	strings.Contains(strings.ToLower(string(status)), "up") {
					containerTable.AddRow(container, constants.Green("up"))
				} else {
					containerTable.AddRow(container, constants.Red("down"))
				}
			}
		}
		containerTable.Render()
		fmt.Println("\n^ If a container is not 'Up', you can run the command `docker logs <container_name>` to see the logs of that container.")
		fmt.Println()
		
		// Logs
		if showLogs {
			fmt.Printf("---- %s ----", constants.Blue("Container Logs"))
			for _, container := range containers {
				logs, err := exec.Command("docker", "logs", "-n", "15", container).Output()
				if err != nil {
					fmt.Printf("%s Failed to get container logs\n", constants.Red("✗"))
					fmt.Printf("Error: %s\n", err)
					os.Exit(1)
				}
				fmt.Println("\n" + constants.Green(container))
				fmt.Printf("\n%s", logs)
			}
			fmt.Println("^ Make sure to remove any personal information from the logs.")
		}
	},
}

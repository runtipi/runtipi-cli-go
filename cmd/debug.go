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
	"syscall"

	"github.com/aquasecurity/table"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/env"
	"github.com/steveiliop56/runtipi-cli-go/internal/system"
	"github.com/steveiliop56/runtipi-cli-go/internal/utils"
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
		fmt.Println("\n⚠️  Make sure you have started tipi before running this command")

		// Containers
		containers := []string{"runtipi", "runtipi-db", "runtipi-redis", "runtipi-reverse-proxy"}

		// Colors
		green := color.New(color.FgGreen).SprintFunc()
		red := color.New(color.FgRed).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()

		// Root folder
		rootFolder, rootFolderErr := os.Getwd()

		if rootFolderErr != nil {
			utils.PrintError("Failed to get root folder")
			fmt.Printf("Error: %s\n", rootFolderErr)
			os.Exit(1)
		}

		// System Info
		fmt.Println("\n--- " + blue("System Info") + " ---")
		operatingSystem := runtime.GOOS
		kernel, kernelErr := exec.Command("uname", "-r").Output()
		if kernelErr != nil {
			utils.PrintError("Failed to run uname command")
			fmt.Printf("Error: %s\n", kernelErr)
			os.Exit(1)
		}
		sysSchema := syscall.Sysinfo_t{}
		sysErr := syscall.Sysinfo(&sysSchema)
		if sysErr != nil {
			utils.PrintError("Failed to get total memory")
			fmt.Printf("Error: %s\n", sysErr)
			os.Exit(1)
		}
		totalMemory := uint64(sysSchema.Totalram) * uint64(sysSchema.Unit)
		totalMemoryGB := fmt.Sprintf("%.2f", float32(totalMemory)/(1<<30))
		arch := system.GetArch()
		sysInfoTable := table.New(os.Stdout)
		sysInfoTable.AddRow("OS", operatingSystem)
		sysInfoTable.AddRow("OS Version", string(kernel[:]))
		sysInfoTable.AddRow("Memory (GB)", totalMemoryGB)
		sysInfoTable.AddRow("Architecture", arch)
		sysInfoTable.Render()

		// User config
		fmt.Println("\n--- " + blue("User Config") + " ---")
		userConfigTable := table.New(os.Stdout)
		if _, err := os.Stat(path.Join(rootFolder, "user-config", "tipi-compose.yml")); errors.Is(err, os.ErrNotExist) {
			userConfigTable.AddRow("Custom tipi docker compose", "false")
		} else {
			userConfigTable.AddRow("Custom tipi docker compose", "true")
		}
		userConfigTable.Render()

		// Settings
		fmt.Println("\n--- " + blue("Settings") + " ---")
		settings, settingsErr := os.ReadFile(path.Join(rootFolder, "state", "settings.json"))
		if settingsErr != nil {
			utils.PrintError("Failed to read settings file")
			fmt.Printf("Error: %s\n", settingsErr)
			os.Exit(1)
		}
		var prettyJson bytes.Buffer
		prettifyErr := json.Indent(&prettyJson, settings, "", "\t")
		if prettifyErr != nil {
			utils.PrintError("Failed to prettify json")
			fmt.Printf("Error: %s\n", prettifyErr)
			os.Exit(1)
		}
		fmt.Println(prettyJson.String())

		// Env
		fmt.Println("\n--- " + blue("Environment") + " ---")
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

		// Containers
		fmt.Println("\n--- " + blue("Container Status") + " ---")
		containerTable := table.New(os.Stdout)
		for _, container := range containers {
			status, err := exec.Command("docker", "ps", "-a", "--filter", "name=" + container, "--format", "{{.Status}}").Output()
			if err != nil {
				containerTable.AddRow(container, red("down"))
			} else {
				if	strings.Contains(strings.ToLower(string(status)), "up") {
					containerTable.AddRow(container, green("up"))
				} else {
					containerTable.AddRow(container, red("down"))
				}
			}
		}
		containerTable.Render()
		fmt.Println("\n^ If a container is not 'Up', you can run the command `docker logs <container_name>` to see the logs of that container.")
	
		// Logs
		if showLogs {
			fmt.Println("\n--- " + blue("Container Logs") + " ---")
			for _, container := range containers {
				logs, err := exec.Command("docker", "logs", "-n", "15", container).Output()
				if err != nil {
					utils.PrintError("Failed to get container logs")
					fmt.Printf("Error: %s\n", err)
					os.Exit(1)
				}
				fmt.Println("\n" + green(container))
				fmt.Printf("\n%s", logs)
			}
			fmt.Println("^ Make sure to remove any personal information from the logs.")
		}
	},
}

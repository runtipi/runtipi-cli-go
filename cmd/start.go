package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/Delta456/box-cli-maker"
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/env"
	"github.com/steveiliop56/runtipi-cli-go/internal/spinner"
	"github.com/steveiliop56/runtipi-cli-go/internal/system"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use: "start",
	Short: "Start runtipi",
	Long: "Use this command to start the runtipi docker stack",
	Run: func(cmd *cobra.Command, args []string) {
		// Docker check
		spinner.SetMessage("Checking user permissions")
		spinner.Start()

		dockerErr := system.EnsureDocker()
		if dockerErr != nil {
			if dockerErr.Error() == "docker-error" {
				spinner.Fail("Docker is not installed or user has not the right permissions. See https://docs.docker.com/engine/install/ for more information")
				spinner.Stop()
				return
			} else if dockerErr.Error() == "compose-error" {
				spinner.Fail("Docker compose plugin is not installed. See https://docs.docker.com/compose/install/linux/ for more information")
				spinner.Stop()
				return
			}
		}
		spinner.Succeed("User permissions are ok")

		// System files
		spinner.SetMessage("Copying system files")
		fileCopyErr := system.CopySystemFiles()
		if fileCopyErr != nil {
			spinner.Fail("Failed to copy system files")
			spinner.Stop()
			fmt.Printf("Error: %s\n", fileCopyErr)
			return
		}
		spinner.Succeed("Copied system files")

		// Env file
		spinner.SetMessage("Generating env file")
		envErr := env.GenerateEnv()
		if envErr != nil {
			spinner.Fail("Failed to generate env file")
			spinner.Stop()
			fmt.Printf("Error: %s\n", envErr)
			return
		}
		spinner.Succeed("Env file generated")

		// Ensure permissions
		spinner.SetMessage("Ensuring permissions...")
		filePermErr := system.EnsureFilePermissions()
		if filePermErr != nil {
			spinner.Fail("Failed to chmod files")
			spinner.Stop()
			fmt.Printf("Error: %s\n", filePermErr)
			return
		}
		spinner.Succeed("File permissions ok")

		// Pull Images
		spinner.SetMessage("Pulling images...")

		rootFolder, rootFolderErr := os.Getwd()

		if rootFolderErr != nil {
			spinner.Fail("Failed to get root folder")
			spinner.Stop()
			fmt.Printf("Error: %s\n", rootFolderErr)
			return
		}

		_, pullError := exec.Command("docker", "compose", "--env-file", path.Join(rootFolder, ".env"), "pull").Output()

		if pullError != nil {
			spinner.Fail("Failed to pull images")
			spinner.Stop()
			fmt.Printf("Error: %s\n", pullError)
			return
		}
		spinner.Succeed("Images pulled")

		// Stop containers
		spinner.SetMessage("Stopping existing containers")

		containersToRm := []string{"runtipi-reverse-proxy", "runtipi-db", "runtipi-redis", "runtipi", "tipi-db", "tipi-redis", "tipi-reverse-proxy", "tipi-docker-proxy", "tipi-dashboard", "tipi-worker"}

		for _, container := range containersToRm {
			exec.Command("docker", "stop", container).Output()
			exec.Command("docker", "rm", container).Output()
		}
		spinner.Succeed("Existing container stopped")

		// Start containers
		spinner.SetMessage("Starting containers")

		baseArgs := []string{"compose", "--project-name", "runtipi", "-f", "docker-compose.yml"}

		if _, err := os.Stat(path.Join(rootFolder, "user-config", "tipi-compose.yml")); err == nil {
			baseArgs = append(baseArgs, "-f", path.Join(rootFolder, "user-config", "tipi-compose.yml"))
		}

		baseArgs = append(baseArgs, "--env-file",  path.Join(rootFolder, ".env"), "up", "--detach", "--remove-orphans")

		_, upErr := exec.Command("docker", baseArgs...).Output()

		if upErr != nil {
			spinner.Fail("Failed to start containers")
			spinner.Stop()
			fmt.Printf("Error: %s\n", upErr)
			return
		}
		spinner.Succeed("Containers started")

		// Finish
		spinner.Stop()

		internalIp, _ := env.GetEnvValue("INTERNAL_IP")
		nginxPort, _ := env.GetEnvValue("NGINX_PORT")

		boxMessage := "Visit http://"  + internalIp + ":" + nginxPort + " to access the dashboard\n\nFind documentation and guides at: https://runtipi.io\n\nTipi is entirely written in TypeScript and we are looking for contributors!"

		Box := box.New(box.Config{Py: 2, Px: 2, Type: "Double", Color: "Green", TitlePos: "Top", ContentAlign: "Center"})
 		Box.Print("Runtipi started successfully 🎉", boxMessage)
	},
}

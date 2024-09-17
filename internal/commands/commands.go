package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"runtipi-cli-go/internal/constants"
	"runtipi-cli-go/internal/env"
	"runtipi-cli-go/internal/spinner"
	"runtipi-cli-go/internal/system"

	"github.com/Delta456/box-cli-maker"
)

func Start(envFile string, noPermissions bool) {
	// Permission Warning
	if noPermissions {
		fmt.Printf("%s No permissions mode enabled, you may face issues with Runtipi\n", constants.Yellow("⚠"))
	}
	
	// Docker check
	spinner.SetMessage("Checking user permissions")
	spinner.Start()

	dockerErr := system.EnsureDocker()
	if dockerErr != nil {
		if dockerErr.Error() == "docker-error" {
			spinner.Fail("Docker is not installed or user has not the right permissions. See https://docs.docker.com/engine/install/ for more information")
			spinner.Stop()
			os.Exit(1)
		} else if dockerErr.Error() == "compose-error" {
			spinner.Fail("Docker compose plugin is not installed. See https://docs.docker.com/compose/install/linux/ for more information")
			spinner.Stop()
			os.Exit(1)
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
		os.Exit(1)
	}
	spinner.Succeed("Copied system files")

	// Env file
	spinner.SetMessage("Generating env file")
	envErr := env.GenerateEnv(envFile)
	if envErr != nil {
		spinner.Fail("Failed to generate env file")
		spinner.Stop()
		fmt.Printf("Error: %s\n", envErr)
		os.Exit(1)
	}
	spinner.Succeed("Env file generated")

	// Ensure permissions
	if !noPermissions {
		spinner.SetMessage("Ensuring permissions...")
		filePermErr := system.EnsureFilePermissions()
		if filePermErr != nil {
			spinner.Fail("Failed to chmod files")
			spinner.Stop()
			fmt.Printf("Error: %s\n", filePermErr)
			os.Exit(1)
		}
		spinner.Succeed("File permissions ok")
	}

	// Pull Images
	spinner.SetMessage("Pulling images...")

	rootFolder, rootFolderErr := os.Getwd()

	if rootFolderErr != nil {
		spinner.Fail("Failed to get root folder")
		spinner.Stop()
		fmt.Printf("Error: %s\n", rootFolderErr)
		os.Exit(1)
	}

	_, pullError := exec.Command("docker", "compose", "--env-file", path.Join(rootFolder, ".env"), "pull").Output()

	if pullError != nil {
		spinner.Fail("Failed to pull images")
		spinner.Stop()
		fmt.Printf("Error: %s\n", pullError)
		os.Exit(1)
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
		os.Exit(1)
	}
	spinner.Succeed("Containers started")

	// Finish
	spinner.Stop()

	internalIp, _ := env.GetEnvValue("INTERNAL_IP")
	nginxPort, _ := env.GetEnvValue("NGINX_PORT")

	boxMessage := fmt.Sprintf("Visit http://%s:%s to access the dashboard\n\nFind documentation and guides at: https://runtipi.io\n\nTipi is entirely written in TypeScript and we are looking for contributors!", internalIp, nginxPort)

	Box := box.New(box.Config{Py: 2, Px: 2, Type: "Double", Color: "Green", TitlePos: "Top", ContentAlign: "Center"})
	Box.Print("Runtipi started successfully 🎉", boxMessage)
}

func Stop() {
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
}
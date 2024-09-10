package release

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/steveiliop56/runtipi-cli-go/internal/schemas"
	"github.com/steveiliop56/runtipi-cli-go/internal/system"
)

func IsMajorBump(newVersion string, currentVersion string) (bool, error) {
	newVersionMajor := strings.Split(strings.Replace(newVersion, "v", "", 1), ".")[0]
	currentVersionMajor := strings.Split(strings.Replace(currentVersion, "v", "", 1), ".")[0]

	newVersionMajorInt, newVersionMajorIntErr := strconv.ParseInt(newVersionMajor, 10, 64)

	if newVersionMajorIntErr != nil {
		return false, newVersionMajorIntErr
	}

	currentVersionMajorInt, currentVersionMajorIntErr := strconv.ParseInt(currentVersionMajor, 10, 64)

	if currentVersionMajorIntErr != nil {
		return false, currentVersionMajorIntErr
	}

	if newVersionMajorInt > currentVersionMajorInt {
		return true, nil
	}

	return false, nil
}

func GetLatestVersion() (string, error) {
	apiUrl := "https://api.github.com/repos/steveiliop56/runtipi-cli-go/releases/latest"

	response, requestErr := http.Get(apiUrl)

	if requestErr != nil {
		return "", requestErr
	}

	defer response.Body.Close()

	release := new(schemas.GithubRelease)

	jsonErr := json.NewDecoder(response.Body).Decode(&release)

	if jsonErr != nil {
		return "", jsonErr
	}

	return release.TagName, nil
}

func ValidateVersion(version string) (bool, error) {
	apiUrl := "https://api.github.com/repos/steveiliop56/runtipi-cli-go/releases/tags/" + version

	response, requestErr := http.Get(apiUrl)

	if requestErr != nil {
		return false, requestErr
	}

	defer response.Body.Close()

	release := new(schemas.GithubRelease)

	jsonErr := json.NewDecoder(response.Body).Decode(&release)

	if jsonErr != nil {
		return false, jsonErr
	}

	if release.Status == "404" {
		return false, nil
	}

	return true, nil
}

func DownloadLatestCLI(version string) (error) {
	arch := system.GetArch()
	assetUrl := fmt.Sprintf("https://github.com/steveiliop56/runtipi-cli-go/releases/download/%s/runtipi-cli-go-%s", version, arch)

	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		return osErr
	}

	cliPath := path.Join(rootFolder, "runtipi-cli-go")

	os.Remove(cliPath)

	create, createErr := os.Create(cliPath)

	if createErr != nil {
		return createErr
	}

	defer create.Close()

	response, requestErr := http.Get(assetUrl)

	if requestErr != nil {
		return requestErr
	}

	defer response.Body.Close()

	_, writeErr := io.Copy(create, response.Body)

	if writeErr != nil {
		return writeErr
	}

	_, chmodErr := exec.Command("chmod", "+x", cliPath).Output()

	if chmodErr != nil {
		return chmodErr
	}

	return nil
}

func BackupCurrentCLI() (error) {
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		return osErr
	}

	cliPath := path.Join(rootFolder, "runtipi-cli-go")
	cliBackupPath := path.Join(rootFolder, "runtipi-cli-go.bak")

	_, copyErr := exec.Command("cp", cliPath, cliBackupPath).Output()

	if copyErr != nil {
		return copyErr
	}

	return nil
}
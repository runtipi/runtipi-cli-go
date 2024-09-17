package backups

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"

	"runtipi-cli-go/internal/system"
)

func CreateBackup() (string, error) {
	// Root folder
	rootFolder, osErr := os.Getwd()

	if osErr != nil {
		return "", osErr
	}

	// Time
	time := time.Now().UnixMilli()

	// Define archive name and path
	archiveName := fmt.Sprintf("%s-%d.tar.gz", filepath.Base(rootFolder), time)
	tmpArchivePath := path.Join("/", "tmp", archiveName)
	archivePath := path.Join(rootFolder, "backups", archiveName)

	// Go to parent root folder
	os.Chdir(path.Dir(rootFolder))

	// Create archive
	_, tarErr := exec.Command("tar", "-czf", tmpArchivePath, filepath.Base(rootFolder)).Output()

	if tarErr != nil {
		return "", tarErr
	}

	// Copy archive to backups folder
	cpErr := system.Copy(tmpArchivePath, archivePath)

	if cpErr != nil {
		return "", cpErr
	}

	return archivePath, nil
}

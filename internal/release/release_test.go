package release_test

import (
	"errors"
	"os"
	"path"
	"testing"

	"github.com/steveiliop56/runtipi-cli-go/internal/release"
)

func init() {
	// Change root folder
	os.Chdir("../..")
}

// Test the major bump validator works
func TestMajorValidator(t *testing.T) {
	// Check major
	isMajor, isMajorErr := release.IsMajorBump("v2.0.0", "v1.0.0")

	// Check for errors
	if isMajorErr != nil {
		t.Fatalf("Major bump validator failed, error: %s\n", isMajorErr)
	}

	// Check result
	if !isMajor {
		t.Fatalf("Major bump validator returned false on major bump!")
	}

	// Run the validator again on feature version
	isMajorFeat, isMajorFeatErr := release.IsMajorBump("v1.1.0", "v1.0.0")

	// Check for errors
	if isMajorFeatErr != nil {
		t.Fatalf("Mahor bump validator failed, error: %s\n", isMajorFeatErr)
	}

	// Check result
	if isMajorFeat {
		t.Fatalf("Major bump validator returned true on feature!")
	}
}

// Test validate version
func TestValidateVersion(t *testing.T) {
	// Try correct version
	validateCheckCorrect, validateCheckCorrectErr := release.ValidateVersion("v0.1.0-alpha.1-runtipi-v3.6.0")

	// Check for errors
	if validateCheckCorrectErr != nil {
		t.Fatalf("Version validater returned an error: %s\n", validateCheckCorrectErr)
	}

	// Check result
	if !validateCheckCorrect {
		t.Fatalf("Version validator returned false on correct version!")
	}

	// Try wrong version
	validateCheckWrong, validateCheckWrongErr := release.ValidateVersion("v0.1.0-alpha.1-runtipi-v3.5.8")

	// Check for errors
	if validateCheckWrongErr != nil {
		t.Fatalf("Version validater returned an error: %s\n", validateCheckWrongErr)
	}

	// Check result
	if validateCheckWrong {
		t.Fatalf("Version validator returned true on wrongs version!")
	}
}

// Test CLI download
func TestCLIDownload(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Define paths
	cliPath := path.Join(rootFolder, "runtipi-cli-go")

	// Delete old CLI
	os.Remove(cliPath)

	// Download new CLI
	downloadErr := release.DownloadLatestCLI("v0.1.0-alpha.1-runtipi-v3.6.0")

	// Check for errors
	if downloadErr != nil {
		t.Fatalf("Failed to download CLI, error: %s\n", downloadErr)
	}

	// Check if CLI got downloaded
	if _, err := os.Stat(cliPath); errors.Is(err, os.ErrNotExist) {
		t.Fatal("CLI doesn't exist!")
	}
}

// Test backup CLI
func TestBackupCLI(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Define paths
	cliPath := path.Join(rootFolder, "runtipi-cli-go")
	cliBackupPath := path.Join(rootFolder, "runtipi-cli-go")

	// Delete old files
	os.Remove(cliPath)
	os.Remove(cliBackupPath)

	// Create empty CLI file
	os.Create(cliPath)

	// Create backup
	backupErr := release.BackupCurrentCLI()

	// Check for errors
	if backupErr != nil {
		t.Fatalf("Error in backing up CLI, error: %s\n", backupErr)
	}

	// Check if backup file exists
	if _, err := os.Stat(cliBackupPath); errors.Is(err, os.ErrNotExist) {
		t.Fatal("CLI backup doesn't exist!")
	}
}
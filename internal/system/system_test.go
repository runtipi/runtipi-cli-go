package system_test

import (
	"errors"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"testing"

	"runtipi-cli-go/internal/seed"
	"runtipi-cli-go/internal/system"
)

func init() {
	// Change back to the root folder
	os.Chdir("../..")
}

// Test docker check is working
func TestDockerCheck(t *testing.T) {
	// Check docker manually
	_, dockerCheckErr := exec.Command("docker", "--version").Output()

	if dockerCheckErr != nil {
		t.Fatal("Docker doesn't seem to be working")
	}

	_, composeCheckErr := exec.Command("docker", "compose", "--version").Output()

	if composeCheckErr != nil {
		t.Fatal("Docker compose doesn't seem to be working")
	}

	// EnsureDocker should return no error now
	err := system.EnsureDocker()
	if err != nil {
		t.Fatalf("Ensure docker returned an error while it shouldn't, error: %s\n", err)
	}
}

// Test get seed is working
func TestGetSeed(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Define paths
	statePath := path.Join(rootFolder, "state")
	seedPath := path.Join(statePath, "seed")

	// Delete state
	os.RemoveAll(statePath)

	// Recreate state path
	os.Mkdir(statePath, 0755)

	// Regenerate seed
	seed.GenerateSeed(rootFolder)

	// Get seed
	seed, err := os.ReadFile(seedPath)

	if err != nil {
		t.Fatalf("Os read file returned an error: %s\n", err)
	}

	// Get seed from thet getSeed function
	seedTest, seedTestErr := system.GetSeed(rootFolder)

	if seedTestErr != nil {
		t.Fatalf("GetSeed returned an error: %s\n", seedTestErr)
	}

	// Check if both seeds match
	if string(seed[:]) != seedTest {
		t.Fatalf("Seeds do not match, test got %s, function got %s\n", seedTest, string(seed[:]))
	}
}

// Checks if the arch is getting picked up correctly
func TestArch(t *testing.T) {
	// Get arch
	arch := runtime.GOARCH

	// Get system helper arch
	archTest := system.GetArch()

	// Check if arch matches
	if arch != archTest {
		t.Fatalf("Arch is incorrect, test got %s, function got %s\n", arch, archTest)
	}
}

// Check if derive entopy is working
func TestDeriveEntopy(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Define paths
	statePath := path.Join(rootFolder, "state")

	// Delete old seed
	os.RemoveAll(statePath)

	// Recreate state apth
	os.Mkdir(statePath, 0755)

	// Regenerate seed
	seed.GenerateSeed(rootFolder)

	// Get seed
	seed, err := system.GetSeed(rootFolder)

	if err != nil {
		t.Fatalf("Failed to get seed, error: %s\n", err)
	}

	// Create new string
	newString := system.DeriveEntopy("test", seed)

	// String should be 64 chars
	if len(newString) != 64 {
		t.Fatalf("Invalid string length, it should be 64 chars, got %s\n", strconv.Itoa(len(newString)))
	}
}

// Test copy system files is working
func TestCopySystemFiles(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Define paths
	paths := []string{
		"docker-compose.yml",
		"VERSION",
		"apps",
		"data",
		"app-data",
		"state",
		"repos",
		"media",
		"traefik",
		"user-config",
		"backups",
		"logs",
	}

	// Delete everything
	for _, p := range paths {
		os.RemoveAll(path.Join(rootFolder, p))
	}

	// Generate files
	system.CopySystemFiles()

	// Check if files got generated
	for _, p := range paths {
		if _, err := os.Stat(path.Join(rootFolder, p)); err != nil {
			t.Fatalf("Cannot find %s", p)
		}
	}
}

// Test ensure file permissions
func TestEnsureFilePermissions(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Generate system files
	system.CopySystemFiles()

	// Define expected permissions
	expectedPermissions := map[string]os.FileMode{
		"state":                    0777,
		"data":                     0777,
		"apps":                     0777,
		"logs":                     0777,
		"traefik":                  0777,
		"repos":                    0777,
		"user-config":              0777,
		"state/settings.json":      0666,
		".env":                     0664,
		"docker-compose.yml":       0664,
		"VERSION":                  0664,
		"traefik/shared/acme.json": 0600,
		"state/seed":               0600,
	}

	// Create some blank test files
	blankFiles := []string{
		"state/settings.json",
		".env",
		"traefik/shared/acme.json",
		"state/seed",
		"docker-compose.yml",
		"VERSION",
	}

	// Make sure traefik shared dir exists
	os.Mkdir(path.Join(rootFolder, "traefik/shared"), 0777)

	// Create blank files
	for _, file := range blankFiles {
		if _, fileErr := os.Stat(path.Join(rootFolder, file)); errors.Is(fileErr, os.ErrNotExist) {
			os.WriteFile(path.Join(rootFolder, file), []byte(""), 0777)	
		}
	}

	// Assign random permissions to these files
	for _, file := range blankFiles {
		os.Chmod(path.Join(rootFolder, file), 0777)
	}

	// Ensure file permissions
	system.EnsureFilePermissions()

	// Check file permissions
	for item, expectedMode := range expectedPermissions {
		info, err := os.Stat(path.Join(rootFolder, item))
		if err != nil {
			t.Fatalf("Failed to get info for file %s, error: %s\n", item, err)
		}
		if info.Mode().Perm() != expectedMode {
			t.Fatalf("File/folder %s didn't get correct permissions, expected %v, got %v\n", item, expectedMode, info.Mode().Perm())
		}
	}
}

// Test copy
func TestCopy(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Define paths
	srcFilePath := path.Join(rootFolder, "cpSrc.txt")
	destFilePath := path.Join(rootFolder, "cpDest.txt")

	// Delete old files
	os.RemoveAll(srcFilePath)
	os.RemoveAll(destFilePath)

	// Create src file
	writeErr := os.WriteFile(srcFilePath, []byte("test"), 0644)

	// Check for errors
	if writeErr != nil {
		t.Fatalf("Failed to write test file, error: %s\n", writeErr)
	}

	// Copy file
	cpErr := system.Copy(srcFilePath, destFilePath)

	// Check for errors
	if cpErr != nil {
		t.Fatalf("Failed to copy test file, error: %s\n", cpErr)
	}

	// Check if file got copied
	destStatus, statErr := os.Stat(destFilePath)

	if statErr != nil {
		t.Fatalf("Test file doesn't exist, error: %s\n", statErr)
	}

	// Check if permissions are the same
	if destStatus.Mode() != os.FileMode(0644) {
		t.Fatalf("Permissions are not the same! Expected -rw-r--r--, got %s\n", destStatus.Mode().String())
	}

	// Check contents
	contents, readErr := os.ReadFile(destFilePath)

	if readErr != nil {
		t.Fatalf("Failed to read destination file, error: %s\n", readErr)
	}

	if string(contents) != "test" {
		t.Fatalf("Contents not the same! Expected test got %s\n", string(contents))
	}
}

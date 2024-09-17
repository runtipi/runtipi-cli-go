package system_test

import (
	"os"
	"os/exec"
	"path"
	"runtime"
	"slices"
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
	os.Remove(statePath)

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
	os.Remove(statePath)

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
	composePath := path.Join(rootFolder, "docker-compose.yml")
	versionPath := path.Join(rootFolder, "VERSION")

	appsPath := path.Join(rootFolder, "apps")
	dataPath := path.Join(rootFolder, "data")
	appDataPath := path.Join(rootFolder, "app-data")
	statePath := path.Join(rootFolder, "state")
	reposPath := path.Join(rootFolder, "repos")
	mediaPath := path.Join(rootFolder, "media")
	traefikPath := path.Join(rootFolder, "traefik")
	userConfigPath := path.Join(rootFolder, "user-config")
	backupsPath := path.Join(rootFolder, "backups")
	logsPath := path.Join(rootFolder, "logs")

	// Delete everything
	os.Remove(composePath)
	os.Remove(versionPath)
	os.Remove(appsPath)
	os.Remove(dataPath)
	os.Remove(appDataPath)
	os.Remove(statePath)
	os.Remove(reposPath)
	os.Remove(mediaPath)
	os.Remove(traefikPath)
	os.Remove(userConfigPath)
	os.Remove(backupsPath)
	os.ReadDir(logsPath)

	// Generate files
	system.CopySystemFiles()

	// Check if files got generated
	if _, err := os.Stat(composePath); err != nil {
		t.Fatal("Cannot find compose file")
	}
	if _, err := os.Stat(versionPath); err != nil {
		t.Fatal("Cannot find version file")
	}
	if _, err := os.Stat(appsPath); err != nil {
		t.Fatal("Cannot find apps folder")
	}
	if _, err := os.Stat(dataPath); err != nil {
		t.Fatal("Cannot find data folder")
	}
	if _, err := os.Stat(appDataPath); err != nil {
		t.Fatal("Cannot find app-data folder")
	}
	if _, err := os.Stat(statePath); err != nil {
		t.Fatal("Cannot find state folder")
	}
	if _, err := os.Stat(reposPath); err != nil {
		t.Fatal("Cannot find repos folder")
	}
	if _, err := os.Stat(mediaPath); err != nil {
		t.Fatal("Cannot find media folder")
	}
	if _, err := os.Stat(traefikPath); err != nil {
		t.Fatal("Cannot find traefik folder")
	}
	if _, err := os.Stat(userConfigPath); err != nil {
		t.Fatal("Cannot find user-config folder")
	}
	if _, err := os.Stat(backupsPath); err != nil {
		t.Fatal("Cannot find backups folder")
	}
	if _, err := os.Stat(logsPath); err != nil {
		t.Fatal("Cannot find log folder")
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

	// Define files and folders permissions
	SevenSevenSevenItems := []string{"state", "data", "apps", "logs", "traefik", "repos", "user-config", "state"}
	SixSixSixItems := []string{"state/settings.json"}
	SixSixFourItems := []string{".env", "docker-compose.yml", "VERSION"}
	SixOOItems := []string{"traefik/shared/acme.json", "state/seed"}

	// Ensure file permissions
	system.EnsureFilePermissions()

	// Check file permissions
	items, itemsErr := os.ReadDir(rootFolder)

	if itemsErr != nil {
		t.Fatalf("Failed to get files and folder in cwd, error: %s\n", itemsErr)
	}

	for _, item := range items {
		info, err := os.Stat(path.Join(rootFolder, item.Name()))
		if err != nil {
			t.Fatalf("Failed to get info for file %s, error: %s\n", item.Name(), err)
		}
		mode := info.Mode().Perm()
		if slices.Contains(SevenSevenSevenItems, item.Name()) {
			if mode != os.FileMode(0777) {
				t.Fatalf("File/folder %s didn't get correct 777 permissions\n", item.Name())
			}
		} else if slices.Contains(SixSixSixItems, item.Name()) {
			if mode != os.FileMode(0666) {
				t.Fatalf("File/folder %s didn't get correct 666 permissions\n", item.Name())
			}
		} else if slices.Contains(SixSixFourItems, item.Name()) {
			if mode != os.FileMode(0664) {
				t.Fatalf("File/folder %s didn't get correct 664 permissions\n", item.Name())
			}
		} else if slices.Contains(SixOOItems, item.Name()) {
			if mode != os.FileMode(0600) {
				t.Fatalf("File/folder %s didn't get correct 600 permissions\n", item.Name())
			}
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
	os.Remove(srcFilePath)
	os.Remove(destFilePath)

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
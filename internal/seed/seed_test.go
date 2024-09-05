package seed_test

import (
	"os"
	"path"
	"strconv"
	"testing"

	"github.com/steveiliop56/runtipi-cli-go/internal/seed"
)

func init() {
	// Change back to the root folder
	os.Chdir("../..")
}

// Seed should be generated correctly
func TestSeedGen(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Define paths
	seedPath := path.Join(rootFolder, "state", "seed")

	// Remove seed
	os.Remove(seedPath)

	// Generate seed
	seed.GenerateSeed(rootFolder)

	// Check seed file exists
	if _, seedCheckErr := os.Stat(seedPath); seedCheckErr != nil {
		t.Fatal("Seed file does not exist!")
	}

	// If seed file exists read it and verify seed
	seed, seedReadErr := os.ReadFile(seedPath)
	if seedReadErr != nil {
		t.Fatalf("Failed to read seed, error: %s\n", seedReadErr)
	}

	// Check if seed is correct
	if len(seed) != 32 {
		t.Fatalf("Seed should be 32 chars, got %s\n", strconv.Itoa(len(seed)))
	}
}
package seed

import (
	"errors"
	"os"
	"path"

	"go.step.sm/crypto/randutil"
)

func GenerateSeed(rootFolder string) (error) {
	if _, err := os.Stat(path.Join(rootFolder, "state", "seed")); errors.Is(err, os.ErrNotExist) {
		seed, seedErr := randutil.Alphanumeric(32)
		if seedErr != nil {
			return seedErr
		}
		writeErr := os.WriteFile(path.Join(rootFolder, "state", "seed"), []byte(seed), 0644)
		if writeErr != nil {
			return writeErr
		}
		return nil
	}
	return nil
}
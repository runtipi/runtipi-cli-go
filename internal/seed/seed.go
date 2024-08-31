package seed

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"os"
	"path"
)

func GenerateSeed(rootFolder string) (error) {
	if _, err := os.Stat(path.Join(rootFolder, "state", "seed")); errors.Is(err, os.ErrNotExist) {
		b := make([]byte, 32)
		if _, err := rand.Read(b); err != nil {
			return err
		}
		seed := base64.URLEncoding.EncodeToString(b)[:32]
		writeErr := os.WriteFile(path.Join(rootFolder, "state", "seed"), []byte(seed), 0644)
		if writeErr != nil {
			return writeErr
		}
		return nil
	}
	return nil
}
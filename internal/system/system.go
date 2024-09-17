package system

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"net"
	"os"
	"os/exec"
	"path"
	"runtime"
	"slices"
	"strings"

	"runtipi-cli-go/internal/constants"
)

func EnsureDocker() (error) {
	_, dockerCheckErr := exec.Command("docker", "--version").Output()

	if dockerCheckErr != nil {
		return errors.New("docker-error")
	}

	_, composeCheckErr := exec.Command("docker", "compose", "--version").Output()

	if composeCheckErr != nil {
		return errors.New("compose-error")
	}

	return nil
}

func CopySystemFiles() (error) {
	rootFolder, err := os.Getwd()

	if err != nil {
		return err
	}

	os.WriteFile(path.Join(rootFolder, "docker-compose.yml"), []byte(constants.Compose), 0664)
	os.WriteFile(path.Join(rootFolder, "VERSION"), []byte(constants.RuntipiVersion), 0644)

	os.Mkdir(path.Join(rootFolder, "apps"), 0755)
	os.Mkdir(path.Join(rootFolder, "data"), 0755)
	os.Mkdir(path.Join(rootFolder, "app-data"), 0755)
	os.Mkdir(path.Join(rootFolder, "state"), 0755)
	os.Mkdir(path.Join(rootFolder, "repos"), 0755)
	os.Mkdir(path.Join(rootFolder, "media"), 0755)
	os.Mkdir(path.Join(rootFolder, "traefik"), 0755)
	os.Mkdir(path.Join(rootFolder, "user-config"), 0755)
	os.Mkdir(path.Join(rootFolder, "backups"), 0755)
	os.Mkdir(path.Join(rootFolder, "logs"), 0755)

	return nil
}

func GetSeed(rootFolder string) (string, error) {
	seed, err := os.ReadFile(path.Join(rootFolder, "state", "seed"))

	if err != nil {
		return "", err
	}

	return string(seed), nil
}

func DeriveEntopy(entopy string, seed string) (string) {
	hasher := sha256.New()
	hasher.Write([]byte(seed))
	hasher.Write([]byte(entopy))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

func GetInternalIp() (string, error) {
	ifaces, ifacesErr := net.Interfaces()

	if ifacesErr != nil {
		return "", ifacesErr
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return "", err
		}
		parsedIp := strings.Split(addrs[0].String(), "/")[0]
		if strings.Count(parsedIp, ":") < 2 {
			if !net.ParseIP(parsedIp).IsLoopback() {
				return parsedIp, nil
			}
		}
	}

	return "", nil
}

func GetArch() (string) {
	return runtime.GOARCH
}

func EnsureFilePermissions() (error) {
	rootFolder, err := os.Getwd()

	if err != nil {
		return err
	}

	SevenSevenSevenItems := []string{"state", "data", "apps", "logs", "traefik", "repos", "user-config", "state"}
	SixSixSixItems := []string{"state/settings.json"}
	SixSixFourItems := []string{".env", "docker-compose.yml", "VERSION"}
	SixOOItems := []string{"traefik/shared/acme.json", "state/seed"}
	
	items, itemsErr := os.ReadDir(rootFolder)

	if itemsErr != nil {
		return itemsErr
	}

	for _, item := range items {
		if slices.Contains(SevenSevenSevenItems, item.Name()) {
			_, execErr := exec.Command("chmod", "-R", "777", item.Name()).Output()
			if execErr != nil {
				return execErr
			}
		} else if slices.Contains(SixSixSixItems, item.Name()) {
			_, execErr := exec.Command("chmod", "-R", "666", item.Name()).Output()
			if execErr != nil {
				return execErr
			}
		} else if slices.Contains(SixSixFourItems, item.Name()) {
			_, execErr := exec.Command("chmod", "-R", "664", item.Name()).Output()
			if execErr != nil {
				return execErr
			}
		} else if slices.Contains(SixOOItems, item.Name()) {
			_, execErr := exec.Command("chmod", "-R", "600", item.Name()).Output()
			if execErr != nil {
				return execErr
			}
		}
	}

	return nil
}

func EnsureTar() (error) {
	_, err := exec.Command("tar", "--version").Output()

	if err != nil {
		return err
	}

	return nil
}

func Copy(src string, dest string) (error) {
	source, sourceErr := os.Open(src)

	if sourceErr != nil {
		return sourceErr
	}

	defer source.Close()

	sourceStat, sourceStatErr := source.Stat()

	if sourceStatErr != nil {
		return sourceStatErr
	}

	destination, destinationErr := os.Create(dest)

	if destinationErr != nil {
		return destinationErr
	}

	defer destination.Close()

	_, copyErr := io.Copy(destination, source)

	if copyErr != nil {
		return copyErr
	}

	chmodErr := os.Chmod(dest, sourceStat.Mode())

	if chmodErr != nil {
		return chmodErr
	}

	return nil
}
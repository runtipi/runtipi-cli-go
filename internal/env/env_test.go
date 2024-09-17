package env_test

import (
	"os"
	"path"
	"runtime"
	"strconv"
	"testing"

	"runtipi-cli-go/internal/constants"
	"runtipi-cli-go/internal/env"

	"github.com/spf13/viper"
)

func init() {
	// Change back to the root folder
	os.Chdir("../..")
}

// It should generate an environment file and settings file if they doesn't exist
func TestEnvGen(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Get env and settings paths
	statePath := path.Join(rootFolder, "state")
	envPath := path.Join(rootFolder, ".env")
	settingsPath := path.Join(statePath, "settings.json")
	envLocalPath := path.Join(rootFolder, ".env.local")

	// Delete env, env local and settigs files
	os.Remove(statePath)
	os.Remove(envPath)
	os.Remove(settingsPath)
	os.Remove(envLocalPath)

	// Create state path
	os.Mkdir(statePath, 0755)

	// Generate env
	env.GenerateEnv("")

	// Check if files did got created
	if _, envCheckErr := os.Stat(envPath); envCheckErr != nil {
		t.Fatalf("Env file didn't get created!")
	}

	if _, settingsCheckErr := os.Stat(settingsPath); settingsCheckErr != nil {
		t.Fatalf("Settings file didn't get created!")
	}
}

// Check that the generated env file is correct
func TestEnvIsCorrect(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Get env and settings paths
	statePath := path.Join(rootFolder, "state")
	envPath := path.Join(rootFolder, ".env")
	settingsPath := path.Join(statePath, "settings.json")
	envLocalPath := path.Join(rootFolder, ".env.local")

	// Delete env, env local and settigs files
	os.Remove(statePath)
	os.Remove(envPath)
	os.Remove(settingsPath)
	os.Remove(envLocalPath)

	// Create state path
	os.Mkdir(statePath, 0755)

	// Generate env
	env.GenerateEnv("")

	// Read env using viper
	viper.SetConfigType("env")
	viper.SetConfigFile(envPath)
	viper.ReadInConfig()

	// Check arch
	arch := viper.GetString("ARCHITECTURE")
	if arch != runtime.GOARCH {
		t.Fatalf("Architecture should be %s, got %s\n", runtime.GOARCH, arch)
	}

	// Check domain
	domain := viper.GetString("DOMAIN")
	if domain != "example.com" {
		t.Fatalf("Domain should be example.com, got %s\n", domain)
	}

	// Check local domain
	localDomain := viper.GetString("LOCAL_DOMAIN")
	if localDomain != "tipi.local" {
		t.Fatalf("Local domain should be tipi.local, got %s\n", localDomain)
	}

	// Check nginx port
	nginxPort := viper.GetString("NGINX_PORT")
	if nginxPort != "80" {
		t.Fatalf("Nginx port should be 80, got %s\n", nginxPort)
	}

	// Check nginx ssl port
	nginxSslPort := viper.GetString("NGINX_PORT_SSL")
	if nginxSslPort != "443" {
		t.Fatalf("Nginx ssl port should be 443, got %s\n", nginxSslPort)
	}

	// Check postgres host
	postgresHost := viper.GetString("POSTGRES_HOST")
	if postgresHost != "runtipi-db" {
		t.Fatalf("Postgres host should be runtipi-db, got %s\n", postgresHost)
	}

	// Check postgres password
	postgresPassword := viper.GetString("POSTGRES_PASSWORD")
	if len(postgresPassword) != 64 {
		t.Fatalf("Postgres password should be 64 chars long, got %s chars\n", strconv.Itoa(len(postgresPassword)))
	}

	// Check postgres username
	postgresUsername := viper.GetString("POSTGRES_USERNAME")
	if postgresUsername != "tipi" {
		t.Fatalf("Postgres username should be tipi, got %s\n", postgresUsername)
	}

	// Check redis host
	redisHost := viper.GetString("REDIS_HOST")
	if redisHost != "runtipi-redis" {
		t.Fatalf("Redis host should be runtipi-redis, got %s\n", redisHost)
	}

	// Check root folder
	rootFolderEnv := viper.GetString("ROOT_FOLDER_HOST")
	if rootFolderEnv != rootFolder {
		t.Fatalf("Root folder should be %s, got %s\n", rootFolder, rootFolderEnv)
	}

	// Check app data path
	appDataPath := viper.GetString("RUNTIPI_APP_DATA_PATH")
	if appDataPath != rootFolder {
		t.Fatalf("App data path should be %s, got %s\n", rootFolder, appDataPath)
	}

	// Check version
	version := viper.GetString("TIPI_VERSION")
	if version != constants.RuntipiVersion {
		t.Fatalf("Tipi version should be %s, got %s\n", constants.RuntipiVersion, version)
	}
}

// Check that .env.local is being recognized and used
func TestLocalEnv(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Get env and settings paths
	statePath := path.Join(rootFolder, "state")
	envPath := path.Join(rootFolder, ".env")
	settingsPath := path.Join(statePath, "settings.json")
	envLocalPath := path.Join(rootFolder, ".env.local")

	// Delete env, env local and settigs files
	os.Remove(statePath)
	os.Remove(envPath)
	os.Remove(settingsPath)
	os.Remove(envLocalPath)

	// Create state path
	os.Mkdir(statePath, 0755)

	// Set testing version in .env.local
	os.WriteFile(envLocalPath, []byte("TIPI_VERSION=testing\n"), 0664)

	// Generate env
	env.GenerateEnv(envLocalPath)

	// Read env using viper
	viper.SetConfigType("env")
	viper.SetConfigFile(envPath)
	viper.ReadInConfig()

	// Check version override
	version := viper.GetString("TIPI_VERSION")
	if version != "testing" {
		t.Fatalf("Tipi version should be testing, got %s\n", version)
	}
}

func TestSettings(t *testing.T) {
	// Get root folder
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		t.Fatalf("Failed to get root folder, error: %s\n", osErr)
	}

	// Get env and settings paths
	statePath := path.Join(rootFolder, "state")
	envPath := path.Join(rootFolder, ".env")
	settingsPath := path.Join(statePath, "settings.json")
	envLocalPath := path.Join(rootFolder, ".env.local")

	// Delete env, env local and settigs files
	os.Remove(statePath)
	os.Remove(envPath)
	os.Remove(envLocalPath)

	// Create state path
	os.Mkdir(statePath, 0755)

	// Set custom domain in settings.json
	os.WriteFile(settingsPath, []byte(`{"domain":"test.com"}`), 0664)

	// Generate env
	env.GenerateEnv("")

	// Read env using viper
	viper.SetConfigType("env")
	viper.SetConfigFile(envPath)
	viper.ReadInConfig()

	// Check if domain got changed
	domain := viper.GetString("DOMAIN")
	if domain != "test.com" {
		t.Fatalf("Domain should be test.com, got %s\n", domain)
	}
}
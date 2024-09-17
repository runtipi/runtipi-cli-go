package env

import (
	"encoding/json"
	"os"
	"path"

	"runtipi-cli-go/internal/constants"
	"runtipi-cli-go/internal/schemas"
	"runtipi-cli-go/internal/seed"
	"runtipi-cli-go/internal/system"

	"github.com/spf13/viper"
)

func GenerateEnv(envFile string) (error) {
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		return osErr
	}

	envPath := path.Join(rootFolder, ".env")

	if _, err := os.Stat(envPath); err != nil {
		os.WriteFile(envPath, []byte(""), 0644)
	}

	settingsPath := path.Join(rootFolder, "state", "settings.json")

	if _, err := os.Stat(settingsPath); err != nil {
		os.WriteFile(settingsPath, []byte("{}"), 0644)
	}

	seedGenErr := seed.GenerateSeed(rootFolder)

	if seedGenErr != nil {
		return seedGenErr
	}

	defaultViper := viper.New()

	defaultViper.SetConfigType("env")
	defaultViper.SetConfigFile(envPath)

	settingsFile, settingsFileErr := os.ReadFile(settingsPath)

	if settingsFileErr != nil {
		return settingsFileErr
	}

	var settings schemas.Settings

	err := json.Unmarshal(settingsFile, &settings)

	if err != nil {
		return err
	}

	version := constants.Version

	seed, seedErr := system.GetSeed(rootFolder)

	if seedErr != nil {
		return seedErr
	}

	defaultViper.Set("POSTGRES_PASSWORD", system.DeriveEntopy("postgres_password", seed))
	defaultViper.Set("REDIS_PASSWORD", system.DeriveEntopy("redis_password", seed))
	
	appDataPath := ""

	if settings.AppDataPath == "" {
		appDataPath = settings.StoragePath
	}

	if appDataPath != "" {
		if _, err := os.Stat(appDataPath); err != nil {
			return err
		}
	}

	internalIp, internalIpErr := system.GetInternalIp()

	if internalIpErr != nil {
		return internalIpErr
	}

	if settings.InternalIp != "" {
		defaultViper.Set("INTERNAL_IP", settings.InternalIp)
	} else {
		defaultViper.Set("INTERNAL_IP", internalIp)	
	}

	defaultViper.Set("ARCHITECTURE", system.GetArch())
	defaultViper.Set("TIPI_VERSION", version)
	defaultViper.Set("ROOT_FOLDER_HOST", rootFolder)

	if settings.Port != nil {
		defaultViper.Set("NGINX_PORT", settings.Port)
	} else {
		defaultViper.Set("NGINX_PORT", 80)
	}

	if settings.SSLPort != nil {
		defaultViper.Set("NGINX_PORT_SSL", settings.SSLPort)
	} else {
		defaultViper.Set("NGINX_PORT_SSL", 443)
	}

	if appDataPath != "" {
		defaultViper.Set("RUNTIPI_APP_DATA_PATH", appDataPath)
	} else {
		defaultViper.Set("RUNTIPI_APP_DATA_PATH", rootFolder)
	}

	defaultViper.Set("POSTGRES_HOST", "runtipi-db")

	if settings.PostgresPort != "" {
		defaultViper.Set("POSTGRES_PORT", settings.PostgresPort)
	} else {
		defaultViper.Set("POSTGRES_PORT", 5432)
	}

	defaultViper.Set("POSTGRES_USERNAME", "tipi")

	defaultViper.Set("REDIS_HOST", "runtipi-redis")

	if settings.Domain != "" {
		defaultViper.Set("DOMAIN", settings.Domain)
	} else {
		defaultViper.Set("DOMAIN", "example.com")
	}

	if settings.LocalDomain != "" {
		defaultViper.Set("LOCAL_DOMAIN", settings.LocalDomain)
	} else {
		defaultViper.Set("LOCAL_DOMAIN", "tipi.local")
	}

	if _, err := os.Stat(envFile); err == nil {
		customViper := viper.New()
		customViper.SetConfigType("env")
		customViper.SetConfigFile(envFile)
		customViper.ReadInConfig()
		overrideKeys := customViper.AllKeys()
		for _, key := range overrideKeys {
			defaultViper.Set(key, customViper.Get(key))
		}
	}

	defaultViper.WriteConfigAs(envPath)
	return nil
}

func GetEnvValue(value string) (string, error) {
	rootFolder, osErr := os.Getwd()
	
	if osErr != nil {
		return "", osErr
	}

	viper.SetConfigType("env")
	viper.SetConfigFile(path.Join(rootFolder, ".env"))
	viper.ReadInConfig()

	return viper.GetString(value), nil
}
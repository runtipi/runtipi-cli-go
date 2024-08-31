package schemas

type Settings struct {
	InternalIp string `json:"internalIp"`
	Port any `json:"port"`
	SSLPort any `json:"sslPort"`

	// Deprecated
	StoragePath string `json:"storagePath"`
	AppDataPath string `json:"appDataPath"`

	PostgresPort any `json:"postgresPort"`
	Domain string `json:"domain"`
	LocalDomain string `json:"localDomain"`
}
package schemas

type Settings struct {
	// Traefik config
	Port any `json:"port"`
	SSLPort any `json:"sslPort"`
	InternalIp string `json:"internalIp"`
	Domain string `json:"domain"`
	LocalDomain string `json:"localDomain"`

	// Deprecated
	StoragePath string `json:"storagePath"`

	// App data path (storagePath)
	AppDataPath string `json:"appDataPath"`

	// Postgres
	PostgresPort any `json:"postgresPort"`
}

type SystemStatus struct {
	// Disk
	DiskUsed float64 `json:"diskUsed"`
	DiskSize float64 `json:"diskSize"`
	PercentUsed float64 `json:"percentUsed"`

	// Cpu
	CpuLoad float64 `json:"cpuLoad"`

	// Memory
	MemoryTotal int `json:"memoryTotal"`
	PercentUsedMemory float64 `json:"percentUsedMemory"`
}

type SystemStatusApi struct {
	// Data
	Data SystemStatus `json:"data"`
}

type GithubRelease struct {
	// Tag Name
	TagName string `json:"tag_name"`

	// Status Code
	Status string `json:"status"`
}
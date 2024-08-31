package constants

import _ "embed"

//go:embed assets/VERSION
var Version string

//go:embed assets/docker-compose.yml
var Compose string
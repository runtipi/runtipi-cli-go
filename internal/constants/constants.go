package constants

import (
	_ "embed"

	"github.com/fatih/color"
)

//go:embed assets/RUNTIPI_VERSION
var RuntipiVersion string

//go:embed assets/CLI_VERSION
var CliVersion string

//go:embed assets/docker-compose.yml
var Compose string

//go:embed assets/neofetch.txt
var Neofetch string

// Colors
var Blue = color.New(color.FgBlue).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
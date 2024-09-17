package constants

import (
	_ "embed"

	"github.com/fatih/color"
)

//go:embed assets/VERSION
var Version string

//go:embed assets/docker-compose.yml
var Compose string

//go:embed assets/neofetch.txt
var Neofetch string

// Colors
var Blue = color.New(color.FgBlue).SprintFunc()
var Green = color.New(color.FgGreen).SprintFunc()
var Red = color.New(color.FgRed).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
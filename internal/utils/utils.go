package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintError(message string) {
	color.Set(color.FgRed)
	fmt.Print("✗ ")
	color.Unset()
	fmt.Println(message)
}

func PrintSuccess(message string) {
	color.Set(color.FgGreen)
	fmt.Print("✓ ")
	color.Unset()
	fmt.Println(message)
}
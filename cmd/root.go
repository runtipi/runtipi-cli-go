package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var noPermissions bool
var envFile string

var rootCmd = &cobra.Command{
	Use:   "runtipi-cli-go",
	Short: "A reimplementation of the runtipi cli in go",
	Long: "A simpler and faster reimplementation of the runtipi cli in the go programming language",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func init() {
	fmt.Println("Welcome to Runtipi CLI in Go ✨")
}
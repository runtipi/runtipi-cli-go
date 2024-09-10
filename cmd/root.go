package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/cmd/app"
	"github.com/steveiliop56/runtipi-cli-go/cmd/system"
)

var noPermissions bool
var envFile string
var showLogs bool

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
	rootCmd.AddCommand(app.AppCmd())
	rootCmd.AddCommand(system.SystemCmd())
}
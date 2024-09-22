package cmd

import (
	"fmt"
	"os"

	"runtipi-cli-go/cmd/app"
	"runtipi-cli-go/cmd/system"

	"github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
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
	coloredcobra.Init(&coloredcobra.Config{
        RootCmd:       rootCmd,
        Headings:      coloredcobra.Yellow + coloredcobra.Bold + coloredcobra.Underline,
        Commands:      coloredcobra.Red + coloredcobra.Bold,
        Example:       coloredcobra.Italic,
        ExecName:      coloredcobra.Bold,
        Flags:         coloredcobra.Bold,
    })
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}


func init() {
	fmt.Print("Welcome to Runtipi CLI in Go ✨\n\n")
	rootCmd.AddCommand(app.AppCmd())
	rootCmd.AddCommand(system.SystemCmd())
}
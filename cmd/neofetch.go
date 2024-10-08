package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"runtipi-cli-go/internal/constants"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(tipiFetchCmd)
}

var tipiFetchCmd = &cobra.Command{
	Use: "neofetch",
	Short: "Print neofetch using the tipi logo (requires neofetch)",
	Long: "This command prints print system info using neofetch but with tipi's logo as ascii art (requires neofetch)",
	Run: func(cmd *cobra.Command, args []string) {
		// Define ascii path
		asciiPath := path.Join("/", "tmp", "tipi-fetch-ascii.txt")
		
		// Write temp ascii file
		if err := os.WriteFile(asciiPath, []byte(constants.Neofetch), 0644); err != nil {
			fmt.Printf("%s Failed to write neofetch ascii\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// Run the neofetch command
		out, err := exec.Command("neofetch", "--ascii", "--ascii_colors", "1", "11", "8", "9", "--source", asciiPath).Output()

		// Check for errors
		if err != nil {
			fmt.Printf("%s Failed to run neofetch command\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// Delete temp file
		if err := os.RemoveAll(asciiPath); err != nil {
			fmt.Printf("%s Failed to remove neofetch ascii\n", constants.Red("✗"))
			fmt.Printf("Error: %s\n", err)
			os.Exit(1)
		}

		// Print output
		fmt.Printf("\n%s", out)
	},
}
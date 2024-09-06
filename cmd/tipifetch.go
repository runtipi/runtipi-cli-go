package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/steveiliop56/runtipi-cli-go/internal/constants"
)

func init() {
	rootCmd.AddCommand(tipiFetchCmd)
}

var tipiFetchCmd = &cobra.Command{
	Use: "tipifetch",
	Short: "Print neofetch using the tipi logo (requires neofetch)",
	Long: "This command prints print system info using neofetch but with tipi's logo as ascii art (requires neofetch)",
	Run: func(cmd *cobra.Command, args []string) {
		// Define ascii path
		asciiPath := path.Join("/", "tmp", "tipi-fetch-ascii.txt")
		
		// Write temp ascii file
		if err := os.WriteFile(asciiPath, []byte(constants.Neofetch), 0644); err != nil {
			color.Set(color.FgRed)
			fmt.Print("✗ ")
			color.Unset()
			fmt.Printf("Failed write neofetch ascii art, error: %s\n", err)
			return
		}

		// Run the neofetch command
		out, err := exec.Command("neofetch", "--ascii", "--ascii_colors", "1", "11", "8", "9", "--source", asciiPath).Output()

		// Check for errors
		if err != nil {
			color.Set(color.FgRed)
			fmt.Print("✗ ")
			color.Unset()
			fmt.Printf("Failed to run neofetch command, error: %s\n", err)
			return
		}

		// Delete temp file
		if err := os.Remove(asciiPath); err != nil {
			color.Set(color.FgRed)
			fmt.Print("✗ ")
			color.Unset()
			fmt.Printf("Failed to remove temp neofetch ascii art, error: %s\n", err)
			return
		}

		// Print output
		fmt.Printf("\n%s", out)
	},
}
package spinner

import (
	"fmt"
	"time"

	"runtipi-cli-go/internal/constants"

	"github.com/briandowns/spinner"
)

var s = spinner.New([]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}, 50*time.Millisecond)

func Start() {
	s.Color("green")
	s.Start()
}

func Stop() {
	s.Stop()
}

func SetMessage(message string) {
	s.Suffix = fmt.Sprintf(" %s", message)
}

func Succeed(message string) {
	s.Stop()
	fmt.Printf("%s %s\n", constants.Green("✓"), message)
	s.Start()
}

func Fail(message string) {
	s.Stop()
	fmt.Printf("%s %s\n", constants.Red("✗"), message)
	s.Start()
}

func Update(message string) {
	s.Stop()
	fmt.Printf("%s %s\n", constants.Blue("↑"), message)
	s.Start()
}

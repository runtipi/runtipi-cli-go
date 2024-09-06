package spinner

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/steveiliop56/runtipi-cli-go/internal/utils"
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
	s.Suffix = " " + message
}

func Succeed(message string) {
	s.Stop()
	utils.PrintSuccess(message)
	s.Start()
}

func Fail(message string) {
	s.Stop()
	utils.PrintError(message)
	s.Start()
}
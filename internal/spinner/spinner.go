package spinner

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
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
	color.Set(color.FgGreen)
	fmt.Print("✓ ")
	color.Unset()
	fmt.Println(message)
	s.Start()
}

func Fail(message string) {
	s.Stop()
	color.Set(color.FgRed)
	fmt.Print("✗ ")
	color.Unset()
	fmt.Println(message)
	s.Start()
}
package main

import (
	"time"

	mainmenu "github.com/mt1976/crt/actions/mainmenu"
	startup "github.com/mt1976/crt/start"
	terminal "github.com/mt1976/crt/support"
	config "github.com/mt1976/crt/support/config"
)

// config is used to store configuration settings for the program, including terminal
// width and height.
//

// Main is the entry point for the program.
func main() {

	C := config.Configuration

	// create a new instance of the Crt
	crt := terminal.NewWithSize(C.TerminalWidth, C.TerminalHeight)
	// set the terminal size
	//crt.SetTerminalSize(config.term_width, config.term_height)

	// start a timer
	start := time.Now()

	// run the startup sequence
	crt.SetDelayInSec(C.Delay)
	startup.Run(&crt)
	crt.ResetDelay()
	// run the main menu
	mainmenu.Run(&crt)

	// stop the timer
	elapsed := time.Since(start)
	// output the elapsed time
	crt.Shout(crt.Bold("DONE") + " " + elapsed.String())

}

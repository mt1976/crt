package main

import (
	"time"

	lang "github.com/mt1976/crt/language"
	term "github.com/mt1976/crt/support"
	conf "github.com/mt1976/crt/support/config"
)

// config is used to store configuration settings for the program, including terminal
// width and height.
//

// Main is the entry point for the program.
func main() {

	C := conf.Configuration

	// create a new instance of the Crt
	crt := term.NewWithSize(C.TerminalWidth, C.TerminalHeight)
	// set the terminal size
	//crt.SetTerminalSize(config.term_width, config.term_height)

	// start a timer
	start := time.Now()

	// run the startup sequence
	crt.SetDelayInSec(C.Delay)
	crt.ResetDelay()
	//godump.Dump(crt)
	//os.Exit(0)
	// run the main menu

	// stop the timer
	elapsed := time.Since(start)
	// output the elapsed time
	crt.Shout(crt.Bold(lang.TxtDone) + lang.Space + elapsed.String())

}

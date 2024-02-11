package main

import (
	"time"

	"github.com/mt1976/admin_me/actions/mainmenu"
	startup "github.com/mt1976/admin_me/start"
	terminal "github.com/mt1976/admin_me/support"
)

// config is used to store configuration settings for the program, including terminal
// width and height.
//
// @property {int} term_width - The term_width property is used to store the width of the terminal
// window.
// @property {int} term_height - The term_height property is used to store the height of the terminal
// window in characters.
type config struct {
	// The config struct is used to store the configuration settings for the program. It has the
	// following fields:
	term_width  int `pkl:"term_width"`
	term_height int `pkl:"term_height"`
}

// Main is the entry point for the program.
func main() {

	// create a new instance of the config struct
	config := config{}
	// set the terminal width and height
	config.term_width = 80
	config.term_height = 20

	// create a new instance of the Crt
	crt := terminal.New()
	// set the terminal size
	crt.SetTerminalSize(config.term_width, config.term_height)

	// start a timer
	start := time.Now()

	// run the startup sequence
	startup.Run(&crt)

	// run the main menu
	mainmenu.Run(&crt)

	// stop the timer
	elapsed := time.Since(start)
	// output the elapsed time
	crt.Shout(crt.Bold("DONE") + " " + elapsed.String())

}

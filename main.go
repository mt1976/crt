package main

import (
	"time"

	menu "github.com/mt1976/admin_me/menu"
	startup "github.com/mt1976/admin_me/start"
	terminal "github.com/mt1976/admin_me/support"
)

type config struct {
	// The `config` struct is used to store the configuration settings for the program. It has the
	// following fields:
	term_width  int `pkl:"term_width"`
	term_height int `pkl:"term_height"`
}

func main() {

	config := config{}
	config.term_width = 80
	config.term_height = 20
	// define a new instance of the Crt
	crt := terminal.NewCrt()
	crt.SetTerminalSize(config.term_width, config.term_height)
	//xx, yy := crt.GetTerminalSize()
	//crt.Println(fmt.Sprintf("Terminal Size: %d*%d", xx, yy))

	// convert the slice to a string, separated by commas
	//

	start := time.Now()

	// //Get Current Path
	// path, err := os.Getwd()
	// if err != nil {

	//crt.Break()

	startup.Run(&crt)

	menu.Run(&crt)
	//	crt.Shout("Unknown action '" + crt.Bold(commandAction) + "'")

	elapsed := time.Since(start)
	crt.Shout(crt.Bold("DONE") + " " + elapsed.String())
	// `elapsed` is a variable of type `time.Duration` that stores
	// the amount of time that has passed since the `start` time. It
	// is used to calculate the total time taken to execute the
	// program

}

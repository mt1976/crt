package main

import (
	"time"

	startup "github.com/mt1976/admin_me/start"
	terminal "github.com/mt1976/admin_me/support"
)

func main() {

	// define a new instance of the Crt
	crt := terminal.NewCrt()

	// convert the slice to a string, separated by commas
	//

	start := time.Now()

	// //Get Current Path
	// path, err := os.Getwd()
	// if err != nil {

	crt.Break()

	startup.Run(&crt)
	//	crt.Shout("Unknown action '" + crt.Bold(commandAction) + "'")

	elapsed := time.Since(start)
	crt.Shout(crt.Bold("DONE") + " " + elapsed.String())
	// `elapsed` is a variable of type `time.Duration` that stores
	// the amount of time that has passed since the `start` time. It
	// is used to calculate the total time taken to execute the
	// program.
}

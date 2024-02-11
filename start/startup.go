package start

import (
	"fmt"

	"github.com/mt1976/admin_me/support"
)

// Run initializes the terminal and runs the main loop.
func Run(crt *support.Crt) {
	// Clear the terminal screen.
	crt.Clear()

	// Display the banner.
	crt.Banner("Starting...")

	// Print a message with the current delay value.
	//crt.Print(fmt.Sprintf("Delay in seconds: %v", crt.GetDelayInSec()))

	// Set the delay to 0.5 seconds.
	//crt.SetDelayInSec(0.5)

	// Print a message.
	crt.Print("Starting Terminal...")

	// Sleep for 2 seconds.
	// Sleep for 250 milliseconds.
	//crt.Sleep(250)

	// Print a message.
	crt.Print("Self Testing...")

	// Print the current date and time.
	crt.Print("Current Date: " + support.DateString())
	crt.Print("Current Time: " + support.TimeString())

	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.Sleep(500)

	// Print a message.
	crt.Print("Please wait...")

	// Check if the terminal has a baud rate set.
	if !crt.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf("Baud Rate Set to %v kbps", crt.Baud())
		crt.Print(msg)
	}

	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.Sleep(500)

	// Print a message.
	crt.Print("Connecting...")

	// Generate a random IP address and port number.
	ip := support.RandomIP()
	port := support.RandomPort()

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf("Dialing... %v:%v", ip, port)
	crt.Print(msg)

	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.DelayIt(500)

	// Print a message.
	crt.Print("Connected.")
}

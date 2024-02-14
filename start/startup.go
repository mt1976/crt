package start

import (
	"fmt"

	"github.com/mt1976/crt/support"
)

// Run initializes the terminal and runs the main loop.
func Run(crt *support.Crt) {
	// Clear the terminal screen.
	crt.Clear()

	// Display the banner.
	crt.Banner(startingText)

	// Print a message with the current delay value.
	//crt.Print(fmt.Sprintf("Delay in seconds: %v", crt.GetDelayInSec()))

	// Set the delay to 0.5 seconds.
	//crt.SetDelayInSec(0.5)

	// Print a message.
	crt.Print(startingTerminalText + newline)

	// Sleep for 2 seconds.
	// Sleep for 250 milliseconds.
	//crt.Sleep(250)

	// Print a message.
	crt.Print(selfTestingText + newline)
	oldDelay := crt.Delay()
	//fmt.Println("Old Delay: ", oldDelay)
	crt.SetDelayInSec(0.25)
	crt.Print(selfTestingText + "Complete" + newline)
	// Print the current date and time.
	crt.SetDelayInMs(oldDelay)
	crt.Print(currentDateText + support.DateString() + newline)
	crt.Print(currentTimeText + support.TimeString() + newline)

	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.Sleep(500)

	// Print a message.
	crt.Print(pleaseWaitText + newline)

	// Check if the terminal has a baud rate set.
	if !crt.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf(baudRateText, crt.Baud())
		crt.Print(msg + newline)
	}

	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.Sleep(500)

	// Print a message.
	crt.Print(connectingText + newline)

	// Generate a random IP address and port number.
	//ip := support.RandomIP()
	//port := support.RandomPort()

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf(dialingText, support.RandomIP(), support.RandomPort())
	crt.Print(msg + newline)
	crt.SetDelayInSec(support.RandomFloat(1, 5))

	if support.CoinToss() {
		crt.Print(dialingFailedText + newline)
		// Print a message with the IP address and port number.
		crt.ResetDelay()
		msg := fmt.Sprintf(dialingText, support.RandomIP(), support.RandomPort())
		crt.Print(msg + newline)
		crt.SetDelayInSec(support.RandomFloat(1, 5))
	}
	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.DelayIt(500)

	// Print a message.
	crt.Print(connectedText + newline)
	crt.SetDelayInMs(oldDelay)
}

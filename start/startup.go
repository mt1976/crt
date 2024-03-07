package start

import (
	"fmt"

	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
)

// Run initializes the terminal and runs the main loop.

var C = config.Configuration

func Run(crt *support.Crt) {
	// Clear the terminal screen.
	crt.Clear()

	// Display the banner.
	crt.Banner(TxtStarting)

	// Print a message with the current delay value.
	//crt.Print(fmt.Sprintf("Delay in seconds: %v", crt.GetDelayInSec()))

	// Set the delay to 0.5 seconds.
	//crt.SetDelayInSec(0.5)

	// Print a message.
	crt.Print(TxtStartingTerminal + SymNewline)

	// Sleep for 2 seconds.
	// Sleep for 250 milliseconds.
	//crt.Sleep(250)

	// Print a message.
	crt.Print(TxtSelfTesting + SymNewline)
	oldDelay := crt.Delay()
	//fmt.Println("Old Delay: ", oldDelay)
	crt.SetDelayInSec(0.25)
	crt.Print(TxtSelfTesting + TxtComplete + SymNewline)
	// Print the current date and time.
	crt.SetDelayInMs(oldDelay)
	crt.Print(TxtCurrentDate + support.DateString() + SymNewline)
	crt.Print(TxtCurrentTime + support.TimeString() + SymNewline)

	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.Sleep(500)

	// Print a message.
	crt.Print(TxtPleaseWait + SymNewline)

	// Check if the terminal has a baud rate set.
	if !crt.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf(TxtBaudRate, crt.Baud())
		crt.Print(msg + SymNewline)
	}

	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.Sleep(500)

	// Print a message.
	crt.Print(TxtConnecting + SymNewline)

	// Generate a random IP address and port number.
	//ip := support.RandomIP()
	//port := support.RandomPort()

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf(TxtDialing, support.RandomIP(), support.RandomPort())
	crt.Print(msg + SymNewline)
	if !C.Debug {
		crt.SetDelayInSec(support.RandomFloat(1, 5))
	}
	if support.CoinToss() && !C.Debug {
		crt.Print(TxtDialingFailed + SymNewline)
		// Print a message with the IP address and port number.
		crt.ResetDelay()
		msg := fmt.Sprintf(TxtDialing, support.RandomIP(), support.RandomPort())
		crt.Print(msg + SymNewline)
		crt.SetDelayInSec(support.RandomFloat(1, 5))
	}
	// Sleep for 2 seconds.
	// Sleep for 500 milliseconds.
	//crt.DelayIt(500)

	// Print a message.
	crt.Print(TxtConnected + SymNewline)
	crt.SetDelayInMs(oldDelay)
}

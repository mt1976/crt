package start

import (
	"fmt"

	t "github.com/mt1976/crt/language"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
)

// Run initializes the terminal and runs the main loop.

var C = config.Configuration

func Run(crt *support.Crt) {
	// Clear the terminal screen.
	crt.Clear()

	// Display the banner.
	crt.Banner(t.TxtStarting)

	// Print a message.
	crt.Print(t.TxtStartingTerminal + t.Newline)

	// Print a message.
	crt.Print(t.TxtSelfTesting + t.Newline)
	oldDelay := crt.Delay()
	//fmt.Println("Old Delay: ", oldDelay)
	crt.SetDelayInSec(0.25)
	crt.Print(t.TxtSelfTesting + t.TxtComplete + t.Newline)
	// Print the current date and time.
	crt.SetDelayInMs(oldDelay)
	crt.Print(t.TxtCurrentDate + support.DateString() + t.Newline)
	crt.Print(t.TxtCurrentTime + support.TimeString() + t.Newline)

	// Print a message.
	crt.Print(t.TxtPleaseWait + t.Newline)

	// Check if the terminal has a baud rate set.
	if !crt.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf(t.TxtBaudRate, crt.Baud())
		crt.Print(msg + t.Newline)
	}

	// Print a message.
	crt.Print(t.TxtConnecting + t.Newline)

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf(t.TxtDialing, support.RandomIP(), support.RandomPort())
	crt.Print(msg + t.Newline)
	if !C.Debug {
		crt.SetDelayInSec(support.RandomFloat(1, 5))
	}
	if support.CoinToss() && !C.Debug {
		crt.Print(t.TxtDialingFailed + t.Newline)
		// Print a message with the IP address and port number.
		crt.ResetDelay()
		msg := fmt.Sprintf(t.TxtDialing, support.RandomIP(), support.RandomPort())
		crt.Print(msg + t.Newline)
		crt.SetDelayInSec(support.RandomFloat(1, 5))
	}

	// Print a message.
	crt.Print(t.TxtConnected + t.Newline)
	crt.SetDelayInMs(oldDelay)
}

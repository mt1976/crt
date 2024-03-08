package start

import (
	"fmt"

	t "github.com/mt1976/crt/language"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
)

// Run initializes the terminal and runs the main loop.

var C = config.Configuration

// Run initializes the terminal and runs the main loop.
func Run(crt *support.Crt) {
	// Clear the terminal screen.
	crt.Clear()

	// Display the banner.
	crt.Banner(t.TxtStarting)

	// Print a message.
	crt.Print(t.TxtStartingTerminal + t.SymNewline)

	// Print a message.
	crt.Print(t.TxtSelfTesting + t.SymNewline)
	oldDelay := crt.Delay()
	//fmt.Println("Old Delay: ", oldDelay)
	crt.SetDelayInSec(0.25)
	crt.Print(t.TxtSelfTesting + t.TxtComplete + t.SymNewline)
	// Print the current date and time.
	crt.SetDelayInMs(oldDelay)
	crt.Print(t.TxtCurrentDate + support.DateString() + t.SymNewline)
	crt.Print(t.TxtCurrentTime + support.TimeString() + t.SymNewline)

	// Print a message.
	crt.Print(t.TxtPleaseWait + t.SymNewline)

	// Check if the terminal has a baud rate set.
	if !crt.NoBaudRate() {
		// Print a message with the current baud rate.
		msg := fmt.Sprintf(t.TxtBaudRate, crt.Baud())
		crt.Print(msg + t.SymNewline)
	}

	// Print a message.
	crt.Print(t.TxtConnecting + t.SymNewline)

	// Print a message with the IP address and port number.
	msg := fmt.Sprintf(t.TxtDialing, support.RandomIP(), support.RandomPort())
	crt.Print(msg + t.SymNewline)
	if !C.Debug {
		crt.SetDelayInSec(support.RandomFloat(1, 5))
	}
	if support.CoinToss() && !C.Debug {
		crt.Print(t.ErrorMessageConnectionFailed + t.SymNewline)
		// Print a message with the IP address and port number.
		crt.ResetDelay()
		msg := fmt.Sprintf(t.TxtDialing, support.RandomIP(), support.RandomPort())
		crt.Print(msg + t.SymNewline)
		crt.SetDelayInSec(support.RandomFloat(1, 5))
	}

	// Print a message.
	crt.Print(t.TxtConnected + t.SymNewline)
	crt.SetDelayInMs(oldDelay)
}

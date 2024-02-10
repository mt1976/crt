package start

import (
	"fmt"

	"github.com/mt1976/admin_me/support"
)

func Run(crt *support.Crt) {
	crt.Clear()
	//crt.SetDelayInSec(0.25) // Set delay in milliseconds
	crt.Banner("Starting...")
	//crt.Print(fmt.Sprintf("Delay in seconds: %v", crt.GetDelayInSec()))
	//crt.SetDelayInSec(0.5) // Set delay in milliseconds
	crt.Print("Starting Terminal...")
	// sleep 2
	// echo "  Self Testing...";
	crt.Print("Self Testing...")
	crt.Print("Current Date: " + support.DateString())
	crt.Print("Current Time: " + support.TimeString())
	// sleep 2
	// echo "  Please wait...";
	crt.Print("Please wait...")
	msg := ""
	if !crt.NoBaudRate() {
		msg = fmt.Sprintf("Baud Rate Set to %v kbps", crt.Baud())
		crt.Print(msg)
	}
	// sleep 2
	// echo "  Connecting...";
	ip := support.RandomIP()
	port := support.RandomPort()
	msg = fmt.Sprintf("Dialing... %v:%v", ip, port)
	crt.Print(msg)
	// sleep 2
	// echo "  Connected.";
	crt.Print("Connected.")
	//crt.Break()
	//crt.Clear()
	//crt.Close()

}

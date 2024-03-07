package skynews

import (
	"github.com/mt1976/crt/support"
	config "github.com/mt1976/crt/support/config"
	page "github.com/mt1976/crt/support/page"
)

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(crt *support.Crt) {

	C := config.Configuration

	crt.Clear()
	//crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Main Menu")
	m := page.New(TxtMenuTitle)
	c := 0
	c++
	m.AddOption(c, TxtTransmission, C.TransmissionURI, "")
	c++
	m.AddOption(c, TxtQTorrent, C.QTorrentURI, "")
	c++

	m.AddAction(page.TxtQuit)

	action, nextLevel := m.Display(crt)

	//log.Println("Action: ", action)
	//log.Println("Next Level: ", nextLevel)
	//pause
	//crt.SetDelayInMin(1)
	//crt.DelayIt()

	if action == page.TxtQuit {
		return
	}

	if support.IsInt(action) {
		switch action {
		case "1":
			Trans(crt, nextLevel.AlternateID, nextLevel.Title)
			action = ""
		case "2":
			//QTor(crt, nextLevel.AlternateID, nextLevel.Title)
			action = ""
		default:
			crt.InputError(page.ErrInvalidAction + "'" + action + "'")
			action = ""
		}
	}
}

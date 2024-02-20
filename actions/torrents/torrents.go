package skynews

import (
	"github.com/mt1976/crt/support"
	config "github.com/mt1976/crt/support/config"
	"github.com/mt1976/crt/support/menu"
)

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
func Run(crt *support.Crt) {

	C := config.Configuration

	crt.Clear()
	//crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Main Menu")
	m := menu.New(menuTitleText)
	c := 0
	c++
	m.Add(c, serviceTransText, C.TransmissionURI, "")
	c++
	m.Add(c, serviceQTorText, serviceQTorURI, "")
	c++

	m.AddAction(menu.Quit)

	ok := false
	for !ok {
		action, nextLevel := m.Display(crt)

		//log.Println("Action: ", action)
		//log.Println("Next Level: ", nextLevel)
		//pause
		//crt.SetDelayInMin(1)
		//crt.DelayIt()

		if action == menu.Quit {
			//	crt.Println("Quitting")
			ok = true
			continue
		}

		if support.IsInt(action) {
			switch action {
			case "1":
				Trans(crt, nextLevel.AlternateID, nextLevel.Title)
				ok = false
				action = ""
			case "2":
				//QTor(crt, nextLevel.AlternateID, nextLevel.Title)
				ok = false
				action = ""
			default:
				crt.InputError(invalidActionErrorText + "'" + action + "'")
				ok = false
				action = ""
			}
		}
	}
}

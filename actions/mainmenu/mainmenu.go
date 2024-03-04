package mainmenu

import (
	"fmt"

	plexmediaserver "github.com/mt1976/crt/actions/plexms"
	"github.com/mt1976/crt/actions/skynews"
	torrents "github.com/mt1976/crt/actions/torrents"
	"github.com/mt1976/crt/actions/weather"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/page"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(crt *support.Crt) {

	m := page.New(mainMenuTitleText)
	//for i := range 11 {
	//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
	//}

	m.AddOption(1, blankText, "", "")
	m.AddOption(2, skyNewsMenuTitleText, "", "")
	m.AddOption(3, bbcNewsMenuTitleText, "", "")
	m.AddOption(4, weatherMenuTitleText, "", "")
	m.AddOption(5, torrentsText, "", "")
	m.AddOption(6, plexmediaserversMenuTitleText, "", "")
	m.AddOption(7, remoteSystemsAccessMenuTitleText, "", "")
	m.AddOption(8, systemsMaintenanceMenuTitleText, "", "")
	m.AddOption(9, blankText, "", "")
	m.AddOption(10, blankText, "", "")
	m.AddAction(page.QuitText)

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		//crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")

		action, _ := m.Display(crt)
		switch action {
		case page.QuitText:
			crt.InfoMessage(quittingMessageText + "\n ")
			ok = true
			continue
		case "1":
			y := page.New(subMenuTitleText)
			for i := range 14 {
				y.AddOption(i, fmt.Sprintf(subMenuTitleText+" %v", action), "", "")
			}
			//action, _ = y.Display(crt)
		case "2":
			skynews.Run(crt)
		case "4":
			weather.Run(crt)
		case "5":
			torrents.Run(crt)
		case "6":
			plexmediaserver.Run(crt)
		default:
			crt.InputError(page.ErrInvalidAction + support.SQuote(action))
		}
	}
}

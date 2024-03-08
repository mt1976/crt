package mainmenu

import (
	"fmt"

	plexmediaserver "github.com/mt1976/crt/actions/plexms"
	"github.com/mt1976/crt/actions/skynews"
	torrents "github.com/mt1976/crt/actions/torrents"
	"github.com/mt1976/crt/actions/weather"
	t "github.com/mt1976/crt/language"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/page"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(crt *support.Crt) {

	m := page.New(t.TxtMainMenuTitle)
	//for i := range 11 {
	//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
	//}

	m.AddOption(1, t.SymBlank, "", "")
	m.AddOption(2, t.TxtSkyNewsMenuTitle, "", "")
	m.AddOption(3, t.TxtBBCNewsMenuTitle, "", "")
	m.AddOption(4, t.TxtWeatherMenuTitle, "", "")
	m.AddOption(5, t.TxtTorrentsMenuTitle, "", "")
	m.AddOption(6, t.TxtPlexMediaServersMenuTitle, "", "")
	m.AddOption(7, t.TxtRemoteSystemsAccessMenuTitle, "", "")
	m.AddOption(8, t.TxtSystemsMaintenanceMenuTitle, "", "")
	m.AddOption(9, t.SymBlank, "", "")
	m.AddOption(10, t.SymBlank, "", "")
	m.AddAction(t.SymActionQuit)

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		//crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")

		action, _ := m.Display(crt)
		switch action {
		case t.SymActionQuit:
			crt.InfoMessage(t.TxtQuittingMessage + t.SymNewline)
			ok = true
			continue
		case "1":
			y := page.New(t.TxtSubMenuTitle)
			for i := range 14 {
				y.AddOption(i, fmt.Sprintf(t.TxtSubMenuTitle+" %v", action), "", "")
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

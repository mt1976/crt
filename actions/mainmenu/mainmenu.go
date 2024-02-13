package mainmenu

import (
	"fmt"

	"github.com/mt1976/admin_me/actions/skynews"
	"github.com/mt1976/admin_me/support"
	menu "github.com/mt1976/admin_me/support/menu"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(crt *support.Crt) {

	m := menu.New(mainMenuTitleText)
	//for i := range 11 {
	//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
	//}

	m.Add(1, "Test", "", "")
	m.Add(2, skyNewsMenuTitleText, "", "")
	m.Add(3, bbcNewsMenuTitleText, "", "")
	m.Add(4, weatherMenuTitleText, "", "")
	m.Add(5, "", "", "")
	m.Add(6, "", "", "")
	m.Add(7, remoteSystemsAccessMenuTitleText, "", "")
	m.Add(8, systemsMaintenanceMenuTitleText, "", "")
	m.AddAction(menu.Quit)

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		//crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")

		action, _ := m.Display(crt)
		switch action {
		case menu.Quit:
			crt.Println(quittingMessageText)
			ok = true
			continue
		case "1":
			y := menu.New(subMenuTitleText)
			for i := range 14 {
				y.Add(i, fmt.Sprintf(subMenuTitleText+" %v", action), "", "")
			}
			//action, _ = y.Display(crt)
		case "2":
			skynews.Run(crt)
		default:
			crt.InputError(invalidActionErrorText + "'" + action + "'")
		}
	}
}

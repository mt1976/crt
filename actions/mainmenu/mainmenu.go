package mainmenu

import (
	"fmt"

	"github.com/mt1976/admin_me/actions/news"
	"github.com/mt1976/admin_me/support"
	menu "github.com/mt1976/admin_me/support/menu"
)

// The Run function displays a main menu and allows the user to navigate through different sub-menus
// and perform various actions.
func Run(crt *support.Crt) {

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")
		m := menu.New(mainMenuTitle)
		//for i := range 11 {
		//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
		//}

		m.Add(1, "Test", "", "")
		m.Add(2, newsMenuTitle, "", "")
		m.Add(3, weatherMenuTitle, "", "")
		m.Add(4, "", "", "")
		m.Add(5, "", "", "")
		m.Add(6, "", "", "")
		m.Add(7, remoteSystemsAccessMenuTitle, "", "")
		m.Add(8, systemsMaintenanceMenuTitle, "", "")
		m.AddAction("Q")

		action, _ := m.Display(crt)
		switch action {
		case "Q":
			crt.Println(quittingMessage)
			ok = true
			continue
		case "1":
			y := menu.New(subMenuTitle)
			for i := range 14 {
				y.Add(i, fmt.Sprintf(subMenuTitle+" %v", action), "", "")
			}
			action, _ = y.Display(crt)
		case "2":
			news.Run(crt)
		}
	}
}

package mainmenu

import (
	"fmt"

	"github.com/mt1976/admin_me/actions/news"
	"github.com/mt1976/admin_me/support"
	menu "github.com/mt1976/admin_me/support/menu"
)

func Run(crt *support.Crt) {

	// loop while ok
	ok := false
	for !ok {

		crt.Clear()
		crt.SetDelayInSec(0.25) // Set delay in milliseconds
		//crt.Header("Main Menu")
		m := menu.NewMenu(mainMenuTitle)
		//for i := range 11 {
		//	m.AddMenuItem(i, fmt.Sprintf("Menu Item %v", i))
		//}

		m.AddMenuItem(1, "Test", "", "")
		m.AddMenuItem(2, newsMenuTitle, "", "")
		m.AddMenuItem(3, weatherMenuTitle, "", "")
		m.AddMenuItem(4, "", "", "")
		m.AddMenuItem(5, "", "", "")
		m.AddMenuItem(6, "", "", "")
		m.AddMenuItem(7, remoteSystemsAccessMenuTitle, "", "")
		m.AddMenuItem(8, systemsMaintenanceMenuTitle, "", "")
		m.AddAction("Q")

		action, _ := m.DisplayMenu(crt)
		switch action {
		case "Q":
			crt.Println(quittingMessage)
			ok = true
			continue
		case "1":
			y := menu.NewMenu(subMenuTitle)
			for i := range 14 {
				y.AddMenuItem(i, fmt.Sprintf(subMenuTitle+" %v", action), "", "")
			}
			action, _ = y.DisplayMenu(crt)
		case "2":
			news.Run(crt)
		}
	}
}

package news

import (
	"log"

	"github.com/mt1976/admin_me/support"
	"github.com/mt1976/admin_me/support/menu"
)

func Run(crt *support.Crt) {

	// Home
	// UK
	// World
	// US
	// Business
	// Politics
	// Technology
	// Entertainment
	// Strange News

	crt.Clear()
	crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Main Menu")
	m := menu.NewMenu("SKY News")
	m.AddMenuItem(1, "Home", "https://feeds.skynews.com/feeds/rss/home.xml", "")
	m.AddMenuItem(2, "UK", "https://feeds.skynews.com/feeds/rss/uk.xml", "")
	m.AddMenuItem(3, "World", "https://feeds.skynews.com/feeds/rss/world.xml", "")
	m.AddMenuItem(4, "US", "https://feeds.skynews.com/feeds/rss/us.xml", "")
	m.AddMenuItem(5, "Business", "https://feeds.skynews.com/feeds/rss/business.xml", "")
	m.AddMenuItem(6, "Politics", "https://feeds.skynews.com/feeds/rss/politics.xml", "")
	m.AddMenuItem(7, "Technology", "https://feeds.skynews.com/feeds/rss/technology.xml", "")
	m.AddMenuItem(8, "Entertainment", "https://feeds.skynews.com/feeds/rss/entertainment.xml", "")
	m.AddMenuItem(9, "Strange News", "https://feeds.skynews.com/feeds/rss/strange.xml", "")
	m.AddAction("Q")

	ok := false
	for !ok {
		action, nextLevel := m.DisplayMenu(crt)

		log.Println("Action: ", action)
		log.Println("Next Level: ", nextLevel)
		//pause
		//crt.SetDelayInMin(1)
		//crt.DelayIt()

		if action == "Q" {
			crt.Println("Quitting")
			ok = true
			continue
		}

		if support.IsInt(action) {
			Topic(crt, nextLevel.AlternateID, nextLevel.Title)
			ok = false
			action = ""
		}
	}
}

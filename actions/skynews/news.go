package skynews

import (
	"log"

	"github.com/mt1976/admin_me/support"
	"github.com/mt1976/admin_me/support/menu"
)

// The Run function displays a menu of news topics and allows the user to select a topic to view the
// news articles related to that topic.
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
	//crt.SetDelayInSec(0.25) // Set delay in milliseconds
	//crt.Header("Main Menu")
	m := menu.New("SKY News")
	m.Add(1, "Home", "https://feeds.skynews.com/feeds/rss/home.xml", "")
	m.Add(2, "UK", "https://feeds.skynews.com/feeds/rss/uk.xml", "")
	m.Add(3, "World", "https://feeds.skynews.com/feeds/rss/world.xml", "")
	m.Add(4, "US", "https://feeds.skynews.com/feeds/rss/us.xml", "")
	m.Add(5, "Business", "https://feeds.skynews.com/feeds/rss/business.xml", "")
	m.Add(6, "Politics", "https://feeds.skynews.com/feeds/rss/politics.xml", "")
	m.Add(7, "Technology", "https://feeds.skynews.com/feeds/rss/technology.xml", "")
	m.Add(8, "Entertainment", "https://feeds.skynews.com/feeds/rss/entertainment.xml", "")
	m.Add(9, "Strange News", "https://feeds.skynews.com/feeds/rss/strange.xml", "")
	m.AddAction("Q")

	ok := false
	for !ok {
		action, nextLevel := m.Display(crt)

		log.Println("Action: ", action)
		log.Println("Next Level: ", nextLevel)
		//pause
		//crt.SetDelayInMin(1)
		//crt.DelayIt()

		if action == "Q" {
			//	crt.Println("Quitting")
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

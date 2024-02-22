package skynews

import (
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
	page "github.com/mt1976/crt/support/page"
)

var C = config.Configuration

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
	m := page.New(menuTitleText)
	c := 0
	c++
	m.AddOption(c, topicHomeText, C.SkyNewsURI+topicHomeURI, "")
	c++
	m.AddOption(c, topicUKText, C.SkyNewsURI+topicUKURI, "")
	c++
	m.AddOption(c, topicWorldText, C.SkyNewsURI+topicWorldURI, "")
	c++
	m.AddOption(c, topicUSText, C.SkyNewsURI+topicUSURI, "")
	c++
	m.AddOption(c, topicBusinessText, C.SkyNewsURI+topicBusinessURI, "")
	c++
	m.AddOption(c, topicPoliticsText, C.SkyNewsURI+topicPoliticsURI, "")
	c++
	m.AddOption(c, topicTechnologyText, C.SkyNewsURI+topicTechnologyURI, "")
	c++
	m.AddOption(c, topicEntertainmentText, C.SkyNewsURI+topicEntertainmentURI, "")
	c++
	m.AddOption(c, topicStrangeText, C.SkyNewsURI+topicStrangeURI, "")
	m.AddAction(page.Quit)

	ok := false
	for !ok {
		action, nextLevel := m.Display(crt)

		//log.Println("Action: ", action)
		//log.Println("Next Level: ", nextLevel)
		//pause
		//crt.SetDelayInMin(1)
		//crt.DelayIt()

		if action == page.Quit {
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

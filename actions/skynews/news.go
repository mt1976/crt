package skynews

import (
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
	m := menu.New(menuTitleText)
	c := 0
	c++
	m.Add(c, topicHomeText, topicHomeURI, "")
	c++
	m.Add(c, topicUKText, topicUKURI, "")
	c++
	m.Add(c, topicWorldText, topicWorldURI, "")
	c++
	m.Add(c, topicUSText, topicUSURI, "")
	c++
	m.Add(c, topicBusinessText, topicBusinessURI, "")
	c++
	m.Add(c, topicPoliticsText, topicPoliticsURI, "")
	c++
	m.Add(c, topicTechnologyText, topicTechnologyURI, "")
	c++
	m.Add(c, topicEntertainmentText, topicEntertainmentURI, "")
	c++
	m.Add(c, topicStrangeText, topicStrangeURI, "")
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
			Topic(crt, nextLevel.AlternateID, nextLevel.Title)
			ok = false
			action = ""
		}
	}
}

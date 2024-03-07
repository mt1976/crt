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
	m.AddOption(c, topicHomeLabel, C.SkyNewsURI+C.SkyNewsHomeURI, "")
	c++
	m.AddOption(c, topicUKLabel, C.SkyNewsURI+C.SkyNewsUKURI, "")
	c++
	m.AddOption(c, topicWorldLabel, C.SkyNewsURI+C.SkyNewsWorldURI, "")
	c++
	m.AddOption(c, topicUSLabel, C.SkyNewsURI+C.SkyNewsUSURI, "")
	c++
	m.AddOption(c, topicBusinessLabel, C.SkyNewsURI+C.SkyNewsBusinessURI, "")
	c++
	m.AddOption(c, topicPoliticsLabel, C.SkyNewsURI+C.SkyNewsPoliticsURI, "")
	c++
	m.AddOption(c, topicTechnologyLabel, C.SkyNewsURI+C.SkyNewsTechnologyURI, "")
	c++
	m.AddOption(c, topicEntertainmentLabel, C.SkyNewsURI+C.SkyNewsEntertainmentURI, "")
	c++
	m.AddOption(c, topicStrangeLabel, C.SkyNewsURI+C.SkyNewsStrangeURI, "")
	m.AddAction(page.TxtQuit)

	action, nextLevel := m.Display(crt)

	if action == page.TxtQuit {
		return
	}
	if support.IsInt(action) {
		Topic(crt, nextLevel.AlternateID, nextLevel.Title)
		action = ""
	}

}

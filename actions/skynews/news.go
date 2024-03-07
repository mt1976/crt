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
	m := page.New(TxtMenuTitle)
	c := 0
	c++
	m.AddOption(c, TxtTopicHome, C.URISkyNews+C.URISkyNewsHome, "")
	c++
	m.AddOption(c, TxtTopicUK, C.URISkyNews+C.URISkyNewsUK, "")
	c++
	m.AddOption(c, TxtTopicWorld, C.URISkyNews+C.URISkyNewsWorld, "")
	c++
	m.AddOption(c, TxtTopicUS, C.URISkyNews+C.URISkyNewsUS, "")
	c++
	m.AddOption(c, TxtTopicBusiness, C.URISkyNews+C.URISkyNewsBusiness, "")
	c++
	m.AddOption(c, TxtTopicPolitics, C.URISkyNews+C.URISkyNewsPolitics, "")
	c++
	m.AddOption(c, TxtTopicTechnology, C.URISkyNews+C.URISkyNewsTechnology, "")
	c++
	m.AddOption(c, TxtTopicEntertainment, C.URISkyNews+C.URISkyNewsEntertainment, "")
	c++
	m.AddOption(c, TxtTopicStrange, C.URISkyNews+C.URISkyNewsStrange, "")
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

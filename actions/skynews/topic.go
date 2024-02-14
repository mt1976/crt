package skynews

import (
	"github.com/mmcdole/gofeed"
	"github.com/mt1976/admin_me/support"
	"github.com/mt1976/admin_me/support/menu"
)

// The function "Topic" takes in a CRT object, a topic, and a title as parameters, and then retrieves
// news items for that topic from an RSS feed, displays them in a menu, and allows the user to select a
// news item to view.
func Topic(crt *support.Crt, topic, title string) {
	//crt.Println("Topic: " + topic + " - " + title)
	// Get the news for the topic
	crt.InfoMessage(topicLoadingText + crt.Bold(title))
	// get the news for the topic from an rss feed
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(topic)

	t := menu.New(feed.Title)
	noNewsItems := len(feed.Items)
	if noNewsItems > menu.MaxMenuItems {
		noNewsItems = menu.MaxMenuItems
	}
	for i := range noNewsItems {
		//log.Println("Adding: ", feed.Items[i].Title, i)
		t.Add(i+1, feed.Items[i].Title, feed.Items[i].Link, feed.Items[i].Published)
	}
	ok := false
	for !ok {
		action, mi := t.Display(crt)

		if action == menu.Quit {
			//crt.Println("Quitting")
			ok = true
			continue
		}
		if support.IsInt(action) {
			Story(crt, mi.AlternateID)
			ok = false
			action = ""
		}

		//log.Println("Action: ", action)
		//log.Println("Next Level: ", mi)

		//spew.Dump(nextLevel)
	}
}

package news

import (
	"log"

	"github.com/mmcdole/gofeed"
	"github.com/mt1976/admin_me/support"
	"github.com/mt1976/admin_me/support/menu"
)

func Topic(crt *support.Crt, topic, title string) {
	crt.Println("Topic: " + topic + " - " + title)
	// Get the news for the topic

	// get the news for the topic from an rss feed
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(topic)

	t := menu.NewMenu(feed.Title)
	noNewsItems := len(feed.Items)
	if noNewsItems > menu.MaxMenuItems {
		noNewsItems = menu.MaxMenuItems
	}
	for i := range noNewsItems {
		log.Println("Adding: ", feed.Items[i].Title, i)
		t.AddMenuItem(i, feed.Items[i].Title, feed.Items[i].Link, feed.Items[i].Published)
	}
	ok := false
	for !ok {
		action, mi := t.DisplayMenu(crt)

		if action == "Q" {
			crt.Println("Quitting")
			ok = true
			continue
		}
		if menu.IsInt(action) {
			Story(crt, mi.AlternateID)
			ok = false
			action = ""
		}

		log.Println("Action: ", action)
		log.Println("Next Level: ", mi)

		//spew.Dump(nextLevel)
	}
}

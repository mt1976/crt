package skynews

import (
	"net/url"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/hekmon/transmissionrpc"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/menu"
)

// The function "Trans" takes in a CRT object, a topic, and a title as parameters, and then retrieves
// news items for that topic from an RSS feed, displays them in a menu, and allows the user to select a
// news item to view.
func Trans(crt *support.Crt, uri, title string) {
	//crt.Println("Topic: " + topic + " - " + title)
	// Get the news for the topic
	crt.InfoMessage(transLoadingTorrentsText)
	p := menu.New("Transmission")
	// get the news for the topic from an rss feed
	endpoint, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}
	tbt, err := transmissionrpc.New(endpoint.Host, "admin", "admin", nil)
	if err != nil {
		panic(err)
	}
	//spew.Dump(tbt)
	//spew.Dump(tbt.RPCVersion())
	spew.Dump(tbt.TorrentGetAll())
	os.Exit(0)
	// torrents, _ := tbt.TorrentGetAll()
	// noTorrents := len(torrents)
	// if noTorrents > menu.MaxMenuItems {
	// 	noTorrents = menu.MaxMenuItems
	// }

	// for i := range noTorrents {
	// 	//log.Println("Adding: ", feed.Items[i].Title, i)
	// 	p.Add(i+1, torrents[i].Name, "", "")
	// }
	ok := false
	for !ok {
		action, _ := p.Display(crt)

		if action == menu.Quit {
			//crt.Println("Quitting")
			ok = true
			continue
		}
		if support.IsInt(action) {
			//	Story(crt, mi.AlternateID)
			ok = false
			action = ""
		}

		//log.Println("Action: ", action)
		//log.Println("Next Level: ", mi)

		//spew.Dump(nextLevel)
	}
}

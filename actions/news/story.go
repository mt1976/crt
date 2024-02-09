package news

import (
	"github.com/mt1976/admin_me/support"
	"github.com/mt1976/admin_me/support/menu"
)

func Story(crt *support.Crt, storyLink string) {
	crt.Println("Story: " + storyLink)

	s := menu.NewMenu("Story")
	s.AddMenuItem(1, storyLink, storyLink, "")
	s.AddAction("Q")
	ok := false
	for !ok {
		x, _ := s.DisplayMenu(crt)
		if x == "Q" {
			ok = true
		}
	}

}

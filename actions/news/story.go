package news

import (
	"github.com/mt1976/admin_me/support"
	"github.com/mt1976/admin_me/support/menu"
)

// The function "Story" displays a story link and allows the user to interact with a menu until they
// choose to quit.
func Story(crt *support.Crt, storyLink string) {
	crt.Println("Story: " + storyLink)

	s := menu.New("Story")
	s.Add(1, storyLink, storyLink, "")
	s.AddAction("Q")
	ok := false
	for !ok {
		x, _ := s.Display(crt)
		if x == "Q" {
			ok = true
		}
	}
}

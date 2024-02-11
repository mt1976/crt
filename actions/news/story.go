package news

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mt1976/admin_me/support"
	page "github.com/mt1976/admin_me/support/page"
)

// The function "Story" displays a story link and allows the user to interact with a menu until they
// choose to quit.
func Story(crt *support.Crt, storyLink string) {
	crt.Println("Story: " + storyLink)

	//s := page.New("Story")
	//s.Add(1, storyLink, storyLink, "")
	//s.AddAction("Q")
	s := buildPage(crt, storyLink)

	ok := false
	for !ok {

		spew.Dump(crt)
		spew.Dump(storyLink)
		x, _ := s.Display(crt, 1)

		if x == "Q" {
			ok = true
		}
	}
}

func buildPage(crt *support.Crt, storyLink string) *page.Page {

	//Get html from storyLink
	//Parse html for story
	//Create page with story
	//Return page
	//c := colly.NewCollector()

	pageTitle := ""
	pageContent := page.Page{}
	//rowNumber := 1
	// Find and visit all links
	// c.OnHTML("title", func(e *colly.HTMLElement) {
	// 	pageTitle = e.Text
	// })

	// c.OnHTML("p", func(e *colly.HTMLElement) {
	// 	rowNumber++
	// 	pageContent.Add(rowNumber, e.Text, "", "")
	// })
	pageTitle = "Test"
	// c.Visit(storyLink)
	for i := 1; i <= 30; i++ {
		pageContent.Add(i, "This is a test", "", "")
	}

	p := page.New(pageTitle)

	spew.Dump(pageTitle)
	spew.Dump(pageContent)
	crt.SetDelayInMin(1)
	crt.DelayIt()
	p.Add(1, storyLink, storyLink, "")
	return p
}

package dashboard

import (
	"time"

	e "github.com/mt1976/crt/errors"
	t "github.com/mt1976/crt/language"
	support "github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
	page "github.com/mt1976/crt/support/page"
)

var C = config.Configuration

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(crt *support.Crt) {

	crt.Clear()
	p := page.New(t.TxtDashboardTitle)

	c := 0
	c++
	p.Add("Testing Server/Service Dashboard", "", time.Now().Format("2006-01-02"))

	p.AddAction(t.SymActionQuit)
	p.AddAction(t.SymActionForward)
	p.AddAction(t.SymActionBack)
	ok := false
	for !ok {

		nextAction, _ := p.Display(crt)
		switch nextAction {
		case t.SymActionForward:
			p.NextPage(crt)
		case t.SymActionBack:
			p.PreviousPage(crt)
		case t.SymActionQuit:
			ok = true
			return
		default:
			crt.InputError(e.ErrInvalidAction + t.SymSingleQuote + nextAction + t.SymSingleQuote)
		}
	}

}

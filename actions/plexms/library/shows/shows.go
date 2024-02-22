package shows

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
	page "github.com/mt1976/crt/support/page"
)

var C = config.Configuration

func Run(crt *support.Crt, mediaVault *plex.Plex, wi *plex.Directory) {
	crt.Shout("Movies-" + wi.Title)

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error("mvLibError", err)
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := page.New(res.MediaContainer.LibrarySectionTitle + " (" + noItems + ")")
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		m.AddOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	//exit := false
	//for !exit {

	nextAction, _ := m.Display(crt)
	switch nextAction {
	case page.Quit:
		//		exit = true
		return
	default:
		if support.IsInt(nextAction) {
			//	Action(crt, mediaVault, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
			Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1], mediaVault)
		} else {
			crt.InputError(page.ErrInvalidAction + "'" + nextAction + "'")
		}
	}
	//}

	//spew.Dump(res)
	//os.Exit(1)
}

func Detail(crt *support.Crt, info plex.Metadata, mediaVault *plex.Plex) {
	p := page.New(info.Title)

	p.AddFieldValuePair(crt, "Title", info.Title)
	p.AddFieldValuePair(crt, "Year", support.ToString(info.Year))
	p.AddFieldValuePair(crt, "Content Rating", info.ContentRating)
	p.AddFieldValuePair(crt, "Released", support.FormatPlexDate(info.OriginallyAvailableAt))
	p.BlankRow()
	p.AddFieldValuePair(crt, "Summary", info.Summary)

	p.AddAction(Seasons) //Drilldown to episodes
	p.SetPrompt(prompt)

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case page.Quit:
		return
	case Seasons:
		SeasonDetails(crt, mediaVault, info)
	}
}

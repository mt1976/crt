package shows

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	e "github.com/mt1976/crt/errors"
	notations "github.com/mt1976/crt/language"
	"github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
	page "github.com/mt1976/crt/support/page"
)

var C = config.Configuration

func Run(crt *support.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error(e.ErrLibraryResponse, err)
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := page.New(res.MediaContainer.LibrarySectionTitle + " " + support.PQuote(noItems))
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		m.AddOption(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	nextAction, _ := m.Display(crt)
	switch nextAction {
	case notations.SymActionQuit:
		return
	default:
		if support.IsInt(nextAction) {
			Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1], mediaVault)
		} else {
			crt.InputError(e.ErrInvalidAction + support.SQuote(nextAction))
		}
	}
}

func Detail(crt *support.Crt, info plex.Metadata, mediaVault *plex.Plex) {
	p := page.New(info.Title)

	p.AddFieldValuePair(crt, notations.TitleLabel, info.Title)
	p.AddFieldValuePair(crt, notations.YearLabel, support.ToString(info.Year))
	p.AddFieldValuePair(crt, notations.ContentLabel, info.ContentRating)
	p.AddFieldValuePair(crt, notations.ReleasedLabel, support.FormatPlexDate(info.OriginallyAvailableAt))
	p.BlankRow()
	p.AddFieldValuePair(crt, notations.SummaryLabel, info.Summary)

	p.AddAction(TxtSeasons) //Drilldown to episodes
	p.SetPrompt(TxtPrompt)

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case notations.SymActionQuit:
		return
	case TxtSeasons:
		SeasonDetails(crt, mediaVault, info)
	}
}

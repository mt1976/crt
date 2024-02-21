package shows

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/support"
	menu "github.com/mt1976/crt/support/menu"
	page "github.com/mt1976/crt/support/page"
)

func Episodes(crt *support.Crt, mediaVault *plex.Plex, seriesTitle string, info plex.Metadata) {
	//spew.Dump(mediaVault, info)

	fmt.Println("Episodes")
	res, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		crt.Error("mvLibError", err)
		os.Exit(1)
	}
	m := menu.New(seriesTitle + " " + info.Title)

	noEps := len(res.MediaContainer.Metadata)
	for i := 0; i < noEps; i++ {
		ep := res.MediaContainer.Metadata[i]
		m.Add(i+1, ep.Title, "", "")
	}

	exit := false
	//m.SetPrompt()
	for !exit {
		nextAction, _ := m.Display(crt)
		switch nextAction {
		case menu.Quit:
			exit = true
			return
		case menu.Forward:
			m.NextPage(crt)
		case menu.Back:
			m.PreviousPage(crt)
		default:
			if support.IsInt(nextAction) {
				EpisodeDetail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1], seriesTitle)
			} else {
				crt.InputError(menu.InvalidActionError + "'" + nextAction + "'")
			}
		}

	}
	spew.Dump(res)
	os.Exit(1)
}

func EpisodeDetail(crt *support.Crt, info plex.Metadata, seriesTitle string) {

	title := seriesTitle + " " + info.ParentTitle + " " + info.Title
	p := page.New(title)
	p.AddFieldValuePair(crt, "Title", info.Title)
	p.AddFieldValuePair(crt, "Season", info.ParentTitle)
	p.AddFieldValuePair(crt, "Summary", info.Summary)
	p.AddFieldValuePair(crt, "Duration", support.FormatPlexDuration(info.Duration))
	p.AddFieldValuePair(crt, "Release", support.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, "Rating", info.ContentRating)

	p.BlankRow()
	p.AddColumns(crt, "Media")
	p.AddColumns(crt, "-----")
	for _, v := range info.Media {
		//p.AddFieldValuePair(crt, "Media", v.Part[0].File)
		p.AddColumns(crt, v.Part[0].File)
	}
	//p.SetPrompt(prompt)

	exit := false
	for !exit {
		nextAction, _ := p.Display(crt)
		switch nextAction {
		case page.Quit:
			exit = true
			return
		case page.Forward:
			p.NextPage(crt)
		case page.Back:
			p.PreviousPage(crt)

		default:
			crt.InputError(menu.InvalidActionError + "'" + nextAction + "'")
		}
	}
}

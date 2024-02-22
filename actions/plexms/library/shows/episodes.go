package shows

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/support"

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
	m := page.New(seriesTitle + " " + info.Title)

	noEps := len(res.MediaContainer.Metadata)
	for i := 0; i < noEps; i++ {
		ep := res.MediaContainer.Metadata[i]
		m.AddOption(i+1, ep.Title, "", "")
	}

	exit := false
	//m.SetPrompt()
	for !exit {
		nextAction, _ := m.Display(crt)
		switch nextAction {
		case page.Quit:
			exit = true
			return
		case page.Forward:
			m.NextPage(crt)
		case page.Back:
			m.PreviousPage(crt)
		default:
			if support.IsInt(nextAction) {
				EpisodeDetail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
			} else {
				crt.InputError(page.ErrInvalidAction + "'" + nextAction + "'")
			}
		}

	}
	//spew.Dump(res)
	//os.Exit(1)
}

func EpisodeDetail(crt *support.Crt, info plex.Metadata) {
	//spew.Dump(info)
	//os.Exit(1)
	title := info.GrandparentTitle + " " + info.ParentTitle + " " + info.Title
	p := page.New(title)
	p.AddFieldValuePair(crt, "Show", info.GrandparentTitle)
	p.AddFieldValuePair(crt, "Season", info.ParentTitle)
	p.AddFieldValuePair(crt, "Episode", info.Title)
	p.AddFieldValuePair(crt, "Summary", info.Summary)
	p.AddFieldValuePair(crt, "Duration", support.FormatPlexDuration(info.Duration))
	p.AddFieldValuePair(crt, "Released", support.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, "Rating", info.ContentRating)
	videoCodec := info.Media[0].VideoCodec
	videoFrameRate := info.Media[0].VideoFrameRate
	videoResolution := info.Media[0].VideoResolution
	videoContainer := info.Media[0].Container
	aspectRatio := info.Media[0].AspectRatio

	p.BlankRow()
	p.AddColumns(crt, "Codec", "Frame Rate", "Resolution", "Container", "Aspect Ratio")
	p.AddColumnsRuler(crt, "Codec", "Frame Rate", "Resolution", "Container", "Aspect Ratio")
	p.AddColumns(crt, videoCodec, videoFrameRate, videoResolution, videoContainer, aspectRatio.String())
	p.BlankRow()
	p.AddColumns(crt, "Media")
	p.AddColumnsRuler(crt, "Media")
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
			crt.InputError(page.ErrInvalidAction + "'" + nextAction + "'")
		}
	}
}

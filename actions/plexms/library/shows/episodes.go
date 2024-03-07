package shows

import (
	"os"

	"github.com/jrudio/go-plex-client"
	notations "github.com/mt1976/crt/actions/plexms/language"
	"github.com/mt1976/crt/support"
	page "github.com/mt1976/crt/support/page"
)

func Episodes(crt *support.Crt, mediaVault *plex.Plex, seriesTitle string, info plex.Metadata) {
	res, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		crt.Error(notations.ErrLibraryResponse, err)
		os.Exit(1)
	}
	m := page.New(seriesTitle + " " + info.Title)

	noEps := len(res.MediaContainer.Metadata)
	for i := 0; i < noEps; i++ {
		ep := res.MediaContainer.Metadata[i]
		m.AddOption(i+1, ep.Title, "", "")
	}

	nextAction, _ := m.Display(crt)
	switch nextAction {
	case page.TxtQuit:
		return
	default:
		if support.IsInt(nextAction) {
			EpisodeDetail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
		} else {
			crt.InputError(notations.ErrInvalidAction + support.SQuote(nextAction))
		}
	}
}

func EpisodeDetail(crt *support.Crt, info plex.Metadata) {

	title := info.GrandparentTitle + " " + info.ParentTitle + " " + info.Title
	p := page.New(title)
	p.AddFieldValuePair(crt, notations.ShowLabel, info.GrandparentTitle)
	p.AddFieldValuePair(crt, notations.SeasonLabel, info.ParentTitle)
	p.AddFieldValuePair(crt, notations.EpisodeLabel, info.Title)
	p.AddFieldValuePair(crt, notations.SummaryLabel, info.Summary)
	p.AddFieldValuePair(crt, notations.DurationLabel, support.FormatPlexDuration(info.Duration))
	p.AddFieldValuePair(crt, notations.ReleasedLabel, support.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, notations.ContentLabel, info.ContentRating)
	videoCodec := info.Media[0].VideoCodec
	videoFrameRate := info.Media[0].VideoFrameRate
	videoResolution := info.Media[0].VideoResolution
	videoContainer := info.Media[0].Container
	aspectRatio := info.Media[0].AspectRatio

	p.BlankRow()
	p.AddColumnsTitle(crt, notations.CodecLabel, notations.FrameRateLabel, notations.ResolutionLabel, notations.ContainerLabel, notations.AspectRatioLabel)
	p.AddColumns(crt, videoCodec, videoFrameRate, videoResolution, videoContainer, aspectRatio.String())
	p.BlankRow()
	p.AddColumnsTitle(crt, notations.MediaLabel)
	for _, v := range info.Media {
		p.AddColumns(crt, v.Part[0].File)
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case page.TxtQuit:
		return
	}
}

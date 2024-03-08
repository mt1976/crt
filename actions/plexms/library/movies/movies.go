package movies

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	notations "github.com/mt1976/crt/language"
	t "github.com/mt1976/crt/language"
	"github.com/mt1976/crt/support"
	page "github.com/mt1976/crt/support/page"
)

func Run(crt *support.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error(notations.ErrLibraryResponse, err)
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
	case t.SymActionQuit:
		return
	default:
		if support.IsInt(nextAction) {
			Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
		} else {
			crt.InputError(notations.ErrInvalidAction + support.SQuote(nextAction))
		}
	}
}

func Detail(crt *support.Crt, info plex.Metadata) {
	p := page.New(info.Title)

	p.AddFieldValuePair(crt, notations.TitleLabel, info.Title)
	p.AddFieldValuePair(crt, notations.ContentLabel, info.ContentRating)
	dur := support.FormatPlexDuration(info.Duration)
	p.AddFieldValuePair(crt, notations.DurationLabel, dur)
	p.AddFieldValuePair(crt, notations.ReleasedLabel, support.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, notations.SummaryLabel, info.Summary)
	//unix time to hrs mins secs
	p.BlankRow()
	for i := 0; i < len(info.Director); i++ {
		poobum := info.Director[i]
		lbl := notations.DirectorLabel
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(crt, lbl, poobum.Tag)
	}

	for i := 0; i < len(info.Writer); i++ {
		poobum := info.Writer[i]
		lbl := notations.WriterLabel
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(crt, lbl, poobum.Tag)
	}

	count := 0
	p.BlankRow()
	p.AddColumnsTitle(crt, notations.ContainerLabel, notations.ResolutionLabel, notations.CodecLabel, notations.AspectRatioLabel, notations.FrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	//range trhough parts
	p.BlankRow()
	p.AddColumnsTitle(crt, notations.MediaLabel)
	for _, v := range info.Media {
		p.AddColumns(crt, v.Part[0].File)
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	default:
		crt.InputError(notations.ErrInvalidAction + support.SQuote(nextAction))
	}

}

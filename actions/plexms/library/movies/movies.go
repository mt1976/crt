package movies

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	e "github.com/mt1976/crt/errors"

	t "github.com/mt1976/crt/language"
	"github.com/mt1976/crt/support"
	page "github.com/mt1976/crt/support/page"
)

func Run(crt *support.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error(e.ErrLibraryResponse, err)
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := page.New(res.MediaContainer.LibrarySectionTitle + t.Space + support.PQuote(noItems))
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
			crt.InputError(e.ErrInvalidAction + support.SQuote(nextAction))
		}
	}
}

func Detail(crt *support.Crt, info plex.Metadata) {
	p := page.New(info.Title)

	p.AddFieldValuePair(crt, t.TitleLabel, info.Title)
	p.AddFieldValuePair(crt, t.ContentLabel, info.ContentRating)
	dur := support.FormatPlexDuration(info.Duration)
	p.AddFieldValuePair(crt, t.DurationLabel, dur)
	p.AddFieldValuePair(crt, t.ReleasedLabel, support.FormatPlexDate(info.OriginallyAvailableAt))
	p.AddFieldValuePair(crt, t.SummaryLabel, info.Summary)
	//unix time to hrs mins secs
	p.BlankRow()
	for i := 0; i < len(info.Director); i++ {
		data := info.Director[i]
		lbl := t.DirectorLabel
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(crt, lbl, data.Tag)
	}

	for i := 0; i < len(info.Writer); i++ {
		poobum := info.Writer[i]
		lbl := t.WriterLabel
		if i > 0 {
			lbl = ""
		}
		p.AddFieldValuePair(crt, lbl, poobum.Tag)
	}

	count := 0
	p.BlankRow()
	p.AddColumnsTitle(crt, t.ContainerLabel, t.ResolutionLabel, t.CodecLabel, t.AspectRatioLabel, t.FrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	//range trhough parts
	p.BlankRow()
	p.AddColumnsTitle(crt, t.MediaLabel)
	for _, v := range info.Media {
		p.AddColumns(crt, v.Part[0].File)
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	default:
		crt.InputError(e.ErrInvalidAction + support.SQuote(nextAction))
	}

}

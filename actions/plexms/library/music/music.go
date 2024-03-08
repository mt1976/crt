package music

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	notations "github.com/mt1976/crt/language"
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
	case page.TxtQuit:
		return
	default:
		if support.IsInt(nextAction) {
			//	Action(crt, mediaVault, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
			Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])

		} else {
			crt.InputError(notations.ErrInvalidAction + support.SQuote(nextAction))
		}
	}
}

func Detail(crt *support.Crt, info plex.Metadata) {

	p := page.New(info.Title)

	p.AddFieldValuePair(crt, notations.TitleLabel, info.Title)
	p.AddFieldValuePair(crt, notations.SummaryLabel, info.Summary)

	count := 0
	p.BlankRow()
	p.AddColumnsTitle(crt, notations.ContainerLabel, notations.ResolutionLabel, notations.CodecLabel, notations.AspectRatioLabel, notations.FrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case page.TxtQuit:
		return
	default:
		crt.InputError(notations.ErrInvalidAction + support.SQuote(nextAction))
	}
}

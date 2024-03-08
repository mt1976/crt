package music

import (
	"fmt"
	"os"

	"github.com/jrudio/go-plex-client"
	e "github.com/mt1976/crt/errors"
	notations "github.com/mt1976/crt/language"
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
			//	Action(crt, mediaVault, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
			Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])

		} else {
			crt.InputError(e.ErrInvalidAction + support.SQuote(nextAction))
		}
	}
}

func Detail(crt *support.Crt, info plex.Metadata) {

	p := page.New(info.Title)

	p.AddFieldValuePair(crt, notations.TxtPlexTitleLabel, info.Title)
	p.AddFieldValuePair(crt, notations.TxtPlexSummaryLabel, info.Summary)

	count := 0
	p.BlankRow()
	p.AddColumnsTitle(crt, notations.TxtPlexContainerLabel, notations.TxtPlexResolutionLabel, notations.TxtPlexCodecLabel, notations.TxtPlexAspectRatioLabel, notations.TxtPlexFrameRateLabel)

	for range info.Media {
		med := info.Media[count]
		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		count++
	}

	nextAction, _ := p.Display(crt)
	switch nextAction {
	case t.SymActionQuit:
		return
	default:
		crt.InputError(e.ErrInvalidAction + support.SQuote(nextAction))
	}
}

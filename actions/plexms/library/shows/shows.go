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

	exit := false
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
				//	Action(crt, mediaVault, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
				Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1], mediaVault)
			} else {
				crt.InputError(page.ErrInvalidAction + "'" + nextAction + "'")
			}
		}
	}

	//spew.Dump(res)
	//os.Exit(1)
}

func Detail(crt *support.Crt, info plex.Metadata, mediaVault *plex.Plex) {
	p := page.New(info.Title)
	//p.Add("Title:"+info.Title, "", "")
	//AddFieldValuePair(crt, p, "Title", info.Title)
	p.AddFieldValuePair(crt, "Title", info.Title)
	//p.Add("ContentRating:"+info.ContentRating, "", "")
	//AddFieldValuePair(crt, p, "ContentRating", info.ContentRating)
	p.AddFieldValuePair(crt, "ContentRating", info.ContentRating)

	//p.Add("Summary:"+info.Summary, "", "")
	//AddFieldValuePair(crt, p, "Summary", info.Summary)
	p.AddFieldValuePair(crt, "Summary", info.Summary)
	//unix time to hrs mins secs
	//xy := time.UnixMilli(int64(info.Duration))
	//dur := xy.Format("15:04:05")
	//p.Add("Duration:"+dur, "", "")
	//AddFieldValuePair(crt, p, "Duration", dur)
	//p.AddFieldValuePair(crt, "Duration", dur)
	count := 0
	p.Add("---", "", "")
	//AddColumns(true, crt, p, "Container", "Resolution", "Codec", "Aspect","FrameRate")
	p.AddColumns(crt, "Container", "Resolution", "Codec", "Aspect", "FrameRate")
	p.AddColumnsRuler(crt, "Container", "Resolution", "Codec", "Aspect", "FrameRate")

	for range info.Media {
		med := info.Media[count]

		p.AddColumns(crt, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String(), med.VideoFrameRate)
		//AddFieldValuePair(crt, p, "AudioCodec", med.AudioCodec)
		//p.Add("FrameRate:"+med.VideoFrameRate, "", "")
		//AddFieldValuePair(crt, p, "FrameRate", med.VideoFrameRate)
		//p.Add("AudioCodec:"+med.AudioCodec, "", "")
		//AddC(crt, p, "AudioChannels", "AudioCodec", "Bitrate", "")
		//AddC(crt, p, fmt.Sprintf("%d", med.AudioChannels), med.AudioCodec, fmt.Sprintf("%d", med.Bitrate), "")
		count++
	}
	p.AddAction(Seasons) //Drilldown to episodes
	p.SetPrompt(prompt)

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
		case Seasons:
			SeasonDetails(crt, mediaVault, info)
		default:
			crt.InputError(page.ErrInvalidAction + "'" + nextAction + "'")
		}
	}
}

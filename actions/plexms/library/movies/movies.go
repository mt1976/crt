package movies

import (
	"fmt"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/support"
	menu "github.com/mt1976/crt/support/menu"
	page "github.com/mt1976/crt/support/page"
)

func Run(crt *support.Crt, mediaVault *plex.Plex, wi plex.Directory) {
	crt.Shout("Movies-" + wi.Title)

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error("mvLibError", err)
		os.Exit(1)
	}

	noItems := fmt.Sprintf("%d", res.MediaContainer.Size)

	m := menu.New(res.MediaContainer.LibrarySectionTitle + " (" + noItems + ")")
	count := 0

	for range res.MediaContainer.Metadata {
		count++
		m.Add(count, res.MediaContainer.Metadata[count-1].Title, "", "")
	}

	exit := false
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
				//	Action(crt, mediaVault, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
				Detail(crt, res.MediaContainer.Metadata[support.ToInt(nextAction)-1])
			} else {
				crt.InputError(menu.InvalidActionError + "'" + nextAction + "'")
			}
		}
	}

	spew.Dump(res)
	os.Exit(1)
}

func Detail(crt *support.Crt, info plex.Metadata) {
	p := page.New(info.Title)
	//p.Add("Title:"+info.Title, "", "")
	AddFieldValuePair(crt, p, "Title", info.Title)
	//p.Add("ContentRating:"+info.ContentRating, "", "")
	AddFieldValuePair(crt, p, "ContentRating", info.ContentRating)

	//p.Add("Summary:"+info.Summary, "", "")
	AddFieldValuePair(crt, p, "Summary", info.Summary)
	//unix time to hrs mins secs
	xy := time.UnixMilli(int64(info.Duration))
	dur := xy.Format("15:04:05")
	//p.Add("Duration:"+dur, "", "")
	AddFieldValuePair(crt, p, "Duration", dur)
	count := 0
	p.Add("---", "", "")
	AddC(crt, p, "Container", "Resolution", "Codec", "Aspect")

	for range info.Media {
		med := info.Media[count]

		//p.Add("Container:"+med.Container, "", "")
		//AddFieldValuePair(crt, p, "Container", med.Container)
		//AddFieldValuePair(crt, p, "VideoResolution", med.VideoResolution)

		//AddFieldValuePair(crt, p, "VideoCodec", med.VideoCodec)
		//AddFieldValuePair(crt, p, "AspectRatio", med.AspectRatio.String())

		AddC(crt, p, med.Container, med.VideoResolution, med.VideoCodec, med.AspectRatio.String())
		//p.Add("", "", "")
		//p.Add("Resolution:"+med.VideoResolution, "", "")
		//AddFieldValuePair(crt, p, "AudioCodec", med.AudioCodec)
		//p.Add("FrameRate:"+med.VideoFrameRate, "", "")
		//AddFieldValuePair(crt, p, "FrameRate", med.VideoFrameRate)
		//p.Add("AudioCodec:"+med.AudioCodec, "", "")
		//AddC(crt, p, "AudioChannels", "AudioCodec", "Bitrate", "")
		//AddC(crt, p, fmt.Sprintf("%d", med.AudioChannels), med.AudioCodec, fmt.Sprintf("%d", med.Bitrate), "")
		count++
	}
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

func AddFieldValuePair(crt *support.Crt, p *page.Page, key string, value string) {
	format := "%-20s : %s\n"
	p.Add(fmt.Sprintf(format, key, crt.Bold(value)), "", "")
	//return p
}

func AddC(crt *support.Crt, p *page.Page, col1 string, col2 string, col3 string, col4 string) {
	format := "%-16s : %-16s : %-16s : %-16s\n"
	result := fmt.Sprintf(format, crt.Bold(col1), crt.Bold(col2), crt.Bold(col3), crt.Bold(col4))

	p.Add(result, "", "")
	//return p
}

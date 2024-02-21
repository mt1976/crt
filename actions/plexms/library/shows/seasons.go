package shows

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/support"
	menu "github.com/mt1976/crt/support/menu"
)

func SeasonDetails(crt *support.Crt, mediaVault *plex.Plex, info plex.Metadata) {

	//key := C.PlexURI + ":" + C.PlexPort + info.Key
	spew.Dump(info)
	yy, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		crt.Error("mvLibError", err)
		os.Exit(1)
	}
	p := menu.New("Seasons - " + info.Title)
	spew.Dump(yy)
	noResps := len(yy.MediaContainer.Metadata)
	for i := 0; i < noResps; i++ {
		season := yy.MediaContainer.Metadata[i]
		p.Add(i+1, season.Title, "", "")
	}

	exit := false
	for !exit {
		na, _ := p.Display(crt)
		switch na {
		case menu.Quit:
			exit = true
			return
		case menu.Forward:
			p.NextPage(crt)
		case menu.Back:
			p.PreviousPage(crt)
		default:
			if support.IsInt(na) {
				Episodes(crt, mediaVault, info.Title, yy.MediaContainer.Metadata[support.ToInt(na)-1])
			} else {
				crt.InputError(menu.InvalidActionError + "'" + na + "'")
			}
		}
	}
}

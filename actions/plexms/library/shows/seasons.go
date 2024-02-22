package shows

import (
	"os"

	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/support"
	page "github.com/mt1976/crt/support/page"
)

func SeasonDetails(crt *support.Crt, mediaVault *plex.Plex, info plex.Metadata) {

	//key := C.PlexURI + ":" + C.PlexPort + info.Key
	//spew.Dump(info)
	yy, err := mediaVault.GetEpisodes(info.RatingKey)
	if err != nil {
		crt.Error("mvLibError", err)
		os.Exit(1)
	}
	p := page.New("Seasons - " + info.Title)
	//spew.Dump(yy)
	noResps := len(yy.MediaContainer.Metadata)
	for i := 0; i < noResps; i++ {
		season := yy.MediaContainer.Metadata[i]
		p.AddOption(i+1, season.Title, "", "")
	}

	//	exit := false
	//for !exit {
	na, _ := p.Display(crt)
	switch na {
	case page.Quit:
		//			exit = true
		return
	default:
		if support.IsInt(na) {
			Episodes(crt, mediaVault, info.Title, yy.MediaContainer.Metadata[support.ToInt(na)-1])
		} else {
			crt.InputError(page.ErrInvalidAction + "'" + na + "'")
		}
	}
	// }
}

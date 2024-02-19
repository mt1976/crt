package movies

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/support"
)

func Run(crt *support.Crt, mediaVault *plex.Plex, wi plex.Directory) {
	crt.Shout("Movies-" + wi.Title)

	res, err := mediaVault.GetLibraryContent(wi.Key, "")
	if err != nil {
		crt.Error("mvLibError", err)
		os.Exit(1)
	}
	spew.Dump(res)
	os.Exit(1)
}

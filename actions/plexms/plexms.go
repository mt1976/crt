package plexmediaserver

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/jrudio/go-plex-client"
	support "github.com/mt1976/crt/support"
	cfg "github.com/mt1976/crt/support/config"
	page "github.com/mt1976/crt/support/page"
)

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(crt *support.Crt) {

	crt.Clear()

	//spew.Dump(cfg.Configuration)
	//os.Exit(1)

	plexConnection, err := plex.New(cfg.Configuration.PlexURI+":"+cfg.Configuration.PlexPort, cfg.Configuration.PlexToken)
	if err != nil {
		crt.Error(plexInitError, err)
		os.Exit(1)
	}

	// Test your connection to your Plex server
	result, err := plexConnection.Test()
	if err != nil || !result {
		crt.Error(plexTestError, err)
		os.Exit(1)
	}

	devices, err := plexConnection.GetServers()
	if err != nil {
		crt.Error(plexInitError, err)
		os.Exit(1)
	}
	//spew.Dump(devices)

	mediaV := 0
	for i := 0; i < len(devices); i++ {
		if devices[i].ClientIdentifier == cfg.Configuration.PlexClientID {
			mediaV = i
		}
	}

	mediaVaultProperties := devices[mediaV]
	spew.Dump(mediaVaultProperties)

	mediaVault, err := plex.New(mediaVaultProperties.Connection[0].URI, cfg.Configuration.PlexToken)
	if err != nil {
		crt.Error(fmt.Sprintf(mvInitError, mediaVaultProperties.Name), err)
		os.Exit(1)
	}

	mvLibraries, err := mediaVault.GetLibraries()
	if err != nil {
		crt.Error(fmt.Sprintf(mvLibError, mediaVaultProperties.Name), err)
		os.Exit(1)
	}

	yy := mvLibraries.MediaContainer.Directory[0]

	res, err := mediaVault.GetLibraryContent(yy.Key, "")
	if err != nil {
		crt.Error(fmt.Sprintf(mvLibError, mediaVaultProperties.Name), err)
		os.Exit(1)
	}

	spew.Dump(res)
	//os.Exit(1)
	//spew.Dump(libs)
	//os.Exit(1)

	p := page.New(plexTitle + " - " + mediaVaultProperties.Name)

	for mvLibrary := range mvLibraries.MediaContainer.Directory {
		xx := mvLibraries.MediaContainer.Directory[mvLibrary]
		p.Add(xx.Title, "", "")
	}

	p.AddAction(page.Quit)
	p.AddAction(page.Forward)
	p.AddAction(page.Back)
	ok := false
	for !ok {

		nextAction, _ := p.Display(crt)
		switch nextAction {
		case page.Forward:
			p.NextPage(crt)
		case page.Back:
			p.PreviousPage(crt)
		case page.Quit:
			ok = true
			return
		default:
			crt.InputError(page.InvalidActionError + "'" + nextAction + "'")
		}
	}

}

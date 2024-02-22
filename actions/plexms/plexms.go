package plexmediaserver

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/actions/plexms/library/movies"
	"github.com/mt1976/crt/actions/plexms/library/music"
	"github.com/mt1976/crt/actions/plexms/library/shows"
	support "github.com/mt1976/crt/support"
	cfg "github.com/mt1976/crt/support/config"
	"github.com/mt1976/crt/support/page"
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
	//spew.Dump(mediaVaultProperties)

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

	//	yy := mvLibraries.MediaContainer.Directory[0]

	// res, err := mediaVault.GetLibraryContent(yy.Key, "")
	// if err != nil {
	// 	crt.Error(fmt.Sprintf(mvLibError, mediaVaultProperties.Name), err)
	// 	os.Exit(1)
	// }

	//spew.Dump(res)
	//os.Exit(1)
	//spew.Dump(libs)
	//os.Exit(1)

	p := page.New(plexTitle + " - " + mediaVaultProperties.Name)
	count := 0
	for mvLibrary := range mvLibraries.MediaContainer.Directory {
		xx := mvLibraries.MediaContainer.Directory[mvLibrary]
		count++
		p.AddOption(count, xx.Title, "", "")
	}

	p.AddAction(page.Quit)
	p.AddAction(page.Forward)
	p.AddAction(page.Back)
	ok := false
	for !ok {

		nextAction, _ := p.Display(crt)
		switch {
		case nextAction == page.Forward:
			p.NextPage(crt)
		case nextAction == page.Back:
			p.PreviousPage(crt)
		case nextAction == page.Quit:
			ok = true
			return
		case support.IsInt(nextAction):
			crt.Error("You selected: "+nextAction, nil)
			naInt, _ := strconv.Atoi(nextAction)
			wi := mvLibraries.MediaContainer.Directory[naInt-1]
			Action(crt, mediaVault, &wi)
			//spew.Dump(wi)
			//os.Exit(1)

		default:
			crt.InputError(page.ErrInvalidAction + "'" + nextAction + "'")
		}
	}

}

func Action(crt *support.Crt, mediaVault *plex.Plex, wi *plex.Directory) {

	switch wi.Type {
	case "movie":
		crt.Shout(wi.Title)
		movies.Run(crt, mediaVault, wi)
	case "show":
		crt.Shout(wi.Title)
		shows.Run(crt, mediaVault, wi)
	case "artist":
		crt.Shout(wi.Title)
		music.Run(crt, mediaVault, wi)
	default:
		crt.Shout(wi.Title)
	}
}

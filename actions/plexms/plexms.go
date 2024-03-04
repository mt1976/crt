package plexmediaserver

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jrudio/go-plex-client"
	"github.com/mt1976/crt/actions/plexms/library/movies"
	"github.com/mt1976/crt/actions/plexms/library/music"
	"github.com/mt1976/crt/actions/plexms/library/shows"
	"github.com/mt1976/crt/actions/plexms/notations"
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
		crt.Error(notations.ErrPlexInit, err)
		os.Exit(1)
	}

	// Test your connection to your Plex server
	result, err := plexConnection.Test()
	if err != nil || !result {
		crt.Error(notations.ErrPlexConnectionTest, err)
		os.Exit(1)
	}

	devices, err := plexConnection.GetServers()
	if err != nil {
		crt.Error(notations.ErrPlexInit, err)
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
		crt.Error(fmt.Sprintf(notations.ErrPlexConnect, mediaVaultProperties.Name), err)
		os.Exit(1)
	}

	mvLibraries, err := mediaVault.GetLibraries()
	if err != nil {
		crt.Error(fmt.Sprintf(notations.ErrLibraryResponse, mediaVaultProperties.Name), err)
		os.Exit(1)
	}

	p := page.New(notations.PlexTitle + notations.DelimiterText + mediaVaultProperties.Name)
	count := 0
	for mvLibrary := range mvLibraries.MediaContainer.Directory {
		xx := mvLibraries.MediaContainer.Directory[mvLibrary]
		count++
		p.AddOption(count, xx.Title, "", "")
	}

	p.AddAction(page.QuitText)
	p.AddAction(page.ForwardText)
	p.AddAction(page.BackText)

	nextAction, _ := p.Display(crt)
	switch {
	case nextAction == page.QuitText:
		return
	case support.IsInt(nextAction):
		crt.Error(notations.InfoYouSelected+nextAction, nil)
		naInt, _ := strconv.Atoi(nextAction)
		wi := mvLibraries.MediaContainer.Directory[naInt-1]
		Action(crt, mediaVault, &wi)

	default:
		crt.InputError(notations.ErrInvalidAction + support.SQuote(nextAction))
	}
	//}

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

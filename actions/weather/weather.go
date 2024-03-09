package weather

import (
	"fmt"
	"os"
	"strings"

	owm "github.com/briandowns/openweathermap"
	e "github.com/mt1976/crt/errors"
	t "github.com/mt1976/crt/language"
	support "github.com/mt1976/crt/support"
	"github.com/mt1976/crt/support/config"
	page "github.com/mt1976/crt/support/page"
)

var C config.Config

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(crt *support.Crt) {

	crt.Clear()
	p := page.New(t.TxtWeatherTitle + t.Space + t.TxtSourceService)

	w, err := owm.NewCurrent(C.OpenWeatherMapApiUnits, C.OpenWeatherMapApiLang, C.OpenWeatherMapApiKey)
	if err != nil {
		crt.Error(fmt.Sprintf(e.ErrOpenWeather, err), err)
		os.Exit(1)
		return
	}

	w.CurrentByCoordinates(
		&owm.Coordinates{Latitude: C.LocationLatitude, Longitude: C.LocationLogitude})
	if err != nil {
		crt.Error(fmt.Sprintf(e.ErrOpenWeather, err), err)
		os.Exit(1)
		return
	}

	c := 0
	c++
	p.Add(fmt.Sprintf(t.SymWeatherFormat2, t.TxtLocationLabel, crt.Bold(w.Name)), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat2, t.TxtConditionsLabel, crt.Bold(w.Weather[0].Main)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtTemperatureLabel, boldFloat(crt, w.Main.Temp)+t.SymDegree, t.TxtFeelsLikeLabel, boldFloat(crt, w.Main.FeelsLike)+t.SymDegree), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtMinLabel, boldFloat(crt, w.Main.TempMin)+t.SymDegree, t.TxtMaxLabel, boldFloat(crt, w.Main.TempMax)+t.SymDegree), "", "")
	//p.Add(hr())
	p.Add(hr(crt), "", "")
	// p.Add(fmt.Sprintf("Feels Like : %v", w.Main.FeelsLike))
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtWindSpeedLabel, boldFloat(crt, w.Wind.Speed), t.TxtWindDirectionLabel, boldFloat(crt, w.Wind.Deg)), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat1, t.TxtCloudCoverLabel, boldInt(crt, w.Clouds.All)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtRain1Hr, boldFloat(crt, w.Rain.OneH), t.TxtRain3Hr, boldFloat(crt, w.Rain.ThreeH)), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtSnow1Hr, boldFloat(crt, w.Snow.OneH), t.TxtSnow3Hr, boldFloat(crt, w.Snow.ThreeH)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtSunriseLabel, crt.Bold(outdate(w.Sys.Sunrise)), t.TxtSunsetLabel, crt.Bold(outdate(w.Sys.Sunset))), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat2, t.TxtSourceLabel, crt.Bold(t.TxtSourceService)), "", "")
	// INSERT CONTENT ABOVE
	p.AddAction(t.SymActionQuit)
	p.AddAction(t.SymActionForward)
	p.AddAction(t.SymActionBack)
	ok := false
	for !ok {

		nextAction, _ := p.Display(crt)
		switch nextAction {
		case t.SymActionForward:
			p.NextPage(crt)
		case t.SymActionBack:
			p.PreviousPage(crt)
		case t.SymActionQuit:
			ok = true
			return
		default:
			crt.InputError(e.ErrInvalidAction + t.SymSingleQuote + nextAction + t.SymSingleQuote)
		}
	}

}

func outdate(t int) string {
	// int to int64
	// unix date to human date
	return support.HumanFromUnixDate(int64(t))
}

// The `hr` function returns a string consisting of a line of dashes.
func hr(crt *support.Crt) string {
	return strings.Repeat(t.SymBreak, 3)
}

func boldFloat(crt *support.Crt, in float64) string {
	return crt.Bold(fmt.Sprintf("%v", in))
}

func boldInt(crt *support.Crt, in int) string {
	return crt.Bold(fmt.Sprintf("%v", in))
}

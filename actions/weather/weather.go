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
	p.Add(fmt.Sprintf(t.SymWeatherFormat2, t.TxtLocation, crt.Bold(w.Name)), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat2, t.TxtConditions, crt.Bold(w.Weather[0].Main)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtTemperature, boldFloat(crt, w.Main.Temp)+t.SymDegree, t.TxtFeelsLike, boldFloat(crt, w.Main.FeelsLike)+t.SymDegree), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtMin, boldFloat(crt, w.Main.TempMin)+t.SymDegree, t.TxtMax, boldFloat(crt, w.Main.TempMax)+t.SymDegree), "", "")
	//p.Add(hr())
	p.Add(hr(crt), "", "")
	// p.Add(fmt.Sprintf("Feels Like : %v", w.Main.FeelsLike))
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtWindSpeed, boldFloat(crt, w.Wind.Speed), t.TxtWindDirection, boldFloat(crt, w.Wind.Deg)), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat1, t.TxtCloudCover, boldInt(crt, w.Clouds.All)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, TxtRain1Hr, boldFloat(crt, w.Rain.OneH), TxtRain3Hr, boldFloat(crt, w.Rain.ThreeH)), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, TxtSnow1Hr, boldFloat(crt, w.Snow.OneH), TxtSnow3Hr, boldFloat(crt, w.Snow.ThreeH)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat4, t.TxtSunrise, crt.Bold(outdate(w.Sys.Sunrise)), t.TxtSunset, crt.Bold(outdate(w.Sys.Sunset))), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(t.SymWeatherFormat2, t.TxtSource, crt.Bold(t.TxtSourceService)), "", "")
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
			crt.InputError(e.ErrInvalidAction + t.SingleQuote + nextAction + t.SingleQuote)
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

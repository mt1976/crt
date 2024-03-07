package weather

import (
	"fmt"
	"os"
	"strings"

	owm "github.com/briandowns/openweathermap"
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
	p := page.New(TxtWeatherTitle + " " + TxtSourceService)

	w, err := owm.NewCurrent(C.OpenWeatherMapApiUnits, C.OpenWeatherMapApiLang, C.OpenWeatherMapApiKey)
	if err != nil {
		crt.Error(fmt.Sprintf(ErrOpenWeather, err), err)
		os.Exit(1)
		return
	}

	w.CurrentByCoordinates(
		&owm.Coordinates{Latitude: C.LocationLatitude, Longitude: C.LocationLogitude})
	if err != nil {
		crt.Error(fmt.Sprintf(ErrOpenWeather, err), err)
		os.Exit(1)
		return
	}

	c := 0
	c++
	p.Add(fmt.Sprintf(SymWeatherFormat2, TxtLocation, crt.Bold(w.Name)), "", "")
	p.Add(fmt.Sprintf(SymWeatherFormat2, TxtConditions, crt.Bold(w.Weather[0].Main)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(SymWeatherFormat4, TxtTemperature, boldFloat(crt, w.Main.Temp)+SymDegree, TxtFeelsLike, boldFloat(crt, w.Main.FeelsLike)+SymDegree), "", "")
	p.Add(fmt.Sprintf(SymWeatherFormat4, TxtMin, boldFloat(crt, w.Main.TempMin)+SymDegree, TxtMax, boldFloat(crt, w.Main.TempMax)+SymDegree), "", "")
	//p.Add(hr())
	p.Add(hr(crt), "", "")
	// p.Add(fmt.Sprintf("Feels Like : %v", w.Main.FeelsLike))
	p.Add(fmt.Sprintf(SymWeatherFormat4, TxtWindSpeed, boldFloat(crt, w.Wind.Speed), TxtWindDirection, boldFloat(crt, w.Wind.Deg)), "", "")
	p.Add(fmt.Sprintf(SymWeatherFormat1, TxtCloudCover, boldInt(crt, w.Clouds.All)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(SymWeatherFormat4, TxtRain1Hr, boldFloat(crt, w.Rain.OneH), TxtRain3Hr, boldFloat(crt, w.Rain.ThreeH)), "", "")
	p.Add(fmt.Sprintf(SymWeatherFormat4, TxtSnow1Hr, boldFloat(crt, w.Snow.OneH), TxtSnow3Hr, boldFloat(crt, w.Snow.ThreeH)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(SymWeatherFormat4, TxtSunrise, crt.Bold(outdate(w.Sys.Sunrise)), TxtSunset, crt.Bold(outdate(w.Sys.Sunset))), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(SymWeatherFormat2, TxtSource, crt.Bold(TxtSourceService)), "", "")
	// INSERT CONTENT ABOVE
	p.AddAction(page.TxtQuit)
	p.AddAction(page.TxtForward)
	p.AddAction(page.TxtBack)
	ok := false
	for !ok {

		nextAction, _ := p.Display(crt)
		switch nextAction {
		case page.TxtForward:
			p.NextPage(crt)
		case page.TxtBack:
			p.PreviousPage(crt)
		case page.TxtQuit:
			ok = true
			return
		default:
			crt.InputError(page.ErrInvalidAction + "'" + nextAction + "'")
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
	return strings.Repeat(" ━━ ", 3)
}

func boldFloat(crt *support.Crt, in float64) string {
	return crt.Bold(fmt.Sprintf("%v", in))
}

func boldInt(crt *support.Crt, in int) string {
	return crt.Bold(fmt.Sprintf("%v", in))
}

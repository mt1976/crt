package weather

import (
	"fmt"
	"os"
	"strings"

	owm "github.com/briandowns/openweathermap"
	support "github.com/mt1976/crt/support"
	page "github.com/mt1976/crt/support/page"
)

// The main function initializes and runs a terminal-based news reader application called StarTerm,
// which fetches news headlines from an RSS feed and allows the user to navigate and open the full news
// articles.
func Run(crt *support.Crt) {

	crt.Clear()
	p := page.New(weatherTitle + " " + sourceServiceText)

	w, err := owm.NewCurrent(apiUnits, apiLang, apiKey)
	if err != nil {
		crt.Error(fmt.Sprintf(owmInitError, err), err)
		os.Exit(1)
		return
	}

	w.CurrentByCoordinates(
		&owm.Coordinates{Latitude: apiLatitude, Longitude: apiLongitude})
	if err != nil {
		crt.Error(fmt.Sprintf(owmInitError, err), err)
		os.Exit(1)
		return
	}

	c := 0
	c++
	p.Add(fmt.Sprintf(weatherFormat2, locationText, crt.Bold(w.Name)), "", "")
	p.Add(fmt.Sprintf(weatherFormat2, conditionsText, crt.Bold(w.Weather[0].Main)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(weatherFormat4, temperatureText, boldFloat(crt, w.Main.Temp)+degreeText, feelsLikeText, boldFloat(crt, w.Main.FeelsLike)+degreeText), "", "")
	p.Add(fmt.Sprintf(weatherFormat4, minText, boldFloat(crt, w.Main.TempMin)+degreeText, maxText, boldFloat(crt, w.Main.TempMax)+degreeText), "", "")
	//p.Add(hr())
	p.Add(hr(crt), "", "")
	// p.Add(fmt.Sprintf("Feels Like : %v", w.Main.FeelsLike))
	p.Add(fmt.Sprintf(weatherFormat4, windSpeedText, boldFloat(crt, w.Wind.Speed), windDirectionText, boldFloat(crt, w.Wind.Deg)), "", "")
	p.Add(fmt.Sprintf(weatherFormat1, cloudCoverText, boldInt(crt, w.Clouds.All)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(weatherFormat4, rain1HrText, boldFloat(crt, w.Rain.OneH), rain3HrText, boldFloat(crt, w.Rain.ThreeH)), "", "")
	p.Add(fmt.Sprintf(weatherFormat4, snow1HrText, boldFloat(crt, w.Snow.OneH), snow3HrText, boldFloat(crt, w.Snow.ThreeH)), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(weatherFormat4, sunriseText, crt.Bold(outdate(w.Sys.Sunrise)), sunsetText, crt.Bold(outdate(w.Sys.Sunset))), "", "")
	p.Add(hr(crt), "", "")
	p.Add(fmt.Sprintf(weatherFormat2, sourceText, crt.Bold(sourceServiceText)), "", "")
	// INSERT CONTENT ABOVE
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

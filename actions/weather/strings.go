package weather

// var linkDisplay string = '\e]8;;{{URL}}\e\\{{TEXT}}\e]8;;\e\\'

const (
	TxtWeatherTitle string = "Weather"

	TxtPrompt string = "Select (Q)uit"

	ErrOpenWeather    string = "failed to initialize OpenWeatherMap: %v"
	SymWeatherFormat2 string = "%-25s | %-15v"
	SymWeatherFormat4 string = "%-25s | %-15v  %-15s : %-15v"
	SymWeatherFormat1 string = "%-25s | %v%%"
	TxtLocation       string = "Location"
	TxtConditions     string = "Conditions"
	TxtTemperature    string = "Temperature"
	TxtFeelsLike      string = "Feels Like"
	TxtMin            string = "Min"
	TxtMax            string = "Max"
	TxtWindSpeed      string = "Wind Speed"
	TxtWindDirection  string = "Wind Direction"
	TxtCloudCover     string = "Cloud Cover"
	TxtRain           string = "Rain"
	TxtSnow           string = "Snow"
	TxtSunrise        string = "Sunrise"
	TxtSunset         string = "Sunset"
	TxtSource         string = "Source"
	TxtSourceService  string = "OpenWeatherMap"
	SymDegree         string = "°"
)

var TxtRain1Hr string = TxtRain + " (1hr)"
var TxtRain3Hr string = TxtRain + " (3hr)"
var TxtSnow1Hr string = TxtSnow + " (1hr)"
var TxtSnow3Hr string = TxtSnow + " (3hr)"

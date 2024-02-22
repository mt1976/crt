package weather

// var linkDisplay string = '\e]8;;{{URL}}\e\\{{TEXT}}\e]8;;\e\\'

const (
	weatherTitle string = "Weather"

	prompt string = "Select (Q)uit"

	ErrOpenWeather     string = "failed to initialize OpenWeatherMap: %v"
	weatherFormat2     string = "%-25s | %-15v"
	weatherFormat4     string = "%-25s | %-15v  %-15s : %-15v"
	weatherFormat1     string = "%-25s | %v%%"
	locationLabel      string = "Location"
	conditionsLabel    string = "Conditions"
	temperatureLabel   string = "Temperature"
	feelsLikeLabel     string = "Feels Like"
	minLabel           string = "Min"
	maxLabel           string = "Max"
	windSpeedLabel     string = "Wind Speed"
	windDirectionLabel string = "Wind Direction"
	cloudCoverLabel    string = "Cloud Cover"
	rainLabel          string = "Rain"
	snowLabel          string = "Snow"
	sunriseLabel       string = "Sunrise"
	sunsetLabel        string = "Sunset"
	sourceLabel        string = "Source"
	sourceServiceText  string = "OpenWeatherMap"
	degreeLabel        string = "°"
)

var rain1HrLabel string = rainLabel + " (1hr)"
var rain3HrLabel string = rainLabel + " (3hr)"
var snow1HrLabel string = snowLabel + " (1hr)"
var snow3HrLabel string = snowLabel + " (3hr)"

package weather

var weatherTitle string = "Weather"

// var linkDisplay string = '\e]8;;{{URL}}\e\\{{TEXT}}\e]8;;\e\\'

var apiKey string = "ab8f78d5469dce541859d896f5483418"
var apiLang string = "EN"
var apiUnits string = "C"
var prompt string = "Select (Q)uit"

var owmInitError string = "failed to initialize OpenWeatherMap: %v"
var weatherFormat2 string = "%-25s | %-15v"
var weatherFormat4 string = "%-25s | %-15v  %-15s : %-15v"
var weatherFormat1 string = "%-25s | %v%%"

var locationText = "Location"
var conditionsText = "Conditions"
var temperatureText = "Temperature"
var feelsLikeText = "Feels Like"
var minText = "Min"
var maxText = "Max"
var windSpeedText = "Wind Speed"
var windDirectionText = "Wind Direction"
var cloudCoverText = "Cloud Cover"
var rainText = "Rain"
var snowText = "Snow"
var rain1HrText = rainText + " (1hr)"
var rain3HrText = rainText + " (3hr)"
var snow1HrText = snowText + " (1hr)"
var snow3HrText = snowText + " (3hr)"
var sunriseText = "Sunrise"
var sunsetText = "Sunset"
var sourceText = "Source"
var sourceServiceText = "OpenWeatherMap"
var degreeText string = "°"

package config

import (
	"github.com/gen2brain/beeep"
	"github.com/spf13/viper"
)

const (
// PlexDateFormat        string = "2006-01-02"
// ApplicationDateFormat string = "02 Jan 2006"
)

type Config struct {
	// 	TermWidth  int    `pkl:"term_width"`
	// 	TermHeight int    `pkl:"term_height"`
	// 	Baud       int    `pkl:"baud"`
	// 	Delay      int    `pkl:"delay"`
	PlexURI                    string  `mapstructure:"PlexURI"`
	PlexPort                   string  `mapstructure:"PlexPort"`
	PlexToken                  string  `mapstructure:"PlexToken"`
	PlexClientID               string  `mapstructure:"PlexClientID"`
	PlexDateFormat             string  `mapstructure:"PlexDateFormat"`
	ApplicationDateFormat      string  `mapstructure:"ApplicationDateFormat"`
	ApplicationDateFormatShort string  `mapstructure:"ApplicationDateFormatShort"`
	ApplicationTimeFormat      string  `mapstructure:"ApplicationTimeFormat"`
	TerminalWidth              int     `mapstructure:"TerminalWidth"`
	TerminalHeight             int     `mapstructure:"TerminalHeight"`
	Delay                      float64 `mapstructure:"Delay"`
	Baud                       int     `mapstructure:"Baud"`
	TransmissionURI            string  `mapstructure:"TransmissionURI"`
	QTorrentURI                string  `mapstructure:"QTorrentURI"`
	MaxContentRows             int     `mapstructure:"MaxContentRows"`
	MaxNoItems                 int     `mapstructure:"MaxNoItems"`
	TitleLength                int     `mapstructure:"TitleLength"`
	Debug                      bool    `mapstructure:"Debug"`

	OpenWeatherMapApiKey   string `mapstructure:"OpenWeatherMapApiKey"`
	OpenWeatherMapApiLang  string `mapstructure:"OpenWeatherMapApiLang"`
	OpenWeatherMapApiUnits string `mapstructure:"OpenWeatherMapApiUnits"`

	LocationLogitude float64 `mapstructure:"LocationLongitude"`
	LocationLatitude float64 `mapstructure:"LocationLatitude"`

	URISkyNews              string `mapstructure:"SkyNewsURI"`
	URISkyNewsHome          string `mapstructure:"SkyNewsHomeURI"`
	URISkyNewsUK            string `mapstructure:"SkyNewsUKURI"`
	URISkyNewsWorld         string `mapstructure:"SkyNewsWorldURI"`
	URISkyNewsUS            string `mapstructure:"SkyNewsUSURI"`
	URISkyNewsBusiness      string `mapstructure:"SkyNewsBusinessURI"`
	URISkyNewsPolitics      string `mapstructure:"SkyNewsPoliticsURI"`
	URISkyNewsTechnology    string `mapstructure:"SkyNewsTechnologyURI"`
	URISkyNewsEntertainment string `mapstructure:"SkyNewsEntertainmentURI"`
	URISkyNewsStrange       string `mapstructure:"SkyNewsStrangeURI"`

	DefaultErrorDelay    float64 `mapstructure:"DefaultErrorDelay"`
	DefaultRandomPortMin int     `mapstructure:"DefaultRandomPortMin"`
	DefaultRandomPortMax int     `mapstructure:"DefaultRandomPortMax"`
	DefaultRandomMACMin  int     `mapstructure:"DefaultRandomMACMin"`
	DefaultRandomMACMax  int     `mapstructure:"DefaultRandomMACMax"`
	DefaultRandomIPMin   int     `mapstructure:"DefaultRandomIPMin"`
	DefaultRandomIPMax   int     `mapstructure:"DefaultRandomIPMax"`
	DefaultBaud          int     `mapstructure:"DefaultBaud"`

	DefaultBeepDuration  int
	DefaultBeepFrequency float64

	ValidBaudRates []int
}

var Configuration = Config{}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Configuration)
	if err != nil {
		panic(err)
	}

	Configuration.DefaultBeepDuration = beeep.DefaultDuration
	Configuration.DefaultBeepFrequency = beeep.DefaultFreq
	Configuration.ValidBaudRates = []int{0, 300, 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200}
	//Configuration.PlexPort = strconv.Itoa(Configuration.plexPortInt)
	// spew.Dump(&Configuration)
	// os.Exit(1)
}

// const (
// 	defaultErrorDelay float64 = 3.0 // default error delay in seconds
// 	//defaultDateFormat    string  = "/06"
// 	//defaultTimeFormat    string  = "15:04:05"
// 	defaultRandomPortMin int = 1
// 	defaultRandomPortMax int = 65535
// 	defaultRandomMACMin  int = 0
// 	defaultRandomMACMax  int = 255
// 	defaultRandomIPMin   int = 1
// 	defaultRandomIPMax   int = 255
// 	defaultBaud          int = 0
// )

// var BaudRates = []int{0, 300, 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200}
// var defaultBeepFrequency float64 = beeep.DefaultFreq
// var defaultBeepDuration int = beeep.DefaultDuration

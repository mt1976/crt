package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// 	TermWidth  int    `pkl:"term_width"`
	// 	TermHeight int    `pkl:"term_height"`
	// 	Baud       int    `pkl:"baud"`
	// 	Delay      int    `pkl:"delay"`
	PlexURI         string  `mapstructure:"PlexURI"`
	PlexPort        string  `mapstructure:"PlexPort"`
	PlexToken       string  `mapstructure:"PlexToken"`
	PlexClientID    string  `mapstructure:"PlexClientID"`
	TerminalWidth   int     `mapstructure:"TerminalWidth"`
	TerminalHeight  int     `mapstructure:"TerminalHeight"`
	Delay           float64 `mapstructure:"Delay"`
	Baud            int     `mapstructure:"Baud"`
	TransmissionURI string  `mapstructure:"TransmissionURI"`
	QTorrentURI     string  `mapstructure:"QTorrentURI"`
	MaxContentRows  int     `mapstructure:"MaxContentRows"`
	MaxNoItems      int     `mapstructure:"MaxNoItems"`
	TitleLength     int     `mapstructure:"TitleLength"`
	Debug           bool    `mapstructure:"Debug"`

	OpenWeatherMapApiKey   string `mapstructure:"OpenWeatherMapApiKey"`
	OpenWeatherMapApiLang  string `mapstructure:"OpenWeatherMapApiLang"`
	OpenWeatherMapApiUnits string `mapstructure:"OpenWeatherMapApiUnits"`

	LocationLogitude float64 `mapstructure:"LocationLongitude"`
	LocationLatitude float64 `mapstructure:"LocationLatitude"`

	SkyNewsURI              string `mapstructure:"SkyNewsURI"`
	SkyNewsHomeURI          string `mapstructure:"SkyNewsHomeURI"`
	SkyNewsUKURI            string `mapstructure:"SkyNewsUKURI"`
	SkyNewsWorldURI         string `mapstructure:"SkyNewsWorldURI"`
	SkyNewsUSURI            string `mapstructure:"SkyNewsUSURI"`
	SkyNewsBusinessURI      string `mapstructure:"SkyNewsBusinessURI"`
	SkyNewsPoliticsURI      string `mapstructure:"SkyNewsPoliticsURI"`
	SkyNewsTechnologyURI    string `mapstructure:"SkyNewsTechnologyURI"`
	SkyNewsEntertainmentURI string `mapstructure:"SkyNewsEntertainmentURI"`
	SkyNewsStrangeURI       string `mapstructure:"SkyNewsStrangeURI"`
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
	//Configuration.PlexPort = strconv.Itoa(Configuration.plexPortInt)
}

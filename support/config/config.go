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
	SkyNewsURI      string  `mapstructure:"SkyNewsURI"`
	MaxContentRows  int     `mapstructure:"MaxContentRows"`
	TitleLength     int     `mapstructure:"TitleLength"`
	Debug           bool    `mapstructure:"Debug"`
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

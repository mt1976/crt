package config

import (
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	// 	TermWidth  int    `pkl:"term_width"`
	// 	TermHeight int    `pkl:"term_height"`
	// 	Baud       int    `pkl:"baud"`
	// 	Delay      int    `pkl:"delay"`
	PlexURI      string `mapstructure:"PlexURI"`
	plexPortInt  int    `mapstructure:"PlexPort"`
	PlexPort     string
	PlexToken    string `mapstructure:"PlexToken"`
	PlexClientID string `mapstructure:"PlexClientID"`
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
	Configuration.PlexPort = strconv.Itoa(Configuration.plexPortInt)
}

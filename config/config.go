package config

import (
	"strconv"
	"strings"

	beep "github.com/gen2brain/beeep"
	viper "github.com/spf13/viper"
)

type Config struct {
	ApplicationDateFormat      string  `mapstructure:"ApplicationDateFormat"`
	ApplicationDateFormatShort string  `mapstructure:"ApplicationDateFormatShort"`
	ApplicationTimeFormat      string  `mapstructure:"ApplicationTimeFormat"`
	Delay                      float64 `mapstructure:"Delay"`
	Baud                       int     `mapstructure:"Baud"`
	MaxContentRows             int     `mapstructure:"MaxContentRows"`
	MaxNoItems                 int     `mapstructure:"MaxNoItems"`
	TitleLength                int     `mapstructure:"TitleLength"`
	Debug                      bool    `mapstructure:"Debug"`
	DefaultErrorDelay          float64 `mapstructure:"DefaultErrorDelay"`
	DefaultRandomPortMin       int     `mapstructure:"DefaultRandomPortMin"`
	DefaultRandomPortMax       int     `mapstructure:"DefaultRandomPortMax"`
	DefaultRandomMACMin        int     `mapstructure:"DefaultRandomMACMin"`
	DefaultRandomMACMax        int     `mapstructure:"DefaultRandomMACMax"`
	DefaultRandomIPMin         int     `mapstructure:"DefaultRandomIPMin"`
	DefaultRandomIPMax         int     `mapstructure:"DefaultRandomIPMax"`
	DefaultBaud                int     `mapstructure:"DefaultBaud"`
	DefaultBeepDuration        int
	DefaultBeepFrequency       float64
	ValidBaudRates             []int
	ValidFileNameCharacters    []string
}

var Configuration = Config{}

// init reads the configuration file and sets up the configuration object
func init() {
	// add the path to the configuration file to the list of paths to search
	viper.AddConfigPath(".")
	// set the name of the configuration file to be loaded
	viper.SetConfigName("support")
	// set the type of the configuration file (in this case, it is an environment file)
	viper.SetConfigType("env")
	// enable automatic environment variable loading
	viper.AutomaticEnv()

	// read in the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		// handle the error
		return
	}

	// unmarshal the configuration file into the configuration object
	err = viper.Unmarshal(&Configuration)
	if err != nil {
		// handle the error
		panic(err)
	}

	// set the default values for the beep duration and frequency
	Configuration.DefaultBeepDuration = beep.DefaultDuration
	Configuration.DefaultBeepFrequency = beep.DefaultFreq
	// set the valid baud rates
	Configuration.ValidBaudRates = []int{0, 300, 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200}
	// set the valid file name characters
	Configuration.ValidFileNameCharacters = []string{" ", "-", "_", ".", "(", ")", "[", "]", "!", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b",
		"c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	// set the default baud rate
	//Configuration.DefaultBaud = 9600
	// set the default beep duration and frequency
	//Configuration.DefaultBeepDuration = beep.DefaultDuration
	//Configuration.DefaultBeepFrequency = beep.DefaultFreq
	// set the default Plex port
	//Configuration.PlexPort = strconv.Itoa(Configuration.plexPortInt)
}

// split splits a string by the given separator.
func split(in string) (r []string) {
	return strings.Split(in, "|")
}

// buildOrder splits a string by the given separator and converts each element to an integer.
func buildOrder(in string) (r []int) {
	s := strings.Split(in, "|")
	r = make([]int, len(s))
	for i := 0; i < len(s); i++ {
		r[i], _ = strconv.Atoi(string(s[i]))
	}
	return
}

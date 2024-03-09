package config

import (
	"fmt"
	"strings"

	"github.com/gen2brain/beeep"
	e "github.com/mt1976/crt/errors"
	"github.com/spf13/viper"
)

type Config struct {
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

	ValidBaudRates          []int
	ValidFileNameCharacters []string

	DashboardURINameIN      string `mapstructure:"DashboardURIName"`
	DashboardURIProtocolIN  string `mapstructure:"DashboardURIProtocol"`
	DashboardURIHostIN      string `mapstructure:"DashboardURIHost"`
	DashboardURIPortIN      string `mapstructure:"DashboardURIPort"`
	DashboardURIQueryIN     string `mapstructure:"DashboardURIQuery"`
	DashboardURIOperationIN string `mapstructure:"DashboardURIOperation"`
	DashboardURISuccessIN   string `mapstructure:"DashboardURISuccess"`

	DashboardURIName         []string
	DashboardURIProtocol     []string
	DashboardURIHost         []string
	DashboardURIPort         []string
	DashboardURIQuery        []string
	DashboardURIOperation    []string
	DashboardURISuccess      []string
	DashboardURIValidActions []string
	DashboardURINoEntries    int
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
	Configuration.ValidFileNameCharacters = []string{" ", "-", "_", ".", "(", ")", "[", "]", "!", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b",
		"c", "d", "e", "f", "g", "h", "i", "j", "k", "l",
		"m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	//Configuration.PlexPort = strconv.Itoa(Configuration.plexPortInt)
	// spew.Dump(&Configuration)
	// os.Exit(1)
	Configuration.DashboardURIName = split(Configuration.DashboardURINameIN)
	Configuration.DashboardURIHost = split(Configuration.DashboardURIHostIN)
	Configuration.DashboardURIProtocol = split(Configuration.DashboardURIProtocolIN)
	Configuration.DashboardURIPort = split(Configuration.DashboardURIPortIN)
	Configuration.DashboardURIQuery = split(Configuration.DashboardURIQueryIN)
	Configuration.DashboardURIOperation = split(Configuration.DashboardURIOperationIN)
	Configuration.DashboardURISuccess = split(Configuration.DashboardURISuccessIN)
	NoEntries := len(Configuration.DashboardURIHost)
	Configuration.DashboardURINoEntries = NoEntries

	if NoEntries != len(Configuration.DashboardURIProtocol) {
		panic(fmt.Sprintf(e.ErrConfigurationColumnMismatch, len(Configuration.DashboardURIProtocol), NoEntries, "DashboardURIProtocol"))
	}
	if NoEntries != len(Configuration.DashboardURIPort) {
		panic(fmt.Sprintf(e.ErrConfigurationColumnMismatch, len(Configuration.DashboardURIPort), NoEntries, "DashboardURIPort"))
	}
	if NoEntries != len(Configuration.DashboardURIQuery) {
		panic(fmt.Sprintf(e.ErrConfigurationColumnMismatch, len(Configuration.DashboardURIQuery), NoEntries, "DashboardURIQuery"))
	}
	if NoEntries != len(Configuration.DashboardURIName) {
		panic(fmt.Sprintf(e.ErrConfigurationColumnMismatch, len(Configuration.DashboardURIName), NoEntries, "DashboardURIName"))
	}
	if NoEntries != len(Configuration.DashboardURIOperation) {
		panic(fmt.Sprintf(e.ErrConfigurationColumnMismatch, len(Configuration.DashboardURIOperation), NoEntries, "DashboardURIOperation"))
	}
	if NoEntries != len(Configuration.DashboardURISuccess) {
		panic(fmt.Sprintf(e.ErrConfigurationColumnMismatch, len(Configuration.DashboardURISuccess), NoEntries, "DashboardURISuccess"))
	}
	Configuration.DashboardURIValidActions = []string{"PING", "HTTP"}
	//spew.Dump(Configuration)
	//os.Exit(1)
}

func split(s string) (r []string) {
	fmt.Printf(s)
	return strings.Split(s, "|")
}

package language

// General
const (
	//	ErrLibraryResponse    string = "library fetch error"
	TitleLabel            string = "Title"
	SummaryLabel          string = "Summary"
	ContainerLabel        string = "Container"
	ResolutionLabel       string = "Resolution"
	CodecLabel            string = "Codec"
	AspectRatioLabel      string = "Aspect Ratio"
	FrameRateLabel        string = "Frame Rate"
	DurationLabel         string = "Duration"
	ReleasedLabel         string = "Released"
	DirectorLabel         string = "Director"
	WriterLabel           string = "Writer"
	MediaLabel            string = "Media"
	ContentLabel          string = "Content Rating"
	PlexTitle             string = "PMS"
	ErrPlexInit           string = "unable to connect to server"
	ErrPlexConnectionTest string = "unable to test connection to server"
	ErrPlexConnect        string = "unable to connect to %v"
	ErrLibraryResponse    string = "unable to get libraries from %v"
	ErrInvalidAction      string = "invalid action "
	ErrAddColumns         string = "too many columns 10 or less"
	InfoYouSelected       string = "you selected: "
	YearLabel             string = "Year"
	SeasonsTitle          string = "Seasons - "
	ShowLabel             string = "Show"
	SeasonLabel           string = "Season"
	EpisodeLabel          string = "Episode"
	DelimiterText         string = " - "
	TxtDone               string = "DONE"
	TxtStarting                  = "Starting..."
	TxtStartingTerminal          = "Starting Terminal..."
	TxtSelfTesting               = "Self Testing..."
	TxtCurrentDate               = "Current Date: "
	TxtCurrentTime               = "Current Time: "
	TxtPleaseWait                = "Please Wait..."
	TxtBaudRate                  = "Baud Rate Set to %v kbps"
	TxtConnecting                = "Connecting..."
	TxtDialing                   = "Dialing... %v:%v"
	TxtConnected                 = "Connected."
	SymNewline                   = "\n"
	TxtDialingFailed             = "Connection failed. Retrying..."
	TxtComplete                  = "Complete"
)

// Mainmenu

const (
	TxtMainMenuTitle                string = "Main Menu"
	TxtSkyNewsMenuTitle             string = "SKY News"
	TxtBBCNewsMenuTitle             string = "BBC News"
	TxtWeatherMenuTitle             string = "Weather"
	TxtRemoteSystemsAccessMenuTitle string = "Remote Systems Access"
	TxtSystemsMaintenanceMenuTitle  string = "Systems Maintenance"
	TxtPlexMediaServersMenuTitle    string = "Plex Media Server"
	TxtQuittingMessage              string = "Quitting"
	TxtSubMenuTitle                 string = "Sub Menu"
	SymBlank                        string = "-"
	TxtTorrentsMenuTitle            string = "Torrents"
)

// SkyNews
const (
	TxtMenuTitle          string = "SKY News"
	TxtTopicHome          string = "Home"
	TxtTopicUK            string = "UK"
	TxtTopicWorld         string = "World"
	TxtTopicUS            string = "US"
	TxtTopicBusiness      string = "Business"
	TxtTopicPolitics      string = "Politics"
	TxtTopicTechnology    string = "Technology"
	TxtTopicEntertainment string = "Entertainment"
	TxtTopicStrange       string = "Strange News"
	TxtLoadingTopic       string = "Loading news for topic: "
	TxtLoadingStory       string = "Loading news for story..."
	HTMLTagTitle          string = "title"
	HTMLTagTagP           string = "p"
)

// Torrents
const (
	TxtTransmission                = "Transmission"
	TxtQTorrent                    = "qTorrent"
	TxtLoadingTorrentsTransmission = "Loading Transmission Torrents..."
	TxtLoadingTorrentsQTor         = "Loading qTorrent Torrents..."
)

// Weather
const (
	TxtWeatherTitle string = "Weather"

	TxtWeatherPrompt string = "Select (Q)uit"

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
	SymBreak          string = " ━━ "
)

package language

// General
const (
	//	ErrLibraryResponse    string = "library fetch error"
	TitleLabel          string = "Title"
	SummaryLabel        string = "Summary"
	ContainerLabel      string = "Container"
	ResolutionLabel     string = "Resolution"
	CodecLabel          string = "Codec"
	AspectRatioLabel    string = "Aspect Ratio"
	FrameRateLabel      string = "Frame Rate"
	DurationLabel       string = "Duration"
	ReleasedLabel       string = "Released"
	DirectorLabel       string = "Director"
	WriterLabel         string = "Writer"
	MediaLabel          string = "Media"
	ContentLabel        string = "Content Rating"
	PlexTitle           string = "PMS"
	InfoYouSelected     string = "you selected: "
	YearLabel           string = "Year"
	SeasonsTitle        string = "Seasons - "
	ShowLabel           string = "Show"
	SeasonLabel         string = "Season"
	EpisodeLabel        string = "Episode"
	DelimiterText       string = " - "
	TxtDone             string = "DONE"
	TxtStarting         string = "Starting..."
	TxtStartingTerminal string = "Starting Terminal..."
	TxtSelfTesting      string = "Self Testing..."
	TxtCurrentDate      string = "Current Date: "
	TxtCurrentTime      string = "Current Time: "
	TxtPleaseWait       string = "Please Wait..."
	TxtBaudRate         string = "Baud Rate Set to %v kbps"
	TxtConnecting       string = "Connecting..."
	TxtDialing          string = "Dialing... %v:%v"
	TxtConnected        string = "Connected."
	Newline             string = "\n"
	TxtDialingFailed    string = "Connection failed. Retrying..."
	TxtComplete         string = "Complete"
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
	TxtTransmission                string = "Transmission"
	TxtQTorrent                    string = "qTorrent"
	TxtLoadingTorrentsTransmission string = "Loading Transmission Torrents..."
	TxtLoadingTorrentsQTor         string = "Loading qTorrent Torrents..."
)

// Weather
const (
	TxtWeatherTitle   string = "Weather"
	TxtWeatherPrompt  string = "Select (Q)uit"
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
	TxtOneHour        string = " (1hr)"
	TxtThreeHour      string = " (3hr)"
	Space             string = " "
)

// Page - Paging

const TxtPrompt string = "Enter (F)orward, (B)ack or (Q)uit"

const (
	SymActionQuit    string = "Q"
	SymActionForward string = "F"
	SymActionBack    string = "B"
	SymActionExit    string = "EX"
	SymTruncate      string = "..."
	SymWildcardBlank string = "{{blank}}"
)

// Plex Seasons

const (
	SymActionSeasons     string = "S"
	TxtPlexSeasonsPrompt string = "Enter (S)easons, (F)orward, (B)ack or (Q)uit"
)

// Support

const (
	TxtOneWord      string = "one"
	TxtOneNumeric   string = "1"
	TxtMinutes      string = "minutes"
	TxtMinutesShort string = "mins"
	TxtHour         string = "hour"
	TxtHourShort    string = "hr"
)

const (
	TxtMillisecondsShort  string = "ms"
	TxtApplicationVersion string = "StarTerm - Utilities 1.0 %s"
	TxtApplicationName    string = "StarTerm"
	SymPromptSymbol       string = "? "
	TxtError              string = "ERROR : "
	TxtInfo               string = "INFO : "
	TxtPaging             string = "Page %v of %v"
)

var ApplicationHeader []string = []string{
	"███████ ████████  █████  ██████  ████████ ███████ ██████  ███    ███ ",
	"██         ██    ██   ██ ██   ██    ██    ██      ██   ██ ████  ████ ",
	"███████    ██    ███████ ██████     ██    █████   ██████  ██ ████ ██ ",
	"     ██    ██    ██   ██ ██   ██    ██    ██      ██   ██ ██  ██  ██ ",
	"███████    ██    ██   ██ ██   ██    ██    ███████ ██   ██ ██      ██ ",
}

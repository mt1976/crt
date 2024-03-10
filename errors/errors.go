package support

const (
	ErrBaudRateError               string = "invalid baud rate"
	ErrTerminalSize                string = "invalid terminal size"
	ErrDurationMismatch            string = "duration mismatch"
	ErrMaxPageRows                 string = "max page rows reached"
	ErrInvalidAction               string = "invalid action specified "
	ErrNoMorePages                 string = "no more pages"
	ErrPlexInit                    string = "unable to connect to server"
	ErrPlexConnectionTest          string = "unable to test connection to server"
	ErrPlexConnect                 string = "unable to connect to %v"
	ErrLibraryResponse             string = "unable to get libraries from %v"
	ErrAddColumns                  string = "too many columns 10 or less"
	ErrConfigurationColumnMismatch string = "column mismatch in configuration got %v wanted %v in %s"
	ErrDashboardNoHost             string = "dashboard: No default host set"
)

// Weather

const (
	ErrOpenWeather string = "failed to initialize OpenWeatherMap: %v"
)

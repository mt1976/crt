package support

const (
	ErrBaudRateError      string = "Invalid Baud Rate"
	ErrTerminalSize       string = "Invalid Terminal Size"
	ErrDurationMismatch   string = "Duration Mismatch"
	ErrMaxPageRows        string = "Max page rows reached"
	ErrInvalidAction      string = "Invalid action specified "
	ErrNoMorePages        string = "No more pages"
	ErrPlexInit           string = "unable to connect to server"
	ErrPlexConnectionTest string = "unable to test connection to server"
	ErrPlexConnect        string = "unable to connect to %v"
	ErrLibraryResponse    string = "unable to get libraries from %v"
	ErrAddColumns         string = "too many columns 10 or less"
)

// Weather

const (
	ErrOpenWeather string = "failed to initialize OpenWeatherMap: %v"
)

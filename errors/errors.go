package support

import "errors"

var (
	ErrTerminalSize                = errors.New("invalid terminal size %v %v")
	ErrBaudRateError               = errors.New("invalid baud rate %v")
	ErrDurationMismatch            = errors.New("duration mismatch")
	ErrMaxPageRows                 = errors.New("max page rows reached")
	ErrInvalidAction               = errors.New("invalid action specified %v")
	ErrNoMorePages                 = errors.New("no more pages")
	ErrAddColumns                  = errors.New("too many columns 10 or less")
	ErrConfigurationColumnMismatch = errors.New("column mismatch in configuration got %v wanted %v in %s")
	ErrDashboardNoHost             = errors.New("dashboard: No default host set")
)

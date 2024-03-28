package errors

import "errors"

var (
	ErrTerminalSize     = errors.New("invalid terminal size %v %v")
	ErrBaudRateError    = errors.New("invalid baud rate %v")
	ErrDurationMismatch = errors.New("duration mismatch")
	ErrMaxPageRows      = errors.New("max page rows reached")
	ErrInvalidAction    = errors.New("invalid action specified. [%v]")
	ErrInvalidActionLen = errors.New("invalid action length. [%v] %v Characters should be %v")

	ErrNoMorePages                 = errors.New("no more pages")
	ErrAddColumns                  = errors.New("too many columns 10 or less")
	ErrConfigurationColumnMismatch = errors.New("column mismatch in configuration got %v wanted %v in %s")
	ErrDashboardNoHost             = errors.New("dashboard: No default host set")
	ErrHostName                    = errors.New("unable to get hostname")
	ErrUserName                    = errors.New("unable to get username")
	ErrSystemInfo                  = errors.New("unable to get machine name")
	ErrInputLengthMinimum          = errors.New("text must be at least %v characters")
	ErrInputLengthMaximum          = errors.New("text must be at most %v characters, is %v")
	ErrNoPromptSpecified           = errors.New("no prompt specified")
)

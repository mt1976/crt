package errors

import "errors"

var (
	ErrTerminalSize                = errors.New("invalid terminal size %v %v")
	ErrBaudRateError               = errors.New("invalid baud rate %v")
	ErrDurationMismatch            = errors.New("duration mismatch")
	ErrMaxPageRows                 = errors.New("max page rows reached")
	ErrNoActionSpecified           = errors.New("no action specified")
	ErrInvalidAction               = errors.New("invalid action specified. [%v]")
	ErrInvalidActionLen            = errors.New("invalid action length. [%v] %v Characters should be %v")
	ErrInputFailure                = errors.New("unable to get input data") // ErrInputFailure is returned when the input fails
	ErrInputScannerFailure         = errors.New("unable to get input data") // ErrInputScannerFailure is returned when the input scanner fails
	ErrNoMorePages                 = errors.New("no more pages")
	ErrAddColumns                  = errors.New("too many columns have %v should be %v or less")
	ErrConfigurationColumnMismatch = errors.New("column mismatch in configuration got %v wanted %v in %s")
	ErrDashboardNoHost             = errors.New("dashboard: No default host set")
	ErrHostName                    = errors.New("unable to get hostname")
	ErrUserName                    = errors.New("unable to get username")
	ErrSystemInfo                  = errors.New("unable to get machine name")
	ErrInputLengthMinimum          = errors.New("text must be at least %v characters")
	ErrInputLengthMaximum          = errors.New("text must be at most %v characters, is %v")
	ErrNoPromptSpecified           = errors.New("no prompt specified")
	ErrNoEmptyDirectories          = errors.New("unable to find empty directories %v")
	ErrUnableToRemoveDirectories   = errors.New("unable to remove empty directories %v")
	ErrUnableToFindFiles           = errors.New("unable to find files %v")
	ErrUnableToResolvePath         = errors.New("unable to resolve path %v")
	ErrNoPathSpecified             = errors.New("no path specification specified %v")
	ErrInvalidPath                 = errors.New("the path provided is not valid %v")
	ErrInvalidPathSpecialDirectory = errors.New("the path provided is the root or home directory")
	ErrFailedToChangeDirectory     = errors.New("failed to change directory to %v %v")
	ErrNotADirectory               = errors.New("%v is not a directory")
	ErrNotAFile                    = errors.New("%v is not a file")
)

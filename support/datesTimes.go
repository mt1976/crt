package support

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

func HumanFromUnixDate(unixTime int64) string {
	// golang date from unixTime
	t := time.Unix(unixTime, 0)
	h := humanize.Time(t)
	return h
}

// The function DateString returns the current date in the format "dd/mm/yy".
func DateString() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format(defaultDateFormat))
}

// The TimeString function returns the current time in the format "15:04:05".
func TimeString() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format(defaultTimeFormat))
}

// The DateTimeString function returns a string that combines the time and date strings.
func DateTimeString() string {
	return TimeString() + " " + DateString()
}

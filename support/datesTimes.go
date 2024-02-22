package support

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/xeonx/timeago"
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

func PlexDateToDate(date string) time.Time {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func PlexDurationToTime(duration int) time.Duration {
	//int to string
	d := fmt.Sprintf("%d", duration)
	//string to int
	i, err := strconv.Atoi(d)
	if err != nil {
		fmt.Println(err)
	}
	if i != duration {
		fmt.Println(errDurationMismatch, i, " ", duration)
	}
	//return time.Duration(i) * time.Second
	t, err := time.ParseDuration(d + milliseconds)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func FormatDate(t time.Time) string {
	return t.Format("02 Jan 2006")
}

func FormatDuration(t time.Duration) string {
	return t.String()
}

func FormatPlexDate(t string) string {
	return humanize.Time(PlexDateToDate(t)) + " (" + FormatDate(PlexDateToDate(t)) + ")"
}

func FormatPlexDuration(t int) string {
	return FormatDuration(PlexDurationToTime(t))
}

func TimeAgo(t string) string {
	// Example time Thu, 25 Jan 2024 09:56:00 +0000
	// Setup a time format and parse the time

	if t != "" {
		mdt, _ := time.Parse(time.RFC1123Z, t)
		return timeago.English.Format(mdt)
	}
	return ""
}

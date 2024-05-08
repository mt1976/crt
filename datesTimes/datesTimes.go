package datestimes

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	conf "github.com/mt1976/crt/config"
	lang "github.com/mt1976/crt/language"
	"github.com/xeonx/timeago"
)

var config = conf.Configuration

func UnixDateToHuman(unixTime int64) string {
	// golang date from unixTime
	t := time.Unix(unixTime, 0)
	h := humanize.Time(t)
	return h
}

// The function dateString returns the current date in the format "dd/mm/yy".
func DateString() string {
	// spew.Dump(c.ApplicationDateFormatShort)
	// spew.Dump(c)
	// os.Exit(1)
	now := time.Now()
	return fmt.Sprintf("%v", now.Format(config.ApplicationDateFormatShort))
}

// The timeString function returns the current time in the format "15:04:05".
func TimeString() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format(config.ApplicationTimeFormat))
}

// The dateTimeString function returns a string that combines the time and date strings.
func DateTimeString() string {
	return TimeString() + lang.Space.Symbol() + DateString()
}

// formatDate returns a formatted date string based on the given time.Time value.
func FormatDate(t time.Time) string {
	return t.Format(config.ApplicationDateFormat)
}

func FormatDuration(t time.Duration) string {
	return t.String()
}

func TimeAgo(t string) string {
	// Example time Thu, 25 Jan 2024 09:56:00 +0000
	// Setup a time format and parse the time
	if t == "" {
		return ""
	}

	if t != "" {
		mdt, _ := time.Parse(time.RFC1123Z, t)
		rtn := timeago.English.Format(mdt)
		rtn = strings.Replace(rtn, lang.OneWord.Text(), lang.OneNumeric.Text(), -1)
		rtn = strings.Replace(rtn, lang.Minutes.Text(), lang.MinutesShort.Text(), -1)
		rtn = strings.Replace(rtn, lang.Hour.Text(), lang.HourShort.Text(), -1)
		//fix len to 10 chars
		if len(rtn) > 10 {
			rtn = rtn[:10]
		}
		if len(rtn) < 10 {
			rtn = strings.Repeat(lang.Space.Symbol(), 10-len(rtn)) + rtn
		}
		return rtn
	}
	return ""
}

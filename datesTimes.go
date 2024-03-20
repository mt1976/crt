package crt

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	conf "github.com/mt1976/crt/config"
	lang "github.com/mt1976/crt/language"
	"github.com/xeonx/timeago"
)

var c = conf.Configuration

func HumanFromUnixDate(unixTime int64) string {
	// golang date from unixTime
	t := time.Unix(unixTime, 0)
	h := humanize.Time(t)
	return h
}

// The function DateString returns the current date in the format "dd/mm/yy".
func DateString() string {
	// spew.Dump(c.ApplicationDateFormatShort)
	// spew.Dump(c)
	// os.Exit(1)
	now := time.Now()
	return fmt.Sprintf("%v", now.Format(c.ApplicationDateFormatShort))
}

// The TimeString function returns the current time in the format "15:04:05".
func TimeString() string {
	now := time.Now()
	return fmt.Sprintf("%v", now.Format(c.ApplicationTimeFormat))
}

// The DateTimeString function returns a string that combines the time and date strings.
func DateTimeString() string {
	return TimeString() + lang.Space + DateString()
}

func FormatDate(t time.Time) string {
	return t.Format(c.ApplicationDateFormat)
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
		rtn = strings.Replace(rtn, lang.TxtOneWord, lang.TxtOneNumeric, -1)
		rtn = strings.Replace(rtn, lang.TxtMinutes, lang.TxtMinutesShort, -1)
		rtn = strings.Replace(rtn, lang.TxtHour, lang.TxtHourShort, -1)
		//fix len to 10 chars
		if len(rtn) > 10 {
			rtn = rtn[:10]
		}
		if len(rtn) < 10 {
			rtn = strings.Repeat(lang.Space, 10-len(rtn)) + rtn
		}
		return rtn
	}
	return ""
}

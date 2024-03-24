package crt

import (
	"time"

	styl "github.com/mt1976/crt/styles"
)

type Helpers struct {
	RandomIP           func() string
	RandomMAC          func() string
	RandomPort         func() int
	RandomFloat        func(min int, max int) float64
	ToInt              func(s string) int
	ToString           func(i int) string
	CoinToss           func() bool
	IsInt              func(i string) bool
	GetHostName        func() string
	GetUsername        func() string
	GetSytemInfo       func() string
	RoundFloatToTwoDPS func(f float64) float64
	IsActionIn         func(action string, actionToCheck ...string) bool
}

type Formatters struct {
	Bold                    func(s string) string
	SQuote                  func(s string) string
	PQuote                  func(s string) string
	DQuote                  func(s string) string
	QQuote                  func(s string) string
	Upcase                  func(s string) string
	Downcase                func(s string) string
	HumanFromUnixDate       func(unixTime int64) string
	HumanDiskSize           func(size uint64) string
	Human                   func(a any) string
	DateString              func() string
	TimeString              func() string
	TrimRepeatingCharacters func(s string, c string) string
	DateTimeString          func() string
	TimeAgo                 func(t string) string
	FormatDate              func(t time.Time) string
	FormatDuration          func(t time.Duration) string
}

type Styles struct {
	Reset     string
	Red       string
	Green     string
	Yellow    string
	Blue      string
	Purple    string
	Cyan      string
	Gray      string
	White     string
	Bold      string
	Underline string
	ClearLine string
}

func initHelpers() *Helpers {
	help := Helpers{
		RandomIP:           randomIP,
		RandomMAC:          randomMAC,
		RandomPort:         randomPort,
		ToInt:              toInt,
		ToString:           toString,
		CoinToss:           coinToss,
		IsInt:              isInt,
		RandomFloat:        randomFloat,
		GetHostName:        getHostName,
		GetUsername:        getUsername,
		GetSytemInfo:       getSystemInfo,
		RoundFloatToTwoDPS: roundFloatToTwoDPS,
		IsActionIn:         isActionIn,
	}
	return &help
}

func initFormatters() *Formatters {
	fmts := Formatters{
		HumanFromUnixDate:       unixDateToHuman,
		HumanDiskSize:           humanDiskSize,
		Human:                   human,
		DateString:              dateString,
		TimeString:              timeString,
		Upcase:                  upcase,
		Downcase:                downcase,
		Bold:                    bold,
		SQuote:                  sQuote,
		PQuote:                  pQuote,
		DQuote:                  dQuote,
		QQuote:                  qQuote,
		TrimRepeatingCharacters: trimRepeatingCharacters,
		DateTimeString:          dateTimeString,
		TimeAgo:                 timeAgo,
		FormatDate:              formatDate,
		FormatDuration:          formatDuration,
	}
	return &fmts
}

func initStyles() *Styles {
	s := Styles{
		Reset:     styl.Reset,
		Red:       styl.Red,
		Green:     styl.Green,
		Yellow:    styl.Yellow,
		Blue:      styl.Blue,
		Purple:    styl.Purple,
		Cyan:      styl.Cyan,
		Gray:      styl.Gray,
		White:     styl.White,
		Bold:      styl.Bold,
		Underline: styl.Underline,
		ClearLine: styl.ClearLine,
	}
	return &s
}

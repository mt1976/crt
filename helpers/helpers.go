package helpers

import (
	"time"

	gtrm "github.com/buger/goterm"
	colr "github.com/fatih/color"
	dttm "github.com/mt1976/crt/datesTimes"
	mock "github.com/mt1976/crt/mock"
	numb "github.com/mt1976/crt/numbers"
	actn "github.com/mt1976/crt/page/actions"
	strg "github.com/mt1976/crt/strings"
	styl "github.com/mt1976/crt/styles"
	supt "github.com/mt1976/crt/support"
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
	GetUserName        func() (string, error)
	GetSytemInfo       func() string
	RoundFloatToTwoDPS func(f float64) float64
	IsActionIn         func(action string, actionToCheck ...*actn.Action) bool
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
	RED       string
	GREEN     string
	YELLOW    string
	BLUE      string
	MAGENTA   string
	CYAN      string
	GREY      string
	GRAY      string
	WHITE     string
	BOLD      string
	UNDERLINE string
	ClearLine string
	Red       func(s string) string
	Green     func(s string) string
	Yellow    func(s string) string
	Blue      func(s string) string
	Magenta   func(s string) string
	Cyan      func(s string) string
	Grey      func(s string) string
	Gray      func(s string) string
	White     func(s string) string
	Bold      func(s string) string
	Underline func(s string) string
}

func InitHelpers() *Helpers {
	help := Helpers{
		RandomIP:           mock.RandomIP,
		RandomMAC:          mock.RandomMAC,
		RandomPort:         mock.RandomPort,
		ToInt:              numb.ToInt,
		ToString:           numb.ToString,
		CoinToss:           numb.CoinToss,
		IsInt:              numb.IsInt,
		RandomFloat:        mock.RandomFloat,
		GetHostName:        supt.GetHostName,
		GetUserName:        supt.GetUserName,
		GetSytemInfo:       supt.GetSystemInfo,
		RoundFloatToTwoDPS: numb.RoundFloatToTwoDPS,
		IsActionIn:         actn.IsActionIn,
	}
	return &help
}

func InitFormatters() *Formatters {
	fmts := Formatters{
		HumanFromUnixDate:       dttm.UnixDateToHuman,
		HumanDiskSize:           strg.HumanDiskSize,
		Human:                   strg.Human,
		DateString:              dttm.DateString,
		TimeString:              dttm.TimeString,
		Upcase:                  strg.Upcase,
		Downcase:                strg.Downcase,
		Bold:                    bold,
		SQuote:                  strg.SQuote,
		PQuote:                  strg.PQuote,
		DQuote:                  strg.DQuote,
		QQuote:                  strg.QQuote,
		TrimRepeatingCharacters: strg.TrimRepeatingCharacters,
		DateTimeString:          dttm.DateTimeString,
		TimeAgo:                 dttm.TimeAgo,
		FormatDate:              dttm.FormatDate,
		FormatDuration:          dttm.FormatDuration,
	}
	return &fmts
}

func InitStyles() *Styles {
	s := Styles{
		Reset:     styl.Reset,
		RED:       styl.Red,
		GREEN:     styl.Green,
		YELLOW:    styl.Yellow,
		BLUE:      styl.Blue,
		MAGENTA:   styl.Purple,
		CYAN:      styl.Cyan,
		GREY:      styl.White,
		WHITE:     styl.White,
		BOLD:      styl.Bold,
		UNDERLINE: styl.Underline,
		ClearLine: styl.ClearLine,
		Red:       red,
		Green:     green,
		Yellow:    yellow,
		Blue:      blue,
		Magenta:   magenta,
		Cyan:      cyan,
		Grey:      grey,
		Gray:      gray,
		White:     white,
		Bold:      bold,
		Underline: underline,
	}
	//fmt.Printf("ansi.Green: %v\n", ansi.Green)
	return &s
}

func red(s string) string {
	return gtrm.Color(s, gtrm.RED)
}

func green(s string) string {
	return gtrm.Color(s, gtrm.GREEN)
}

func yellow(s string) string {
	return gtrm.Color(s, gtrm.YELLOW)
}

func blue(s string) string {
	return gtrm.Color(s, gtrm.BLUE)
}

func magenta(s string) string {
	return gtrm.Color(s, gtrm.MAGENTA)
}

func cyan(s string) string {
	return gtrm.Color(s, gtrm.CYAN)
}

func grey(s string) string {
	return gray(s)
}

func gray(s string) string {
	gr := colr.New(colr.FgWhite, colr.Faint)
	return gr.Sprint(s)
}

func white(s string) string {
	return gtrm.Color(s, gtrm.WHITE)
}

func bold(s string) string {
	return gtrm.Bold(s)
}

func underline(s string) string {
	und := colr.New(colr.Underline)
	return und.Sprint(s)
	//return colr.UnderlineString(s)
}

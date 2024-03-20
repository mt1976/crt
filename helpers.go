package crt

type Helpers struct {
	RandomIP   func() string
	RandomMAC  func() string
	RandomPort func() int

	ToInt    func(s string) int
	ToString func(i int) string
	CoinToss func() bool
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
	DateString              func() string
	TrimRepeatingCharacters func(s string, c string) string
}

func initHelpers() *Helpers {
	help := Helpers{
		RandomIP:   randomIP,
		RandomMAC:  randomMAC,
		RandomPort: randomPort,
		ToInt:      toInt,
		ToString:   toString,
		CoinToss:   coinToss,
	}
	return &help
}

func initFormatters() *Formatters {
	help := Formatters{
		HumanFromUnixDate:       humanFromUnixDate,
		DateString:              dateString,
		Upcase:                  upcase,
		Downcase:                downcase,
		Bold:                    bold,
		SQuote:                  sQuote,
		PQuote:                  pQuote,
		DQuote:                  dQuote,
		QQuote:                  qQuote,
		TrimRepeatingCharacters: trimRepeatingCharacters,
	}
	return &help
}

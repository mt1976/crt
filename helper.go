package crt

type Helpers struct {
	RandomIP                func() string
	RandomMAC               func() string
	RandomPort              func() int
	HumanFromUnixDate       func(unixTime int64) string
	DateString              func() string
	Upcase                  func(s string) string
	ToInt                   func(s string) int
	ToString                func(i int) string
	Bold                    func(s string) string
	SQuote                  func(s string) string
	CoinToss                func() bool
	PQuote                  func(s string) string
	TrimRepeatingCharacters func(s string, c string) string
}

func initHelpers() *Helpers {
	help := Helpers{
		RandomIP:                RandomIP,
		RandomMAC:               RandomMAC,
		RandomPort:              RandomPort,
		HumanFromUnixDate:       HumanFromUnixDate,
		DateString:              DateString,
		Upcase:                  Upcase,
		ToInt:                   ToInt,
		ToString:                ToString,
		Bold:                    Bold,
		SQuote:                  SQuote,
		CoinToss:                CoinToss,
		PQuote:                  PQuote,
		TrimRepeatingCharacters: TrimRepeatingCharacters,
	}
	return &help
}

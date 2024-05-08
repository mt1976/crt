package strings

import (
	"fmt"
	"strings"

	colour "github.com/fatih/color"
	lang "github.com/mt1976/crt/language"
	numb "github.com/mt1976/crt/numbers"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

// The upcase function in Go converts a string to uppercase.
// upcase converts a string to uppercase.
func Upcase(s string) string {
	return strings.ToUpper(s)
}

func Downcase(s string) string {
	return strings.ToLower(s)
}

// The function `trimRepeatingCharacters` takes a string `s` and a character `c` as input, and returns
// a new string with all consecutive occurrences of `c` trimmed down to a single occurrence.
func TrimRepeatingCharacters(s string, c string) string {

	result := ""
	lenS := len(s)

	for i := 0; i < lenS; i++ {
		if i == 0 {
			result = string(s[i])
		} else {
			if string(s[i]) != c || string(s[i-1]) != c {
				result = result + string(s[i])
			}
		}
	}
	return result
}

// italic returns a string with the italic style applied to it.
func Italic(s string) string {
	return colour.New(colour.Italic).Sprintf("%v", s)
}

// boldInt returns a bolded string representation of an integer.
func BoldInt(s int) string {
	return fmt.Sprintf("%d", s)
}

// sQuote returns a string with the single quote symbol around it.
func SQuote(s string) string {
	return lang.SingleQuote.Symbol() + s + lang.SingleQuote.Symbol()
}

// pQuote returns a string with the square bracket symbol around it.
func PQuote(s string) string {
	return lang.SquareQuoteOpen.Symbol() + s + lang.SquareQuoteClose.Symbol()
}

// dQuote returns a string with the double quote symbol around it.
func DQuote(s string) string {
	return lang.DoubleQuote.Symbol() + s + lang.DoubleQuote.Symbol()
}

// qQuote returns a string with the square quote symbol around it.
func QQuote(s string) string {
	return lang.SquareQuoteOpen.Symbol() + s + lang.SquareQuoteClose.Symbol()
}

// cleanContent removes unwanted characters from the rowContent string
func CleanContent(msg string) string {
	// replace \n, \r, \t, and " with empty strings
	msg = strings.Replace(msg, lang.Newline.Symbol(), "", -1)
	msg = strings.Replace(msg, lang.CarridgeReturn.Symbol(), "", -1)
	msg = strings.Replace(msg, lang.Tab.Symbol(), "", -1)
	msg = strings.Replace(msg, lang.DoubleQuote.Symbol(), lang.Space.Symbol(), -1)
	return msg
}

// The `humanNumber` method of the `Crt` struct is used to convert a value to a human-readable string. It
// takes a parameter `v` of type `any`, which means it can accept any type of value.
func Human(v any) string {
	if v == nil {
		return ""
	}

	//T.Basic(fmt.Sprintf("Type: %T", v))

	p := message.NewPrinter(language.English)

	switch v.(type) {
	case int, int8, int16, uint, uint8, uint16, int32, int64, uint64:
		return p.Sprintf("%d", number.Decimal(v))
	case float32, float64:
		return p.Sprintf("%.2f", number.Decimal(v))
	case string:
		return fmt.Sprintf("%s", v)
	default:
		//T.Special(fmt.Sprintf("Type: %T", v))
	}

	return fmt.Sprintf("%v", v)
}

// The function `humanDiskSize` converts a given input value representing disk size in bytes to a
// human-readable format in gigabytes (GB) and terabytes (TB).
func HumanDiskSize(input uint64) string {

	// convert input to float64
	val := float64(input)

	// input is in bytes
	// convert bytes to GB
	val = val / 1024 / 1024 / 1024
	//fmt.Println(val)
	tbs := val / 1024
	//fmt.Println(tbs)
	//fmt.Println(val, tbs)
	value := fmt.Sprintf("%.2fGB (%.2fTB)", numb.RoundFloatToTwoDPS(val), numb.RoundFloatToTwoDPS(tbs))
	return value
}

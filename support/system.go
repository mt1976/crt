package support

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func GetSystemInfo(crt Crt) string {
	unameCmd := exec.Command("uname", "-n")
	output, err := unameCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to get machine name: %v", CHspecial, err))
		crt.Error("Unable to get machine name", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}

func GetHostname(crt Crt) string {
	hostnameCmd := exec.Command("hostname")
	output, err := hostnameCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to get hostname: %v", CHspecial, err))
		crt.Error("Unable to get hostname", err)
		return ""
	}
	// remove ".local" from hostname
	output = []byte(strings.ReplaceAll(string(output), ".local", ""))
	return strings.TrimSpace(string(output))
}

func GetUsername(crt Crt) string {
	whoamiCmd := exec.Command("whoami")
	output, err := whoamiCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to get username: %v", CHspecial, err))
		crt.Error("Unable to get username", err)
		return ""
	}
	return strings.TrimSpace(string(output))
}

// The function GetHostName retrieves the hostname of the current machine.
func GetHostName() (string, error) {
	hostName, err := os.Hostname()
	if err != nil {
		fmt.Printf("%s Error getting hostname: %v\n", chNormal, err)
		return "", err
	}
	return hostName, nil
}

// The `Human` method of the `Crt` struct is used to convert a value to a human-readable string. It
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

// The function `DiskSizeHuman` converts a given input value representing disk size in bytes to a
// human-readable format in gigabytes (GB) and terabytes (TB).
func DiskSizeHuman(input uint64) string {

	// convert input to float64
	val := float64(input)

	// input is in bytes
	// convert bytes to GB
	val = val / 1024 / 1024 / 1024
	//fmt.Println(val)
	tbs := val / 1024
	//fmt.Println(tbs)
	//fmt.Println(val, tbs)
	value := fmt.Sprintf("%.2fGB (%.2fTB)", roundFloatToTwo(val), roundFloatToTwo(tbs))
	return value
}

// The function roundFloatToTwo rounds a float64 number to two decimal places.
func roundFloatToTwo(input float64) float64 {
	// round float64 to 2 decimal places
	rtnVal := math.Round(input*100) / 100

	return rtnVal
}

// The function `TrimRepeatingCharacters` takes a string `s` and a character `c` as input, and returns
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

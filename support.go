package crt

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"

	errs "github.com/mt1976/crt/errors"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/number"
)

func getSystemInfo() string {
	unameCmd := exec.Command("uname", "-n")
	output, err := unameCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to get machine name: %v", CHspecial, err))
		fmt.Println(errs.ErrSystemInfo.Error(), err.Error())
		return ""
	}
	return strings.TrimSpace(string(output))
}

func getUsername() string {
	whoamiCmd := exec.Command("whoami")
	output, err := whoamiCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to get username: %v", CHspecial, err))
		fmt.Println(errs.ErrUserName.Error(), err.Error())
		return ""
	}
	return strings.TrimSpace(string(output))
}

// The function getHostName retrieves the hostname of the current machine.
func getHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		//	fmt.Printf("ERROR: getting hostname: %v\n", err.Error())
		fmt.Println(errs.ErrHostName.Error(), err.Error())
		return ""
	}
	return hostName
}

// The `humanNumber` method of the `Crt` struct is used to convert a value to a human-readable string. It
// takes a parameter `v` of type `any`, which means it can accept any type of value.
func human(v any) string {
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
func humanDiskSize(input uint64) string {

	// convert input to float64
	val := float64(input)

	// input is in bytes
	// convert bytes to GB
	val = val / 1024 / 1024 / 1024
	//fmt.Println(val)
	tbs := val / 1024
	//fmt.Println(tbs)
	//fmt.Println(val, tbs)
	value := fmt.Sprintf("%.2fGB (%.2fTB)", roundFloatToTwoDPS(val), roundFloatToTwoDPS(tbs))
	return value
}

// The function roundFloatToTwoDPS rounds a float64 number to two decimal places.
func roundFloatToTwoDPS(input float64) float64 {
	// round float64 to 2 decimal places
	rtnVal := math.Round(input*100) / 100

	return rtnVal
}

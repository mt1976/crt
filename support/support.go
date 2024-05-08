package support

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	errs "github.com/mt1976/crt/errors"
)

func GetSystemInfo() string {
	unameCmd := exec.Command("uname", "-n")
	output, err := unameCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to get machine name: %v", CHspecial, err))
		fmt.Println(errs.ErrSystemInfo.Error(), err.Error())
		return ""
	}
	return strings.TrimSpace(string(output))
}

func GetUserName() (string, error) {
	whoamiCmd := exec.Command("whoami")
	output, err := whoamiCmd.Output()
	if err != nil {
		//log.Fatal(fmt.Sprintf("%s Unable to get username: %v", CHspecial, err))
		fmt.Println(errs.ErrUserName.Error(), err.Error())
		return "", errs.ErrUserName
	}
	return strings.TrimSpace(string(output)), nil
}

// The function getHostName retrieves the hostname of the current machine.
func GetHostName() string {
	hostName, err := os.Hostname()
	if err != nil {
		//	fmt.Printf("ERROR: getting hostname: %v\n", err.Error())
		fmt.Println(errs.ErrHostName.Error(), err.Error())
		return ""
	}
	return hostName
}

func GetUserHome() (string, error) {
	// Function gets the home directory of the current user, or returns an error if it cant.
	//
	// Returns:
	// The home directory of the current user, or an error if it cant.
	return os.UserHomeDir()
}

package disksize

import (
	"fmt"
	"os"

	"github.com/mt1976/admin_me/support"
)

const (
	Help = `
Prints file sizes in bytes, kilobytes, megabytes, and gigabytes
Usage: sz <file> <file> <file>
`
)

var crt support.Crt
var debugMode bool

func Run(crtIn support.Crt, debug bool, args []string) {
	crt = crtIn
	if len(args) < 2 {
		crt.Shout("No files specified")
		crt.Print(Help)
		return
	}
	crt.Special(fmt.Sprintf("File Sizes of %v files", len(args)-1))
	crt.Break()
	for _, file := range args[1:] {
		printFileSize(file)
	}
}

func printFileSize(path string) {
	fi, err := os.Stat(path)
	if os.IsNotExist(err) {
		msg := fmt.Sprintf("File %s does not exist", path)
		crt.Error(msg, err)
		return
	}
	if err != nil {
		//fmt.Print("Something went wrong\n")
		crt.Error("Error getting file info", err)
		return
	}
	size := fi.Size()
	sizeKb := float64(size) / 1024
	sizeMb := sizeKb / 1024
	sizeGb := sizeMb / 1024
	sizeTb := sizeGb / 1024
	fileName := crt.Bold(fi.Name())
	duMsg := fmt.Sprintf("%5d b | %5.2f kb | %5.2f mb | %5.2f gb | %5.2f tb | %s",
		size,
		sizeKb,
		sizeMb,
		sizeGb,
		sizeTb,
		fileName,
	)
	crt.Print(duMsg)
}

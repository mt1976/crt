package styles

import (
	"runtime"
)

var Reset string = "\033[0m"
var Red string = "\033[1;31m"
var Green string = "\033[32m"
var Yellow string = "\033[33m"
var Blue string = "\033[34m"
var Purple string = "\033[35m"
var Cyan string = "\033[1;36m"
var Gray string = "\033[1;37m"
var White string = "\033[97m"
var Bold string = "\033[1m"
var Underline string = "\033[4m"
var ClearLine string = "\033[2K"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Green = ""
		Yellow = ""
		Blue = ""
		Purple = ""
		Cyan = ""
		Gray = ""
		White = ""
		Bold = ""
		Underline = ""
		ClearLine = ""
	}
}

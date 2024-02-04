package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	catalog "github.com/mt1976/admin_me/actions/catalog"
	clean "github.com/mt1976/admin_me/actions/clean"
	"github.com/mt1976/admin_me/actions/disksize"
	alive "github.com/mt1976/admin_me/actions/pushover"
	tidy "github.com/mt1976/admin_me/actions/tidy"
	terminal "github.com/mt1976/admin_me/support"
)

const actionClean = "clean"
const actionHeartbeat = "message"
const actionTidy = "tidy"
const actionCatalog = "catalog"
const actionDu = "du"
const actionDebug = true
const cmdLineFlagAction = "action"
const cmdLineFlagMessage string = "message"
const cmdLineFlagTitle string = "title"
const cmdLineFlagBody string = "body"
const cmdLineFlagDebug string = "debug"

var debugMode bool = false

//var COUNT int = 0

var commandAction string = ""
var commandMessage string = ""
var commandTitle string = ""
var commandContent string = ""

func main() {

	// define a new instance of the Crt
	crt := terminal.NewCrt()

	actions := []string{actionClean, actionHeartbeat, actionTidy, actionCatalog, actionDu}
	// convert the slice to a string, separated by commas
	//
	actions_txt := fmt.Sprintf("%v", actions)
	action_message := "Action to perform " + crt.Bold(actions_txt)

	flag.BoolVar(&debugMode, cmdLineFlagDebug, actionDebug, "Debug Mode")
	flag.StringVar(&commandAction, cmdLineFlagAction, actionClean, action_message)
	flag.StringVar(&commandMessage, cmdLineFlagMessage, "", "Type of message "+crt.Bold("[up, heartbeat, other, message]"))
	flag.StringVar(&commandTitle, cmdLineFlagTitle, "", "Title of message")
	flag.StringVar(&commandContent, cmdLineFlagBody, "", "Body of message")
	flag.Parse()

	start := time.Now()

	if commandAction == "" {
		//crt.Blurt("No action specified\n", &T)
		crt.Shout("No action specified")
		return
	}

	//Get Current Path
	path, err := os.Getwd()
	if err != nil {
		crt.Error("Error getting current path", err)
		return
	}
	debugMessage := ""
	if debugMode {
		debugMessage = "in " + crt.Bold(crt.Underline("debug")) + " mode."
	}
	//fmt.Println(pr(BOLD+"StarTerm 1.0 - Utilities"+RESET+msg, T))
	crt.Banner(debugMessage)

	// Current folder as the base
	//get a list of all files in the current directory, and store in a slice
	// Clean the file name
	// Case 1: Remove all characters that are not in the ValidChars list
	// Case 2: Remove send a notification to the user
	crt.Print(crt.Bold("ACTION : ") + commandAction)
	crt.Break()

	switch commandAction {
	case actionClean:
		clean.Run(crt, debugMode, path, debugMessage)
	case actionHeartbeat:
		alive.Run(crt, debugMode, commandMessage, commandTitle, commandContent)
	case actionTidy:
		tidy.Run(crt, debugMode, path)
	case actionCatalog:
		catalog.Run(crt, debugMode, path)
	case actionDu:
		myArgs := os.Args[1:]
		//crt.Print(fmt.Sprintf("Args: %s", myArgs[1:]))
		disksize.Run(crt, debugMode, myArgs)
	default:
		crt.Shout("Unknown action '" + crt.Bold(commandAction) + "'")
		return
	}

	elapsed := time.Since(start)
	crt.Shout(crt.Bold("DONE") + " " + elapsed.String())
	// `elapsed` is a variable of type `time.Duration` that stores
	// the amount of time that has passed since the `start` time. It
	// is used to calculate the total time taken to execute the
	// program.
}

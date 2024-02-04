package pushover

import (
	flags "flag"
	"time"

	"github.com/mt1976/admin_me/support"
)

// Init Objects

// var (
//
//	msgType  string
//	msgTitle string
//	msgBody  string
//
// )
var messages = make(map[string]string)
var titles = make(map[string]string)
var hostName string
var crt support.Crt

//var debugMode bool

type MessageTypes struct {
	Up        string
	Heartbeat string
	Other     string
	Message   string
}

// End Init Objects
func init() {

	var types = MessageTypes{}
	types.Up = "-up"
	types.Heartbeat = "-heartbeat"
	types.Other = "-other"
	types.Message = "-message"

	hostName, _ = support.GetHostName()

	messages[types.Up] = hostName + " has started - System is available"
	messages[types.Heartbeat] = hostName + " - Heartbeat - OK"
	messages[types.Other] = hostName + " - Unknown Message"
	messages[types.Message] = hostName + " - Message"

	titles[types.Up] = "System"
	titles[types.Heartbeat] = "Heartbeat"
	titles[types.Other] = "Other"
	titles[types.Message] = "This Message"
}

func Run(crtIn support.Crt, debug bool, messageType, messageTitle, messageBody string) {

	crt = crtIn
	//debugMode = debug

	flags.Parse()

	crt.Print("Starting Alive")
	crt.Print("Message Type: " + messageType)
	crt.Print("Message Title: " + messageTitle)
	crt.Print("Message Body: " + messageBody)
	crt.Blank()

	//argsWithoutProg := os.Args[1:]
	if messageType != "" {
		//fmt.Println("Args: ", argsWithoutProg)
		//CONFIG.DebugCFG()
		//Get Time
		now := time.Now().Format("2006-01-02 15:04:05")

		//L.WithFields(xlg.Fields{"args": argsWithoutProg, "msgType": msgType}).Info("Arguments")
		switch messageType {
		case "up":
			//xlg.Info("ACTION=UP")
			sendMessage("System Started", hostName+" started at "+now)
		case "heartbeat":
			//xlg.Info("ACTION=HEARTBEAT")
			sendMessage("System Online", hostName+" online at "+now)
		case "other":
			//xlg.Info("ACTION=OTHER")
			sendMessage("Unknown Message", "An unknown message was received at "+now)
		default:
			//xlg.Info("ACTION=RAW")
			sendMessage(messageTitle, messageBody)
		}
	}
}

func sendMessage(inMessage, inTitle string) {
	//L.WithFields(xlg.Fields{"Message": inMessage, "Title": inTitle}).Info("Sending Message")
	crt.Print("Sending Message")
	crt.Print("Message: " + inMessage)
	crt.Print("Title: " + inTitle)
	Normal(inMessage, inTitle)
	crt.Print("Message Sent")
}

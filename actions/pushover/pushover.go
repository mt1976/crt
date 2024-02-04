package pushover

import (
	"strings"
	"time"

	"github.com/gregdel/pushover"
	"github.com/mt1976/admin_me/support"
)

type cConfig struct {
	PushoverKey   string `mapstructure:"pushoverkey"`
	PushoverToken string `mapstructure:"pushovertoken"`
	AppPort       string
	HostName      string
	AppName       string
}

// The init function initializes global variables for the application.
func newPushoverConfiguration() cConfig {

	cTokens := cConfig{}
	cTokens.HostName, _ = support.GetHostName()
	cTokens.AppName = "admin_me"
	cTokens.AppPort = "8080"
	cTokens.PushoverKey = "acxw2ety975n7ux83wkzpp47jzq42q"
	cTokens.PushoverToken = "uyosdopsu9wxxo7b264bmnnhbfz8nj"
	return cTokens
}

// The function "newPushoverMessage" creates a newPushoverMessage pushover.Message object with the given title, body, and priority,
// along with other optional parameters.
func newPushoverMessage(C cConfig, title string, body string, priority int) *pushover.Message {
	return &pushover.Message{
		Message:     body,
		Title:       title,
		Priority:    priority,
		URL:         "http://" + C.HostName + ":" + C.AppPort + "/",
		URLTitle:    C.HostName,
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(C.HostName, ".", "_"),
		CallbackURL: "http://" + C.HostName + ":" + C.AppPort + "/ACK",
		Sound:       pushover.SoundCosmic,
	}
}

// The function builds a push notification message with a title, body, and priority using the Pushover
// library in Go.
func buildPushoverMessage(messageBody string, messageTitle string, priority int) (*pushover.Pushover, *pushover.Recipient, *pushover.Message) {
	C := newPushoverConfiguration()
	//spew.Dump(C)
	app := pushover.New(C.PushoverKey)

	recipient := pushover.NewRecipient(C.PushoverToken)

	messageTitle = messageTitle + " [" + C.HostName + "]"

	message := newPushoverMessage(C, messageTitle, messageBody, priority)
	return app, recipient, message
}

// The function sends a Pushover message using the provided app, message, and recipient.
func sendNotification(app *pushover.Pushover, message *pushover.Message, recipient *pushover.Recipient) {
	//C := setup()
	//xlogs.Info("Sending Pushover Message")
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		panic(err)
	}
}

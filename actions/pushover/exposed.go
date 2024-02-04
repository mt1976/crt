package pushover

import (
	"github.com/gregdel/pushover"
)

// The function "Emergency" sends an emergency message with a specified title and body using the
// Pushover service.
func Emergency(messageTitle string, messageBody string) {
	// Build Message
	app, recipient, message := buildPushoverMessage(messageBody, messageTitle, pushover.PriorityEmergency)
	// Send the message to the recipient
	sendNotification(app, message, recipient)
}

// The Normal function sends a push notification with a normal priority level.
func Normal(messageTitle string, messageBody string) {
	// Build Message
	app, recipient, message := buildPushoverMessage(messageBody, messageTitle, pushover.PriorityNormal)
	// Send the message to the recipient
	sendNotification(app, message, recipient)
}

// The function "WithURL" sends a push notification with a message title, message body, and a URL to a
// recipient using the Pushover service.
func WithURL(messageTitle string, messageBody string, url string) {
	// Build Message
	app, recipient, message := buildPushoverMessage(messageBody, messageTitle, pushover.PriorityNormal)
	// Add URI/URL
	message.URL = message.URL + url
	// Send the message to the recipient
	sendNotification(app, message, recipient)
}

// The High function sends a high priority message with a title and body to a recipient using the
// Pushover service.
func High(messageTitle string, messageBody string) {
	// Build Message
	app, recipient, message := buildPushoverMessage(messageBody, messageTitle, pushover.PriorityHigh)
	// Send the message to the recipient
	sendNotification(app, message, recipient)
}

// The Low function sends a low priority push notification with a given title and body.
func Low(messageTitle string, messageBody string) {
	//cfg, _ := notification_GetConfig()
	app, recipient, message := buildPushoverMessage(messageBody, messageTitle, pushover.PriorityLow)
	// Send the message to the recipient
	sendNotification(app, message, recipient)
}

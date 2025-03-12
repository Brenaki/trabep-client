package notification

import (
	"github.com/gen2brain/beeep"
)

// AppName is the name shown in notifications
const AppName = "Time Tracker"

// NotifySessionStart shows a notification when a new time tracking session starts
func NotifySessionStart() error {
	return beeep.Notify(
		AppName,
		"Time tracking started. Run again to complete the session.",
		"",
	)
}

// NotifySessionComplete shows a notification when a time tracking session is successfully completed
func NotifySessionComplete() error {
	return beeep.Notify(
		AppName,
		"Time record successfully sent to API!",
		"",
	)
}

// NotifyError shows an alert notification when an error occurs
func NotifyError(message string) error {
	return beeep.Alert(
		AppName+" Error",
		message,
		"",
	)
}
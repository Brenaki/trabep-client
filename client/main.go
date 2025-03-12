package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Brenaki/trabep-client/client/api"
	"github.com/Brenaki/trabep-client/client/config"
	"github.com/Brenaki/trabep-client/client/notification"
	"github.com/Brenaki/trabep-client/client/session"
)

func main() {
	// Get current time in the required format for the API
	currentTime := time.Now().Format("02/01/2006, 15:04:05")
	fmt.Printf("Current time: %s\n", currentTime)

	// Check if there's an existing session
	startTime, err := session.CheckExistingSession()
	if err != nil {
		fmt.Printf("Error checking session: %v\n", err)
		os.Exit(1)
	}

	if startTime == "" {
		// No existing session, create a new one
		fmt.Println("Starting new time tracking session...")
		
		// Save the current time as the session start time
		if err := session.SaveSession(currentTime); err != nil {
			fmt.Printf("Error saving session: %v\n", err)
			os.Exit(1)
		}
		
		// Show notification for session start
		if err := notification.NotifySessionStart(); err != nil {
			fmt.Println("Error showing notification:", err)
		}
		
		fmt.Println("Session started. Run the program again to complete time tracking.")
	} else {
		// Existing session found, complete it
		fmt.Printf("Continuing session started at: %s\n", startTime)
		fmt.Println("Completing time tracking session...")
		
		// Send time record to API
		if err := api.SendToAPI(config.GetUsername(), startTime, currentTime); err != nil {
			fmt.Printf("Error sending to API: %v\n", err)
			
			// Show error notification
			notification.NotifyError("Failed to send time record to API.")
		} else {
			fmt.Println("Successfully sent time record to API!")
			
			// Show success notification
			if err := notification.NotifySessionComplete(); err != nil {
				fmt.Println("Error showing notification:", err)
			}
			
			// Delete the session file
			if err := session.DeleteSession(); err != nil {
				fmt.Printf("Error deleting session: %v\n", err)
			} else {
				fmt.Println("Session file deleted successfully.")
			}
		}
	}

	// Keep console open for a moment so user can see the output
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}
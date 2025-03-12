package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Brenaki/trabep-client/api"
	"github.com/Brenaki/trabep-client/models"
	"github.com/Brenaki/trabep-client/queue"
	"github.com/gen2brain/beeep"
)

func main() {
	// Get queue file path
	queueFilePath, err := queue.GetQueueFilePath()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		os.Exit(1)
	}

	// Get current time in the required format
	currentTime := time.Now().Format("02/01/2006, 15:04:05")
	fmt.Printf("Recording time entry: %s\n", currentTime)

	// Load existing queue
	timeQueue, err := queue.LoadQueue(queueFilePath)
	if err != nil {
		// If file doesn't exist, create a new queue
		timeQueue = models.TimeQueue{
			Entries: []models.TimeEntry{},
			User:    queue.DefaultUser,
		}
	}

	// Add current time to queue
	queue.AddEntry(&timeQueue, currentTime)
	fmt.Printf("Added to queue. Queue now has %d entries.\n", len(timeQueue.Entries))

	// Save updated queue
	if err := queue.SaveQueue(queueFilePath, timeQueue); err != nil {
		fmt.Println("Error saving queue:", err)
		os.Exit(1)
	}

	// Show notification for first entry
	if len(timeQueue.Entries) == 1 {
		err = beeep.Notify("Time Tracker", "First time entry recorded. Run again to complete the time tracking.", "")
		if err != nil {
			fmt.Println("Error showing notification:", err)
		}
	}

	// If we have 2 or more entries, process them
	if queue.HasTwoOrMoreEntries(timeQueue) {
		fmt.Println("Queue has 2 entries. Processing...")
		
		// Take the first two entries
		startTime := timeQueue.Entries[0].Timestamp
		endTime := timeQueue.Entries[1].Timestamp
		
		// Send to API
		if err := api.SendToAPI(timeQueue.User, startTime, endTime); err != nil {
			fmt.Printf("Error sending to API: %v\n", err)
			
			// Show error notification
			beeep.Alert("Time Tracker Error", "Failed to send time record to API.", "")
		} else {
			fmt.Println("Successfully sent time record to API!")
			
			// Show success notification
			beeep.Notify("Time Tracker", "Time record successfully sent to API!", "")
			
			// Clear the queue after successful submission
			queue.ClearQueue(&timeQueue)
			if err := queue.SaveQueue(queueFilePath, timeQueue); err != nil {
				fmt.Println("Error clearing queue:", err)
			} else {
				fmt.Println("Queue cleared successfully.")
			}
		}
	} else {
		fmt.Println("Waiting for one more entry before sending to API.")
	}

	// Keep console open for a moment so user can see the output
	fmt.Println("\nPress Enter to exit...")
	fmt.Scanln()
}
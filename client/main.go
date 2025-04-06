package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Brenaki/trabep-client/client/api"
	"github.com/Brenaki/trabep-client/client/config"
	"github.com/Brenaki/trabep-client/client/notification"
	"github.com/Brenaki/trabep-client/client/session"
	"github.com/Brenaki/trabep-client/client/system" // Novo pacote para gerenciar eventos do sistema
)

var (
	logger        *log.Logger
	flagInstall   = flag.Bool("install", false, "Install as autostart service")
	flagUninstall = flag.Bool("uninstall", false, "Uninstall autostart service")
	flagBackground = flag.Bool("background", false, "Run in background mode")
	flagVersion    = flag.Bool("version", false, "Show version information")
)

const (
	appVersion = "1.1.0"
)

// setupLogger configures the application logger
func setupLogger() {
	// Get executable directory for log file
	logFile, err := os.OpenFile("timetracker.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		logFile = os.Stderr // Fallback to stderr
	}
	
	// Create logger that writes to the file
	logger = log.New(logFile, "TimeTracker: ", log.LstdFlags)
	logger.Println("Logger initialized")
}

// startTracking starts a new tracking session
func startTracking() (string, error) {
	currentTime := time.Now().Format("02/01/2006, 15:04:05")
	logger.Printf("Starting tracking session at %s", currentTime)
	
	// Save the current time as the session start time
	if err := session.SaveSession(currentTime); err != nil {
		return "", fmt.Errorf("error saving session: %w", err)
	}
	
	// Show notification for session start
	if err := notification.NotifySessionStart(); err != nil {
		logger.Printf("Error showing notification: %v", err)
	}
	
	return currentTime, nil
}

// endTracking ends the current tracking session and sends data to API
func endTracking(startTime string) error {
	currentTime := time.Now().Format("02/01/2006, 15:04:05")
	logger.Printf("Ending tracking session at %s", currentTime)
	
	// Send time record to API
	if err := api.SendToAPI(config.GetUsername(), startTime, currentTime); err != nil {
		return fmt.Errorf("error sending to API: %w", err)
	}
	
	logger.Println("Successfully sent time record to API")
	
	// Delete the session file
	if err := session.DeleteSession(); err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}
	
	return nil
}

// handleShutdown is called when a shutdown/suspend event is detected
func handleShutdown() {
	logger.Println("System shutdown/suspend detected")
	
	// Check if we have an active session
	startTime, err := session.CheckExistingSession()
	if err != nil {
		logger.Printf("Error checking session: %v", err)
		return
	}
	
	// If we have a session, end it and send data
	if startTime != "" {
		if err := endTracking(startTime); err != nil {
			logger.Printf("Error ending tracking: %v", err)
			notification.NotifyError("Failed to send time record: " + err.Error())
		} else {
			notification.NotifySessionComplete()
			logger.Println("Successfully completed tracking session on shutdown")
		}
	}
}

// runBackground runs the app in background mode
func runBackground() {
	logger.Println("Starting in background mode")
	
	// Check if there's an existing session
	startTime, err := session.CheckExistingSession()
	if err != nil {
		logger.Fatalf("Error checking session: %v", err)
	}
	
	// If no existing session, start a new one
	if startTime == "" {
		startTime, err = startTracking()
		if err != nil {
			logger.Fatalf("Failed to start tracking: %v", err)
		}
		logger.Println("Time tracking started successfully")
		
		// Show notification
		notification.NotifySessionStart()
	} else {
		logger.Printf("Continuing existing session started at: %s", startTime)
		
		// Show notification
		notification.NotifySessionComplete()
	}
	
	// Setup shutdown handler
	err = system.SetupShutdownHandler(logger, handleShutdown)
	if err != nil {
		logger.Printf("Warning: Failed to setup shutdown handler: %v", err)
	}
	
	// Setup signal handling for manual termination
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	
	logger.Println("Waiting for shutdown signals...")
	fmt.Println("Time Tracker running in background. Press Ctrl+C to exit.")
	
	// Wait for shutdown signal
	<-sigs
	
	// Handle graceful shutdown
	handleShutdown()
	logger.Println("Time Tracker shutdown complete")
}

func main() {
	// Parse command-line flags
	flag.Parse()
	
	// Setup logger
	setupLogger()
	
	// Load environment variables from .env file
	if err := config.LoadEnv(); err != nil {
		logger.Printf("Warning: Error loading .env file: %v", err)
	}
	
	// Handle version flag
	if *flagVersion {
		fmt.Printf("Time Tracker version %s\n", appVersion)
		return
	}
	
	// Handle install flag
	if *flagInstall {
		fmt.Println("Installing Time Tracker as autostart service...")
		if err := system.InstallAutostart(); err != nil {
			fmt.Printf("Error installing autostart: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Time Tracker installed successfully.")
		return
	}
	
	// Handle background flag
	if *flagBackground {
		runBackground()
		return
	}
	
	// No special flags, run in normal mode
	// Check if there's an existing session
	startTime, err := session.CheckExistingSession()
	if err != nil {
		fmt.Printf("Error checking session: %v\n", err)
		os.Exit(1)
	}

	currentTime := time.Now().Format("02/01/2006, 15:04:05")
	fmt.Printf("Current time: %s\n", currentTime)

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
		fmt.Println("To run in background mode, use the -background flag.")
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
package session

import (
	"encoding/gob"
	"os"
	"path/filepath"
)

const (
	// SessionFile is the name of the file that stores the current time tracking session
	SessionFile = "current_session.bin"
)

// TimeSession represents a time tracking session with start time
type TimeSession struct {
	StartTime string
}

// GetSessionFilePath returns the absolute path to the session file
// The file is stored in the same directory as the executable
func GetSessionFilePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, SessionFile), nil
}

// CheckExistingSession checks if there's an existing time session
// Returns the start time if a session exists, empty string otherwise
func CheckExistingSession() (string, error) {
	filePath, err := GetSessionFilePath()
	if err != nil {
		return "", err
	}

	// Check if file exists
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", nil // File doesn't exist
	}

	// Open and decode the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var session TimeSession
	decoder := gob.NewDecoder(file)
	if err := decoder.Decode(&session); err != nil {
		return "", err
	}

	return session.StartTime, nil
}

// SaveSession saves the current time as a session start
// This creates a binary file with the encoded session data
func SaveSession(startTime string) error {
	filePath, err := GetSessionFilePath()
	if err != nil {
		return err
	}

	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode and save the session
	session := TimeSession{StartTime: startTime}
	encoder := gob.NewEncoder(file)
	return encoder.Encode(session)
}

// DeleteSession removes the time session file
// This is typically called after a session is completed
func DeleteSession() error {
	filePath, err := GetSessionFilePath()
	if err != nil {
		return err
	}

	// Remove the file if it exists
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return os.Remove(filePath)
	}
	return nil
}
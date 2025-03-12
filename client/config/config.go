package config

import (
	"os"
	"os/user"
)

// GetUsername returns the current computer username
// Falls back to "AutoTracker" if username cannot be determined
func GetUsername() string {
	// Try to get username from environment variable first
	if username := os.Getenv("USERNAME"); username != "" {
		return username
	}

	// If environment variable is not available, try using os/user package
	currentUser, err := user.Current()
	if err == nil && currentUser.Username != "" {
		return currentUser.Username
	}

	// Fall back to default name if all else fails
	return "Unknown"
}
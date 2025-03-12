package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// LoadEnv loads environment variables from .env file
func LoadEnv() error {
	// Get executable directory
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	exeDir := filepath.Dir(exePath)
	
	// Try to open .env file from executable directory
	envPath := filepath.Join(exeDir, ".env")
	file, err := os.Open(envPath)
	if err != nil {
		// If file doesn't exist, just return without error
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split by first equals sign
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		
		// Remove quotes if present
		value = strings.Trim(value, "\"'")
		
		// Set environment variable
		os.Setenv(key, value)
	}

	return scanner.Err()
}
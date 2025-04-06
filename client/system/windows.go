//go:build windows
// +build windows

package system

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Windows-specific implementation for shutdown detection
// In a real implementation, you would use the Windows API to register for shutdown events
// This is a simplified version that sets up the infrastructure

// setupWindowsShutdownDetection registers for Windows power events
func setupWindowsShutdownDetection(logger *log.Logger, handler EventHandler) error {
	logger.Println("Setting up Windows shutdown detection")
	
	// This is where you would register for Windows power events
	// Using the Win32 API RegisterPowerSettingNotification
	
	// For demonstration purposes, we're just setting up a goroutine
	// that will check system status periodically
	go func() {
		logger.Println("Windows shutdown monitor running")
		// In a real implementation, this would wait for Windows events
		// instead of just exiting
	}()
	
	return nil
}

// CreateWindowsStartupShortcut creates a shortcut in the Windows Startup folder
func CreateWindowsStartupShortcut() error {
	// Get path to executable
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	
	// Get path to Startup folder
	startupFolder := filepath.Join(os.Getenv("APPDATA"), 
		"Microsoft", "Windows", "Start Menu", "Programs", "Startup")
	
	// Create shortcut name
	shortcutPath := filepath.Join(startupFolder, "TimeTracker.lnk")
	
	// Simple PowerShell script to create shortcut
	psScript := `
$WshShell = New-Object -ComObject WScript.Shell
$Shortcut = $WshShell.CreateShortcut("` + shortcutPath + `")
$Shortcut.TargetPath = "` + exePath + `"
$Shortcut.Arguments = "-background"
$Shortcut.Save()
	`
	
	// Replace any potential escaped characters
	psScript = strings.ReplaceAll(psScript, "\r\n", "\n")
	
	// Execute PowerShell script
	cmd := exec.Command("powershell", "-Command", psScript)
	return cmd.Run()
}

// WindowsRegistryAutostart adds the app to Windows registry autostart
func WindowsRegistryAutostart() error {
	// Get path to executable
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	
	// Simple PowerShell script to add registry key
	psScript := `
$appPath = "` + exePath + ` -background"
Set-ItemProperty -Path "HKCU:\SOFTWARE\Microsoft\Windows\CurrentVersion\Run" -Name "TimeTracker" -Value $appPath
	`
	
	// Execute PowerShell script
	cmd := exec.Command("powershell", "-Command", psScript)
	return cmd.Run()
}
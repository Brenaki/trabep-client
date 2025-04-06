//go:build linux
// +build linux

package system

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

// Linux implementation for shutdown detection
// In a real implementation, this would use D-Bus to listen for login1 events

// setupLinuxShutdownDetection registers for Linux power events via D-Bus
func setupLinuxShutdownDetection(logger *log.Logger, handler EventHandler) error {
	logger.Println("Setting up Linux shutdown detection")
	
	// This is a simplified version
	// In a real implementation, you would use godbus/dbus to listen for
	// org.freedesktop.login1.Manager.PrepareForSleep and PrepareForShutdown signals
	
	// For demonstration purposes, we're just setting up a goroutine
	// that simply acknowledges we're monitoring
	go func() {
		logger.Println("Linux shutdown monitor running")
		// In a real implementation, this would wait for D-Bus signals
		// instead of just exiting
	}()
	
	return nil
}

// Linux autostart implementation using systemd user service
func CreateLinuxSystemdService() error {
	// Get path to executable
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	
	// Create systemd user directory if it doesn't exist
	userServiceDir := filepath.Join(os.Getenv("HOME"), ".config", "systemd", "user")
	if err := os.MkdirAll(userServiceDir, 0755); err != nil {
		return err
	}
	
	// Service template
	serviceTemplate := `[Unit]
Description=Time Tracker Background Service
After=network.target

[Service]
Type=simple
ExecStart={{.ExePath}} -background
Restart=on-failure
Environment=DISPLAY=:0

[Install]
WantedBy=default.target
`

	// Create template
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return err
	}
	
	// Create service file
	serviceFile := filepath.Join(userServiceDir, "timetracker.service")
	file, err := os.Create(serviceFile)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Execute template
	err = tmpl.Execute(file, struct{
		ExePath string
	}{
		ExePath: exePath,
	})
	if err != nil {
		return err
	}
	
	// Enable and start the service
	cmd := exec.Command("systemctl", "--user", "enable", "timetracker.service")
	if err := cmd.Run(); err != nil {
		return err
	}
	
	cmd = exec.Command("systemctl", "--user", "start", "timetracker.service")
	return cmd.Run()
}

// CreateLinuxDesktopAutostart creates a .desktop file in ~/.config/autostart
func CreateLinuxDesktopAutostart() error {
	// Get path to executable
	exePath, err := os.Executable()
	if err != nil {
		return err
	}
	
	// Create autostart directory if it doesn't exist
	autostartDir := filepath.Join(os.Getenv("HOME"), ".config", "autostart")
	if err := os.MkdirAll(autostartDir, 0755); err != nil {
		return err
	}
	
	// Desktop file template
	desktopTemplate := `[Desktop Entry]
Type=Application
Name=Time Tracker
Exec={{.ExePath}} -background
Terminal=false
X-GNOME-Autostart-enabled=true
`

	// Create template
	tmpl, err := template.New("desktop").Parse(desktopTemplate)
	if err != nil {
		return err
	}
	
	// Create desktop file
	desktopFile := filepath.Join(autostartDir, "timetracker.desktop")
	file, err := os.Create(desktopFile)
	if err != nil {
		return err
	}
	defer file.Close()
	
	// Execute template
	err = tmpl.Execute(file, struct{
		ExePath string
	}{
		ExePath: exePath,
	})
	
	return err
}
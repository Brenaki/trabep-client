#!/bin/bash

# Detect OS
case "$(uname -s)" in
    Linux*)     OS=Linux;;
    Darwin*)    OS=Mac;;
    CYGWIN*)    OS=Windows;;
    MINGW*)     OS=Windows;;
    *)          OS="UNKNOWN";;
esac

echo "Detected OS: $OS"
echo "Installing Time Tracker..."

# Build the application
go build -o timetracker

if [ $? -ne 0 ]; then
    echo "Error: Build failed!"
    exit 1
fi

# Install based on OS
if [ "$OS" = "Linux" ]; then
    # Create directory if it doesn't exist
    sudo mkdir -p /usr/local/bin
    
    # Copy binary
    sudo cp timetracker /usr/local/bin/
    
    # Setup systemd service for the current user
    mkdir -p ~/.config/systemd/user/
    cp timetracker.service ~/.config/systemd/user/
    
    # Enable and start the service
    systemctl --user daemon-reload
    systemctl --user enable timetracker.service
    systemctl --user start timetracker.service
    
    echo "Time Tracker installed as a systemd user service"
    echo "Service status: systemctl --user status timetracker.service"

elif [ "$OS" = "Mac" ]; then
    # Create directory if it doesn't exist
    sudo mkdir -p /usr/local/bin
    
    # Copy binary
    sudo cp timetracker /usr/local/bin/
    
    # Create LaunchAgent plist
    mkdir -p ~/Library/LaunchAgents
    cat > ~/Library/LaunchAgents/com.timetracker.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.timetracker</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/bin/timetracker</string>
        <string>-background</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
</dict>
</plist>
EOF
    
    # Load the LaunchAgent
    launchctl load ~/Library/LaunchAgents/com.timetracker.plist
    
    echo "Time Tracker installed as a LaunchAgent"
    echo "To check status: launchctl list | grep timetracker"

elif [ "$OS" = "Windows" ]; then
    # Create directory if it doesn't exist
    mkdir -p "$APPDATA/TimeTracker"
    
    # Copy binary
    cp timetracker.exe "$APPDATA/TimeTracker/"
    
    # Create shortcut in Startup folder
    powershell -Command "
    \$WshShell = New-Object -comObject WScript.Shell;
    \$Shortcut = \$WshShell.CreateShortcut(\"\$env:APPDATA\Microsoft\Windows\Start Menu\Programs\Startup\TimeTracker.lnk\");
    \$Shortcut.TargetPath = \"\$env:APPDATA\TimeTracker\timetracker.exe\";
    \$Shortcut.Arguments = \"-background\";
    \$Shortcut.Save();
    "
    
    echo "Time Tracker installed in startup folder"
    echo "Application will start on next login"
    
    # Start the application now
    start "" "$APPDATA/TimeTracker/timetracker.exe" -background
else
    echo "Unsupported OS: $OS"
    echo "Manual installation required"
    exit 1
fi

echo "Installation complete!"
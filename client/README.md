# Time Tracker Client

A simple desktop application for tracking time spent on activities. The client works with the Time Tracking API to store and manage time records.

## Features

- Simple start/stop time tracking
- Automatic user identification using computer username
- Desktop notifications for session events
- Seamless integration with Time Tracking API
- Cross-platform support (Windows, macOS, Linux)

## Prerequisites

- [Go](https://golang.org/dl/) (version 1.21 or higher)
- Running instance of the Time Tracking API (default: http://localhost:3000)

## Installation

1. Clone the repository
2. Navigate to the client directory
3. Build the application:

```bash
go build -o timetracker.exe
```

## Usage

The Time Tracker client is designed to be simple to use:

1. **Start a session**: Run the application to start tracking time
   ```bash
   .\timetracker.exe
   ```
   You'll see a notification that time tracking has started.

2. **End a session**: Run the application again to end the current session
   ```bash
   .\timetracker.exe
   ```
   The application will calculate the time spent and send it to the API.

3. **View time records**: Access the API's web interface or use the API endpoints to view your time records.

## How It Works

The Time Tracker client uses a simple session-based approach:

1. When you start the application for the first time, it creates a session file with the current timestamp
2. When you run it again, it reads the existing session file, calculates the time difference, and sends the data to the API
3. After successfully sending the data, it deletes the session file

## Configuration

The client automatically uses your computer's username for tracking. No additional configuration is required.

The API endpoint is set to `http://localhost:3000` by default. If you need to change this, modify the `BaseURL` constant in `api/client.go`.

## Project Structure

```
client/
├── api/            # API client code
├── config/         # Configuration utilities
├── models/         # Data models
├── notification/   # Desktop notification handlers
├── session/        # Session management
├── main.go         # Main application entry point
└── go.mod          # Go module definition
```

## Building for Different Platforms

### Windows
```bash
go build -o timetracker.exe
```

### macOS
```bash
GOOS=darwin GOARCH=amd64 go build -o timetracker
```

### Linux
```bash
GOOS=linux GOARCH=amd64 go build -o timetracker
```

## Troubleshooting

- **API Connection Issues**: Ensure the Time Tracking API is running and accessible at the configured URL
- **Session File Problems**: If you encounter issues with sessions, delete the `current_session.bin` file in the application directory
- **Notification Errors**: Desktop notifications require system support; they may not work in all environments

## License

This project is licensed under the MIT License.
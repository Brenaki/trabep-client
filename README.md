![Time Tracker](https://img.shields.io/badge/Time-Tracker-blue) ![API](https://img.shields.io/badge/API-Bun%20%2B%20TypeScript-green) ![Client](https://img.shields.io/badge/Client-Go-cyan)
# Time Tracker Project

A complete time tracking solution consisting of a backend API and a desktop client application. This project allows users to track time spent on activities with automatic user identification and detailed time calculations.

## Project Overview

The Time Tracker project consists of two main components:

1. **API Server**: A RESTful API built with Bun and TypeScript that handles time record storage and calculations
2. **Desktop Client**: A Go-based desktop application that provides a simple interface for starting and stopping time tracking

## Features

- **Simple Time Tracking**: Start and stop tracking with minimal interaction
- **Automatic User Detection**: Uses the computer's username for tracking
- **Detailed Time Calculations**: Tracks hours, minutes, and seconds spent
- **Desktop Notifications**: Provides feedback on tracking events
- **Data Persistence**: Stores all time records in a SQLite database
- **API Documentation**: Includes Swagger UI for exploring the API

## Repository Structure

```
trabep-client/
    ├── README.md
    ├── api                                                # Backend API server
    │   ├── .dockerignore
    │   ├── .gitignore
    │   ├── Dockerfile
    │   ├── README.md
    │   ├── bun.lockb
    │   ├── package.json                                  # Dependencies and scripts
    │   ├── src                                           # Source code
    │   │   ├── db                                        # Database operations
    │   │   │   ├── create-database.ts
    │   │   │   └── save-user-time-to-database.ts
    │   │   ├── index.ts
    │   │   ├── routes                                    # API endpoints
    │   │   │   ├── delete-time.ts
    │   │   │   ├── get-times.ts
    │   │   │   ├── swagger-ui.ts
    │   │   │   ├── user-time-data.dto.ts
    │   │   │   └── user-times.ts
    │   │   ├── swagger                                   # API documentation
    │   │   │   └── swagger.ts
    │   │   └── utils                                     # Utility functions
    │   │   │   └── calculate-time-difference.ts
    │   └── tsconfig.json                                 # TypeScript configuration
    └── client                                             # Desktop client application
        ├── .gitignore
        ├── README.md
        ├── api                                           # API client code
        │   └── client.go
        ├── config                                        # Configuration utilities
        │   ├── config.go
        │   └── env.go
        ├── go.mod                                        # Go module definition
        ├── go.sum
        ├── main.go                                       # Main entry point
        ├── models                                        # Data models
        │   └── types.go
        ├── notification                                  # Desktop notifications
        │   └── notifier.go
        └── session                                       # Session management
            └── manager.go
```


## Getting Started

### Prerequisites

- [Bun](https://bun.sh/) (v1.0.0 or higher) for the API
- [Go](https://golang.org/dl/) (v1.21 or higher) for the client
- Windows, macOS, or Linux operating system

### Installation

1. Clone this repository:
   ```bash
   git clone https://github.com/Brenaki/trabep-client.git
   cd trabep-client
   ```

2. Set up the API:
   ```bash
   cd api
   bun install
   ```

3. Build the client:
   ```bash
   cd ../client
   go build -o timetracker.exe
   ```

## Running the Application

### Start the API Server

```bash
cd api
bun run dev
```

The API will be available at http://localhost:3000. You can access the Swagger UI documentation at http://localhost:3000/api-docs.

### Use the Client Application

1. **Start tracking time**:
   ```bash
   cd client
   .\timetracker.exe
   ```
   You'll see a notification that time tracking has started.

2. **Stop tracking and record time**:
   Run the application again to end the current session and send the time record to the API.
   ```bash
   .\timetracker.exe
   ```

## How It Works

1. When you run the client for the first time, it creates a session file with the current timestamp
2. When you run it again, it reads the existing session file, calculates the time difference, and sends the data to the API
3. The API calculates the hours, minutes, and seconds spent and stores the record in the database
4. You can view all time records through the API endpoints

## API Endpoints

- `GET /user-times`: Get all time records
- `POST /user-times`: Create a new time record
- `DELETE /user-times?id=1`: Delete a time record by ID

## Building for Production

### API

```bash
cd api
bun run build
```

### Client

#### Windows
```bash
cd client
go build -o timetracker.exe
```

#### macOS
```bash
cd client
GOOS=darwin GOARCH=amd64 go build -o timetracker
```

#### Linux
```bash
cd client
GOOS=linux GOARCH=amd64 go build -o timetracker
```

## Troubleshooting

- **API Connection Issues**: Ensure the API server is running on http://localhost:3000
- **Session File Problems**: If you encounter issues with the client, delete the `current_session.bin` file
- **Database Issues**: Check the `api/data` directory for database files

## License

This project is licensed under the MIT License.

## Contributors

- [Brenaki](https://github.com/Brenaki)

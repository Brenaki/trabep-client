# Time Tracking API

A simple and efficient API for tracking time spent on activities. Built with Bun and TypeScript.

## Features

- Track time spent on activities with user identification
- Calculate hours, minutes, and seconds spent between start and end times
- Store time records in a SQLite database
- RESTful API with JSON responses
- Swagger UI documentation
- Docker support for easy deployment

## Prerequisites

- [Bun](https://bun.sh/) (v1.0.0 or higher)
- [Docker](https://www.docker.com/) (optional, for containerized deployment)

## Installation

1. Clone the repository
2. Navigate to the API directory
3. Install dependencies:

```bash
bun install
```

## Running the API

### Development mode (with auto-reload)

```bash
bun run dev
```

### Production mode

```bash
bun run start
```

The API will be available at http://localhost:3000.

## Docker Deployment

The API includes a Dockerfile for containerized deployment:

### Building the Docker image

```bash
docker build -t time-tracking-api .
```

### Running the Docker container

```bash
docker run -p 3000:3000 -d --name time-tracking-api time-tracking-api
```

### Persisting data with Docker volumes

To persist the database data between container restarts:

```bash
docker run -p 3000:3000 -v time-tracker-data:/app/data -d --name time-tracking-api time-tracking-api
```

This creates a Docker volume named `time-tracker-data` that will store your database files.

## API Documentation

Swagger UI documentation is available at http://localhost:3000/api-docs when the server is running.

## API Endpoints

### GET /user-times

Returns all time records stored in the database.

**Response:**

```json
{
  "success": true,
  "count": 2,
  "data": [
    {
      "id": 1,
      "user": "John",
      "start_time": "12/03/2023, 09:00:00",
      "end_time": "12/03/2023, 17:00:00",
      "hours_spent": 8,
      "minutes_spent": 0,
      "seconds_spent": 0,
      "created_at": "2023-03-12 17:01:00"
    },
    {
      "id": 2,
      "user": "Jane",
      "start_time": "13/03/2023, 10:00:00",
      "end_time": "13/03/2023, 15:30:00",
      "hours_spent": 5,
      "minutes_spent": 30,
      "seconds_spent": 0,
      "created_at": "2023-03-13 15:31:00"
    }
  ]
}
```

### POST /user-times

Creates a new time record.

**Request Body:**

```json
{
  "user": "John",
  "startTime": "12/03/2023, 09:00:00",
  "endTime": "12/03/2023, 17:00:00"
}
```

**Response:**

```json
{
  "success": true,
  "message": "Data received successfully",
  "data": {
    "user": "John",
    "startTime": "12/03/2023, 09:00:00",
    "endTime": "12/03/2023, 17:00:00"
  },
  "timeSpent": {
    "hours": 8,
    "minutes": 0,
    "seconds": 0,
    "formatted": "8h 0m 0s"
  },
  "savedToDatabase": true
}
```

### DELETE /user-times?id=1

Deletes a time record by ID.

**Response:**

```json
{
  "success": true,
  "message": "Record with ID 1 deleted successfully"
}
```

## Project Structure

```
api/
├── data/                  # Database files
├── src/
│   ├── db/                # Database operations
│   ├── routes/            # API route handlers
│   ├── swagger/           # Swagger documentation
│   ├── utils/             # Utility functions
│   └── index.ts           # Main application entry point
├── package.json
└── tsconfig.json
```

## Database Schema

The API uses a SQLite database with the following schema:

```sql
CREATE TABLE IF NOT EXISTS user_times (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user TEXT NOT NULL,
    start_time TEXT NOT NULL,
    end_time TEXT NOT NULL,
    hours_spent INTEGER NOT NULL,
    minutes_spent INTEGER NOT NULL,
    seconds_spent INTEGER NOT NULL,
    created_at TEXT DEFAULT CURRENT_TIMESTAMP
);
```

## Building for Production

To build the project for production:

```bash
bun run build
```

This will create a `dist` directory with the compiled JavaScript files.

## License

This project is licensed under the MIT License.

package models

// TimeEntry represents a single time entry
type TimeEntry struct {
	Timestamp string `json:"timestamp"`
}

// TimeQueue represents the queue of time entries
type TimeQueue struct {
	Entries []TimeEntry `json:"entries"`
	User    string      `json:"user"`
}

// CreateRecordRequest represents the request to create a new time record
type CreateRecordRequest struct {
	User      string `json:"user"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// TimeSpent represents the time spent information returned by the API
type TimeSpent struct {
	Hours     int    `json:"hours"`
	Minutes   int    `json:"minutes"`
	Seconds   int    `json:"seconds"`
	Formatted string `json:"formatted"`
}

// UserData represents the user data in the API response
type UserData struct {
	User      string `json:"user"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

// CreateRecordResponse represents the response from creating a record
type CreateRecordResponse struct {
	Success         bool      `json:"success"`
	Message         string    `json:"message"`
	Data            UserData  `json:"data"`
	TimeSpent       TimeSpent `json:"timeSpent"`
	SavedToDatabase bool      `json:"savedToDatabase"`
}
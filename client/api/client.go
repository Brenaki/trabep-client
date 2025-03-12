package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Brenaki/trabep-client/client/models"
)

const (
	// BaseURL is the API endpoint base URL
	BaseURL = "http://localhost:3000"
)

// SendToAPI sends the time record to the API
// Parameters:
// - user: The username for the time record
// - startTime: The start time in format "DD/MM/YYYY, HH:MM:SS"
// - endTime: The end time in format "DD/MM/YYYY, HH:MM:SS"
// Returns an error if the API request fails
func SendToAPI(user, startTime, endTime string) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Prepare request data
	requestData := models.CreateRecordRequest{
		User:      user,
		StartTime: startTime,
		EndTime:   endTime,
	}

	// Convert request data to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	// Send POST request to API
	resp, err := client.Post(BaseURL+"/user-times", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		
		// Try to parse error message from JSON response
		var errorResp struct {
			Error string `json:"error"`
		}
		if jsonErr := json.Unmarshal(body, &errorResp); jsonErr == nil && errorResp.Error != "" {
			return fmt.Errorf("API error: %s", errorResp.Error)
		}
		
		return fmt.Errorf("API returned error: %s - %s", resp.Status, string(body))
	}

	// Parse response
	var response models.CreateRecordResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	// Check if API operation was successful
	if !response.Success {
		return fmt.Errorf("API returned unsuccessful response")
	}

	// Display time spent information
	fmt.Printf("Time spent: %s (%dh %dm %ds)\n", 
		response.TimeSpent.Formatted, 
		response.TimeSpent.Hours,
		response.TimeSpent.Minutes,
		response.TimeSpent.Seconds)
	return nil
}
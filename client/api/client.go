package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Brenaki/trabep-client/client/models"
)

// getBaseURL returns the API base URL from environment variable or default
func getBaseURL() string {
	// Try to get URL from environment variable
	if envURL := os.Getenv("URL"); envURL != "" {
		return envURL
	}
	// Fall back to default localhost URL
	return "http://localhost:3000"
}

// getAPIEndpoint returns the API endpoint path from environment variable or default
func getAPIEndpoint() string {
	// Try to get endpoint from environment variable
	if endpoint := os.Getenv("API_ENDPOINT"); endpoint != "" {
		return endpoint
	}
	// Fall back to default endpoint
	return "/user-times"
}

// SendToAPI sends time tracking data to the API
func SendToAPI(user, startTime, endTime string) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Get base URL from environment or default
	baseURL := getBaseURL()
	
	// Ensure the base URL doesn't end with a trailing slash
	baseURL = strings.TrimSuffix(baseURL, "/")
	
	// Get API endpoint from environment or default
	endpoint := getAPIEndpoint()
	
	// Ensure endpoint starts with a slash
	if !strings.HasPrefix(endpoint, "/") {
		endpoint = "/" + endpoint
	}
	
	// Construct the full URL
	fullURL := baseURL + endpoint

	// Prepare request data
	requestData := models.CreateRecordRequest{
		User:      user,
		StartTime: startTime,
		EndTime:   endTime,
	}

	// Convert request data to JSON
	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	// Send POST request to API
	resp, err := client.Post(fullURL, "application/json", bytes.NewBuffer(jsonData))
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
	var response struct {
		Success         bool      `json:"success"`
		Message         string    `json:"message"`
		Data            models.UserData  `json:"data"`
		TimeSpent       models.TimeSpent `json:"timeSpent"`
		SavedToDatabase interface{} `json:"savedToDatabase"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	// Check if API operation was successful
	if !response.Success {
		return fmt.Errorf("API returned unsuccessful response")
	}

	// Convert response to our model
	result := models.CreateRecordResponse{
		Success:   response.Success,
		Message:   response.Message,
		Data:      response.Data,
		TimeSpent: response.TimeSpent,
	}

	// Handle savedToDatabase which might be a boolean or an object
	switch v := response.SavedToDatabase.(type) {
	case bool:
		result.SavedToDatabase = v
	case map[string]interface{}:
		// If it's an object, consider it successful if it exists
		result.SavedToDatabase = true
	default:
		// For any other type, default to false
		result.SavedToDatabase = false
	}

	// Display time spent information
	fmt.Printf("Time spent: %s (%dh %dm %ds)\n", 
		result.TimeSpent.Formatted, 
		result.TimeSpent.Hours,
		result.TimeSpent.Minutes,
		result.TimeSpent.Seconds)
	return nil
}
package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Brenaki/trabep-client/models"
)

const (
	BaseURL = "http://localhost:3000"
)

// SendToAPI sends the time record to the API
func SendToAPI(user, startTime, endTime string) error {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	requestData := models.CreateRecordRequest{
		User:      user,
		StartTime: startTime,
		EndTime:   endTime,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := client.Post(BaseURL+"/user-times", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API returned error: %s - %s", resp.Status, string(body))
	}

	var response models.CreateRecordResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if !response.Success {
		return fmt.Errorf("API returned unsuccessful response")
	}

	fmt.Printf("Time spent: %s (%d minutes)\n", response.TimeSpent.Formatted, response.TimeSpent.Minutes)
	return nil
}
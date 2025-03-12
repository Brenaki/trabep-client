package queue

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Brenaki/trabep-client/models"
)

const (
	QueueFile   = "time_queue.json"
	DefaultUser = "AutoTracker"
)

// GetQueueFilePath returns the path to the queue file
func GetQueueFilePath() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, QueueFile), nil
}

// LoadQueue loads the time queue from a file
func LoadQueue(filePath string) (models.TimeQueue, error) {
	var queue models.TimeQueue

	data, err := os.ReadFile(filePath)
	if err != nil {
		return queue, err
	}

	err = json.Unmarshal(data, &queue)
	return queue, err
}

// SaveQueue saves the time queue to a file
func SaveQueue(filePath string, queue models.TimeQueue) error {
	data, err := json.MarshalIndent(queue, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

// AddEntry adds a new time entry to the queue
func AddEntry(queue *models.TimeQueue, timestamp string) {
	queue.Entries = append(queue.Entries, models.TimeEntry{Timestamp: timestamp})
}

// ClearQueue removes all entries from the queue
func ClearQueue(queue *models.TimeQueue) {
	queue.Entries = []models.TimeEntry{}
}

// HasTwoOrMoreEntries checks if the queue has at least two entries
func HasTwoOrMoreEntries(queue models.TimeQueue) bool {
	return len(queue.Entries) >= 2
}
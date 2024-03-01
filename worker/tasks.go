package worker

import (
	"fmt"
	"strconv"
	"strings"
)

type Tasker interface {
	ConfirmReceived(task Task) (bool, error)
	ConfirmCompleted(task Task) (bool, error)
	ConfirmError(task Task, err error) (bool, error)
	DoTask(task Task) error
}

func execute(t Tasker, task Task) error {
	err := t.DoTask(task)
	return err
}

type SlayEnemyWorker struct {
	Enemies string
	Reward  string
}

func (w *SlayEnemyWorker) DoTask(t Task) error {
	return nil
}

func (w *SlayEnemyWorker) ConfirmReceived(t Task) (bool, error) {
	return true, nil
}

func (w *SlayEnemyWorker) ConfirmCompleted(t Task) (bool, error) {
	return true, nil
}

func (w *SlayEnemyWorker) ConfirmError(t Task, err error) (bool, error) {
	return true, nil
}

type Task struct {
	ID       int
	Type     string
	Payload  map[string]string
	Priority string
	Metadata map[string]string
}

func parseTaskFromString(taskStr string) (*Task, error) {
	lines := strings.Split(taskStr, "\n")
	if len(lines) < 5 {
		return nil, fmt.Errorf("invalid task string format")
	}

	// Parse ID
	idStr := strings.TrimPrefix(lines[0], "Task ID: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid ID format: %v", err)
	}

	// Type
	taskType := strings.TrimPrefix(lines[1], "Type: ")

	// Parse Payload
	payloadStr := strings.TrimPrefix(lines[2], "Payload: ")
	payloadStr = strings.Trim(payloadStr, "[]")
	payloadItems := strings.Split(payloadStr, " ")
	payload := make(map[string]string)
	for _, item := range payloadItems {
		keyValue := strings.Split(item, ":")
		if len(keyValue) == 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.Trim(strings.TrimSpace(keyValue[1]), ",")
			payload[key] = value
		}
	}

	// Priority
	priority := strings.TrimPrefix(lines[3], "Priority: ")

	// Parse Metadata
	metadataStr := strings.TrimPrefix(lines[4], "Metadata: ")
	metadataStr = strings.Trim(metadataStr, "[]")
	metadataItems := strings.Split(metadataStr, " ")
	metadata := make(map[string]string)
	for _, item := range metadataItems {
		keyValue := strings.Split(item, ":")
		if len(keyValue) == 2 {
			key := strings.TrimSpace(keyValue[0])
			value := strings.Trim(strings.TrimSpace(keyValue[1]), ",")
			metadata[key] = value
		} else {
			return nil, fmt.Errorf("incorrect formatting of metada. It can not include ':'")
		}
	}

	return &Task{
		ID:       id,
		Type:     taskType,
		Payload:  payload,
		Priority: priority,
		Metadata: metadata,
	}, nil
}

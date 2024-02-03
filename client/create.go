package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type Task struct {
	Type     string            `json:"Type"`
	Payload  map[string]string `json:"Payload"`
	Priority string            `json:"Priority"`
	Metadata map[string]string `json:"Metadata"`
}

func genRandomTask() Task {
	taskTypes := []string{"SlayEnemy", "CollectHerbs"}
	priorities := []string{"High", "Medium", "Low"}

	// Generating a random task
	task := Task{
		Type:     taskTypes[rand.Intn(len(taskTypes))],
		Payload:  map[string]string{"data": fmt.Sprintf("random_data_%d", rand.Intn(100))},
		Priority: priorities[rand.Intn(len(priorities))],
		Metadata: map[string]string{"timestamp": time.Now().Format(time.RFC3339)},
	}

	return task
}

func taskToJSON(task Task) ([]byte, error) {
	taskJSON, err := json.Marshal(task)
	if err != nil {
		return nil, err
	}
	return taskJSON, nil
}

func main() {
	url := "http://localhost:25520/new-job"
	method := "POST"

	client := &http.Client{}
	payload, err := taskToJSON(genRandomTask())
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

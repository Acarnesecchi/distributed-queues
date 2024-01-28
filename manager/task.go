package manager

import (
	"fmt"
	"strings"
)

type Task struct {
	ID       int               `json:"-"`
	Type     string            `json:"Type"`
	Payload  map[string]string `json:"Payload"`
	Priority string            `json:"Priority"`
	Metadata map[string]string `json:"Metadata"`
}

func (t *Task) String() string {
	payload := strings.TrimPrefix(fmt.Sprintf("%v", t.Payload), "map")
	metadata := strings.TrimPrefix(fmt.Sprintf("%v", t.Metadata), "map")

	return fmt.Sprintf("Task ID: %d\nType: %s\nPayload: %s\nPriority: %s\nMetadata: %s",
		t.ID, t.Type, payload, t.Priority, metadata)
}

type TaskSlice struct {
	Tasks          []Task
	CompletedTasks []Task
}

func (s *TaskSlice) AddTask(t Task) {
	s.Tasks = append(s.Tasks, t)
}

func (s *TaskSlice) CompleteTask(ID int) bool {
	for i, t := range s.Tasks {
		if t.ID == ID {
			s.Tasks = append(s.Tasks[:i], s.Tasks[i+1:]...)
			s.CompletedTasks = append(s.CompletedTasks, t)
			return true
		}
	}
	return false
}

func (s *TaskSlice) SyncTasks() {
	// etcd, CRD, Operator or PV to interact with a kubernetes resource to sync the Task list
	// alternatively, if not using k8s can be sync'd on memory using mutex
	// for the moment this is nothing more than a happy idea :)
}

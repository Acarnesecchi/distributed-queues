package worker

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"
)

func StartConnection(c Config) {
	ok := validateConf(c.Tasks...)
	if !ok {
		log.Fatalf("tasks must start in uppercase and not contain special characters")
	}
	conn, err := net.Dial(c.ConnMode, c.TargetAddr)
	if err != nil {
		fmt.Printf("Error connection to %s: %v", c.TargetAddr, err)
		return
	}
	id := register(conn, c.Tasks)
	for {
		ok := waitForTasks(conn, id)
		if ok {
			break
		}
		// log.Fatalf("we good")
	}
	defer conn.Close()
}

func waitForTasks(c net.Conn, id string) bool {
	taskBytes := make([]byte, 1024)
	_, err := c.Read(taskBytes)
	if err != nil {
		return false
	}
	taskStr := string(bytes.Trim(taskBytes[:], "\x00"))

	T, err := parseTaskFromString(string(taskStr))
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	var w *KillRatWorker
	execute(w, *T)
	return true
}

func register(c net.Conn, tasks []string) string {
	taskString := "tasks: " + strings.Join(tasks, ",")
	_, err := c.Write([]byte(taskString))
	if err != nil {
		fmt.Printf("could not register, retrying in %d seconds", 10)
	}
	id := make([]byte, 36)
	_, err = c.Read(id)
	if err != nil {
		fmt.Printf("could not assign ID, retrying in %d seconds. %v \n", 10, err)
	}
	return string(bytes.Trim(id[:], "\x00"))
}

var taskRegex = regexp.MustCompile("^[A-Z][a-zA-Z]*$")

func validateConf(t ...string) bool {
	if len(t) == 0 {
		log.Fatalf("worker must be able to do at least one task")
	}
	for i := range t {
		if !taskRegex.MatchString(t[i]) {
			return false
		}
	}
	return true
}

package worker

import (
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
	register(conn, c.Tasks)
	defer conn.Close()
}

func register(c net.Conn, tasks []string) {
	taskString := "tasks: " + strings.Join(tasks, ", ")
	_, err := c.Write([]byte(taskString))
	if err != nil {
		fmt.Printf("could not register, retrying in %d seconds", 10)
	}
	id := make([]byte, 1024)
	_, err = c.Read(id)
	if err != nil {
		fmt.Printf("could not assign ID, retrying in %d seconds. %v \n", 10, err)
	}
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

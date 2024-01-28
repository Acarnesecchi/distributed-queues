package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

var (
	TotalTasks int    = 0
	MemServer  Server //for testing only
)

type Worker struct {
	Connection  net.Conn
	Tasks       []string
	ConnectedAt time.Time
}

type Server struct {
	config     Config
	WorkerList map[string]Worker
	StartTime  time.Time
}

func NewServer(c Config) Server {
	MemServer.config = c
	return MemServer
	// return Server{
	// 	config: c,
	// }
}

func StartServer(s Server) {
	addr := s.config.ListenAddr
	network := s.config.ConnMode
	s.StartTime = time.Now()
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatalf("could not listen on: %s\n", addr)
	}
	defer listener.Close()

	fmt.Printf("Listening on %s\n", addr)

	go waitForConnections(&s, listener)
	waitForJobs()
}

func handleClient(s *Server, conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		return
	}
	msg := string(bytes.Trim(buf[:], "\x00"))
	fmt.Print(msg)

	firstWord := strings.Fields(msg)[0]
	fmt.Println(firstWord)

	switch firstWord {
	case "tasks:":
		id, _ := uuid.NewRandom()
		fmt.Println(id)
		taskList := strings.Split(strings.TrimPrefix(msg, "tasks:"), ",")
		w := Worker{
			Connection:  conn,
			Tasks:       taskList,
			ConnectedAt: time.Now(),
		}
		if s.WorkerList == nil {
			s.WorkerList = make(map[string]Worker)
		}
		s.WorkerList[id.String()] = w
		MemServer = *s
		idBytes := []byte(id.String())
		conn.Write(idBytes)
	case "id":
		fmt.Print("aa")
	default:
		fmt.Println("nothing happened")
	}
}

func waitForJobs() {
	receiveJobs := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var task Task
		task.ID = assignID()
		d := json.NewDecoder(r.Body)
		err := d.Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Printf("Received task: %+v\n", task)
		distributeTask(task)
	}
	http.HandleFunc("/new-job", receiveJobs)
	if err := http.ListenAndServe(":25520", nil); err != nil {
		log.Fatalf("Error starting server on port 25520: %v", err)
	}
}

func assignID() int {
	// this should check on the etcd or a PV
	// rather than being in one manager memory
	TotalTasks++
	return TotalTasks
}

func distributeTask(t Task) {
	availableWorkers := make([]Worker, 0)

	for _, worker := range MemServer.WorkerList {
		if contains(worker.Tasks, t.Type) {
			availableWorkers = append(availableWorkers, worker)
		}
	}

	conn := availableWorkers[0].Connection
	conn.Write([]byte(t.String()))
}

func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

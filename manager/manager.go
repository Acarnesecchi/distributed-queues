package manager

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Worker struct {
	Address     string
	Tasks       []string
	ConnectedAt time.Time
}

type Server struct {
	config     Config
	WorkerList map[string]Worker
	StartTime  time.Time
}

func NewServer(c Config) Server {
	return Server{
		config: c,
	}
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

	waitForConnections(&s, listener)

}

func handleClient(s *Server, conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		return
	}
	msg := string(buf)
	fmt.Print(msg)

	firstWord := strings.Fields(msg)[0]
	fmt.Println(firstWord)

	switch firstWord {
	case "tasks:":
		id, _ := uuid.NewRandom()
		taskList := strings.Split(strings.TrimPrefix(msg, "tasks:"), ",")
		w := Worker{
			Address:     conn.RemoteAddr().String(),
			Tasks:       taskList,
			ConnectedAt: time.Now(),
		}
		if s.WorkerList == nil {
			s.WorkerList = make(map[string]Worker)
		}
		s.WorkerList[id.String()] = w
		conn.Write(id[:])
	case "id":
		fmt.Print("aa")
	default:
		fmt.Println("nothing happened")
	}
	defer conn.Close()
}

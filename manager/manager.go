package manager

import (
	"fmt"
	"log"
	"net"
	"time"
)

type Worker struct {
	Address     string
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

	waitForConnections(s.config.Timeout, listener)

}

func handleClient(conn net.Conn) {
	defer conn.Close()
}

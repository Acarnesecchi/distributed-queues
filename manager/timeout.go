package manager

import (
	"fmt"
	"net"
)

func waitForConnections(s *Server, l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Printf("Established a connection with %s\n", conn.LocalAddr().String())
		go handleClient(s, conn)
	}
}

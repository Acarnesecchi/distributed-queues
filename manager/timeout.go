package manager

import (
	"fmt"
	"log"
	"net"
	"time"
)

func waitForConnections(t time.Duration, l net.Listener) {
	con := make(chan string)

	go func() {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error:", err)
		}
		con <- conn.LocalAddr().String()
		handleClient(conn)
	}()

	select {
	case addr := <-con:
		fmt.Printf("Established a connection with %s\n", addr)
	case <-time.After(t):
		log.Fatal("server timeout. Could not establish a connection\n")
	}
}

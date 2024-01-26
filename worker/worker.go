package worker

import (
	"fmt"
	"net"
)

func StartConnection(c Config) {
	conn, err := net.Dial(c.ConnMode, c.TargetAddr)
	if err != nil {
		fmt.Printf("Error connection to %s: %v", c.TargetAddr, err)
		return
	}
	defer conn.Close()
}

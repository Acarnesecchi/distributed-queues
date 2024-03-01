package main

import "github.com/Acarnesecchi/distributed-queues/manager"

func main() {
	c := manager.NewConfig().WithListenAddr("localhost:15255")
	s := manager.NewServer(c)
	manager.StartServer(s)
}

package main

import "github.com/Acarnesecchi/distributed-queue/manager"

func main() {
	conf := manager.NewConfig()
	server := manager.NewServer(conf)
	manager.StartServer(server)
}

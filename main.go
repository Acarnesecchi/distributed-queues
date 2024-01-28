package main

import (
	"time"

	"github.com/Acarnesecchi/distributed-queue/manager"
	"github.com/Acarnesecchi/distributed-queue/worker"
)

func main() {
	conf := manager.NewConfig()
	server := manager.NewServer(conf)
	go func() {
		time.Sleep(2 * time.Second)
		worker.StartConnection(worker.NewConfig().WithTasks("JamonAsado", "Profiler", "KillRat"))
	}()
	manager.StartServer(server)
}

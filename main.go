package main

import (
	"time"

	"github.com/Acarnesecchi/distributed-queues/manager"
	"github.com/Acarnesecchi/distributed-queues/worker"
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

package main

import "github.com/Acarnesecchi/distributed-queues/worker"

func main() {
	c := worker.NewConfig().WithConnMode("tcp").WithTargetAddr("localhost:15255").WithTasks("MeteoData").WithTimeout(30)
	worker.StartConnection(c)
}

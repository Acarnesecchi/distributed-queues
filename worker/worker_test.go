package worker_test

import (
	"testing"

	"github.com/Acarnesecchi/distributed-queues/worker"
)

func TestConn(t *testing.T) {
	worker.StartConnection(worker.NewConfig().WithTasks("KillRat", "Bingchilling", "Bruh"))
}

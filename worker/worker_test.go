package worker_test

import (
	"testing"

	"github.com/Acarnesecchi/distributed-queue/worker"
)

func TestConn(t *testing.T) {
	worker.StartConnection(worker.NewConfig().WithTasks("KillRat", "Bingchilling", "Bruh"))
}

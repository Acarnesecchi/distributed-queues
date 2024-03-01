package main

import "github.com/Acarnesecchi/distributed-queues/worker"

type Clicker struct {
	Clicks int
	Reward string
}

func (c *Clicker) DoTask(t worker.Task) error {
	return nil
}

func (c *Clicker) ConfirmReceived(t worker.Task) (bool, error) {
	return true, nil
}

func (c *Clicker) ConfirmCompleted(t worker.Task) (bool, error) {
	return true, nil
}

func (c *Clicker) ConfirmError(t worker.Task, err error) (bool, error) {
	return true, nil
}

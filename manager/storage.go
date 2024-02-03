package manager

import "fmt"

func (s *Server) addWorker(w Worker, ID string) {
	switch s.config.Storage {
	case "memory":
		s.WorkerList[ID] = w
	case "kubernetes":
		fmt.Println("k8s storage not implemented yet")
	}
}

func (s *Server) removeWorker(ID string) {
	switch s.config.Storage {
	case "memory":
		delete(s.WorkerList, ID)
	case "kubernetes":
		fmt.Println("k8s storage not implemented yet")
	}
}

func (s *Server) countTasks() int {
	switch s.config.Storage {
	case "memory":
		return len(s.TaskList.Tasks)
	case "kubernetes":
		fmt.Println("k8s storage not implemented yet")
		return 0
	default:
		return -1
	}
}

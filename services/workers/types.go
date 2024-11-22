package workers

import "sync"

type Service interface {
	Work(task Task)
	Shutdown()
}

type Task func()

type Impl struct {
	taskChan chan Task
	wg       sync.WaitGroup
}

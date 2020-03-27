package scheduler

import "Project/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (scheduler *SimpleScheduler) WorkerChan() chan engine.Request {
	return scheduler.workerChan
}

func (scheduler *SimpleScheduler) WorkReady(chan engine.Request) {
}

func (scheduler *SimpleScheduler) Run() {
	scheduler.workerChan = make(chan engine.Request)
}

func (scheduler *SimpleScheduler) Submit(request engine.Request) {
	// send request down to worker chan
	go func() {
		scheduler.workerChan <- request
	}()
}

package scheduler

import "Project/crawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (scheduler *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (scheduler *QueuedScheduler) Submit(request engine.Request) {
	scheduler.requestChan <- request
}

func (scheduler *QueuedScheduler) WorkReady(workerChan chan engine.Request) {
	scheduler.workerChan <- workerChan
}

func (scheduler *QueuedScheduler) ConfigureMasterWorkerChan(chan engine.Request) {
	panic("implement me")
}

func (scheduler *QueuedScheduler) Run() {
	scheduler.workerChan = make(chan chan engine.Request)
	scheduler.requestChan = make(chan engine.Request)

	go func() {
		var requestQueue []engine.Request
		var workerQueue []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}

			select {
			case request := <-scheduler.requestChan:
				requestQueue = append(requestQueue, request)
			case workerChan := <-scheduler.workerChan:
				workerQueue = append(workerQueue, workerChan)
			case activeWorker <- activeRequest:
				workerQueue = workerQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}

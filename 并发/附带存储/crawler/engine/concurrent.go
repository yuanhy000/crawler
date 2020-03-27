package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// Separate the method as a Independent interface
type ReadyNotifier interface {
	WorkReady(chan Request)
}

func (engine *ConcurrentEngine) Run(seeds ...Request) {
	outputChan := make(chan ParseResult)
	engine.Scheduler.Run()

	for i := 0; i < engine.WorkCount; i++ {
		// simple: every worker use one channel
		// queue: every worker use a different channel
		// Determined by scheduler
		createWorker(engine.Scheduler.WorkerChan(), outputChan, engine.Scheduler)
	}

	for _, request := range seeds {
		engine.Scheduler.Submit(request)
	}

	for {
		// get the city list from out channel, {ParseResult}
		result := <-outputChan
		for _, item := range result.Items {
			go func() { engine.ItemChan <- item }()
		}

		// fetching request, add request into the scheduler
		for _, request := range result.Requests {
			engine.Scheduler.Submit(request)
		}
	}
}

func createWorker(inputChan chan Request, outputChan chan ParseResult,
	ready ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i'm ready
			ready.WorkReady(inputChan)

			request := <-inputChan
			result, err := worker(request)
			if err != nil {
				continue
			}
			outputChan <- result
		}
	}()
}

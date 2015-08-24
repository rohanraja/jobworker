package jobworker

import "time"

func Run() {

	jobRequestQueue := make(chan JobRequest, 1000)
	jobResultQueue := make(chan JobResult, 5000)

	reqprocessExit := make(chan int)
	jobFetcherSignal := make(chan int)
	resultDispatcherSignal := make(chan int)

	go JobsFetcher(jobRequestQueue, jobResultQueue, jobFetcherSignal)

	for i := 0; i < Config.NumWorkers; i++ {
		go ProcessQueue(jobRequestQueue, reqprocessExit)
	}

	go ResultsDispatcher(jobResultQueue, resultDispatcherSignal)

	time.Sleep(10000 * time.Second)
	// SignalPoller()

}

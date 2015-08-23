package jobworker

import "time"

func Run() {

	jobRequestQueue := make(chan JobRequest)
	jobResultQueue := make(chan JobResult)

	reqprocessExit := make(chan int)
	jobFetcherSignal := make(chan int)
	resultDispatcherSignal := make(chan int)

	go JobsFetcher(jobRequestQueue, jobResultQueue, jobFetcherSignal)

	for i := 0; i < Config.NumWorkers; i++ {
		go ProcessQueue(jobRequestQueue, reqprocessExit)
	}

	go ResultsDispatcher(jobResultQueue, resultDispatcherSignal)

	time.Sleep(200 * time.Second)
	// SignalPoller()

}

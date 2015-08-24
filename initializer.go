package jobworker

var reqprocessExit chan int

func Run() {

	jobRequestQueue := make(chan JobRequest, 1000)
	jobResultQueue := make(chan JobResult, Config.NumWorkers)

	reqprocessExit = make(chan int)
	jobFetcherSignal := make(chan int)
	resultDispatcherSignal := make(chan int)

	go JobsFetcher(jobRequestQueue, jobResultQueue, jobFetcherSignal)

	for i := 0; i < Config.NumWorkers; i++ {
		go ProcessQueue(jobRequestQueue, reqprocessExit)
	}

	go ResultsDispatcher(jobResultQueue, resultDispatcherSignal)

	StartWebServer()
	// time.Sleep(10000 * time.Second)
	// SignalPoller()

}

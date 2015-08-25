package jobworker

func Run() {

	jobRequestQueue := make(chan JobRequest, 1000)
	jobResultQueue := make(chan JobResult, Config.NumWorkers)

	// reqprocessExit = make(chan int)
	jobFetcherSignal := make(chan int)
	resultDispatcherSignal := make(chan int)

	go JobsFetcher(jobRequestQueue, jobResultQueue, jobFetcherSignal)

	workForce = NewWorkForce(jobRequestQueue)
	workForce.StartWorking()

	// for i := 0; i < Config.NumWorkers; i++ {
	// 	go ProcessQueue(jobRequestQueue, reqprocessExit)
	// }

	go ResultsDispatcher(jobResultQueue, resultDispatcherSignal)

	HandleSignals()

	StartWebServer()
	// time.Sleep(10000 * time.Second)
	// SignalPoller()

}

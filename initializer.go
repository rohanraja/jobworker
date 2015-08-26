package jobworker

func Run() {

	jobRequestQueue := make(chan JobRequest, 1000)
	jobResultQueue := make(chan JobResult, Config.NumWorkers)

	jobFetcherSignal := make(chan int)
	resultDispatcherSignal := make(chan int)

	go JobsFetcher(jobRequestQueue, jobResultQueue, jobFetcherSignal)

	workForce = NewWorkForce(jobRequestQueue)
	workForce.StartWorking()

	go ResultsDispatcher(jobResultQueue, resultDispatcherSignal)

	HandleSignals()

	StartWebServer()

}

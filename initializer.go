package jobworker

func Run() {

	jobRequestQueue := make(chan JobRequest, Config.RequestQueueSize)
	jobResultQueue := make(chan JobResult, Config.NumWorkers)

	jobFetcherSignal := make(chan int)
	resultDispatcherSignal := make(chan int)

	go JobsFetcher(jobRequestQueue, jobResultQueue, jobFetcherSignal)

	workForce = NewWorkForce(jobRequestQueue)
	workForce.StartWorking()

	go ResultsDispatcher(jobResultQueue, resultDispatcherSignal)

	HandleSignals()

	go StartWebServer()

	go PeriodicInfoUpdater()

	RunTerminalUI()

	GracefullQuit()
}

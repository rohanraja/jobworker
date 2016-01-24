package jobworker

func Run() {

	jobRequestQueue := make(chan JobRequest, Config.RequestQueueSize)
	jobResultQueue := make(chan JobResult, Config.NumWorkers)

	jobFetcherSignal := make(chan int)
	resultDispatcherSignal := make(chan int)

	workForce = NewWorkForce(jobRequestQueue, jobResultQueue)
	workForce.StartWorking()
	workForce.LoadActiveJobs()

	go JobsFetcher(jobRequestQueue, jobResultQueue, jobFetcherSignal)

	go ResultsDispatcher(jobResultQueue, resultDispatcherSignal)

	HandleSignals()

	go StartWebServer()

	go PeriodicInfoUpdater()

	// time.Sleep(20 * time.Second)
	RunTerminalUI()

	GracefullQuit()
}

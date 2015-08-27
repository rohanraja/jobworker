package jobworker

type WorkForce struct {
	ExitSignalChannel chan int
	NumWorkers        int
	JobRequestQueue   chan JobRequest
}

var workForce *WorkForce

func NewWorkForce(jrq chan JobRequest) *WorkForce {

	reqSignal := make(chan int)

	wf := WorkForce{reqSignal, Config.NumWorkers, jrq}

	return &wf

}

func (w *WorkForce) ChangeNumWorkers(num int) {

	if num > w.NumWorkers {
		w.NumWorkers = num - w.NumWorkers
		w.StartWorking()

	} else {

		w.NumWorkers = w.NumWorkers - num
		w.StopJobWorkers()
	}
	w.NumWorkers = num

}

func (w *WorkForce) StartWorking() {

	for i := 0; i < w.NumWorkers; i++ {
		go ProcessQueue(w.JobRequestQueue, w.ExitSignalChannel)
	}
}

func (w *WorkForce) StopJobWorkers() {

	for i := 0; i < w.NumWorkers; i++ {

		w.ExitSignalChannel <- 1
	}
}

func ProcessQueue(queue chan JobRequest, exitSignaler chan int) {

	defer func() {
		// color.Red("\nExiting JobWorker")
	}()

	// color.Green("\nStarting Jobworker")

	for {

		select {

		case job := <-queue:
			job.ProcessRequest()

			select {
			case <-exitSignaler:
				return

			default:

			}

		case <-exitSignaler:
			return

		}

	}

}

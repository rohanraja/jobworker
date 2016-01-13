package jobworker

type WorkForce struct {
	ExitSignalChannel chan int
	NumWorkers        int
	JobRequestQueue   chan JobRequest
	ActiveJobs        map[*JobRequest]bool
}

var workForce *WorkForce

func NewWorkForce(jrq chan JobRequest) *WorkForce {

	reqSignal := make(chan int)

	var jobsmap map[*JobRequest]bool
	wf := WorkForce{reqSignal, Config.NumWorkers, jrq, jobsmap}

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

	w.ActiveJobs = make(map[*JobRequest]bool)

	for i := 0; i < w.NumWorkers; i++ {
		go w.ProcessQueue()
	}
}

func (w *WorkForce) StopJobWorkers() {

	for i := 0; i < w.NumWorkers; i++ {
		w.ExitSignalChannel <- 1
	}
}

func (w *WorkForce) ProcessQueue() {

	for {
		select {

		case job := <-w.JobRequestQueue:

			w.ActiveJobs[&job] = true
			job.ProcessRequest()
			delete(w.ActiveJobs, &job)

			select {
			case <-w.ExitSignalChannel:
				return

			default:
			}

		case <-w.ExitSignalChannel:
			return

		}

	}

}

func (w *WorkForce) GetStatusAll() (res string) {

	res = ""

	for job := range w.ActiveJobs {
		outb, _ := job.GetStatus()
		res = res + string(outb) + "\n"
	}
	return
}
func (w *WorkForce) ExitGracefullyAll() (res string) {

	res = ""

	for job := range w.ActiveJobs {
		_ = job.ExitGracefully()
	}
	return
}

package jobworker

import "github.com/fatih/color"

func StopJobWorkers() {

	for i := 0; i < Config.NumWorkers; i++ {

		reqprocessExit <- 1
	}
}

func ProcessQueue(queue chan JobRequest, exitSignaler chan int) {

	defer func() {
		color.Yellow("\nExiting JobWorker")
	}()

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

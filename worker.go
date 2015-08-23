package jobworker

func ProcessQueue(queue chan JobRequest, exitSignaler chan int) {

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

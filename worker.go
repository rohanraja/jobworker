package jobworker

func processQueue(queue chan JobRequest, exitSignaler chan bool) {

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

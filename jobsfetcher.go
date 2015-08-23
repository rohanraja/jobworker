package jobworker

import "time"

func JobsFetcher(reqQueue chan JobRequest, resultQueue chan JobResult, signaler chan int) {

	for i := 0; i < 20; i++ {

		requests := FetchRequests()

		for _, r := range requests {
			r.ResultsChannel = resultQueue
			reqQueue <- r
		}

		select {
		case <-signaler:
			return
		default:
		}

		time.Sleep(Config.FetchPollDelay)

	}
}

func FetchRequests() []JobRequest {

	var jinfo JobInfo
	var jrequest JobRequest

	jinfo.BinaryKey = "parsebinary"
	jrequest.Jobinfo = jinfo

	out := []JobRequest{jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest}

	return out
}

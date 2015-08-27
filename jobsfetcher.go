package jobworker

import (
	"encoding/json"
	"time"
)

func JobsFetcher(reqQueue chan JobRequest, resultQueue chan JobResult, signaler chan int) {

	for i := 0; i < 100; i++ {

		requests := FetchRequests(Config.Fetch_Binkey)
		Messages = append(Messages, "Got new jobs")
		// color.Yellow("\nGot %d new jobs", len(requests))
		if len(requests) == 0 {
			time.Sleep(15 * time.Second)
		}

		for _, r := range requests {
			r.ResultsChannel = resultQueue
			reqQueue <- r
		}
		// color.White("\nAdded to queue, len=%d", len(resultQueue))

		select {
		case <-signaler:
			return
		default:
		}

		time.Sleep(Config.FetchPollDelay)

	}
}

func FetchRequests(binkey string) (reqs []JobRequest) {

	for i := 0; i < Config.NumFetches; i++ {

		var jinfo JobInfo
		var jrequest JobRequest

		jinfoStr, _ := FetchJob(binkey)

		if jinfoStr == "" {
			return
		}

		json.Unmarshal([]byte(jinfoStr), &jinfo)

		jrequest.Jobinfo = jinfo

		// color.Green("Got Jinfo: %v", jinfo)

		reqs = append(reqs, jrequest)
	}

	return

}

func FetchRequests_Mock() []JobRequest {

	var jinfo JobInfo
	var jrequest JobRequest

	jinfo.BinaryKey = "parsebinary"
	jrequest.Jobinfo = jinfo

	out := []JobRequest{jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest}

	return out
}

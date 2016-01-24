package jobworker

import (
	"encoding/json"
	"time"
)

func JobsFetcher(reqQueue chan JobRequest, resultQueue chan JobResult, signaler chan int) {

	for i := 0; i < 100; i++ {

		if len(workForce.JobRequestQueue) > (Config.RequestQueueSize / 7) {
			time.Sleep(Config.FetchPollDelay)
			continue
		}

		requests := FetchRequests(Config.Fetch_Binkey)
		// requests := FetchRequests_Mock()
		// Messages = append(Messages, "Got new jobs")
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

	for i := 0; i < Config.RequestQueueSize; i++ {

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

	jinfo.BinaryKey = "a.out"
	jinfo.Args = "/Users/rraja/code/cgt_distributed/examples/"
	jinfo.Args = "4_0"
	jinfo.Jobid = "mockJobId"

	jrequest.Jobinfo = jinfo

	out := []JobRequest{jrequest, jrequest}
	// out := []JobRequest{jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest, jrequest}

	return out
}

package jobworker

import "github.com/fatih/color"
import "time"

var stTime time.Time

func ResultsDispatcher(resultQueue chan JobResult, signaler chan int) {

	stTime = time.Now()
	for {

		select {

		case jr := <-resultQueue:
			ProcessResult(&jr)

		case <-signaler:
			return

		}
	}
}

var cnt int = 0

func ProcessResult(jresult *JobResult) {

	cnt += 1
	elapsed := time.Since(stTime)

	rate := float64(cnt) / elapsed.Seconds()

	color.Cyan("#%d - Rate: %f jobs/seconds\n", cnt, rate)
	if jresult.Status != 0 {
		color.Red("Error: %s", jresult.ErrorMsg)
	}

	DispatchResult(jresult)

}

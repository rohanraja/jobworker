package jobworker_test

import (
	"testing"
	"time"

	"github.com/fatih/color"
	"github.com/rohan1020/jobworker"
)

func initReq() jobworker.JobRequest {

	var jinfo jobworker.JobInfo
	var jrequest jobworker.JobRequest

	jinfo.BinaryKey = "a.out"
	jinfo.Args = "/Users/rraja/code/cgt_distributed/examples/"
	jrequest.Jobinfo = jinfo

	return jrequest
}

func TestExecutingRequest(t *testing.T) {

	jrequest := initReq()

	execu := func() {
		str, err := jrequest.Executor()

		color.Green("%v", str)
		color.Red("%v", err)
		if err != nil {
			t.Error("Exoected Error")
		}
	}

	go execu()

	time.Sleep(5 * time.Second)

	stat, _ := jrequest.GetStatus()

	color.Green("\nStatus:\n%s", stat)

	time.Sleep(5 * time.Second)

	stat, _ = jrequest.GetStatus()

	color.Green("\nStatus:\n%s", stat)

	time.Sleep(25 * time.Second)

}

func TesTProcessingRequest(t *testing.T) {

	j := initReq()
	resChan := make(chan jobworker.JobResult)
	j.ResultsChannel = resChan

	go j.ProcessRequest()

	jr := <-resChan

	color.Yellow("%v", jr)
}

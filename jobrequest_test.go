package jobworker_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/rohan1020/jobworker"
)

func initReq() jobworker.JobRequest {

	var jinfo jobworker.JobInfo
	var jrequest jobworker.JobRequest

	jinfo.BinaryKey = "parsebinary"
	jrequest.Jobinfo = jinfo

	return jrequest
}

func TesTExecutingRequest(t *testing.T) {

	jrequest := initReq()
	str, err := jrequest.Executor()

	color.Green("%v", str)
	color.Red("%v", err)

	if err != nil {
		t.Error("Exoected Error")
	}

}

func TesTProcessingRequest(t *testing.T) {

	j := initReq()
	resChan := make(chan jobworker.JobResult)
	j.ResultsChannel = resChan

	go j.ProcessRequest()

	jr := <-resChan

	color.Yellow("%v", jr)
}

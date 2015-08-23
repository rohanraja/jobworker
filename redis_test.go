package jobworker_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/rohan1020/jobworker"
)

func TesTGettingRandomJids(t *testing.T) {

	out := jobworker.GetPendingJids("parsebin")

	color.Green("%v", out)
}
func TesTMovingJid(t *testing.T) {

	jid := jobworker.GetPendingJids("parsebin")

	jobworker.MoveJidToProcessingSet("parsebin", jid)

	color.Green("%v", jid)
}
func TesTGettingJobInfo(t *testing.T) {

	jid := jobworker.GetPendingJids("parsebin")

	out := jobworker.GetJobInfo("parsebin", jid)

	color.Cyan(out)

}
func TesTFetchingJob(t *testing.T) {

	out, _ := jobworker.FetchJob("parsebin")

	color.Cyan(out)

}

func TesTDispatchingResult(t *testing.T) {

	jr := jobworker.JobResult{}
	jr.BinaryKey = "parsebin"

	jobworker.DispatchResult(&jr)

}

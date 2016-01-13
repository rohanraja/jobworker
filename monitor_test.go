package jobworker_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/rohan1020/jobworker"
)

func TestJobMonitor(t *testing.T) {

	jinfos := jobworker.GetActiveJobs()

	color.Green("%+v", jinfos)
}

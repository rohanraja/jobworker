package jobworker_test

import (
	"testing"

	"github.com/rohan1020/jobworker"
)

func TestServing(t *testing.T) {

	jobworker.StartWebServer()
}

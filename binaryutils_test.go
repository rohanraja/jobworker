package jobworker_test

import (
	"testing"

	"github.com/rohan1020/jobworker"
)

func TestWorkingBinaryExecutor(t *testing.T) {

	_, err := jobworker.ExecuteBinary("ls", "-a")

	if err != nil {
		t.Error(err)
	}
}
func TestFailingBinaryExecutor(t *testing.T) {

	_, err := jobworker.ExecuteBinary("lssdsd", "-a")

	if err == nil {
		t.Error("Expected an error.")
	}
}

func TestCheckingFileExistance(t *testing.T) {

	fileStr := "/Users/rohanraja/code/nlppro/data5.csv"

	out := jobworker.CheckIfFileExists(fileStr)

	if out == false {
		t.Error("File should exist")
	}

	fileStr = "/Users/rohanraja/code/nlppro/data5xx.csv"

	out = jobworker.CheckIfFileExists(fileStr)

	if out == true {
		t.Error("File should not exist")
	}

}

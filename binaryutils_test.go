package jobworker_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/rohan1020/jobworker"
)

func TestBinaryExecutor(t *testing.T) {

	color.Blue("test")

	out, err := jobworker.ExecuteBinary("ls", "-a")

	color.Yellow("\n%v\n%v", out, err)

}

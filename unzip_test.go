package jobworker_test

import (
	"testing"

	"github.com/fatih/color"
	"github.com/rohan1020/jobworker"
)

func TesTUnzipping(t *testing.T) {

	inpFile := "test.zip"
	outp := "outzip"
	err := jobworker.Unzip(inpFile, outp)

	if err != nil {
		color.Red("%v", err)
		t.Error("Exoected Error")
	}

}
func TesTZipping(t *testing.T) {

	inpFile := "test2.zip"
	outp := "outzip"
	err := jobworker.Zipit(outp, inpFile)

	if err != nil {
		color.Red("%v", err)
		t.Error("Exoected Error")
	}

}

package jobworker

import (
	"log"
	"os"
)

var Log *log.Logger

func init() {

	Log = log.New(os.Stdout, "", 0)

}

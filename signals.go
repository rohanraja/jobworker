package jobworker

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fatih/color"
)

func GracefullQuit() {

	fmt.Println("\nQuitting Gracefully..")
	DispatchMassResults()
	workForce.SaveActiveJobs()
	workForce.JobRequestQueue = make(chan JobRequest)
	workForce.ExitGracefullyAll()
	workForce.StopJobWorkers()
	os.Exit(0)
}

func DisplayInfo() {

	color.Yellow("\n\n******* GO WORKER INFORMATION *******\n\n")

	infoObj := GetInfoObj()

	color.Green("%+v", infoObj)

	color.Yellow("\n\n*************************************\n\n")

}

func HandleSignals() {
	sigs := make(chan os.Signal, 1)
	infosigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(infosigs, syscall.SIGTSTP)

	go func() {

		quitInProgress := false

		for {

			<-sigs
			if quitInProgress == true {
				os.Exit(1)
			}
			go GracefullQuit()
			quitInProgress = true

		}

	}()
	go func() {

		for {

			<-infosigs
			DisplayInfo()

		}

	}()

}

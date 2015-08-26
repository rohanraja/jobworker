package jobworker

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func GracefullQuit() {

	fmt.Println("\nQuitting Gracefully..")
	DispatchMassResults()
	workForce.StopJobWorkers()
	os.Exit(0)
}

func HandleSignals() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
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

}

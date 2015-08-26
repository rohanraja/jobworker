package jobworker

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func HandleSignals() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		for {

			<-sigs
			fmt.Println("\nQuitting..")
			DispatchMassResults()
			workForce.StopJobWorkers()
			os.Exit(3)

		}

	}()

}

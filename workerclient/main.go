package main

import (
	"github.com/fatih/color"
	"github.com/rohan1020/jobworker"
)

func main() {

	color.Yellow("Job Worker Client")

	jobworker.Run()

}

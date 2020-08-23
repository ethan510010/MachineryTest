package main

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1"
	"github.com/ethan510010/go_machinery_learning/utils"
)

func StartWorker(taskserver *machinery.Server) error {
	worker := taskserver.NewWorker("machinery_worker", 0)
	if err := worker.Launch(); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("start consuming tasks")
	taskserver := utils.GetMachineryServer()
	StartWorker(taskserver)
}

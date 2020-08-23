package main

import (
	"net/http"

	"github.com/RichardKnop/machinery/v1"
	"github.com/ethan510010/go_machinery_learning/server"
	"github.com/ethan510010/go_machinery_learning/utils"
)

var taskServer *machinery.Server

func init() {
	taskServer = utils.GetMachineryServer()
}

func main() {
	router := server.StartServer(taskServer)
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err)
	}

	// start consumer
	// worker.StartWorker(taskServer)
}

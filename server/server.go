package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/gorilla/mux"
)

var taskServer *machinery.Server

func StartServer(taskserver *machinery.Server) *mux.Router {

	taskServer = taskserver

	router := mux.NewRouter()

	subRouter := router.PathPrefix("/").Subrouter()
	subRouter.HandleFunc("/sendMail", sendMailHandler).Methods("POST")

	return router
}

func sendMailHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 處理 https request 送進來的參數
	decoder := json.NewDecoder(r.Body)
	var params map[string]string
	decoder.Decode(&params)

	targetEmailAddress := params["emailTo"]

	task := &tasks.Signature{
		Name: "send_mail",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: targetEmailAddress,
			},
		},
	}

	asyncResult, _ := taskServer.SendTask(task)

	res := map[string]string{
		"task_uuid": asyncResult.GetState().TaskUUID,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		fmt.Print(err)
		return
	}
}

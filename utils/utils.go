package utils

import (
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/ethan510010/go_machinery_learning/tasks"
	// "honnef.co/go/tools/config"
)

func GetMachineryServer() *machinery.Server {
	taskServer, err := machinery.NewServer(&config.Config{
		Broker:        "redis://127.0.0.1:6379",
		ResultBackend: "redis://127.0.0.1:6379",
	})
	if err != nil {
		panic(err)
	}

	taskServer.RegisterTask("send_mail", tasks.SendMail)

	return taskServer
}

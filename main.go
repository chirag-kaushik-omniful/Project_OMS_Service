package main

import (
	"context"
	"fmt"
	"time"

	"oms/routes"
	"oms/utils/dbconn"
	"oms/utils/intsrv"
	"oms/utils/kafka"
	"oms/utils/sqs"

	"github.com/omniful/go_commons/config"
	server "github.com/omniful/go_commons/http"
)

func main() {
	err := config.Init(15 * time.Second)
	if err != nil {
		panic(err)
	}

	srv := server.InitializeServer(config.GetString(context.Background(), "server.port"), 0, 0, 0)
	srv.Use(config.Middleware())

	routes.GetRouter(srv)
	dbconn.Connect(config.GetString(context.Background(), "mongodb.url"))
	intsrv.InitInterSrvClient()
	sqs.InitSQS()
	kafka.InitKafka()

	err = srv.StartServer("oms")
	if err != nil {
		fmt.Println("Server error:", err)
	}

}

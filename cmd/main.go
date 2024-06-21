package main

import (
	"fmt"
	"log"
	"net/http"
	agent "practice/agentservice"
	controller "practice/controllerservice"
	nats "practice/nats"
	payload "practice/payloadservice"
	"time"
)

func main() {
	nats.Init()

	go controller.OperateInstances()
	payload.InstanceHandler()

	handler := agent.NewRouter(&agent.Controller{})

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		WriteTimeout: 10 * time.Second,
	}
	errors := make(chan error, 1)
	go func() {
		fmt.Println("Запущен агент")
		errors <- server.ListenAndServe()
	}()

	if err := <-errors; err != nil {
		log.Fatalf("Агент остановил работу! : %v", err)
	}

}

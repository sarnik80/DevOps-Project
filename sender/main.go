package main

import (
	"fmt"
	"log"
	"net/http"
	"p1-sender/broker"
	"p1-sender/handlers"
)

const port string = ":9090"

func main() {

	// rabbitMQ config
	var (
		user       = "guest"
		password   = "guest"
		host       = "rbt"
		rabbitPort = "5672"
		queueName  = "hello"
	)

	conn := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, host, rabbitPort)
	messageBroker, err := broker.Init(conn, queueName)
	if err != nil {
		log.Fatal("faild to connect to rabbitMQ")
	}
	defer messageBroker.Channel.Close()
	handlers := handlers.New(messageBroker)

	server := http.Server{
		Handler: GetRoutes(handlers),
		Addr:    port,
	}
	log.Printf("server is up on port%s\n", port)
	err = server.ListenAndServe()
	log.Fatal(err.Error())
}

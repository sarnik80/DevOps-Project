package broker

import (
	"encoding/json"
	"log"
	"p1-reciever/dbRepo/repository"
	"p1-reciever/models"

	amqp "github.com/rabbitmq/amqp091-go"
)

type rabbit struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

var rabbitMQ *rabbit

func InitRabbit(connectionString, channelName string) error {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return err
	}
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	q, err := ch.QueueDeclare(
		channelName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	queue := &rabbit{
		conn: conn,
		ch:   ch,
		q:    q,
	}
	rabbitMQ = queue
	return nil
}

func Listen(postgreRepo *repository.PostgresRepo) error {
	defer rabbitMQ.conn.Close()
	defer rabbitMQ.ch.Close()
	msgs, err := rabbitMQ.ch.Consume(
		rabbitMQ.q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for d := range msgs {
		message := models.Message{}
		err := json.Unmarshal(d.Body, &message)
		if err != nil {
			log.Println("invalid json message")
			continue
		}
		postgreRepo.InsertMessage(&message)
	}
	return nil
}

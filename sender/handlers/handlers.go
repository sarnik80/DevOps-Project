package handlers

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"io"
	"log"
	"net/http"
	"p1-sender/broker"
	"p1-sender/models"
	"time"
)

type Repository struct {
	messageBroker *broker.RabbitMQ
}

func New(broker *broker.RabbitMQ) *Repository {
	return &Repository{messageBroker: broker}
}

func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := &models.Message{
		Payload: "server is up",
	}

	response, _ := json.Marshal(message)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (repo *Repository) Send(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		message := &models.Message{
			Payload: "Invalid jbody payload",
		}
		response, _ := json.Marshal(message)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(response)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Millisecond)
	defer cancel()
	err = repo.messageBroker.Channel.PublishWithContext(ctx,
		"",
		repo.messageBroker.Queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "appliction/json",
			Body:        requestBody,
		})
	if err != nil {
		log.Println(err.Error())
		message := &models.Message{
			Payload: "Internal error",
		}
		response, _ := json.Marshal(message)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}
	message := &models.Message{
		Payload: "Message Sent",
	}
	response, _ := json.Marshal(message)
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

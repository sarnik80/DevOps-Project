package broker

import amqp "github.com/rabbitmq/amqp091-go"

type RabbitMQ struct {
	Queue   *amqp.Queue
	Channel *amqp.Channel
}

func Init(connString, queueName string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(connString)
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	q, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	result := &RabbitMQ{
		Queue:   &q,
		Channel: ch,
	}

	return result, nil
}

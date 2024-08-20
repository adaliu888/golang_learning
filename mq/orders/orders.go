package main

import (
	"context"
	"encoding/json"
	"log"

	common "golang_learning/mq/common"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	user = "guest"
	pass = "guest"
	host = "localhost"
	port = "5672"
	//queueName = "test"
)

func main() {

	ch, close := common.ConnectAmqp(user, pass, host, port)
	defer func() {
		close()
		ch.Close()
	}()

	q, err := ch.QueueDeclare(common.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	marshaledorder, err := json.Marshal(common.Order{
		ID: "order-1",
		Items: []common.Item{
			{ID: "item-1", Quantity: 2},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	// Set Context
	/*
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err = ch.PublishWithContext( // errors here
			ctx,    // context
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        marshaledorder,
			})
	*/

	err = ch.PublishWithContext(context.Background(), "",
		q.Name, false, false, amqp.Publishing{
			ContentType: "application/json",
			Body:        marshaledorder,
		})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Order published")

}

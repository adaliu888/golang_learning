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
	//创建一个新的通道ch
	ch, close := common.ConnectAmqp(user, pass, host, port)
	defer func() {
		close()
		ch.Close()
	}()
	// 声明一个queue，声明时需要指定是否持久化，是否排他，是否自动删除，是否等待队列非空等参数。
	// 若queueName不存在，RabbitMQ会创建一个新的queue。
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
	// Publish order to queue
	// 若Publish函数返回error，那么err参数将包含该error。
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

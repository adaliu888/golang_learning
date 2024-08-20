package main

import (
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

func CreatePaymentLink() (string, error) {
	//TODO: implement payment link creation logic
	return "payment-link-123", nil
}
func main() {
	ch, close := common.ConnectAmqp(user, pass, host, port)
	defer func() {
		close()
		ch.Close()
	}()

	listen(ch)
}

func listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(common.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// 处理消息体
			o := &common.Order{}
			if err := json.Unmarshal(d.Body, &o); err != nil {
				d.Nack(false, false) // 拒绝消息，不重新排队
				log.Printf("Error unmarshalling JSON: %v", err)
				continue
			}

			// 业务逻辑处理
			paymentLink, err := CreatePaymentLink()
			if err != nil {
				d.Nack(false, false) // 拒绝消息，不重新排队
				log.Printf("Error creating payment link: %v", err)
				continue
			}

			// 确认消息处理成功
			d.Ack(false)

			log.Printf("Payment link is %v", paymentLink)
			// 这里可以添加更多的业务逻辑处理
		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

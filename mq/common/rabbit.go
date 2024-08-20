package common

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

//创建一个MQ生产者客户端

/*
消息队列的作用：
异步，将同步的消息变为异步，例如我们可以使用rpc调用另一个服务，但是我们必须等待返回（同步），用mq可以变异步
解耦，将单体服务拆分多个微服务，实现了分布式部署，单个服务的修改、增加或删除，不影响其他服务，不需要全部服务关闭重启
抗压，由于是异步，解耦的，高并发请求到来时，我们不直接发送给服务，而是发给MQ，让服务决定什么时候接收消息，提供服务，这样就缓解了服务的压力
图示：
用户注册后发邮件和虚拟币：
异步解耦图：
————————————————

                            版权声明：本文为博主原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接和本声明。

原文链接：https://blog.csdn.net/weixin_50134791/article/details/120851969
*/

const (
	OrderCreatedEvent = "Order.Created"
)

func ConnectAmqp(user, pass, host, port string) (*amqp.Channel, func() error) {

	// 建立一个rabbit连接
	addr := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, pass, host, port)
	Conn, err := amqp.Dial(addr)
	defer Conn.Close()
	log.Fatal("Failed to connect to RabbitMQ", err)
	//在RabbitMQ中，通道是进行消息交换的地方。需要为每个任务或操作创建一个通道：
	Channel, err := Conn.Channel()
	log.Fatal(err, "Failed to open a channel")
	defer Channel.Close()

	//声明队列和交换机：在通道上声明队列和交换机，以及它们之间的绑定关系：

	err = Channel.ExchangeDeclare(
		"Exchange", // 队列名称
		"direct",   // 交换机类型
		true,       // 是否持久化
		false,      // 是否自动删除
		false,      // 是否排他性
		false,      // 是否阻塞
		nil,        // 额外属性
	)
	if err != nil {
		log.Fatal(err, "Failed to declare an exchange")
	}

	err = Channel.ExchangeDeclare(
		"OrderCreatedEvent", // 队列名称
		"fanout",            // 交换机类型
		true,                // 是否持久化
		false,               // 是否自动删除
		false,               // 是否排他性
		false,               // 是否阻塞
		nil,                 // 额外属性
	)
	if err != nil {
		log.Fatal(err, "Failed to declare an exchange")
	}
	return Channel, Conn.Close
	/*// 将队列绑定到交换机，这里使用空字符串作为绑定键，因为fanout交换机不检查绑定键
	err = ch.QueueBind(
		"",                  // 队列名称
		"",                  // 绑定键
		"OrderCreatedEvent", // 交换机名称
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to bind a queue to an exchange: %s", err)
	}
	*/

	/*err = ch.Publish(
		"q.name",     // 交换机名称
		q.Name, // 路由键
		false,  // 是否强制
		false,  // 是否立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("Hello RabbitMQ"),
		},
	)
	*/

}

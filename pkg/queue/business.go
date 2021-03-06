package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"golang-dts/business"
	"golang-dts/pkg/logging"
)

// 初始化连接exchange 和 routingKey
func NewRabbitMQRouting(queueInfo int) *RabbitMQ {
	return NewRabbitMQ(queueInfo)
}

func (r *RabbitMQ) PublishMessage (message string) {
	err := r.channel.ExchangeDeclare(
		r.ExchangeName,
		r.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		logging.Fatal("交换机声明失败："+err.Error(), message)
		panic("交换机声明失败："+err.Error())
	}

	/** 消息推送 */
	err = r.channel.Publish(r.ExchangeName, r.RouteName, false, false, amqp.Publishing{
		ContentType : "text/plain",
		Body:[]byte(message),
		ContentEncoding: "UTF-8",
	})

	if err != nil {
		logging.Fatal("消息推送失败："+err.Error(), message)
		panic("消息推送失败："+err.Error())
	}
}

func (r *RabbitMQ) ReceiveRouting(business business.InterfaceQueue) {
	// 试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.ExchangeName,
		r.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		logging.Fatal("交换机声明失败："+err.Error())
		panic("交换机声明失败："+err.Error())
	}

	_, err = r.channel.QueueDeclare(
		r.QueueName,
		true,
		false,
		false,
		false,
		nil,
		)
	if err != nil {
		logging.Fatal("队列声明失败："+err.Error())
		panic("队列声明失败："+err.Error())
	}

	// 绑定队列到exchange中
	err = r.channel.QueueBind(
		r.QueueName,
		r.RouteName,
		r.ExchangeName,
		false,
		nil,
		)
	if err != nil {
		logging.Fatal("队列绑定交换机失败："+err.Error())
		panic("队列绑定交换机失败："+err.Error())
	}

	// 获取队列消息进行消费
	msg, err := r.channel.Consume(
		r.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
		)
	forever := make(chan bool)

	// 协程处理 队列监听消费
	go func() {
		for info := range msg {
			business.Delivery(info)
			logging.Info(string([]byte(info.Body)))
			fmt.Printf("%s",info.Body)
			fmt.Printf("%T",info.Body)
		}
	}()
	<-forever
}



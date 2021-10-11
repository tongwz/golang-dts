package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"golang-dts/pkg/logging"
	"golang-dts/pkg/setting"
	"golang-dts/pkg/e"
	"time"
)

// RabbitMQ结构体
type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	QueueName string // 队列名
	ExchangeName string // 交换机
	ExchangeType string // 交换机类型
	RouteName string // 路由
	MqUrl string // 连接信息
}

// 初始化rabbitmq对象 进行后续操作
func NewRabbitMQ(queueInfo int) *RabbitMQ {
	// 获取rabbitmq的连接配置
	sec, err := setting.Cfg.GetSection("rabbitmq")
	if err != nil {
		logging.Info("获取rabbitmq连接配置失败", err)
	}
	// 连接字符串
	configMq := fmt.Sprintf(
		"amqp://%s:%s@%s:%s", // 账号:密码@地址：端口/vhost
				sec.Key("USER").MustString("admin"),
				sec.Key("PASSWORD").MustString("123456"),
				sec.Key("HOST").MustString("127.0.0.1"),
				sec.Key("PORT").MustString("5672"),
				// sec.Key("VHOST").MustString("/"),
		)
	// 连接配置
	amqpConfig := amqp.Config{
		Vhost: sec.Key("VHOST").MustString("/"),
		Heartbeat: 1 * time.Minute,
	}
	// 结构体实例
	rabbitMQ := &RabbitMQ{
		ExchangeName:e.PublishQueueMap[queueInfo].ExchangeName,
		RouteName:e.PublishQueueMap[queueInfo].RouteName,
		QueueName:e.PublishQueueMap[queueInfo].QueueName,
		ExchangeType: "direct",
		MqUrl: configMq,
	}
	rabbitMQ.conn, err = amqp.DialConfig(rabbitMQ.MqUrl, amqpConfig)
	if err != nil {
		logging.Info("rabbitmq连接失败:", err.Error())
		panic("rabbitmq连接失败:"+ err.Error())
	}

	rabbitMQ.channel, err = rabbitMQ.conn.Channel()
	if err != nil {
		logging.Info("rabbitmq channel创建失败:", err.Error())
		panic("rabbitmq channel创建失败:"+ err.Error())
	}
	return rabbitMQ
}

func (r *RabbitMQ) Destroy() {
	r.channel.Close()
	r.conn.Close()
}

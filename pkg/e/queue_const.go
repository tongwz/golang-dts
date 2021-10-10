package e

const (
	RTest = iota

)
// 队列信息结构体
type publishQueue struct {
	ExchangeName string
	RouteName string
	QueueName string
}

var PublishQueueMap = map[int] publishQueue {
	RTest : {
		ExchangeName: "transfer",
		RouteName: "test",
		QueueName: "test-queue",
	},
}


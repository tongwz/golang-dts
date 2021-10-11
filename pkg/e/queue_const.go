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

// 对应命令行执行queue监听
var CommandString = map[string]int {
	"test-queue" : RTest,
	"user-upload" :1,
}


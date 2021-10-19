package e

import (
	"golang-dts/business"
	"golang-dts/business/scrm"
)

const (
	RTest = iota
	AppNewUserReport
)
// 队列信息结构体
type publishQueue struct {
	ExchangeName string
	RouteName string
	QueueName string

	businessQueue business.InterfaceQueue
}

var PublishQueueMap = map[int] publishQueue {
	// 测试
	RTest : {
		ExchangeName: "transfer",
		RouteName: "test",
		QueueName: "test-queue",
		businessQueue: nil,
	},
	// 用户上报
	AppNewUserReport : {
		ExchangeName: "user",
		RouteName: "scrm-app-new-user-report",
		QueueName: "scrm-app-new-user-report",

		businessQueue:new(scrm.UserReportHandle),
	},
}

// 对应命令行执行queue监听
var 	CommandString = map[string]int {
	"test-queue" : RTest,
	"scrm-app-new-user-report" :AppNewUserReport,
}

// 消费逻辑实例
func Receive(businessInt int) business.InterfaceQueue {
	return PublishQueueMap[businessInt].businessQueue
}


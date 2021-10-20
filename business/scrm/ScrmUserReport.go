package scrm

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"golang-dts/business"
	"golang-dts/business/dataForm"
	"golang-dts/pkg/logging"
)

// interface 的方法实现必须要通过new(struct) 我们直接继承通用结构体 来实现InterfaceQueue
type UserReportHandle struct {
	business.SQueue
}

func (m *UserReportHandle) SetLogName() string {
	return "scrmUserReport"
}

func (m *UserReportHandle) Delivery(msg amqp.Delivery) {
	message := new(dataForm.ScrmUserReportFrom)
	//message := make(map[string]interface{})
	//isAck := true
	if err := json.Unmarshal(msg.Body, &message); err != nil {
		logging.Fatal("scrm用户上报数据解析失败：", err.Error())
		return
	}
	userInfo := message.Data.Data
	regSource, ok := userInfo["reg_source"]
	if ok {
		fmt.Println(regSource)
	}

	fmt.Println(message.Data.Data)
	logging.Info([]byte(msg.Body))
	return
}


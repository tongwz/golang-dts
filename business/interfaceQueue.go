package business

import "github.com/streadway/amqp"

type InterfaceQueue interface {
	SetLogName() string
	Delivery(msg amqp.Delivery)
}

type SQueue struct {

}

func (i *SQueue) SetLogName() string {
	return ""
}

func (i *SQueue) Delivery( msg *amqp.Delivery){

}

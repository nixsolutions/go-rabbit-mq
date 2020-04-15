package config

import "github.com/streadway/amqp"

type ConsumerConfig struct {
	Queue     string
	Consumer  string
	AutoAck   bool
	Exclusive bool
	NoLocal   bool
	NoWait    bool
	Args      amqp.Table
}

func NewConsumerConfigByQueueName(name string) *ConsumerConfig {
	return &ConsumerConfig{
		Queue:     name,
		Consumer:  name + "_consumer",
		AutoAck:   false,
		Exclusive: false,
		NoLocal:   false,
		NoWait:    false,
		Args:      nil,
	}
}

func (cc *ConsumerConfig) GetConfig() (string, string, bool, bool, bool, bool, amqp.Table) {
	return cc.Queue, cc.Consumer, cc.AutoAck, cc.Exclusive, cc.NoLocal, cc.NoWait, cc.Args
}

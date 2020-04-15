package config

import "github.com/streadway/amqp"

type QueueConfig struct {
	Name       string
	Durable    bool // survive or not after server restarts
	AutoDelete bool // delete or not when has`t activities on queues
	Exclusive  bool // redeclare if same names
	NoWait     bool
	Args       amqp.Table
}

func NewQueueConfig(name string) *QueueConfig {
	return &QueueConfig{
		Name:       name,
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil,
	}
}

func (qc *QueueConfig) GetConfig() (string, bool, bool, bool, bool, amqp.Table) {
	return qc.Name, qc.Durable, qc.AutoDelete, qc.Exclusive, qc.NoWait, qc.Args
}

type BindQueueConfig struct {
	QueueName  string
	RoutingKey string
	Exchange   string
	NoWait     bool
	Args       amqp.Table
}

func NewBindQueueConfigByQueueName(name string) *BindQueueConfig {
	return &BindQueueConfig{
		QueueName:  name,
		RoutingKey: name,
		Exchange:   name,
		NoWait:     false,
		Args:       nil,
	}
}

func (bqc *BindQueueConfig) GetConfig() (string, string, string, bool, amqp.Table) {
	return bqc.QueueName, bqc.RoutingKey, bqc.Exchange, bqc.NoWait, bqc.Args
}

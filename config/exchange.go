package config

import "github.com/streadway/amqp"

type ExchangeConfig struct {
	Name       string
	Type       string
	Durable    bool
	AutoDelete bool
	Internal   bool
	NoWait     bool
	Args       amqp.Table
}

func NewExchangeConfigByQueueName(name string) *ExchangeConfig {
	return &ExchangeConfig{
		Name:       name,
		Type:       "topic",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil,
	}
}

func (ec *ExchangeConfig) GetConfig() (string, string, bool, bool, bool, bool, amqp.Table) {
	return ec.Name, ec.Type, ec.Durable, ec.AutoDelete, ec.Internal, ec.NoWait, ec.Args
}

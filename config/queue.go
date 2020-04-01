package config

import "github.com/streadway/amqp"

type QueueConfig struct {
	Name       string
	Durable    bool  // survive or not after server restarts
	AutoDelete bool  // delete or not when has`t activities on queues
	Exclusive  bool  // redeclare if same names
	NoWait     bool
	Args       amqp.Table
}


package channel

import (
	"github.com/streadway/amqp"
	"gitlab.nixdev.co/golang-general/rabbit-mq-go/config"
	"gitlab.nixdev.co/golang-general/rabbit-mq-go/connection"
)

type ChannelInterface interface {
	Create(connection.RConnectionInterface) error
	Close() error
	QueueDeclare(config.QueueConfig) error
	ExchangeDeclare(config.ExchangeConfig) error
	BindQueue(string) error
	Publish(string, amqp.Publishing) error
}

type RChannelInterface interface {
	QueueDeclare(string, bool, bool, bool, bool, amqp.Table) (amqp.Queue, error)
	ExchangeDeclare(string, string, bool, bool, bool, bool, amqp.Table) error
	QueueBind(string, string, string, bool, amqp.Table) error
	Publish(string, string, bool, bool, amqp.Publishing) error
	Qos(int, int, bool) error
	Consume(string, string, bool, bool, bool, bool, amqp.Table) (<-chan amqp.Delivery, error)
	Close() error
}

type Channel struct {
	channel   RChannelInterface
	QueueName string
}

func (ch *Channel) GetChannel() RChannelInterface {
	return ch.channel
}

func (ch *Channel) Create(conn connection.RConnectionInterface) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}
	ch.channel = channel
	return nil
}

func (ch *Channel) Close() error {
	return ch.channel.Close()
}

func (ch *Channel) QueueDeclare(queueConfig config.QueueConfig) error {
	q, err := ch.channel.QueueDeclare(
		queueConfig.Name,
		queueConfig.Durable,
		queueConfig.AutoDelete,
		queueConfig.Exclusive,
		queueConfig.NoWait,
		queueConfig.Args,
	)
	if err != nil {
		return err
	}
	ch.QueueName = q.Name

	return nil
}

func (ch *Channel) ExchangeDeclare(config config.ExchangeConfig) error {
	return ch.channel.ExchangeDeclare(
		config.Name,
		config.Type,
		config.Durable,
		config.AutoDelete,
		config.Internal,
		config.NoWait,
		config.Args,
	)
}

func (ch *Channel) BindQueue(exchangeName string) error {
	return ch.channel.QueueBind(
		ch.QueueName, // queue name
		ch.QueueName, // routing key
		exchangeName, // exchange
		false,
		nil)
}

func (ch *Channel) Publish(exchangeName string, body amqp.Publishing) error {
	return ch.channel.Publish(
		exchangeName,
		ch.QueueName,
		false, // mandatory
		false, // immediate
		body,
	)
}

package channel

import (
	"github.com/nixsolutions/go-rabbit-mq/config"
	"github.com/nixsolutions/go-rabbit-mq/connection"
	"github.com/streadway/amqp"
)

type ChannelInterface interface {
	GetChannel() RChannelInterface
	Create(connection.RConnectionInterface) error
	Close() error
	QueueDeclare(string) error
	QueueDeclareByConfig(config.QueueConfig) error
	ExchangeDeclare() error
	ExchangeDeclareByConfig(config.ExchangeConfig) error
	BindQueue() error
	BindQueueByConfig(bqc config.BindQueueConfig) error
	Publish(amqp.Publishing) error
	PublishWithParams(exchange, routingKey string, mandatory, immediate bool, body amqp.Publishing) error
	Consume() (<-chan amqp.Delivery, error)
	ConsumeByConfig(cc config.ConsumerConfig) (<-chan amqp.Delivery, error)
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

func (ch *Channel) QueueDeclare(name string) error {
	qc := config.NewQueueConfig(name)
	q, err := ch.channel.QueueDeclare(qc.GetConfig())
	if err != nil {
		return err
	}
	ch.QueueName = q.Name

	return nil
}

func (ch *Channel) QueueDeclareByConfig(qc config.QueueConfig) error {
	q, err := ch.channel.QueueDeclare(qc.GetConfig())
	if err != nil {
		return err
	}
	ch.QueueName = q.Name

	return nil
}

func (ch *Channel) ExchangeDeclare() error {
	ec := config.NewExchangeConfigByQueueName(ch.QueueName)
	return ch.channel.ExchangeDeclare(ec.GetConfig())
}

func (ch *Channel) ExchangeDeclareByConfig(ec config.ExchangeConfig) error {
	return ch.channel.ExchangeDeclare(ec.GetConfig())
}

func (ch *Channel) BindQueue() error {
	bqc := config.NewBindQueueConfigByQueueName(ch.QueueName)
	return ch.channel.QueueBind(bqc.GetConfig())
}

func (ch *Channel) BindQueueByConfig(bqc config.BindQueueConfig) error {
	return ch.channel.QueueBind(bqc.GetConfig())
}

func (ch *Channel) Publish(body amqp.Publishing) error {
	return ch.channel.Publish(
		ch.QueueName,
		ch.QueueName,
		false, // mandatory
		false, // immediate
		body,
	)
}

func (ch *Channel) PublishWithParams(exchange, routingKey string, mandatory, immediate bool, body amqp.Publishing) error {
	return ch.channel.Publish(
		exchange,
		routingKey,
		mandatory,
		immediate,
		body,
	)
}

func (ch *Channel) Consume() (<-chan amqp.Delivery, error) {
	cc := config.NewConsumerConfigByQueueName(ch.QueueName)
	return ch.channel.Consume(cc.GetConfig())
}

func (ch *Channel) ConsumeByConfig(cc config.ConsumerConfig) (<-chan amqp.Delivery, error) {
	return ch.channel.Consume(cc.GetConfig())
}

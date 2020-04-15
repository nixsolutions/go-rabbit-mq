package channel

import (
	"github.com/golang/mock/gomock"
	"github.com/nixsolutions/go-rabbit-mq/config"
	"github.com/nixsolutions/go-rabbit-mq/connection"
	"github.com/streadway/amqp"
	"testing"
)

func TestBindQueue(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockRChannelInterface(ctrl)
	ch := Channel{channel: m, QueueName: "Test"}
	m.EXPECT().QueueBind("Test", "Test", "Test", false, nil).Return(nil)
	err := ch.BindQueue()
	if err != nil {
		t.Error("Bind queue error: ", err)
	}
}

func TestQueueDeclare(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockRChannelInterface(ctrl)
	ch := Channel{channel: m, QueueName: "Test"}
	c := config.QueueConfig{
		Name:       "Test",
		Durable:    true,
		AutoDelete: false,
		Exclusive:  false,
		NoWait:     false,
		Args:       nil}
	m.EXPECT().QueueDeclare(c.Name, c.Durable, c.AutoDelete, c.Exclusive, c.NoWait, c.Args).Return(amqp.Queue{Name: c.Name}, nil)
	_ = ch.QueueDeclare("Test")
	if ch.QueueName != "Test" {
		t.Error("Declared queue name incorrect, expected: Test, Actual: ", ch.QueueName)
	}
}

func TestExchangeDeclare(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockRChannelInterface(ctrl)
	ch := Channel{channel: m, QueueName: "Test"}
	c := config.ExchangeConfig{
		Name:       "Test",
		Type:       "topic",
		Durable:    true,
		AutoDelete: false,
		Internal:   false,
		NoWait:     false,
		Args:       nil}
	m.EXPECT().ExchangeDeclare(c.Name, c.Type, c.Durable, c.AutoDelete, c.Internal, c.NoWait, c.Args).Return(nil)
	err := ch.ExchangeDeclare()
	if err != nil {
		t.Error("Exchange declaring error: ", err)
	}
}

func TestPublish(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockRChannelInterface(ctrl)
	ch := Channel{channel: m, QueueName: "Test"}
	p := amqp.Publishing{}
	m.EXPECT().Publish("Test", "Test", false, false, p).Return(nil)
	err := ch.Publish(p)
	if err != nil {
		t.Error("Publishing to queue error: ", err)
	}
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ch := Channel{channel: nil}
	mc := connection.NewMockRConnectionInterface(ctrl)
	mc.EXPECT().Channel().Return(&amqp.Channel{}, nil)
	_ = ch.Create(mc)
	if ch.channel == nil {
		t.Error("Created channel is nil")
	}
}

func TestClose(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockRChannelInterface(ctrl)
	ch := Channel{channel: m}
	m.EXPECT().Close().Return(nil)
	err := ch.Close()
	if err != nil {
		t.Error("Close channel error: ", err)
	}
}

func TestGetChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockRChannelInterface(ctrl)
	ch := Channel{channel: m}
	if m != ch.GetChannel() {
		t.Error("Get channel return incorrect channel")
	}
}

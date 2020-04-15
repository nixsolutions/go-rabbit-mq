package connection

import (
	"github.com/nixsolutions/go-rabbit-mq/config"
	"github.com/streadway/amqp"
)

type ConnectionInterface interface {
	ConnectByCredentials(config.Credentials) error
	ConnectByDialString(string) error
	Close() error
	GetConnector() RConnectionInterface
}

type RConnectionInterface interface {
	Channel() (*amqp.Channel, error)
	Close() error
}

type Connection struct {
	connection RConnectionInterface
}

func (con *Connection) ConnectByCredentials(credentials config.Credentials) error {
	return con.ConnectByDialString(credentials.GetAMQPDialString())
}

func (con *Connection) ConnectByDialString(url string) error {
	connection, err := amqp.Dial(url)
	con.connection = connection
	return err
}

func (con *Connection) Close() error {
	return con.connection.Close()
}

func (con *Connection) GetConnector() RConnectionInterface {
	return con.connection
}

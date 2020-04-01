package connection

import (
	"github.com/streadway/amqp"
	"rabbit-mq-go/config"
)

type ConnectionInterface interface {
	Connect(config.Credentials) error
	Close() error
}

type RConnectionInterface interface {
	Channel() (*amqp.Channel, error)
	Close() error
}

type Connection struct {
	connection RConnectionInterface
}

func (conn *Connection) Connect(credentials config.Credentials) error {
	connector, err := amqp.Dial("amqp://" +
		credentials.Username + ":" +
		credentials.Password + "@" +
		credentials.ClusterName + ":" +
		credentials.Port + "/")

	if err != nil {
		return err
	}
	conn.connection = connector

	return nil
}

func (conn *Connection) Close() error {
	return conn.connection.Close()
}

func (conn *Connection) GetConnector() RConnectionInterface {
	return  conn.connection
}

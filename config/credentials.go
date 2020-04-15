package config

import "os"

type Credentials struct {
	username    string
	password    string
	port        string
	clusterName string
}

func NewEnvCredentials() *Credentials {
	return &Credentials{
		username:    os.Getenv("RABBIT_USERNAME"),
		password:    os.Getenv("RABBIT_PASSWORD"),
		port:        os.Getenv("RABBIT_PORT"),
		clusterName: os.Getenv("RABBIT_CLUSTER_NAME"),
	}
}

func NewCredentials(username, password, port, clusterName string) *Credentials {
	return &Credentials{
		username:    username,
		password:    password,
		port:        port,
		clusterName: clusterName,
	}
}

func (c *Credentials) GetAMQPDialString() string {
	return "amqp://" +
		c.username + ":" +
		c.password + "@" +
		c.clusterName + ":" +
		c.port + "/"
}

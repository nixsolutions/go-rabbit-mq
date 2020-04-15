## Overview

The wrapper provides convenient access to the RabbitMQ from applications written in Golang language.

This module allows you to get rid of code duplication from services using RabbitMQ.

It includes predefined configs, you can use them with minimum code in your source code.

Although this module has structures for custom configs and the methods for using them, so it doesn't limit you.

### How to use

Init credentials:
```
    credentials := config.NewCredentials{
        os.Getenv("RABBIT_USERNAME"),
        os.Getenv("RABBIT_PASSWORD"),
        os.Getenv("RABBIT_PORT"),
        os.Getenv("RABBIT_CLUSTER_NAME"),
    }
    // or
    credentials := config.NewEnvCredentials()
```

Create connection and channel:

```
    con := connection.Connection{}
    ch := channel.Channel{}
    err := con.ConnectByCredentials(credentials)
    if err != nil {
        // handle error
    }
    err = ch.Create(con.GetConnector())
    if err != nil {
        // handle error
    }
```

Configure channel:

```
    _ = ch.GetChannel().Qos(1, 0, false)
    _ = ch.QueueDeclare("queueName")
    _ = ch.ExchangeDeclare()
    _ = ch.BindQueue()
```

Start consume:

```
    stream, err := ch.Consume()
    if err != nil {
        // handle error ...
    }

    for d := range stream {
        // handle messages ...
    }
```

Publish message:
```
    data := amqp.Publishing{
    	ContentType: "text/plain",
    	Body:        []byte("message"),
    }

    err := ch.Publish(data)
    // or with params
    err := ch.PublishWithParams(
        "exchange", 
        "routingKey", 
        false, // mandatory 
        false, // immediate 
        data)
    if err != nil {
        // handle error ...
    }
```

Close connection and channel:

```
    _ = ch.Close()
    _ = con.Close()
```

Create custom configs:

```
    qc := config.QueueConfig{...}
    ec := config.ExchangeConfig{...}
    bqc := config.BindQueueConfig{...}
    cc := config.ConsumerConfig{...}

```

Take methods below for using your configs:
```
    err := ch.QueueDeclareByConfig(qc)
    err := ch.ExchangeDeclareByConfig(ec)
    err := ch.BindQueueByConfig(bqc)
    err := ch.ConsumeByConfig(cc)
```

## License
The project is developed by [NIX Solutions](http://nixsolutions.com) Go team and distributed under [MIT LICENSE](https://raw.github.com/nixsolutions/go-rabbit-mq/master/LICENSE)

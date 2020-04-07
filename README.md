## Overview

Wrapper for RabbitMQ for easily using

### How use

Init credentials:
```
    credentials := config.Credentials{
        Username:    os.Getenv("RABBIT_USERNAME"),
        Password:    os.Getenv("RABBIT_PASSWORD"),
        Port:        os.Getenv("RABBIT_PORT"),
        ClusterName: os.Getenv("RABBIT_CLUSTER_NAME"),
    }
```

Create connection and channel:

```
    con := connection.Connection{}
    ch := channel.Channel{}
    err := connector.Connect(credentials)
    if err != nil {
        // handle error
    }
    err = ch.Create(con.GetConnector())
    if err != nil {
        // handle error
    }
```

Init configs:

```
    exchangeConfig := config.ExchangeConfig{
        Name:       "name",
        Type:       "topic",
        Durable:    true,
        AutoDelete: false,
        Internal:   false,
        NoWait:     false,
        Args:       nil,
    }
    queueConfig := config.QueueConfig{
        Name:       "name",
        Durable:    true,
        AutoDelete: false,
        Exclusive:  false,
        NoWait:     false,
        Args:       nil,
    }
```

Configure channel:

```
    _ = ch.GetChannel().Qos(1, 0, false)
    _ = ch.ExchangeDeclare(exchangeConfig)
    _ = ch.QueueDeclare(queueConfig)
    _ = ch.BindQueue(exchangeConfig.Name)
```

Start consume:

```
    stream, err := ch.GetChannel().Consume(
        ch.QueueName,
        "consumer_index,
        false,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        // handle error ...
    }

    for d := range stream {
        // handle messages ...
    }
```

For close connection and channel:

```
    _ = ch.Close()
    _ = connector.Close()
```

package main

import (
	"context"
	"fmt"

	"github.com/ResulCelik0/rabbitmq-golang-demo/pkg/rabbitmq"
)

func main() {
	fmt.Println("Hello from publisher!")
	mq := rabbitmq.NewRabbitMQ("amqp://guest:guest@localhost:5672/")
	err := mq.Connect()
	if err != nil {
		panic(err)
	}
	defer mq.Close()
	err = mq.CreateChannel("merhabaRabbitMQ")
	if err != nil {
		panic(err)
	}
	defer mq.CloseChannel("merhabaRabbitMQ")
	err = mq.PublishText(context.Background(), "merhabaRabbitMQ", "Merhaba RabbitMQ!")
	if err != nil {
		panic(err)
	}

}

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ResulCelik0/rabbitmq-golang-demo/pkg/rabbitmq"
)

func main() {
	fmt.Println("Hello from consumer!")
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
	mess, err := mq.Consume("merhabaRabbitMQ")
	if err != nil {
		panic(err)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	select {
	case m := <-mess:
		log.Println(string(m.Body))
	case <-c:
		fmt.Println("Exiting..")
		break
	}
}

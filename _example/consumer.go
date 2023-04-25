package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://admin:passowrd101@localhost:5672/")
	if err != nil {
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
	}

	msgs, err := ch.Consume("blog", "", true, false, false, false, amqp.Table{})

	if err != nil {
		panic(err)
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
		}
	}()
}

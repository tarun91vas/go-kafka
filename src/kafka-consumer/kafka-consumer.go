package main

import (
	"fmt"

	kafka "github.com/Shopify/sarama"
)

func main() {
	consumer, err := kafka.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Successful !!")

	messages, err := consumer.ConsumePartition("test-topic", 0, kafka.OffsetNewest)
	if err != nil {
		panic(err)
	}

	fmt.Println("Printing consumed messages....")
	for i := 0; i < 10; i++ {
		select {
		case message := <-messages.Messages():
			fmt.Println(message.Offset, string(message.Value))
		case err := <-messages.Errors():
			fmt.Println(err)
		}
	}

	consumer.Close()
}

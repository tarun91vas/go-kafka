package main

import (
	"fmt"
	"strconv"

	kafka "github.com/Shopify/sarama"
)

func main() {
	config := kafka.NewConfig()
	config.Producer.Flush.Messages = 10
	config.Producer.Return.Successes = true
	producer, err := kafka.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection Successful !!")

	for i := 0; i < 10; i++ {
		producer.Input() <- &kafka.ProducerMessage{
			Topic: "test-topic",
			Key:   nil,
			Value: kafka.StringEncoder("Event: " + strconv.Itoa(i)),
		}
	}

	for i := 0; i < 10; i++ {
		select {
		case msg := <-producer.Errors():
			fmt.Println(msg.Err)
		case success := <-producer.Successes():
			fmt.Println(success.Topic, success.Value)
		}
	}

	producer.Close()
}

package main

import (
	"fmt"
	"log"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/joho/godotenv"
	kafka2 "github.com/lucascmreis/codelivery/gps-simulator/app/kafka"
	"github.com/lucascmreis/codelivery/gps-simulator/infra/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("oie", "readtest", producer)

	// for{
	// 	_=1
	// }
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	//put it in another tread - Go routing - async
	go consumer.Consume()

	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go kafka2.Produce(msg)
	}
}
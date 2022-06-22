package main

import (
	"fmt"
	routeModule "github.com/lucascmreis/codelivery/gps-simulator/app/route"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("error loading .env file")
// 	}
// }

func main() {
	// msgChan := make(chan *ckafka.Message)
	// consumer := kafka.NewKafkaConsumer(msgChan)
	// go consumer.Consume()
	// for msg := range msgChan {
	// 	fmt.Println(string(msg.Value))
	// 	go kafka2.Produce(msg)
	// }

	route := routeModule.Route{
		ID: "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
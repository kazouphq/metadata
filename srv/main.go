package main

import (
	"log"

	"github.com/kazouphq/metadata/srv/handler"
	"github.com/kazouphq/metadata/srv/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"

	example "github.com/kazouphq/metadata/srv/proto/example"
)

var (
	topic = "topic.go.micro.srv.metadata"
)

func main() {
	// Init broker
	if err := broker.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	// Connect broker
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker Connert error: %v", err)
	}
	broker.Subscribe(
		topic,
		subscriber.Handle,
		broker.Queue(topic),
	)
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.metadata"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// // Register Struct as Subscriber
	// service.Server().Subscribe(
	// 	service.Server().NewSubscriber("topic.go.micro.srv.metadata", new(subscriber.Example)),
	// )
	//
	// // Register Function as Subscriber
	// service.Server().Subscribe(
	// 	service.Server().NewSubscriber("topic.go.micro.srv.metadata", subscriber.Handler),
	// )

	// Initialise service
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

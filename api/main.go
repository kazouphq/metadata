package main

import (
	"log"

	"github.com/kazouphq/metadata/api/client"
	"github.com/kazouphq/metadata/api/handler"
	"github.com/micro/go-micro"

	example "github.com/kazouphq/metadata/api/proto/example"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.metadata"),
		micro.Version("latest"),
	)

	// Register Handler
	example.RegisterMetadataHandler(service.Server(), new(handler.Example))

	// Initialise service
	service.Init(
		// create wrap for the Metadata srv client
		micro.WrapHandler(client.ExampleWrapper(service)),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

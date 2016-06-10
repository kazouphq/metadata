package handler

import (
	"encoding/json"
	"log"

	"github.com/kazouphq/metadata/api/client"
	example "github.com/kazouphq/metadata/srv/proto/example"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"

	"golang.org/x/net/context"
)

type Example struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// Example.Send is called by the API as /metadata/example/send with post body {"fileName": "foo","filePath":"bar"}
func (e *Example) Send(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Example.Call request")

	// extract the client from the context
	exampleClient, ok := client.ExampleFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.metadata.example.call", "example client not found")
	}

	// make request
	response, err := exampleClient.Send(ctx, &example.Request{
		FileName: extractValue(req.Post["fileName"]),
		FilePath: extractValue(req.Post["filePath"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.metadata.example.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}

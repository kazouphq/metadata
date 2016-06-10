package handler

import (
	"log"

	example "github.com/kazouphq/metadata/srv/proto/example"
	"github.com/kazouphq/metadata/srv/publisher"
	"golang.org/x/net/context"
)

type Example struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Example) Send(ctx context.Context, req *example.Request, rsp *example.Response) error {
	log.Printf("Received Send.Call request %v", req)
	var header = make(map[string]string)

	header["topic"] = "topic.go.micro.srv.metadata"
	header["id"] = "1"
	msg := &example.Message{
		FileName: req.FileName,
		FilePath: req.FilePath,
	}

	err := publisher.PublishData(header, []byte(msg.String()))

	if err != nil {
		// TODO: com.kazoup.srv.publish
		return err
	}
	rsp.Msg = "Thanks from service !"
	return nil

}

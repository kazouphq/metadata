package subscriber

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/micro/go-micro/broker"
)

// type Example struct{}
//
// func (e *Example) Handle(ctx context.Context, msg *example.Message) error {
// 	log.Print("Handler Received message: ", msg.FileName)
// 	return nil
// }
//
// func Handler(ctx context.Context, msg *example.Message) error {
// 	log.Print("Function Received message: ", msg.FileName)
// 	return nil
// }

func Handle(p broker.Publication) error {
	// if err := elastic.Create(p.Message().Header["id"], p.Message().Body); err != nil {
	// 	return err
	// }
	spew.Dump(p.Message())
	return nil
}

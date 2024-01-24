package main

import (
	"context"
	"fmt"
	"github.com/kubemq-io/kubemq-go"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	commandsClients, err := kubemq.NewCommandsClient(ctx,
		kubemq.WithAddress("localhost", 50000))

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = commandsClients.Close()
	}()
	commandsSub := &kubemq.CommandsSubscription{
		Channel:  "c1",
		Group:    "",
		ClientId: "",
	}

	err = commandsClients.Subscribe(ctx, commandsSub, func(msg *kubemq.CommandReceive, err error) {
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Receiver - Commands Received:\nCommandID: %s\nBody: %s\n", msg.Id, msg.Body)
			response := kubemq.NewResponse().
				SetRequestId(msg.Id).
				SetExecutedAt(time.Now()).
				SetResponseTo(msg.ResponseTo)
			if err := commandsClients.Response(ctx, response); err != nil {
				log.Fatal(err)
			}
		}
	})
	time.Sleep(100 * time.Millisecond)
	resp, err := commandsClients.Send(ctx, kubemq.NewCommand().
		SetChannel("c1").
		SetTimeout(time.Second).
		SetBody([]byte("hello kubemq - sending command")))
	if err != nil {
		log.Fatal(fmt.Sprintf("error sending command: %s", err.Error()))
	}
	log.Printf("Sender - Response Received:\nCommandID: %s\nIsExecuted: %t\n", resp.CommandId, resp.Executed)
	time.Sleep(time.Second)

}

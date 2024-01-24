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
	queriesClients, err := kubemq.NewQueriesClient(ctx,
		kubemq.WithAddress("localhost", 50000))

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = queriesClients.Close()
	}()
	queriesSub := &kubemq.QueriesSubscription{
		Channel:  "q1",
		Group:    "",
		ClientId: "",
	}

	err = queriesClients.Subscribe(ctx, queriesSub, func(msg *kubemq.QueryReceive, err error) {
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Receiver - Query Received:\nQueryID: %s\nBody: %s\n", msg.Id, msg.Body)
			response := kubemq.NewResponse().
				SetRequestId(msg.Id).
				SetExecutedAt(time.Now()).
				SetResponseTo(msg.ResponseTo).
				SetBody([]byte("some response"))
			if err := queriesClients.Response(ctx, response); err != nil {
				log.Fatal(err)
			}
		}
	})
	time.Sleep(100 * time.Millisecond)
	resp, err := queriesClients.Send(ctx, kubemq.NewQuery().
		SetChannel("q1").
		SetTimeout(time.Second).
		SetBody([]byte("hello kubemq - sending query")))
	if err != nil {
		log.Fatal(fmt.Sprintf("error sending query: %s", err.Error()))
	}
	log.Printf("Sender - Response Received:\nQueryID: %s\nIsExecuted: %t\nBody:%s\n", resp.QueryId, resp.Executed, resp.Body)
	time.Sleep(time.Second)

}

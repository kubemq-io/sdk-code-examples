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
	eventsClient, err := kubemq.NewEventsClient(ctx,
		kubemq.WithAddress("localhost", 50000))

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = eventsClient.Close()
	}()
	eventsSub := &kubemq.EventsSubscription{
		Channel:  "e1",
		Group:    "",
		ClientId: "",
	}

	err = eventsClient.Subscribe(ctx, eventsSub, func(msg *kubemq.Event, err error) {
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Receiver - Event Received:\nEventID: %s\nBody: %s\n", msg.Id, msg.Body)
		}
	})
	time.Sleep(100 * time.Millisecond)
	err = eventsClient.Send(ctx, kubemq.NewEvent().
		SetChannel("e1").
		SetBody([]byte("hello kubemq - sending event")))
	if err != nil {
		log.Fatal(fmt.Sprintf("error sending event: %s", err.Error()))
	}
	time.Sleep(time.Second)

}

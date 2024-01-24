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
	eventsStoreClient, err := kubemq.NewEventsStoreClient(ctx,
		kubemq.WithAddress("localhost", 50000))

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = eventsStoreClient.Close()
	}()
	eventsStoreSub := &kubemq.EventsStoreSubscription{
		Channel:          "es1",
		Group:            "",
		ClientId:         fmt.Sprintf("%d", time.Now().Nanosecond()),
		SubscriptionType: kubemq.StartFromFirstEvent(),
	}

	err = eventsStoreClient.Subscribe(ctx, eventsStoreSub, func(msg *kubemq.EventStoreReceive, err error) {
		if err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Receiver - Event Received:\nEventID: %s\nBody: %s\n", msg.Id, msg.Body)
		}
	})
	time.Sleep(100 * time.Millisecond)
	_, err = eventsStoreClient.Send(ctx, kubemq.NewEventStore().
		SetChannel("es1").
		SetClientId("sender").
		SetBody([]byte("hello kubemq - sending event store")))
	if err != nil {
		log.Fatal(fmt.Sprintf("error sending event store: %s", err.Error()))
	}
	time.Sleep(time.Second)

}

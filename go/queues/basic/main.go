package main

import (
	"context"
	client "github.com/kubemq-io/kubemq-go/queues_stream"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	senderClient, err := client.NewQueuesStreamClient(ctx,
		client.WithAddress("localhost", 50000))
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = senderClient.Close()
	}()

	receiverClient, err := client.NewQueuesStreamClient(ctx,
		client.WithAddress("localhost", 50000))

	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = receiverClient.Close()
	}()

	msg := client.NewQueueMessage().
		SetChannel("q1").
		SetBody([]byte("hello kubemq - sending single message"))

	_, err = senderClient.Send(ctx, msg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("message sent")
	pollRequest := client.NewPollRequest().
		SetChannel("q1").
		SetMaxItems(1).
		SetWaitTimeout(1).
		SetAutoAck(true)

	receiveResult, err := receiverClient.Poll(ctx, pollRequest)
	if err != nil {
		log.Fatal(err)
	}
	for _, msg := range receiveResult.Messages {
		log.Printf("MessageID: %s, Body: %s", msg.MessageID, string(msg.Body))
	}
}

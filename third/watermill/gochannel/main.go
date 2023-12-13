package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main() {
	pubsub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	message, err := pubsub.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(message)

	publishMessages(pubsub)
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))
		msg.Ack()
	}
}

func publishMessages(publisher message.Publisher) {
	for {
		msg :=	message.NewMessage(watermill.NewShortUUID(), []byte("hello, world!"))

		if err := publisher.Publish("example.topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second*3)
	}
}
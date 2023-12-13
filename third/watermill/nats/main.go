package main

import (
	"context"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-nats/pkg/nats"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/nats-io/stan.go"
)

func main() {
	subsriber, err := nats.NewStreamingSubscriber(
		nats.StreamingSubscriberConfig {
			ClusterID: "test-cluster",
			ClientID: "example-subscriber",
			QueueGroup: "example",
			DurableName: "my-durable",
			SubscribersCount: 4,
			CloseTimeout: time.Minute,
			AckWaitTimeout: time.Second * 30,
			StanOptions: []stan.Option{
				stan.NatsURL("localhost://nats-stream:4222"),
			},
			Unmarshaler: nats.GobMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	messages, err := subsriber.Subscribe(context.Background(), "example.topic")
	if err != nil {
		panic(err)
	}

	go process(messages)

	publisher, err := nats.NewStreamingPublisher(
		nats.StreamingPublisherConfig{
			ClusterID: "test-cluster",
			ClientID: "example-publisher",
			StanOptions: []stan.Option {
				stan.NatsURL("localhost://nats-stream:4222"),
			},
			Marshaler: nats.GobMarshaler{},
		},
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}
	publishMessages(publisher)
}

func process(messages <-chan *message.Message) {
	for msg := range messages {
		log.Printf("received message: %s, payload: %s", msg.UUID, string(msg.Payload))

		// we need to Acknowledge that we received and processed the message,
		// otherwise, it will be resent over and over again.
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
package main

import (
	"context"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

func produceMessages(routerClosed chan struct{}, publisher message.Publisher) {
	for {
		select {
		case <- routerClosed:
			return
		default:
			//go on
		}

		time.Sleep(time.Millisecond*50 + time.Duration(random.Intn(50))*time.Millisecond) 
		msg := message.NewMessage(watermill.NewUUID(), []byte{})
		_= publisher.Publish("sub_topic", msg)
	}
}

func consumeMessages(subscriber message.Subscriber) {
	messages, err := subscriber.Subscribe(context.Background(), "pub_topic")
	if err != nil {
		panic(err)
	}

	for msg := range messages {
		msg.Ack()
	}
}
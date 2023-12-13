package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

var (
	logger = watermill.NewStdLogger(false, false)
)

func main() {

	pubsub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	go publishMessages(pubsub)


	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		panic(err)
	}

	router.AddPlugin(plugin.SignalsHandler)

	router.AddMiddleware(
		middleware.CorrelationID,
		middleware.Retry {
			MaxRetries: 3,
			InitialInterval: time.Millisecond * 100,
			Logger: logger,
		}.Middleware,
		middleware.Recoverer,
	)

	router.AddNoPublisherHandler(
		"print_incoming_message",
		"incoming_messages_topic",
		pubsub,
		printMessages,
	)

	router.AddNoPublisherHandler(
		"print_outgoing_messages",
		"outgoing_messages_topic",
		pubsub,
		printMessages,
	)

	handler := router.AddHandler(
		"struct_handler",
		"incoming_messages_topic",
		pubsub,
		"outgoing_messages_topic",
		pubsub,
		structHandler{}.Handler,
	)

	handler.AddMiddleware(func(h message.HandlerFunc) message.HandlerFunc {
		return func(msg *message.Message) ([]*message.Message, error) {
			log.Println("executing handler specific middleware for ", msg.UUID)
			return h(msg)
		}
	})

	if err := router.Run(context.Background()); err != nil {
		panic(err)
	}
}

func publishMessages(publisher message.Publisher) {
	time.Sleep(time.Second * 5)
	for {
		msg :=	message.NewMessage(watermill.NewShortUUID(), []byte("hello, world!"))
		middleware.SetCorrelationID(watermill.NewUUID(), msg)

		log.Printf("sending message %s, correlation id: %s\n", msg.UUID, middleware.MessageCorrelationID(msg))

		if err := publisher.Publish("incoming_messages_topic", msg); err != nil {
			panic(err)
		}

		time.Sleep(time.Second*100)
	}
}

func printMessages(msg *message.Message) error {
	fmt.Printf(
		"\n> Received message: %s\n> %s\n> metadata: %v\n\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}



package main

import (
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
)

type structHandler struct {}

func (s structHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	log.Println("structHandler received message", msg.UUID)

	message.NewMessage(watermill.NewShortUUID(), []byte("message produced by structHandler"))
	return message.Messages{msg}, nil
}
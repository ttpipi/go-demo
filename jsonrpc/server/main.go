package main

import (
	"demo/jsonrpc/services"
	"log"
)

func main() {
	log.Fatal(services.StartRpcService(services.DivService{}, ":1234"))
}

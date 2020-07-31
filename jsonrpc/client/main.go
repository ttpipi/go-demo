package main

import (
	"demo/jsonrpc/services"
	"fmt"
	"log"
)

func main() {
	client, err := services.NewRpcClient(":1234")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	var ret services.DivRespon
	err = client.Call("DivService.Div", services.DivRequest{A: 10, B: 2}, &ret)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(ret.Ret)
}

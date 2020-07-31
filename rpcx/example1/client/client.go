package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	"golang/rpcx-example"
	"log"
)

var (
	addr = flag.String("addr", "localhost:12345", "server address")
)

func main() {
	flag.Parse()
	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	args := rpcx_example.Args{
		A: 10,
		B: 20,
	}

	reply := &rpcx_example.Reply{}

	//同步调用
	//err := xclient.Call(context.Background(), "Mul", args, reply)
	//if err != nil {
	//	log.Fatalf("Fail call Arith: %v", err)
	//}

	//异步调用
	call, err := xclient.Go(context.Background(), "Mul", args, reply, nil)
	if err != nil {
		log.Fatalf("Fail call Arith: %v", err)
	}
	replyCall := <- call.Done
	if replyCall.Error != nil {
		log.Fatalf("Fail cal Arith: %v", replyCall.Error)
	}


	log.Printf("%d*%d=%d", args.A, args.B, reply.C)
}


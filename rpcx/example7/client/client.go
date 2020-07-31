package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	rpcx_example "golang/rpcx-example"
	"log"
	"math/rand"
	"time"
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

	go func(){
		for {
			reply := &rpcx_example.Reply{}
			err := xclient.Call(context.Background(), "Mul", args, reply)
			if err != nil {
				log.Fatalf("failed to call: %v", err)
			}

			time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
		}
	}()

	go func(){
		for {
			reply := &rpcx_example.Reply{}
			err := xclient.Call(context.Background(), "Add", args, reply)
			if err != nil {
				log.Fatalf("failed to call: %v", err)
			}

			time.Sleep(time.Duration(rand.Intn(100))*time.Millisecond)
		}
	}()


}

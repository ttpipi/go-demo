package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/client"
	rpcx_example "golang/rpcx-example"
	"log"
	"time"
)

var (
	addr1 = flag.String("addr1", "localhost:12345", "server1 address")
	addr2 = flag.String("addr2", "localhost:12346", "server2 addresss")
)

func main() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1}, {Key: *addr2}})
	xclient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, client.DefaultOption)
	defer xclient.Close()

	args := rpcx_example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &rpcx_example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Fatalf("Faile call Arith: %v", err)
		}

		log.Printf("%d*%d=%d", args.A, args.B, reply.C)
		time.Sleep(1e9)
	}
}

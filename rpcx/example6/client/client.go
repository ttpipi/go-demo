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
	addr2 = flag.String("addr2", "localhost:12346", "server2 address")
)


func main() {
	flag.Parse()

	d := client.NewMultipleServersDiscovery([]*client.KVPair{{Key: *addr1, Value: "state=inactive"}, {Key: *addr2, Value: "group=test"}})
	option := client.DefaultOption
	option.Group = "test"
	xclient := client.NewXClient("Arith", client.Failtry, client.RoundRobin, d, option)
	defer xclient.Close()
	
	args := rpcx_example.Args{
		A: 10,
		B: 20,
	}

	for {
		reply := &rpcx_example.Reply{}
		err := xclient.Call(context.Background(), "Mul", args, reply)
		if err != nil {
			log.Printf("failed to call: %v", err)
		}

		log.Printf("%d*%d=%d", args.A, args.B, reply.C)
		time.Sleep(time.Second*2)
	}
}

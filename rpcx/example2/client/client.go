package main

import (
	"context"
	"crypto/tls"
	"flag"
	"github.com/smallnest/rpcx/client"
	rpcx_example "golang/rpcx-example"
	"log"
)

var (
	addr = flag.String("addr", "localhost:12345", "server address")
)

func main() {
	flag.Parse()

	d := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	option := client.DefaultOption
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}
	option.TLSConfig = conf

	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, option)
	defer xclient.Close()

	args := rpcx_example.Args{
		A: 10,
		B: 20,
	}

	reply := &rpcx_example.Reply{}

	err := xclient.Call(context.Background(), "Mul", args, reply)
	if err != nil {
		log.Fatalf("Failed call Arith")
	}

	log.Printf("%d*%d=%d", args.A, args.B, reply.C)
}

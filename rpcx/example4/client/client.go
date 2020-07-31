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
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd address")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	d := client.NewEtcdDiscovery(*basePath, "Arith", []string{*etcdAddr}, nil)
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
			log.Printf("Faile Call: %v\n", err)
			time.Sleep(time.Second *5)
			continue
		}

		log.Printf("%d*%d=%d", args.A, args.B, reply.C)
		time.Sleep(time.Second * 5)
	}
}

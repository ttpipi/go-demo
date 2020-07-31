package main

import (
	"context"
	"flag"
	"github.com/smallnest/rpcx/server"
	rpcx_example "golang/rpcx-example"
)

var (
	addr1 = flag.String("addr1", "localhost:12345", "server1 address")
	addr2 = flag.String("addr2", "localhost:12346", "server2 address")
)

type Arith int
func (t *Arith) Mul(ctx context.Context, args *rpcx_example.Args, reply *rpcx_example.Reply) error {
	reply.C = args.A * args.B * 100
	return nil
}

func main() {
	flag.Parse()

	go createServer1(*addr1, "state=inactive")
	go createServer2(*addr2, "group=test")
	select {}
}

func createServer1(addr string, meta string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(rpcx_example.Arith), meta)
	s.Serve("tcp", addr)
}

func createServer2(addr string, meta string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), meta)
	s.Serve("tcp", addr)
}


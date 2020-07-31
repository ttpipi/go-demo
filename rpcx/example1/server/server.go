package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	"golang/rpcx-example"
)

var (
	addr = flag.String("addr", "localhost:12345", "server address")
)

func main() {
	flag.Parse()
	s := server.NewServer()
	s.RegisterName("Arith", new(rpcx_example.Arith), "")
	s.Serve("tcp", *addr)
}

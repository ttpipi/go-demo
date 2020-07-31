package main

import (
	"flag"
	"github.com/smallnest/rpcx/server"
	rpcx_example "golang/rpcx-example"
)

var (
	addr1 = flag.String("addr1", "localhost:12345", "server1 address")
	addr2 = flag.String("addr2", "localhost:12346", "server2 address")
)

func main() {
	flag.Parse()

	go createService(*addr1)
	go createService(*addr2)

	select {

	}
}

func createService(addr string) {
	s := server.NewServer()
	s.RegisterName("Arith", new(rpcx_example.Arith), "")
	s.Serve("tcp", addr)
}
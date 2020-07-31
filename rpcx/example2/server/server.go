package main

import (
	"crypto/tls"
	"flag"
	"github.com/smallnest/rpcx/server"
	rpcx_example "golang/rpcx-example"
	"log"
)

var (
	addr = flag.String("addr", "localhost:12345", "server address")
)

func main() {
	flag.Parse()

	cert, err := tls.LoadX509KeyPair("./server.pem", "./server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}

	s := server.NewServer(server.WithTLSConfig(config))
	s.RegisterName("Arith", new(rpcx_example.Arith), "")
	s.Serve("tcp", *addr)
}

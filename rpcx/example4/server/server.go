package main

import (
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	rpcx_example "golang/rpcx-example"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:12345", "server address")
	etcdAddr = flag.String("etcdAddr", "localhost:2379", "etcd addresss")
	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	flag.Parse()

	s := server.NewServer()
	addRegistryPlugin(s)

	s.RegisterName("Arith", new(rpcx_example.Arith), "")
	go s.Serve("tcp", *addr)

	time.Sleep(time.Minute)

	err := s.UnregisterAll()
	if err != nil {
		panic(err)
	}
}

func addRegistryPlugin(s *server.Server) {
	r := &serverplugin.EtcdRegisterPlugin{
		ServiceAddress: "tcp@"+*addr,
		EtcdServers: []string{*etcdAddr},
		BasePath: *basePath,
		Metrics: metrics.NewRegistry(),
		UpdateInterval: time.Minute,
	}
	err := r.Start()
	if err != nil {
		log.Fatal(err)
	}

	s.Plugins.Add(r)
}
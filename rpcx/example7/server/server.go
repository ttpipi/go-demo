package main

import (
	"flag"
	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/serverplugin"
	graphite "github.com/cyberdelia/go-metrics-graphite"
	rpcx_example "golang/rpcx-example"
	"net"
	"time"
)

var (
	addr = flag.String("addr", "localhost:12345", "server address")
)

func main() {
	flag.Parse()

	s := server.NewServer()

	p := serverplugin.NewMetricsPlugin(metrics.DefaultRegistry)
	s.Plugins.Add(p)
	startMetrics()

	s.Register(new(rpcx_example.Arith), "")
	s.Serve("tcp", *addr)
}

func startMetrics() {
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)
	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, time.Second)

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.2003")
	go graphite.Graphite(metrics.DefaultRegistry, 1e9, "rpcx.services.host.127_0_0_1", addr)
}
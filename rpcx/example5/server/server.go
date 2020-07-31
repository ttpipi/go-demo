package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
	rpcx_example "golang/rpcx-example"
)

var (
	addr = flag.String("addr", "localhost:12345", "server address")
)

type Arith int

func (t *Arith) Mul(ctx context.Context, args *rpcx_example.Args, reply *rpcx_example.Reply) error {
	reqMeta := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	resMeta := ctx.Value(share.ResMetaDataKey).(map[string]string)
	fmt.Printf("Recevied meta: %v\n", reqMeta)
	resMeta["echo"] = "from serer"

	reply.C = args.A * args.B * 100
	return nil
}

func main() {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", *addr)
}

package main

import (
	"grpc/helper"
	"grpc/services"
	"net"

	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
)

func main() {

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	services.RegisterOrderServiceServer(rpcServer, new(services.OrderService))
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))

	lser, err := net.Listen("tcp", ":9090")
	if err != nil {
		panic(err)
	}
	defer lser.Close()

	log.Fatal(rpcServer.Serve(lser))
}

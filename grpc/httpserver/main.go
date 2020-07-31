package main

import (
	"context"
	"demo/grpc/helper"
	"demo/grpc/services"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	gwMux := runtime.NewServeMux()
	grpcServer := "localhost:9090"
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwMux, grpcServer, opt)
	if err != nil {
		log.Fatal(err)
	}

	err = services.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwMux, grpcServer, opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwMux,
	}
	log.Fatal(httpServer.ListenAndServe())
}

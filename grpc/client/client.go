package main

import (
	"context"
	"demo/grpc/helper"
	. "demo/grpc/services"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	ctx := context.Background()
	// client := NewProdServiceClient(conn)
	// result, err := client.GetProdInfo(ctx, &ProdRequest{Id: 0, Area: ProdAreas_B})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)

	// result, err := client.GetProdStocks(ctx, &services.QuerySize{Size: 10})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result.Stocks)

	// client := NewOrderServiceClient(conn)
	// t := timestamp.Timestamp{Seconds: time.Now().Unix()}
	// result, err := client.NewOrder(ctx, &OrderRequest{
	// 	Order: &OrderMain{
	// 		Id:     1001,
	// 		Number: "1111",
	// 		Money:  100,
	// 		Time:   &t,
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(result)

	client := NewUserServiceClient(conn)
	stream, err := client.GetUserScoreByTWS(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for j := 1; j <= 3; j++ {
		req := UserScoreRequest{}
		req.Users = make([]*UserInfo, 0)
		for i := 1; i <= 5; i++ {
			req.Users = append(req.Users, &UserInfo{Id: int32(i + 5*(j-1))})
		}
		err := stream.Send(&req)
		if err != nil {
			log.Println(err)
		}

		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		fmt.Println(res.Users)
	}

}

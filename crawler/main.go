package main

import (
	"crawler/comm"
	"crawler/engine"
	"crawler/myrpc"
	"crawler/persist"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/rpc"
	"strings"
)

var (
	seedUrl = "http://localhost:8080/mock/www.zhenai.com/zhenghun"
	//FIXME: 应该做服务发现
	workerServiceHosts = flag.String("whosts", "", "WorkerService hosts")
	//FIXME: 应该做服务发现
	saverServiceHost = flag.Int("shost", 0, "SaverService host")
)

func CreateWorkerServiceCientPool(hosts []string) (chan *rpc.Client, error) {
	var clients []*rpc.Client
	for _, host := range hosts {
		client, err := myrpc.NewClient(fmt.Sprintf(":%s", host))
		if err != nil {
			log.Printf("连接WorkerService失败, host=%s, err=%v", host, err)
			continue
		}
		clients = append(clients, client)
	}

	if len(clients) == 0 {
		return nil, errors.New("没有WorkerService可用")
	}

	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				out <- client
			}
		}
	}()
	return out, nil
}

func main() {
	flag.Parse()
	if *saverServiceHost == 0 {
		log.Println("must specify SaverService host")
		return
	}

	saver, err := persist.CreateSaver(fmt.Sprintf(":%d", *saverServiceHost))
	if err != nil {
		panic(err)
	}

	pool, err := CreateWorkerServiceCientPool(strings.Split(*workerServiceHosts, ","))
	if err != nil {
		panic(err)
	}

	e := engine.NewEngine(saver, pool)

	e.Run(comm.Request{
		Url:        seedUrl,
		ParserName: "cityList",
	})
}

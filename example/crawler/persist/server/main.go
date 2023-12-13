package main

import (
	"demo/example/crawler/config"
	"demo/example/crawler/myrpc"
	"flag"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

var (
	host = "http://ttlin.com:9200"
	port = flag.Int("port", 0, "SaverService port")
)

func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must specity SaverService port")
	}

	client, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	log.Fatal(myrpc.StartService(
		&myrpc.SaverService{
			Client: client,
			Index:  config.ElasticIndex,
		},
		fmt.Sprintf(":%d", *port)))
}

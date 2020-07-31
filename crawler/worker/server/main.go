package main

import (
	"crawler/myrpc"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "WorkerService port")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Println("must specify the port")
		return
	}
	log.Fatal(myrpc.StartService(myrpc.WorkerService{}, fmt.Sprintf(":%d", *port)))
}

package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello"))
	if err != nil {
		log.Println(err)
		return
	}

	var data [0xFFF]byte
	n, err := conn.Read(data[:])
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(data[:n]))
}

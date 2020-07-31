package main

import (
	"io"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	var data [0xFFF]byte
	for {
		n, err := conn.Read(data[:])

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		_, err = conn.Write(data[:n])
		if err != nil {
			log.Println(err)
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}
}

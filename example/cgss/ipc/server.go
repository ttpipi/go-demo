package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Methon string `json:"methon"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handle(methon, params string) *Response
}

type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("invalid request format:", request)
			}
			resp := server.Handle(req.Methon, req.Params)
			b, err := json.Marshal(resp)

			c <- string(b)
		}
		fmt.Println("session closed.")
	}(session)
	fmt.Println("a new session has been create successfully.")
	return session
}

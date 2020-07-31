package ipc

import "encoding/json"

type Client struct {
	conn chan string
}

func NewClient(server *IpcServer) *Client {
	c := server.Connect()
	return &Client{c}
}

func (client *Client) Call(methon, params string) (resp *Response, err error) {
	req := &Request{methon, params}
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}

	client.conn <- string(b)
	str := <-client.conn

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1
	return
}

func (client *Client) Close() {
	client.conn <- "CLOSE"
}

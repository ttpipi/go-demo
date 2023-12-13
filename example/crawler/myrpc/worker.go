package myrpc

import (
	"demo/example/crawler/comm"
	"demo/example/crawler/parser"
	"net/http"
	"net/http/httputil"
)

type WorkerService struct{}

func (WorkerService) Work(r comm.Request, respon *comm.Result) error {
	html, err := fetch(r.Url)
	if err != nil {
		return err
	}

	f := parser.ParserMap[r.ParserName]
	*respon = f(html, r.Url)
	return nil
}

func fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return httputil.DumpResponse(resp, true)
}

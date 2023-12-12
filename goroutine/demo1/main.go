package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"
	"time"
)

/**
1.只要pprof, app有一个退出了, 就全退出
2.退出前优雅处理
3.优雅退出, 处理未完成的事, 如日志打印, 数据库写入, 数据上报等
*/

type Reporter struct {
	worker   int
	messages chan string
	wg       sync.WaitGroup
	closed   bool
}

func NewReporter(worker, buffer int) *Reporter {
	return &Reporter{worker: worker, messages: make(chan string, buffer)}
}

func (r *Reporter) run(stop <-chan struct{}) {
	go func() {
		<-stop
		r.shutdown()
	}()

	for i := 0; i < r.worker; i++ {
		r.wg.Add(1)
		go func() {
			for msg := range r.messages {
				time.Sleep(10 * time.Second)
				fmt.Printf("report: %s\n", msg)
			}
			r.wg.Done()
		}()
	}
	r.wg.Wait()
}

func (r *Reporter) shutdown() {
	r.closed = true
	close(r.messages)
}

func (r *Reporter) report(data string) {
	if r.closed {
		return
	}
	r.messages <- data
}

func main() {
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})
	slog.SetDefault(slog.New(h))

	done := make(chan error, 3)
	stop := make(chan struct{})
	go func() {
		done <- pprof(stop)
	}()
	go func() {
		done <- app(stop)
	}()

	report := NewReporter(2, 1)
	go func() {
		report.run(stop)
		done <- fmt.Errorf("report close")
	}()

	report.report("开始")
	report.report("开始")
	report.report("开始")
	report.report("开始")
	report.report("开始")
	report.report("开始")
	var stoped bool
	for i := 0; i < cap(done); i++ {
		fmt.Println("into")
		err := <-done
		slog.Info("server exit err", "err", err)
		if !stoped {
			stoped = true
			close(stop)
		}
	}
}

func server(handler http.Handler, addr string, stop <-chan struct{}) error {
	s := http.Server{
		Handler: handler,
		Addr:    addr,
	}

	go func() {
		<-stop
		slog.Info("server will exiting", "addr", addr)
		s.Shutdown(context.Background())
	}()

	return s.ListenAndServe()
}

func pprof(stop <-chan struct{}) error {
	go func() {
		server(http.DefaultServeMux, ":8081", stop)
	}()
	time.Sleep(3 * time.Second)
	return fmt.Errorf("mock pprof exit")
}

func app(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	return server(mux, ":8080", stop)
}

package worker

import (
	"crawler/comm"
	"crawler/scheduler"
	"log"
	"net/rpc"
)

func CreateWorker(out chan comm.Result, s scheduler.Scheduler, pool chan *rpc.Client) error {
	go func() {
		in := make(chan comm.Request)
		for {
			s.ReceiveWorker(in)
			r := <-in

			var result comm.Result
			client := <-pool
			//FIXME: 这里应该是异步调用, 不然并分量大的时候会卡住等待远程服务返回
			err := client.Call("WorkerService.Work", r, &result)
			if err != nil {
				log.Println(err)
				continue
			}

			out <- result
		}
	}()
	return nil
}

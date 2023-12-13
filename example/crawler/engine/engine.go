package engine

import (
	"demo/example/crawler/comm"
	"demo/example/crawler/config"
	"demo/example/crawler/scheduler"
	"demo/example/crawler/worker"
	"net/rpc"
)

type Engine struct {
	workerCount             int
	scheduler               scheduler.Scheduler
	saver                   chan comm.Profile
	WorkerServiceClientPool chan *rpc.Client
}

func NewEngine(saver chan comm.Profile, pool chan *rpc.Client) *Engine {
	return &Engine{
		workerCount:             config.WorkerCount,
		saver:                   saver,
		WorkerServiceClientPool: pool,
	}
}

func (e *Engine) Run(seedReqs ...comm.Request) {
	out := make(chan comm.Result)
	e.scheduler.Schedule()

	for i := 0; i < e.workerCount; i++ {
		err := worker.CreateWorker(out, e.scheduler, e.WorkerServiceClientPool)
		if err != nil {
			panic(err)
		}
	}

	for _, r := range seedReqs {
		e.scheduler.ReceiveRequest(r)
	}

	for {
		result := <-out

		if result.Profile.Data != nil {
			go func() {
				e.saver <- result.Profile
			}()
		}

		for _, r := range result.Requests {
			e.scheduler.ReceiveRequest(r)
		}
	}

}

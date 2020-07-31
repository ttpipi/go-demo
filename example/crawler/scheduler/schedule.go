package scheduler

import (
	"crawler/comm"
)

type Scheduler struct {
	reqChan    chan comm.Request
	workerChan chan chan comm.Request
}

func (s *Scheduler) ReceiveRequest(r comm.Request) {
	s.reqChan <- r
}

func (s *Scheduler) ReceiveWorker(w chan comm.Request) {
	s.workerChan <- w
}

func (s *Scheduler) Schedule() {
	s.reqChan = make(chan comm.Request)
	s.workerChan = make(chan chan comm.Request)
	go func() {
		var requestQ []comm.Request
		var workerQ []chan comm.Request
		for {
			var activeRequest comm.Request
			var activeWorker chan comm.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-s.reqChan:
				requestQ = append(requestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

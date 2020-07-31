package bloomfilter

type BloomFilter struct {
}

func isDuplicate(url string) bool {
	if redis.Get(url) == 1 {
		return true
	}
}

func (e *ConcurrentEngine) createWorker() {
	go func() {
		for {
			request := <-in

			//去重
			if isDuplicate(request.Url) {
				continue
			}

			result, err := e.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

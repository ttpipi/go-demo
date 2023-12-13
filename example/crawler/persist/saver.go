package persist

import (
	"demo/example/crawler/comm"
	"demo/example/crawler/myrpc"
	"fmt"
	"log"
)

func CreateSaver(host string) (chan comm.Profile, error) {
	client, err := myrpc.NewClient(host)
	if err != nil {
		return nil, err
	}

	saver := make(chan comm.Profile)
	go func() {
		for {
			item := <-saver

			var result string
			//FIXME: 这里应该是异步调用
			err = client.Call("SaverService.Save", item, &result)
			if err != nil {
				log.Println(err)
				continue
			}

			fmt.Println(item)
		}
	}()
	return saver, nil
}

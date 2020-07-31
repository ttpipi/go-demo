package singleton

import (
	"sync"
)

type Singleton struct{}

var (
	once     sync.Once
	instance *Singleton
)

func GetInstance() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
	})
	return instance
}

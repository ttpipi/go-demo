package main

import "fmt"

/*
	1.不用维护一个全局的track映射, 代价是每次添加元素都要实例化一个新的映射并且每个元素都要拷贝一下
	2.用container包装一下, 使用比较简单
*/

type defInfo struct {
	key   string
	scope int
}

type tracked map[string]int

func (s tracked) add(info defInfo) tracked {
	newList := make(tracked, len(s))
	for k, v := range s {
		newList[k] = v
	}
	newList[info.key] = info.scope
	return newList
}

type container struct {
	t tracked
}

func (c *container) builder(info defInfo) *container {
	return &container{
		t: c.t.add(info),
	}
}

func main() {
	info := []defInfo{
		{key: "app", scope: 0},
		{key: "app2", scope: 100},
	}

	c := &container{}

	n := c.
		builder(info[0]).
		builder(info[1])

	fmt.Println(n)
}

package zhenai

import (
	"demo/example/crawler/comm"
	"regexp"
)

var (
	cityListRe = regexp.MustCompile(`href="(.*www\.zhenai\.com/zhenghun/[0-9a-z]+)"`)
)

func ParseCityList(text []byte, _ string) comm.Result {
	matches := cityListRe.FindAllSubmatch(text, -1)
	var result comm.Result
	for _, m := range matches {
		result.Requests = append(result.Requests, comm.Request{
			Url:        string(m[1]),
			ParserName: "cityUser",
		})
	}
	return result
}

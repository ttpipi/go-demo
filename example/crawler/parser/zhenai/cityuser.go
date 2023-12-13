package zhenai

import (
	"demo/example/crawler/comm"
	"regexp"
)

var (
	cityUserRe         = regexp.MustCompile(`<a href="(.*album.zhenai\.com/u/[0-9]+)">([^<]+)</a>`)
	nextPageCityUserRe = regexp.MustCompile(`<a href="(.*www.zhenai\.com/zhenghun/[0-9a-z]+/[0-9]+)">`)
)

func ParseCityUser(html []byte, _ string) comm.Result {
	var result comm.Result

	matches := cityUserRe.FindAllSubmatch(html, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, comm.Request{
			Url:        string(m[1]),
			ParserName: "user",
		})
	}

	matches = nextPageCityUserRe.FindAllSubmatch(html, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, comm.Request{
			Url:        string(m[1]),
			ParserName: "cityUser",
		})
	}
	return result
}

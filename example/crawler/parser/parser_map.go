package parser

import (
	"demo/example/crawler/comm"
	"demo/example/crawler/parser/zhenai"
)

var ParserMap = map[string]func([]byte, string) comm.Result{
	"cityList": zhenai.ParseCityList,
	"cityUser": zhenai.ParseCityUser,
	"user":     zhenai.ParseUser,
}

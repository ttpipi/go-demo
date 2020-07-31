package parser

import (
	"crawler/comm"
	"crawler/parser/zhenai"
)

var ParserMap = map[string]func([]byte, string) comm.Result{
	"cityList": zhenai.ParseCityList,
	"cityUser": zhenai.ParseCityUser,
	"user":     zhenai.ParseUser,
}

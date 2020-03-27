package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:           "http://www.zhenai.com/zhenghun",
		ParseFunction: parser.ParseCityList,
	})

}

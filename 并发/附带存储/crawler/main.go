package main

import (
	"Project/crawler/engine"
	"Project/crawler/scheduler"
	"Project/crawler/storage"
	"Project/crawler/zhenai/parser"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:           "http://www.zhenai.com/zhenghun",
	//	ParseFunction: parser.ParseCityList,
	//})

	//concurrentEngine := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkCount: 100,
	//}

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
		ItemChan:  storage.ItemSaver(),
	}
	//concurrentEngine.Run(engine.Request{
	//	Url:           "http://www.zhenai.com/zhenghun/shanghai",
	//	ParseFunction: parser.ParseCity,
	//})
	concurrentEngine.Run(engine.Request{
		Url:           "http://www.zhenai.com/zhenghun",
		ParseFunction: parser.ParseCityList,
	})

}

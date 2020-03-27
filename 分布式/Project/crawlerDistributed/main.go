package main

import (
	"../crawler/engine"
	"../crawler/scheduler"
	"../crawler/zhenai/parser"
	"../crawlerDistributed/config"
	"../crawlerDistributed/rpcSupport"
	itemSaverClient "../crawlerDistributed/storage/client"
	workerClient "../crawlerDistributed/worker/client"
	"flag"
	"log"
	"net/rpc"
	"strings"
)

var (
	itemSaverHost = flag.String(
		"item_saver_host", "", "item saver host")

	workerHosts = flag.String(
		"worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := itemSaverClient.ItemSaver(*itemSaverHost)
	if err != nil {
		panic(err)
	}

	pool := createClientPool(strings.Split(*workerHosts, ","))
	processor := workerClient.CreateProcessor(pool)

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkCount:        100,
		ItemChan:         itemChan,
		RequestProcessor: processor,
	}
	concurrentEngine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client

	for _, host := range hosts {
		client, err := rpcSupport.NewClient(host)
		if err == nil {
			clients = append(clients, client)
			log.Printf("Connected to %s", host)
		} else {
			log.Printf("Error connecting to %s: %v", host, err)
		}
	}

	output := make(chan *rpc.Client)
	go func() {
		for {
			for _, client := range clients {
				output <- client
			}
		}
	}()
	return output
}

package client

import (
	"../../../crawler/engine"
	"../../../crawlerDistributed/config"
	"../../../crawlerDistributed/worker"
	"net/rpc"
)

func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(request engine.Request) (result engine.ParseResult, err error) {
		serializeRequest := worker.SerializeRequest(request)

		var serializeResult worker.ParseResult
		client := <-clientChan
		err = client.Call(config.CrawlServiceRpc, serializeRequest, &serializeResult)

		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(serializeResult), nil
	}
}

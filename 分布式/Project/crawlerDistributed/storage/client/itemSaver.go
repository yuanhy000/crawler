package client

import (
	"Project/crawler/engine"
	"Project/crawlerDistributed/config"
	"Project/crawlerDistributed/rpcSupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcSupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	outputChan := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-outputChan
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			itemCount++

			// call RPC to save item
			result := ""
			err = client.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Item Saver: error"+"saving item %v :%v", item, err)
			}
		}
	}()
	return outputChan, nil
}

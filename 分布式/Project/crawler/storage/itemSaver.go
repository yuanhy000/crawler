package storage

import (
	"Project/crawler/engine"
	"log"
)

func ItemSaver() chan engine.Item {
	outputChan := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-outputChan
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			itemCount++

			_ = Save(item)
		}
	}()
	return outputChan
}

func Save(item engine.Item) error {
	// save in database
	// save in elasticSearch
	log.Printf("save item success")
	return nil
}

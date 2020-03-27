package storage

import "log"

func ItemSaver() chan interface{} {
	outputChan := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-outputChan
			log.Printf("Item Saver: got item "+"#%d: %v", itemCount, item)
			itemCount++

			save(item)
		}
	}()
	return outputChan
}

func save(item interface{}) {
	// save in database
	// save in elasticSearch
}

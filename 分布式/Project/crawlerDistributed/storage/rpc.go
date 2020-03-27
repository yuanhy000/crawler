package storage

import (
	"../../crawler/engine"
	"../../crawler/storage"
	"log"
)

type ItemSaverService struct {
}

func (saveService *ItemSaverService) Save(item engine.Item, result *string) error {
	// save in database
	// save in elasticSearch
	err := storage.Save(item)
	log.Printf("Item %v saved", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}

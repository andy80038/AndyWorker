package persist

import (
	"log"

	"github.com/andy80038/AndyWorker/crawler/engine"
	"github.com/andy80038/AndyWorker/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
)

type ItemSaveService struct {
	Clinet *elastic.Client
	Index  string
}

func (s *ItemSaveService) Save(item engine.Item, result *string) error {

	err := persist.Save(s.Clinet, s.Index, item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v:%v", item, err)
	}
	log.Printf("Item %v saved.", item)
	return nil
}

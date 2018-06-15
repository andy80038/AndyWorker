package persist

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/andy80038/AndyWorker/crawler/engine"
	"gopkg.in/olivere/elastic.v5"
)

const indexString = "dating_profile"

func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //Must turn off sniff in docker
	if err != nil {
		return nil, err
	}
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, indexString, item)
			if err != nil {
				log.Printf("Item Saver:error  "+"saving item %v:%v", item, err)
			}
		}
	}()
	return out, nil

}
func Save(client *elastic.Client, index string, item engine.Item) (err error) {

	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)

	return nil
}

package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/andy80038/AndyWorker/crawler_distributed/config"
	"github.com/andy80038/AndyWorker/crawler_distributed/persist"
	"github.com/andy80038/AndyWorker/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
)

var port = flag.Int("port", 0, "the port for me to listen")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a  port")
	}
	log.Fatal(serveRpc(
		fmt.Sprintf(":%d", *port), config.ElasticIndex))

}
func serveRpc(host, index string) error {
	clinet, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return err
	}
	return rpcsupport.ServeRpc(host, &persist.ItemSaveService{
		Clinet: clinet,
		Index:  index,
	})

}

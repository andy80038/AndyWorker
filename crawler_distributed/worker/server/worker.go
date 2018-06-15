package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/andy80038/AndyWorker/crawler_distributed/rpcsupport"
	"github.com/andy80038/AndyWorker/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a  port")
	}

	log.Fatal(
		rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}

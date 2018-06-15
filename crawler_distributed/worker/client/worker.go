package client

import (
	"net/rpc"

	"github.com/andy80038/AndyWorker/crawler_distributed/worker"

	"github.com/andy80038/AndyWorker/crawler/engine"
	"github.com/andy80038/AndyWorker/crawler_distributed/config"
)

func CreateProcessor(clientsChan chan *rpc.Client) engine.Processor {

	return func(req engine.Request) (engine.ParseResult, error) {
		sReq := worker.SerializeRequest(req)
		var sResuit worker.ParseResult
		c := <-clientsChan
		err := c.Call(config.CrawlServiceRpc, sReq, &sResuit)
		if err != nil {
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(sResuit), nil
	}
}

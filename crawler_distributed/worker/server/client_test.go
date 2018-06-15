package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/andy80038/AndyWorker/crawler_distributed/config"
	"github.com/andy80038/AndyWorker/crawler_distributed/rpcsupport"
	"github.com/andy80038/AndyWorker/crawler_distributed/worker"
)

func TestCrawlService(t *testing.T) {
	const host = ":9000"
	go rpcsupport.ServeRpc(host, worker.CrawlService{})
	time.Sleep(time.Second * 1)

	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://album.zhenai.com/u/108072009",
		Parser: worker.SerializedParser{
			Name: config.ParseProfile,
			Args: "心灵的声音",
		},
	}
	var result worker.ParseResult
	err = client.Call(config.CrawlServiceRpc, req, &result)
	if err != nil {
		t.Errorf("ERRR")
	} else {
		fmt.Println(result)
	}
}

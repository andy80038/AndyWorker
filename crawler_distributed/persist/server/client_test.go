package main

import (
	"testing"
	"time"

	"github.com/andy80038/AndyWorker/crawler/engine"
	"github.com/andy80038/AndyWorker/crawler/model"
	"github.com/andy80038/AndyWorker/crawler_distributed/config"
	"github.com/andy80038/AndyWorker/crawler_distributed/rpcsupport"
)

func TestItemSaver(t *testing.T) {
	const host = ":123"
	go serveRpc(host, "test1")
	time.Sleep(5 * time.Second)
	item := engine.Item{
		Url:  "http://123123",
		Type: "zfdadadadadadada",
		Id:   "108165161",
		Payload: model.Profile{
			Age: 34,
		},
	}
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	result := ""
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result :%s; err:%s", result, err)
	}

}

package main

import (
	"github.com/andy80038/AndyWorker/crawler/engine"
	"github.com/andy80038/AndyWorker/crawler/persist"
	"github.com/andy80038/AndyWorker/crawler/scheduler"
	"github.com/andy80038/AndyWorker/crawler/zhenai/parser"
)

/*
使用http.Get獲取內容
使用Encoding來轉換編碼 gbk>utf8
使用charset.DetermineEncoding來判斷編碼

*/

/*
獲取城市名稱連結
使用css選擇棄
使用xpath
使用正則
*/
/*
解析器 Parser
輸入 utf-8 編碼文本
輸出request{URL,對應Parser}列表,Item列表
*/

func main() {
	itemChan, err := persist.ItemSaver()
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(
			parser.ParseCityList, "ParseCityList"),
	})

	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun",
	// 	ParserFunc: parser.ParseCityList,
	// })

}

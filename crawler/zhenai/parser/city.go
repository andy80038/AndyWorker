package parser

import (
	"regexp"

	"github.com/andy80038/AndyWorker/crawler/engine"
	"github.com/andy80038/AndyWorker/crawler_distributed/config"
)

var profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^<]*">([^<]+)</a></`)
var cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)">下一页</a>`)

func ParseCity(content []byte, _ string) engine.ParseResult {
	//re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}

	for _, m := range matches {
		name := string(m[2])
		//result.Items = append(result.Items, "User "+name)
		url := string(m[1])
		result.Requests = append(result.Requests, engine.Request{
			Url:    url,
			Parser: NewProfileParser(name),
		})
		//fmt.Printf("URL:%s   City:%s\n", m[1], m[2])

	}
	matches = cityUrlRe.FindAllSubmatch(content, -1)
	for _, m := range matches {
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			Parser: engine.NewFuncParser(
				ParseCity, config.ParseCity),
		})

	}

	return result
}

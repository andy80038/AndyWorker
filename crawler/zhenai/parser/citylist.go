package parser

import (
	"regexp"

	"github.com/andy80038/AndyWorker/crawler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-zA-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(content []byte, url string) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		//result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})
		//fmt.Printf("URL:%s   City:%s\n", m[1], m[2])

	}
	return result
}

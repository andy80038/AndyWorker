package parser

import (
	"log"
	"regexp"
	"strconv"

	"github.com/andy80038/AndyWorker/crawler/engine"
	"github.com/andy80038/AndyWorker/crawler_distributed/config"

	"github.com/andy80038/AndyWorker/crawler/model"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([0-9]+)岁</td>`)
var marriageRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([\d]+)CM</td>`)
var weightRe = regexp.MustCompile(`<span class="label">体重：</span><span field="">([0-9]+)KG</span>`)
var incomeRe = regexp.MustCompile(` <td><span class="label">月收入：</span>([^<]+)</td>`)
var educationRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var carRe = regexp.MustCompile(`<span class="label">是否购车：</span><span field="">([^<]+)</span>`)

//var occupationRe = regexp.MustCompile(`<td><span class="label">工作地：</span>四川阿坝</td>`)
//var a = regexp.MustCompile(` <td><span class="label">有无孩子：</span>有，我们住在一起</td>`)
var hokouRe = regexp.MustCompile(`<td><span class="label">籍贯：</span>([^<])</td>`)
var idUrlRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

func parseProfile(contents []byte, name string, url string) engine.ParseResult {
	profile := model.Profile{}

	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		log.Printf("name:%s  age is null", name)
	}
	profile.Age = age
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		log.Printf("name:%s  height is null", name)
	}
	weight, err := strconv.Atoi(extractString(contents, weightRe))
	if err != nil {
		log.Printf("name:%s  weight is null", name)
	}
	profile.Height = height
	profile.Weight = weight
	profile.Marriage = extractString(contents, marriageRe)
	profile.Hokou = extractString(contents, hokouRe)
	profile.Education = extractString(contents, educationRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Occupation = extractString(contents, occupationRe)
	profile.Car = extractString(contents, carRe)
	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Url:     url,
				Type:    "zhenai",
				Id:      extractString([]byte(url), idUrlRe),
				Payload: profile,
			},
		},
	}
	return result

}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}

}

type ProfileParser struct {
	userName string
}

func (p *ProfileParser) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.userName)
}

func (p *ProfileParser) Serialize() (name string, args interface{}) {
	return config.ParseProfile, p.userName
}

func NewProfileParser(name string) *ProfileParser {
	return &ProfileParser{
		userName: name,
	}
}

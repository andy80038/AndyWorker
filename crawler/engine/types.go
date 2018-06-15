package engine

type ParserFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url    string
	Parser Parser
}
type ParseResult struct {
	Requests []Request
	Items    []Item
}
type Parser interface {
	Parse(contents []byte, url string) ParseResult
	Serialize() (name string, args interface{})
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}
type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) Serialize() (name string, args interface{}) {
	return "NilParser", nil
}

type FuncParser struct {
	parser ParserFunc
	Name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.Name, nil
}

func NewFuncParser(p ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		Name:   name,
	}
}

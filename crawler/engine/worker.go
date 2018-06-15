package engine

import (
	"log"

	"github.com/andy80038/AndyWorker/crawler/fetcher"
)

func Worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher :error fetcging url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	log.Printf("Fetching : %s", r.Url)
	return r.Parser.Parse(body, r.Url), nil
}

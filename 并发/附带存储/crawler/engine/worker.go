package engine

import (
	"Project/crawler/fetcher"
	"log"
)

func worker(request Request) (ParseResult, error) {
	log.Printf("Fetching %s", request.Url)
	// get the website html
	body, err := fetcher.Fetch(request.Url)
	if err != nil {
		log.Printf("Fetcher error: fetching url %s: %v", request.Url, err)
		return ParseResult{}, err
	}

	// parse the html, regexp to get the target info, {ParseResult struct}
	return request.ParseFunction(body), nil
}

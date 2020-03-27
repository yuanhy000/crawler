package engine

import (
	"Project/crawler/fetcher"
	"log"
)

type SimpleEngine struct {
}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	// add initial seed into queue
	for _, request := range seeds {
		requests = append(requests, request)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		parseResult, err := worker(request)
		if err != nil {
			continue
		}

		// add seeds into queue
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

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

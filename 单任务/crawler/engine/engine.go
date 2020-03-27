package engine

import (
	"awesomeProject/crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	// add initial seed into queue
	for _, request := range seeds {
		requests = append(requests, request)
	}

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", request.Url)
		// get the website html
		body, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Printf("Fetcher error: fetching url %s: %v", request.Url, err)
			continue
		}

		// parse the html, regexp to get the target info, {ParseResult struct}
		parseResult := request.ParseFunction(body)
		// add seeds into queue
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

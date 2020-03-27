package parser

import (
	"Project/crawler/engine"
	"regexp"
)

const cityListRegexp = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// parse the html, get target useful info, {City name, Http request}
func ParseCityList(contents []byte) engine.ParseResult {
	result := regexp.MustCompile(cityListRegexp)
	matches := result.FindAllSubmatch(contents, -1)

	// init the target struct
	parseResult := engine.ParseResult{}
	for _, match := range matches {
		//parseResult.Items = append(parseResult.Items, "City: "+string(match[2])) // City name
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:           string(match[1]),
			ParseFunction: ParseCity,
		})

	}
	// return many Result as slice
	return parseResult
}

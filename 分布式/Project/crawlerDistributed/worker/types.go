package worker

import (
	"../../crawler/engine"
	"../../crawler/zhenai/parser"
	"../../crawlerDistributed/config"
	"errors"
	"fmt"
	"log"
)

type SerializedParser struct {
	Name string
	Args interface{}
}

type Request struct {
	Url    string
	Parser SerializedParser
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeRequest(request engine.Request) Request {
	name, args := request.Parser.Serialize()
	return Request{
		Url: request.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func SerializeResult(parseResult engine.ParseResult) ParseResult {
	result := ParseResult{
		Items: parseResult.Items,
	}
	for _, request := range parseResult.Requests {
		result.Requests = append(result.Requests, SerializeRequest(request))
	}
	return result
}

func DeserializeRequest(request Request) (engine.Request, error) {
	p, err := DeserializeParser(request.Parser)
	if err != nil {
		return engine.Request{}, err
	}
	return engine.Request{
		Url:    request.Url,
		Parser: p,
	}, nil
}

func DeserializeResult(parseResult ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: parseResult.Items,
	}
	for _, request := range parseResult.Requests {
		engineRequest, err := DeserializeRequest(request)
		if err != nil {
			log.Printf("error deserializing request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineRequest)
	}
	return result
}

func DeserializeParser(p SerializedParser) (engine.Parser, error) {
	//log.Printf("%s", p.Name)
	switch p.Name {
	case config.ParseCity:
		return engine.NewFuncParser(parser.ParseCity, config.ParseCity), nil
	case config.ParseCityList:
		return engine.NewFuncParser(parser.ParseCityList, config.ParseCityList), nil
	case config.NilParser:
		return engine.NilParser{}, nil
	case config.ParseProfile:
		if userGender, ok := p.Args.(string); ok {
			return parser.NewProfileParser(userGender), nil
		} else {
			return nil, fmt.Errorf("invaild arg: %v", p.Args)
		}
	default:
		return nil, errors.New("unknown parser name")
	}
}

package worker

import "../../crawler/engine"

type CrawlService struct {
}

func (CrawlService) Process(request Request, result *ParseResult) error {
	engineRequest, err := DeserializeRequest(request)
	if err != nil {
		return err
	}

	engineResult, err := engine.Worker(engineRequest)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}

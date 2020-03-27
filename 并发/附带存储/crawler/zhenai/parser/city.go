package parser

import (
	"Project/crawler/engine"
	"regexp"
)

//const cityRegexp = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`
//const genderRegexp = `<span class="grayL">性别：</span>([^<]+)</td>`

var (
	cityRegexp = regexp.MustCompile(
		`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRegexp = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
	genderRegexp = regexp.MustCompile(
		`<span class="grayL">性别：</span>([^<]+)</td>`)
)

func ParseCity(contents []byte, ) engine.ParseResult {
	cityMatches := cityRegexp.FindAllSubmatch(contents, -1)

	genderList := ParseGender(contents)
	// init the target struct
	parseResult := engine.ParseResult{}

	for index, match := range cityMatches {
		userName := string(match[2])
		userGender := genderList[index]
		//parseResult.Items = append(parseResult.Items, "User: "+userName)
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url: string(match[1]),
			// Closure function, pass the name
			ParseFunction: func(contents []byte) engine.ParseResult {
				return ParseProfile(contents, userName, userGender)
			},
		})

	}

	cityUrlMatches := cityUrlRegexp.FindAllSubmatch(contents, -1)
	for _, match := range cityUrlMatches {
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			Url:           string(match[1]),
			ParseFunction: ParseCity,
		})
	}

	// return many Result as slice
	return parseResult
}

func ParseGender(contents []byte) []string {
	genderMatches := genderRegexp.FindAllSubmatch(contents, -1)

	var genderList []string
	for _, match := range genderMatches {
		genderList = append(genderList, string(match[1]))
	}
	return genderList
}

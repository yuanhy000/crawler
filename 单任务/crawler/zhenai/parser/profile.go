package parser

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/model"
	"fmt"
	"regexp"
	"strings"
)

var profileRegexp = regexp.MustCompile(`<div class="des f-cl" [^>]*>([^<]+)</div>`)

func ParseProfile(contents []byte, name string, gender string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name
	profile.Gender = gender

	match := profileRegexp.FindSubmatch(contents)
	targetProfile := strings.Split(string(match[1]), " | ")

	profile.Address = targetProfile[0]
	profile.Age = targetProfile[1]
	profile.Education = targetProfile[2]
	profile.Marriage = targetProfile[3]
	profile.Height = targetProfile[4]
	profile.Income = targetProfile[5]

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, reg *regexp.Regexp) string {
	matches := reg.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		fmt.Printf("%s ", match[1])
	}
	if len(matches) >= 2 {
		//return string(match[1])
		return ""
	} else {
		return ""
	}
}

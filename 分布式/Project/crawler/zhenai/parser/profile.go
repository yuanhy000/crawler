package parser

import (
	"Project/crawler/engine"
	"Project/crawler/model"
	"regexp"
	"strings"
)

var nameRegexp = regexp.MustCompile(`<h1 class="nickName" [^>]*>([^<]+)</h1>`)
var profileRegexp = regexp.MustCompile(`<div class="des f-cl" [^>]*>([^<]+)</div>`)

func parseProfile(contents []byte, url, gender string) engine.ParseResult {
	profile := model.Profile{}
	name := nameRegexp.FindSubmatch(contents)
	profile.Name = string(name[1])
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
		Items: []engine.Item{
			{
				Url:     url,
				Payload: profile,
			},
		},
	}
	return result
}

type ProfileParse struct {
	gender string
}

func (p *ProfileParse) Parse(contents []byte, url string) engine.ParseResult {
	return parseProfile(contents, url, p.gender)
}

func (p *ProfileParse) Serialize() (name string, args interface{}, ) {
	return "ProfileParse", p.gender
}

func NewProfileParser(gender string) *ProfileParse {
	return &ProfileParse{
		gender: gender,
	}
}

package util

import (
	"regexp"
)

func Match(source, regex string) []string {
	regexMatch := regexp.MustCompile(regex)
	params := regexMatch.FindStringSubmatch(source)
	return params
}

package shorten

import (
	"regexp"
)

func AddHTTP(url string) string {
	matched := match("https://", url)
	if matched {
		return url
	} else {
		return "https://" + url
	}
}

func match(pattern string, text string) bool {
	matched, _ := regexp.Match(pattern, []byte(text))
	if matched {
		return true
	} else {
		return false
	}
}

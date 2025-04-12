package config

import (
	"regexp"
	"testing"
)

func TestAllowedOrigins(t *testing.T) {
	allowedOrigins := App.AllowOrigins

	var regexPatterns []*regexp.Regexp

	// 预编译正则表达式
	for _, origin := range allowedOrigins {
		re := regexp.MustCompile(origin)
		regexPatterns = append(regexPatterns, re)
	}

	origins := []string{
		"http://127.0.0.1",
		"https://127.0.0.1",
		"http://127.0.0.1:80",
		"http://127.0.0.1:80",
		"http://127.0.0.1:8080",
		"http://127.0.0.1:8080",

		"http://localhost",
		"https://localhost",
		"http://localhost:80",
		"https://localhost:80",
		"http://localhost:8080",
		"https://localhost:8080",

		"http://oupro.cn",
		"https://oupro.cn",
		"http://oupro.cn:80",
		"https://oupro.cn:80",
		"http://oupro.cn:8080",
		"https://oupro.cn:8080",

		"http://rapid.oupro.cn",
		"https://rapid.oupro.cn",
		"http://rapid.oupro.cn:80",
		"https://rapid.oupro.cn:80",
		"http://rapid.oupro.cn:8080",
		"https://rapid.oupro.cn:8080",
	}

	for _, origin := range origins {
		isAllowed := false

		for _, re := range regexPatterns {
			if re.MatchString(origin) {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			t.Errorf("%s is not allowed", origin)
		}
	}
}

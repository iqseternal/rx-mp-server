package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
	"rx-mp/config"
	"rx-mp/internal/biz"
	"rx-mp/internal/pkg/rx"
	"strings"
)

// 判断是否为正则表达式（按需实现）
func isRegexPattern(s string) bool {
	return strings.Contains(s, "*") || strings.Contains(s, `\.`) || strings.Contains(s, "^")
}

// DomainWhitelistMiddleware 校验域名的白名单, 拒绝不符合要求的请求
func DomainWhitelistMiddleware() gin.HandlerFunc {
	allowedOrigins := config.App.AllowOrigins

	var regexPatterns []*regexp.Regexp

	// 预编译正则表达式
	for _, origin := range allowedOrigins {
		if isRegexPattern(origin) {
			re := regexp.MustCompile(origin)
			regexPatterns = append(regexPatterns, re)
		}
	}

	return rx.WrapHandler(func(c *rx.Context) {
		origin := c.Request.Header.Get("Origin")

		log.Printf("来源 %s", origin)

		if origin == "" {
			c.Next() // 允许无 Origin 的请求
			return
		}

		// 检查精确匹配或正则匹配
		isAllowed := false
		for _, allowed := range allowedOrigins {
			if allowed == origin {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			for _, re := range regexPatterns {
				if re.MatchString(origin) {
					isAllowed = true
					break
				}
			}
		}

		if !isAllowed {
			c.AbortFinish(http.StatusForbidden, &rx.R{
				Code:    biz.UnknownOrigin,
				Message: biz.Message(biz.UnknownOrigin),
				Data:    nil,
			})
			return
		}

		c.Next()
	})
}

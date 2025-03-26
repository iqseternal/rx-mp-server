package common

import (
	"fmt"
	"strings"
)

// ParseBearerAuthorizationToken 解析 Bearer Authorization 请求头中的 Bearer 格式 token
func ParseBearerAuthorizationToken(bearerAuthorization string) (string, error) {

	if bearerAuthorization == "" {
		return "", fmt.Errorf("bearerAuthorization is empty")
	}

	if !strings.HasPrefix(bearerAuthorization, "Bearer ") {
		return "", fmt.Errorf("is not a valid bearerAuthorization")
	}

	bearerAuthorization = bearerAuthorization[7:]

	if bearerAuthorization == "" {
		return "", fmt.Errorf("bearer token is empty")
	}

	return bearerAuthorization, nil
}

package common

import (
	"testing"
)

func TestParseBearerAuthorizationToken(t *testing.T) {
	bearerAuthorization := "Bearer A"

	t.Run("检测是否能够正确处理 Bearer Auth 头", func(t *testing.T) {
		_, err := ParseBearerAuthorizationToken(bearerAuthorization)
		if err != nil {
			t.Errorf("错误的处理")
		}

		_, err = ParseBearerAuthorizationToken("Bearer ")
		if err == nil {
			t.Errorf("错误的处理")
		}

		_, err = ParseBearerAuthorizationToken("")
		if err == nil {
			t.Errorf("错误的处理")
		}
	})
}

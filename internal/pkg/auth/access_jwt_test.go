package auth

import "testing"

func TestVerifyAccessToken(t *testing.T) {
	token, err := GenerateAccessToken("1")
	if err != nil {
		t.Error(err)
		return
	}

	claims, err := VerifyAccessToken(token)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("claims:", claims)
}

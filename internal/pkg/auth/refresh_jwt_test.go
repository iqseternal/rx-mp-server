package auth

import "testing"

func TestVerifyRefershToken(t *testing.T) {
	token, err := GenerateRefreshToken("1")
	if err != nil {
		t.Error(err)
		return
	}

	claims, err := VerifyRefreshToken(token)

	if err != nil {
		t.Error(err)
		return
	}

	t.Log("claims:", claims)
}

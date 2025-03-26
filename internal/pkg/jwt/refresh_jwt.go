package jwt

import (
	"fmt"
	"rx-mp/internal/constants"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RefreshJwtClaims struct {
	UserId string `json:"user_id"`
	*jwt.RegisteredClaims
}

func GenerateRefershToken(user_id string) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodES256, RefreshJwtClaims{
		UserId: user_id,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.RefreshJwtExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    constants.RefreshJwtIssuer,
			ID:        user_id,
			Audience:  []string{},
		},
	})

	secret, err := ParseECDSAPemToPrivateKey(constants.RefreshJwtPrivateSecret)
	if err != nil {
		return "", err
	}

	token, err := tokenStruct.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyRefershToken(tokenString string) (*RefreshJwtClaims, error) {
	tokenObj, err := jwt.ParseWithClaims(tokenString, &RefreshJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret, err := ParseECDSAPemToPublicKey(constants.RefreshJwtPublicSecret)
		if err != nil {
			return nil, err
		}

		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenObj.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := tokenObj.Claims.(*RefreshJwtClaims)

	if !ok {
		return nil, fmt.Errorf("invalid claims type")
	}

	if claims.Issuer != constants.RefreshJwtIssuer {
		return nil, fmt.Errorf("invalid issuer")
	}

	return claims, nil
}

package jwt

import (
	"fmt"
	"rx-mp/internal/constants"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AccessJwtClaims struct {
	UserId string `json:"user_id"`
	*jwt.RegisteredClaims
}

func GenerateAccessToken(user_id string) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodES256, AccessJwtClaims{
		UserId: user_id,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.AccessJwtExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    constants.AccessJwtIssuer,
			ID:        user_id,
			Audience:  []string{},
		},
	})

	secret, err := ParseECDSAPemToPrivateKey(constants.AccessJwtSecret)
	if err != nil {
		return "", err
	}

	token, err := tokenStruct.SignedString(secret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyAccessToken(tokenString string) (*AccessJwtClaims, error) {
	tokenObj, err := jwt.ParseWithClaims(tokenString, &RefreshJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret, err := ParseECDSAPemToPrivateKey(constants.AccessJwtSecret)
		if err != nil {
			return "", err
		}

		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokenObj.Claims.(*AccessJwtClaims)

	if !ok || !tokenObj.Valid {
		fmt.Printf("%v %v\n", claims.UserId, claims.RegisteredClaims)
		return claims, nil
	}

	if claims.Issuer != constants.AccessJwtIssuer {
		return nil, fmt.Errorf("invalid issuer")
	}

	return nil, err
}

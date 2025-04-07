package auth

import (
	"fmt"
	"rx-mp/internal/constants"
	"time"

	pkgjwt "rx-mp/internal/pkg/jwt"

	"github.com/golang-jwt/jwt/v5"
)

type AccessJwtClaims struct {
	UserId string `json:"user_id"`
	*jwt.RegisteredClaims
}

func GenerateAccessToken(userId string) (string, error) {
	tokenStruct := jwt.NewWithClaims(constants.AccessJwtSigningMethod, AccessJwtClaims{
		UserId: userId,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(constants.AccessJwtExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    constants.AccessJwtIssuer,
			ID:        userId,
			Audience:  []string{},
		},
	})

	secret, err := pkgjwt.ParseECDSAPemToPrivateKey(constants.AccessJwtPrivateSecret)
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
	tokenObj, err := jwt.ParseWithClaims(tokenString, &AccessJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret, err := pkgjwt.ParseECDSAPemToPublicKey(constants.AccessJwtPublicSecret)
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

	claims, ok := tokenObj.Claims.(*AccessJwtClaims)

	if !ok {
		return nil, fmt.Errorf("invalid claims type")
	}

	if claims.Issuer != constants.AccessJwtIssuer {
		return nil, fmt.Errorf("invalid issuer")
	}

	return claims, nil
}

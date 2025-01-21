package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const jwtSecret = "secret"

type Claims struct {
	Id string
	*jwt.RegisteredClaims
}

func GenerateToken(id string) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodES256, Claims{
		Id: id,
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "RAPID",
			ID:        id,
			Audience:  []string{},
		},
	})

	token, err := tokenStruct.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	tokenObj, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(*Claims); ok && tokenObj.Valid {
		fmt.Printf("%v %v\n", claims.Id, claims.RegisteredClaims)
		return claims, nil
	}

	return nil, err
}

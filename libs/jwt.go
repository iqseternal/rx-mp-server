package libs

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JwtStruct struct {
}

type JwtClaims struct {
	Id string
	*jwt.RegisteredClaims
}

const jwtSecret = "secret"

func NewJwt() *JwtStruct {
	jwtStruct := new(JwtStruct)
	return jwtStruct
}

func (jwtStruct *JwtStruct) GenerateToken(id string) (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodES256, JwtClaims{
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

func (jwtStruct *JwtStruct) VerifyToken(tokenString string) (*JwtClaims, error) {
	tokenObj, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenObj.Claims.(*JwtClaims); ok && tokenObj.Valid {
		fmt.Printf("%v %v\n", claims.Id, claims.RegisteredClaims)
		return claims, nil
	}

	return nil, err
}

package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"
)

func TestGenerateJwtAccessTokenSecret(t *testing.T) {

}

func TestGenerateJwtRefreshTokenSecret(t *testing.T) {

}

func TestGenerateES256Secret(t *testing.T) {
	key := make([]byte, 32)

	if _, err := rand.Read(key); err != nil {
		panic("密钥生成失败：" + err.Error())
	}

	priKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic("私钥生成失败：" + err.Error())
	}

	fmt.Println(priKey)
}

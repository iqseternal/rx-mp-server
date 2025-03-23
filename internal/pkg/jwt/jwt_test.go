package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"testing"
)

func TestGenerateJwtAccessTokenSecret(t *testing.T) {

}

func TestGenerateJwtRefreshTokenSecret(t *testing.T) {

}

func TestGenerateECDSSecret(t *testing.T) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		panic("生成密钥失败")
	}

	der, err := x509.MarshalECPrivateKey(privateKey)

	if err != nil {
		log.Fatalf("无法序列化私钥: %v", err)
	}

	pemBlock := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: der,
	}

	fmt.Println(string(pem.EncodeToMemory(pemBlock)))
}

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
	// 1. 生成私钥（P-256曲线）
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("生成私钥失败: %v", err)
	}

	// 2. 序列化私钥为DER格式
	derPrivate, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Fatalf("序列化私钥失败: %v", err)
	}

	// 3. 编码私钥为PEM格式
	pemPrivate := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: derPrivate,
	})
	fmt.Println("私钥(PEM):\n", string(pemPrivate))

	// 4. 生成公钥
	publicKey := &privateKey.PublicKey
	derPublic, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalf("序列化公钥失败: %v", err)
	}

	// 5. 编码公钥为PEM格式
	pemPublic := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublic,
	})
	fmt.Println("公钥(PEM):\n", string(pemPublic))
}

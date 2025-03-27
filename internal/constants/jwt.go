package constants

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const RefreshJwtExpire = 24 * time.Hour
const RefreshJwtIssuer = "rapid_rj"

var RefreshJwtPublicSecret string
var RefreshJwtPrivateSecret string

var RefreshJwtSigningMethod = jwt.SigningMethodES384

const AccessJwtExpire = 2 * time.Hour
const AccessJwtIssuer = "rapid_aj"

var AccessJwtPublicSecret string
var AccessJwtPrivateSecret string

var AccessJwtSigningMethod = jwt.SigningMethodES256

func init() {
	refreshPublicSecret, refreshPrivateSecret, err := GenerateRefershSecretPair()
	if err != nil {
		return
	}

	RefreshJwtPublicSecret = refreshPublicSecret
	RefreshJwtPrivateSecret = refreshPrivateSecret

	accessPublicSecret, accessPrivateSecret, err := GenerateAccessSecretPair()
	if err != nil {
		return
	}

	AccessJwtPublicSecret = accessPublicSecret
	AccessJwtPrivateSecret = accessPrivateSecret
}

// GenerateRefershSecretPair 生成密钥对
func GenerateRefershSecretPair() (string, string, error) {
	// 1. 生成私钥（P-256曲线）
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("生成私钥失败: %v", err)
		return "", "", err
	}

	// 2. 序列化私钥为DER格式
	derPrivate, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Fatalf("序列化私钥失败: %v", err)
		return "", "", err
	}

	// 3. 编码私钥为PEM格式
	pemPrivate := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: derPrivate,
	})

	// 4. 生成公钥
	publicKey := &privateKey.PublicKey
	derPublic, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalf("序列化公钥失败: %v", err)
		return "", "", err
	}

	// 5. 编码公钥为PEM格式
	pemPublic := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublic,
	})

	return string(pemPublic), string(pemPrivate), nil
}

// GenerateAccessSecretPair 生成密钥对
func GenerateAccessSecretPair() (string, string, error) {
	// 1. 生成私钥（P-256曲线）
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("生成私钥失败: %v", err)
		return "", "", err
	}

	// 2. 序列化私钥为DER格式
	derPrivate, err := x509.MarshalECPrivateKey(privateKey)
	if err != nil {
		log.Fatalf("序列化私钥失败: %v", err)
		return "", "", err
	}

	// 3. 编码私钥为PEM格式
	pemPrivate := pem.EncodeToMemory(&pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: derPrivate,
	})

	// 4. 生成公钥
	publicKey := &privateKey.PublicKey
	derPublic, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		log.Fatalf("序列化公钥失败: %v", err)
		return "", "", err
	}

	// 5. 编码公钥为PEM格式
	pemPublic := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPublic,
	})

	return string(pemPublic), string(pemPrivate), nil
}

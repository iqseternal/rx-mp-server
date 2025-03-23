package jwt

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func GenerateJwtToken() {

}

// ParseECDSAPrivateKey 解析 ECDSA Pem 字符串的 key, 加载为 *ecdsa.PrivateKey
func ParseECDSAPemToPrivateKey(pemKey string) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, fmt.Errorf("无效的 PEM 数据")
	}

	secret, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

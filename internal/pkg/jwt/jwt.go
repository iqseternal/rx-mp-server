package jwt

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"
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

// ParseECDSAPemToPublicKey 解析 ECDSA Pem 字符串的 key, 加载为 *ecdsa.PublicKey
func ParseECDSAPemToPublicKey(pemKey string) (*ecdsa.PublicKey, error) {
	// 清理可能的空白字符
	pemKey = strings.TrimSpace(pemKey)

	block, _ := pem.Decode([]byte(pemKey))
	if block == nil {
		return nil, fmt.Errorf("无效的PEM数据：未找到PEM块")
	}

	// 统一处理两种公钥格式
	var pub interface{}
	var err error
	switch block.Type {
	case "PUBLIC KEY":
		pub, err = x509.ParsePKIXPublicKey(block.Bytes)
	default:
		return nil, fmt.Errorf("不支持的PEM类型: %s", block.Type)
	}

	if err != nil {
		return nil, fmt.Errorf("解析公钥失败: %v", err)
	}

	// 类型断言 + 安全校验
	if ecdsaPub, ok := pub.(*ecdsa.PublicKey); ok {
		if ecdsaPub.Curve == nil {
			return nil, fmt.Errorf("无效的公钥：曲线参数为空")
		}
		return ecdsaPub, nil
	}
	return nil, fmt.Errorf("非ECDSA公钥类型: %T", pub)
}

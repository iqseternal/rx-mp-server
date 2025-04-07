package constants

import (
	"crypto/elliptic"
	pkgjwt "rx-mp/internal/pkg/jwt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const RefreshJwtExpire = 24 * time.Hour
const RefreshJwtIssuer = "rx_rj"

var RefreshJwtPublicSecret string
var RefreshJwtPrivateSecret string
var RefreshJwtSigningMethod = jwt.SigningMethodES384

const AccessJwtExpire = 2 * time.Hour
const AccessJwtIssuer = "rapid_aj"

var AccessJwtPublicSecret string
var AccessJwtPrivateSecret string
var AccessJwtSigningMethod = jwt.SigningMethodES256

func init() {
	refreshPublicSecret, refreshPrivateSecret, err := pkgjwt.GenerateSecretPair(elliptic.P384())
	if err != nil {
		panic(err.Error())
		return
	}

	RefreshJwtPublicSecret = refreshPublicSecret
	RefreshJwtPrivateSecret = refreshPrivateSecret

	accessPublicSecret, accessPrivateSecret, err := pkgjwt.GenerateSecretPair(elliptic.P256())
	if err != nil {
		panic(err.Error())
		return
	}

	AccessJwtPublicSecret = accessPublicSecret
	AccessJwtPrivateSecret = accessPrivateSecret
}

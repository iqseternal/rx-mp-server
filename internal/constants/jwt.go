package constants

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	RefreshJwtExpire = 24 * time.Hour
	RefreshJwtIssuer = "rapid_rj"
	RefreshJwtSecret = "abcdefghijklmnopqrstsrhxhfhfgrtjrruvwxyz"
)

var RefreshJwtSigningMethod = jwt.SigningMethodES384

const (
	AccessJwtExpire = 2 * time.Hour
	AccessJwtIssuer = "rapid_aj"
	AccessJwtSecret = "htaxewahhttyuyuiyuydsdsd"
)

var AccessJwtSigningMethod = jwt.SigningMethodES256

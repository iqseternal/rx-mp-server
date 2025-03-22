package constants

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	RefreshJwtExpire = 24 * time.Hour
	RefreshJwtIssuer = "rapid_rj"
	RefreshJwtSecret = "K5IZEQPOR11sVbTr5D1DuznNsIsZG6AsW6OK1slrORQ="
)

var RefreshJwtSigningMethod = jwt.SigningMethodES384

const (
	AccessJwtExpire = 2 * time.Hour
	AccessJwtIssuer = "rapid_aj"
	AccessJwtSecret = "K5IZEQPOR11sVbTr5D1DuznNsIsZG6AsW6OK1slrORQ="
)

var AccessJwtSigningMethod = jwt.SigningMethodES256

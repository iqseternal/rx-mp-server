package constants

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const RefreshJwtExpire = 24 * time.Hour
const RefreshJwtIssuer = "rapid_rj"
const RefreshJwtPublicSecret = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE8nYNMT87AYcRdMBeA+5aK3IGh87Z
WZbw4M4eJCswptYgPjh8oXFtiQ4jdx1cur1hZh+sNcEqLgRqBMJerAPmwA==
-----END PUBLIC KEY-----`
const RefreshJwtPrivateSecret = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIA5OpghSKgGPTNQo6aNxyUAZDBetxrTVpTFevO+v+v8joAoGCCqGSM49
AwEHoUQDQgAE8nYNMT87AYcRdMBeA+5aK3IGh87ZWZbw4M4eJCswptYgPjh8oXFt
iQ4jdx1cur1hZh+sNcEqLgRqBMJerAPmwA==
-----END EC PRIVATE KEY-----`

var RefreshJwtSigningMethod = jwt.SigningMethodES384

const AccessJwtExpire = 2 * time.Hour
const AccessJwtIssuer = "rapid_aj"
const AccessJwtPublicSecret = `-----BEGIN PUBLIC KEY-----
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEqTW776+uuycYhBzDmLZzqWHUbdDs
6vIDJemDy4a03KLTjzRPScWObfRWg9sNRx9zOPjdF9HyIZ2VV0QGsdOmmQ==
-----END PUBLIC KEY-----`
const AccessJwtPrivateSecret = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEICHKtgHr1yuA9j4Yt+WZYwYLrg4kt6+nggrRIDIEz9hgoAoGCCqGSM49
AwEHoUQDQgAEqTW776+uuycYhBzDmLZzqWHUbdDs6vIDJemDy4a03KLTjzRPScWO
bfRWg9sNRx9zOPjdF9HyIZ2VV0QGsdOmmQ==
-----END EC PRIVATE KEY-----`

var AccessJwtSigningMethod = jwt.SigningMethodES256

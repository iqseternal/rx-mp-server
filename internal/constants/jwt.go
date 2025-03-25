package constants

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const RefreshJwtExpire = 24 * time.Hour
const RefreshJwtIssuer = "rapid_rj"
const RefreshJwtSecret = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIK8JMWU8KUAhVlmtZt0L4unxvZSwFQupZ6IysZy5LdH+oAoGCCqGSM49
AwEHoUQDQgAEntROnEO9ugNR4RKKjuUJKm9Mqh5zy1CF2HGXdOu5QxjtGqDvpsIg
VywumnrD0mMr8PiDbf+RxiI/xyVHV9BnLQ==
-----END EC PRIVATE KEY-----`

var RefreshJwtSigningMethod = jwt.SigningMethodES384

const AccessJwtExpire = 2 * time.Hour
const AccessJwtIssuer = "rapid_aj"
const AccessJwtSecret = `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIAQQSgl+Y7BXj6/SQrRn68B4G0F3KGbm0jlSN9B+U5zgoAoGCCqGSM49
AwEHoUQDQgAE/8m/RSFWhhY7LJEPzXHzPUzvCbkfWDUPw9eD7y6TjuWfr3asjwSC
jlhy7Ym3OkC4DoNIb6TQw4qZSyMDuCcB3A==
-----END EC PRIVATE KEY-----`

var AccessJwtSigningMethod = jwt.SigningMethodES256

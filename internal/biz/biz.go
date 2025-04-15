package biz

import "net/http"

// 预定义, 通常情况下, 0 => 表示成功, 其余表示错误或者异常行为
const (
	Success = 0
	Failure = -1
)

// 400 +
const (
	Unauthorized     = http.StatusUnauthorized
	Forbidden        = http.StatusForbidden
	RequestTimeout   = http.StatusRequestTimeout
	MethodNotAllowed = http.StatusMethodNotAllowed
)

// 500 +
const (
	InternalServerError = http.StatusInternalServerError
	NotImplemented      = http.StatusNotImplemented
	BadGateway          = http.StatusBadGateway
	ServiceUnavailable  = http.StatusServiceUnavailable
)

// 错误化状态码：< -1000
const (
	BearerAuthorizationInvalid = -1001
	UnknownOrigin              = -1002
	ParameterError             = -1003

	NotCarryResourceAccessToken = -2001
	AccessTokenInvalid          = -2002
	RefreshTokenInvalid         = -2003
	DatabaseQueryError          = -2004
	AccessTokenExpired          = -2005
	RefreshTokenExpired         = -2006

	AttemptDeleteInValidData = -2007
	AttemptUpdateInValidData = -2008
)

// 业务化状态码: >1000
const (
	UserNotHasAdminRole = 1001
	UserNotExists       = 1010
)

var bizMessage = map[int]string{
	Success: "",
	Failure: "",

	// 400 状态码内置
	Unauthorized:     http.StatusText(Unauthorized),
	Forbidden:        http.StatusText(Forbidden),
	RequestTimeout:   http.StatusText(RequestTimeout),
	MethodNotAllowed: http.StatusText(MethodNotAllowed),

	// 500 状态码
	InternalServerError: http.StatusText(InternalServerError),
	NotImplemented:      http.StatusText(NotImplemented),
	BadGateway:          http.StatusText(BadGateway),
	ServiceUnavailable:  http.StatusText(ServiceUnavailable),

	// 错误化状态码：< -1000
	BearerAuthorizationInvalid: "Bearer Authorization Invalid",
	UnknownOrigin:              "Unknown Origin",
	ParameterError:             "Parameter Error",

	NotCarryResourceAccessToken: "Not CarryResourceAccessToken",
	AccessTokenInvalid:          "AccessToken Invalid",
	RefreshTokenInvalid:         "RefreshToken Invalid",
	DatabaseQueryError:          "database query error",
	AccessTokenExpired:          "AccessToken expired",
	RefreshTokenExpired:         "RefreshToken expired",

	AttemptDeleteInValidData: "Attempt Delete In Valid Data",
	AttemptUpdateInValidData: "Attempt Update In Valid Data",

	// 业务化状态码: >1000
	UserNotHasAdminRole: "User does not have admin role",
}

// Message 获取预定义状态码内置的错误消息返回
func Message(code int) *string {

	text := bizMessage[code]

	if text == "" {
		return nil
	}

	return &text
}

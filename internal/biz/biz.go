package biz

import "net/http"

const (
	// 预定义, 通常情况下, 0 => 表示成功, 其余表示错误或者异常行为
	BizSuccess = 0
	BizFailure = -1

	// 400 +
	BizUnauthorized    = http.StatusUnauthorized
	BizForbidden       = http.StatusForbidden
	BizRequestTimeout  = http.StatusRequestTimeout
	BizMehodNotAllowed = http.StatusMethodNotAllowed

	// 500 +
	BizInternalServerError = http.StatusInternalServerError
	BizNotImplemented      = http.StatusNotImplemented
	BizBadGateway          = http.StatusBadGateway

	// 错误化状态码：< -1000
	BizBearerAuthorizationInvalid = -1001

	BizAccessTokenInvalid  = -2002
	BizRefreshTokenInvalid = -2003
	BizDatabaseQueryError  = -2004
	BizAccessTokenExpired  = -2005
	BizRefreshTokenExpired = -2006

	// 业务化状态码: >1000
	BizUserNotHasAdminRole = 1001
)

var bizMessage = map[int]string{
	BizSuccess: "",
	BizFailure: "",

	// 400 状态码内置
	BizUnauthorized:    http.StatusText(BizUnauthorized),
	BizForbidden:       http.StatusText(BizForbidden),
	BizRequestTimeout:  http.StatusText(BizRequestTimeout),
	BizMehodNotAllowed: http.StatusText(BizMehodNotAllowed),

	// 500 状态码
	BizInternalServerError: http.StatusText(BizInternalServerError),
	BizNotImplemented:      http.StatusText(BizNotImplemented),
	BizBadGateway:          http.StatusText(BizBadGateway),

	// 错误化状态码：< -1000
	BizBearerAuthorizationInvalid: "Bearer Authorization Invalid",
	BizAccessTokenInvalid:         "AccessToken Invalid",
	BizRefreshTokenInvalid:        "RefreshToken Invalid",
	BizDatabaseQueryError:         "database query error",
	BizAccessTokenExpired:         "AccessToken expired",
	BizRefreshTokenExpired:        "RefreshToken expired",

	// 业务化状态码: >1000
	BizUserNotHasAdminRole: "User does not have admin role",
}

// BizMessage 获取预定义状态码内置的错误消息返回
func BizMessage(code int) *string {

	text := bizMessage[code]

	if text == "" {
		return nil
	}

	return &text
}

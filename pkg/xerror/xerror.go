// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package xerror

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alimy/mir/v5"
)

var _ mir.Error = (*Error)(nil)

// codes = map[int]string{}

type Error struct {
	code    int
	msg     string
	details []string
}

type ValidError struct {
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

func NewError(code int, msg string) *Error {
	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息: %s", e.StatusCode(), e.Msg())
}

func (e *Error) StatusCode() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []any) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)

	return &newError
}

func HttpStatusCode(e error) (statusCode int, code int) {
	var ok bool
	statusCode = http.StatusInternalServerError
	if code, ok = mir.HttpStatusCode(e); !ok {
		return
	}
	switch {
	case code == Success.StatusCode():
		statusCode = http.StatusOK
	case code == ServerError.StatusCode():
		statusCode = http.StatusInternalServerError
	case code == InvalidParams.StatusCode():
		statusCode = http.StatusBadRequest
	case code == UnauthorizedAuthNotExist.StatusCode(),
		code == UnauthorizedAuthFailed.StatusCode(),
		code == UnauthorizedTokenError.StatusCode(),
		code == UnauthorizedTokenGenerate.StatusCode(),
		code == UnauthorizedTokenTimeout.StatusCode():
		statusCode = http.StatusUnauthorized
	case code >= 10009 && code < 10100, code >= 20000:
		// Forbidden(10009) 及自定义权限错误码(2xxxx) → 403
		statusCode = http.StatusForbidden
	case code == TooManyRequests.StatusCode():
		statusCode = http.StatusTooManyRequests
	}
	return
}

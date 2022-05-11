package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("error code %d already defined", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}
func (err *Error) Code() int {
	return err.Code()
}
func (err *Error) Msg() string {
	return err.msg
}
func (err *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(err.msg, args...)
}
func (err *Error) Details() []string {
	return err.details
}
func (err *Error) Error() string {
	return fmt.Sprintf("error code: %d, error message: %s", err.Code(), err.Msg())
}

func (err *Error) WithDetails(details ...string) *Error {
	newErr := *err
	newErr.details = []string{}
	for _, d := range details {
		newErr.details = append(newErr.details, d)
	}
	return &newErr
}
func (err *Error) StatusCode() int {
	switch err.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}

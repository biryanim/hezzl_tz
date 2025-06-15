package errors

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrGoodsNotFound      = errors.New("goods not found")
	ErrGoodsAlreadyExists = errors.New("goods already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidInput       = errors.New("invalid input")
	ErrInternal           = errors.New("internal server error")
)

const (
	CodeAlreadyExists = 1
	CodeInvalidCreds  = 2
	CodeNotFound      = 3
	CodeInvalidInput  = 4
	CodeInternalError = 5
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.Err.Error())
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(code int, msg, details string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
		Details: details,
	}
}

func Wrap(err error, code int, msg, details string) error {
	return &AppError{
		Code:    code,
		Message: msg,
		Err:     err,
		Details: details,
	}
}

func FromError(err error) *AppError {
	if err == nil {
		return nil
	}

	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	switch {
	case errors.Is(err, ErrGoodsAlreadyExists):
		return New(CodeAlreadyExists, "error.common.alreadyExists", "")
	case errors.Is(err, ErrInvalidCredentials):
		return New(CodeInvalidCreds, "error.common.invalidCredentials", "")
	case errors.Is(err, ErrGoodsNotFound):
		return New(CodeNotFound, "error.common.notFound", "")
	case errors.Is(err, ErrInvalidInput):
		return New(CodeInvalidInput, "error.common.invalidInput", "")
	default:
		return New(CodeInternalError, "error.common.internalError", "")
	}

}

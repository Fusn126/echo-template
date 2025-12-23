package utils

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AppError 应用错误
type AppError struct {
	Code    int
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Message
}

// NewAppError 创建应用错误
func NewAppError(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

// ErrBadRequest 400 错误
func ErrBadRequest(message string) *AppError {
	return NewAppError(http.StatusBadRequest, message, nil)
}

// ErrNotFound 404 错误
func ErrNotFound(message string) *AppError {
	return NewAppError(http.StatusNotFound, message, nil)
}

// ErrInternal 500 错误
func ErrInternal(message string, err error) *AppError {
	return NewAppError(http.StatusInternalServerError, message, err)
}

// HandleError 处理错误并返回响应
func HandleError(c echo.Context, err error) error {
	var appErr *AppError
	if errors.As(err, &appErr) {
		return Error(c, appErr.Code, appErr.Message)
	}
	// 默认返回 500 错误
	return ErrorInternal(c, "内部服务器错误")
}

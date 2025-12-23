package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response 通用响应结构
// @Description 通用API响应
type Response struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg" example:"success"`
}

// SuccessResponse 成功响应（用于删除等操作）
// @Description 成功响应
type SuccessResponse struct {
	Code int    `json:"code" example:"200"`
	Msg  string `json:"msg" example:"操作成功"`
}

// ErrorResponse 错误响应
// @Description 错误响应
type ErrorResponse struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"msg" example:"操作失败"`
}

// Success 成功响应（通用，适用于任何数据类型）
func Success(c echo.Context, data interface{}, msg string) error {
	return c.JSON(http.StatusOK, Response{
		Code: http.StatusOK,
		Data: data,
		Msg:  msg,
	})
}

// SuccessCreated 创建成功响应（通用）
func SuccessCreated(c echo.Context, data interface{}, msg string) error {
	return c.JSON(http.StatusCreated, Response{
		Code: http.StatusCreated,
		Data: data,
		Msg:  msg,
	})
}

// SuccessNoContent 无内容成功响应（用于删除等操作）
func SuccessNoContent(c echo.Context, msg string) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Code: http.StatusOK,
		Msg:  msg,
	})
}

// Error 错误响应
func Error(c echo.Context, statusCode int, msg string) error {
	return c.JSON(statusCode, ErrorResponse{
		Code: statusCode,
		Msg:  msg,
	})
}

// ErrorBadRequest 400 错误响应
func ErrorBadRequest(c echo.Context, msg string) error {
	return Error(c, http.StatusBadRequest, msg)
}

// ErrorNotFound 404 错误响应
func ErrorNotFound(c echo.Context, msg string) error {
	return Error(c, http.StatusNotFound, msg)
}

// ErrorInternal 500 错误响应
func ErrorInternal(c echo.Context, msg string) error {
	return Error(c, http.StatusInternalServerError, msg)
}

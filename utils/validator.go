package utils

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// ParseUintParam 解析路径中的 uint 参数
func ParseUintParam(c echo.Context, paramName string) (uint, error) {
	idStr := c.Param(paramName)
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, ErrBadRequest("Invalid " + paramName)
	}
	return uint(id), nil
}

// BindAndValidate 绑定并验证请求体
func BindAndValidate(c echo.Context, dest interface{}) error {
	if err := c.Bind(dest); err != nil {
		return ErrBadRequest("Invalid request body")
	}
	// 可以在这里添加验证逻辑，例如使用 validator 库
	return nil
}

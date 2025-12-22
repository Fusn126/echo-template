package routes

import (
	"echo-template/app/routes/v1"
	"echo-template/middleware"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo) {
	// Swagger 文档
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// API 路由组
	api := e.Group("/api")

	// 应用中间件
	api.Use(middleware.CORS())

	// 注册版本路由
	v1.RegisterRoutes(api)

	// 健康检查
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]interface{}{
			"status": "ok",
		})
	})
}


package v1

import (
	"echo-template/app/controllers"

	"github.com/labstack/echo/v4"
)

// RegisterRoutes 注册 v1 版本的路由
func RegisterRoutes(api *echo.Group) {
	v1 := api.Group("/v1")

	// 用户路由
	userController := controllers.NewUserController()
	users := v1.Group("/users")
	{
		users.GET("", userController.GetUsers)
		users.GET("/:id", userController.GetUser)
		users.POST("", userController.CreateUser)
		users.PUT("/:id", userController.UpdateUser)
		users.DELETE("/:id", userController.DeleteUser)
	}
}


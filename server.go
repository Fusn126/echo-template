package main

import (
	"echo-template/app/models"
	"echo-template/app/routes"
	"echo-template/config"
	"echo-template/database"
	"echo-template/docs"
	"echo-template/middleware"
	"log"

	"github.com/labstack/echo/v4"
)

func init() {
	// 初始化 Swagger 文档
	docs.SwaggerInfo.Host = "localhost:1323"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

// @title           Echo Template API
// @version         1.0
// @description     基于 Echo 框架的 MVC 架构 API 文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:1323
// @BasePath  /api

// @schemes   http https

func main() {
	// 加载配置
	if err := config.LoadConfig(); err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 初始化数据库
	if err := database.InitDB(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 自动迁移数据库表
	db := database.GetDB()
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// 创建 Echo 实例
	e := echo.New()

	// 注册中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 注册路由
	routes.InitRoutes(e)

	// 启动服务器
	serverAddr := config.AppConfig.Server.Host + ":" + config.AppConfig.Server.Port
	log.Printf("Server starting on %s", serverAddr)
	e.Logger.Fatal(e.Start(serverAddr))
}

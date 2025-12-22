# Echo MVC 脚手架

基于 Echo 框架的 MVC 架构脚手架，支持前后端分离，集成 GORM ORM。

## 项目结构

```
.
├── app/             # MVC 应用核心代码
│   ├── controllers/ # 控制器层
│   ├── models/      # 数据模型
│   ├── routes/      # 路由配置
│   │   ├── v1/      # v1 版本路由
│   │   └── v2/      # v2 版本路由
│   └── services/    # 业务逻辑层
├── config/          # 配置文件
├── database/        # 数据库连接
├── middleware/      # 中间件
├── server.go        # 主入口文件
└── .env.example     # 环境变量示例
```

## 功能特性

- ✅ MVC 架构模式（集中在 app 目录）
- ✅ API 版本控制（v1、v2）
- ✅ 前后端分离（CORS 支持）
- ✅ GORM ORM 集成
- ✅ 支持 PostgreSQL、MySQL、SQLite
- ✅ 自动数据库迁移
- ✅ 中间件支持（日志、恢复、CORS）
- ✅ RESTful API 示例
- ✅ Swagger API 文档自动生成

## 快速开始

### 1. 安装依赖

```bash
go mod download
```

### 2. 配置环境变量

复制 `.env.example` 为 `.env` 并修改配置：

```bash
cp .env.example .env
```

### 3. 生成 Swagger 文档（首次运行或修改 API 后）

```bash
swag init -g server.go -o docs --parseDependency --parseInternal
```

或者使用 go run：

```bash
go run github.com/swaggo/swag/cmd/swag@latest init -g server.go -o docs --parseDependency --parseInternal
```

### 4. 运行项目

```bash
go run server.go
```

服务器将在 `http://localhost:1323` 启动

访问 `http://localhost:1323/swagger/index.html` 查看 API 文档

## API 示例

### 用户管理 API (v1)

- `GET /api/v1/users` - 获取用户列表
- `GET /api/v1/users/:id` - 获取单个用户
- `POST /api/v1/users` - 创建用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 用户管理 API (v2)

- `GET /api/v2/users` - 获取用户列表
- `GET /api/v2/users/:id` - 获取单个用户
- `POST /api/v2/users` - 创建用户
- `PUT /api/v2/users/:id` - 更新用户
- `DELETE /api/v2/users/:id` - 删除用户

### 健康检查

- `GET /health` - 健康检查接口

### Swagger API 文档

- `GET /swagger/index.html` - Swagger UI 文档界面

启动服务器后，访问 `http://localhost:1323/swagger/index.html` 查看完整的 API 文档。

## 数据库支持

项目支持多种数据库：

- PostgreSQL（默认）
- MySQL
- SQLite

通过 `DB_TYPE` 环境变量切换数据库类型。

## 开发指南

### 添加新的模型

1. 在 `app/models/` 目录创建模型文件
2. 在 `server.go` 的 `AutoMigrate` 中添加模型

### 添加新的控制器

1. 在 `app/controllers/` 目录创建控制器
2. 在 `app/services/` 目录创建对应的服务
3. 在 `app/routes/v1/routes.go` 或 `app/routes/v2/routes.go` 中注册路由

### API 版本管理

- v1 版本路由：`app/routes/v1/routes.go`
- v2 版本路由：`app/routes/v2/routes.go`
- 主路由文件：`app/routes/routes.go` 负责注册所有版本的路由

每个版本可以有不同的实现，便于 API 的演进和向后兼容。

### Swagger 文档

项目集成了 Swagger，可以自动生成 API 文档。

**添加 Swagger 注释：**

在控制器方法上添加注释：

```go
// GetUsers 获取用户列表
// @Summary      获取用户列表
// @Description  获取所有用户信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Router       /v1/users [get]
func (uc *UserController) GetUsers(c echo.Context) error {
    // ...
}
```

**生成文档：**

修改 API 后，需要重新生成 Swagger 文档：

```bash
swag init -g server.go -o docs --parseDependency --parseInternal
```

**查看文档：**

启动服务器后访问：`http://localhost:1323/swagger/index.html`

## 依赖

- Echo v4 - Web 框架
- GORM - ORM 框架
- godotenv - 环境变量管理
- Swagger/OpenAPI - API 文档生成


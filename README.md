# Echo MVC 脚手架

基于 Echo 框架的 MVC 架构脚手架，支持前后端分离，集成 GORM ORM 和 Swagger 文档。

## 项目结构

```
.
├── app/                    # MVC 应用核心代码
│   ├── controllers/       # 控制器层
│   ├── models/           # 数据模型
│   ├── routes/           # 路由配置
│   │   └── v1/           # v1 版本路由
│   └── services/         # 业务逻辑层（接口化设计）
├── config/               # 配置文件
├── database/            # 数据库连接
├── docs/                # Swagger 文档
├── middleware/          # 中间件
├── utils/               # 工具包
│   ├── errors.go        # 统一错误处理
│   ├── response.go      # 统一响应处理
│   └── validator.go     # 参数验证
├── server.go           # 主入口文件
└── .env.example        # 环境变量示例
```

## 功能特性

- ✅ **MVC 架构模式** - 清晰的代码分层，集中在 `app` 目录
- ✅ **API 版本控制** - 支持多版本 API（v1, v2...）
- ✅ **前后端分离** - CORS 中间件支持
- ✅ **GORM ORM** - 支持 PostgreSQL、MySQL、SQLite
- ✅ **自动数据库迁移** - 启动时自动创建表结构
- ✅ **统一响应格式** - 通用响应工具，适用于所有业务
- ✅ **统一错误处理** - 自定义错误类型，集中处理
- ✅ **参数验证** - 统一的参数解析和验证工具
- ✅ **服务层接口化** - 便于测试和扩展
- ✅ **Swagger 文档** - 自动生成 API 文档

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

编辑 `.env` 文件，配置数据库连接等信息。

### 3. 生成 Swagger 文档

首次运行或修改 API 后需要生成文档：

```bash
# 方式1：使用 swag 命令（需要先安装）
swag init -g server.go -o docs --parseDependency --parseInternal

# 方式2：使用 go run（推荐，无需安装）
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

### 其他接口

- `GET /health` - 健康检查
- `GET /swagger/index.html` - Swagger UI 文档界面

## 响应格式

项目使用统一的响应格式：

**成功响应（有数据）：**
```json
{
  "code": 200,
  "data": {...},
  "msg": "操作成功"
}
```

**成功响应（无数据）：**
```json
{
  "code": 200,
  "msg": "操作成功"
}
```

**错误响应：**
```json
{
  "code": 400,
  "msg": "错误信息"
}
```

## 开发指南

### 添加新的模型

1. 在 `app/models/` 目录创建模型文件
2. 在 `server.go` 的 `AutoMigrate` 中添加模型：

```go
db.AutoMigrate(&models.User{}, &models.NewModel{})
```

### 添加新的控制器和服务

1. **创建服务接口**（`app/services/interfaces.go`）：
```go
type ProductServiceInterface interface {
    GetAllProducts() ([]models.Product, error)
    GetProductByID(id uint) (*models.Product, error)
}
```

2. **实现服务**（`app/services/product_service.go`）：
```go
type ProductService struct {
    db *gorm.DB
}

func NewProductService() *ProductService {
    return &ProductService{db: database.GetDB()}
}

func (ps *ProductService) GetAllProducts() ([]models.Product, error) {
    // 实现逻辑
}
```

3. **创建控制器**（`app/controllers/product_controller.go`）：
```go
type ProductController struct {
    productService services.ProductServiceInterface
}

func (pc *ProductController) GetProducts(c echo.Context) error {
    products, err := pc.productService.GetAllProducts()
    if err != nil {
        return utils.HandleError(c, err)
    }
    return utils.Success(c, products, "获取产品列表成功")
}
```

4. **注册路由**（`app/routes/v1/routes.go`）：
```go
productController := controllers.NewProductController()
v1.GET("/products", productController.GetProducts)
```

### 使用工具函数

**统一响应：**
```go
// 成功响应（有数据）
return utils.Success(c, data, "操作成功")

// 创建成功
return utils.SuccessCreated(c, data, "创建成功")

// 删除成功（无数据）
return utils.SuccessNoContent(c, "删除成功")
```

**统一错误处理：**
```go
// 服务层返回错误
return utils.ErrNotFound("资源不存在")
return utils.ErrBadRequest("参数错误")
return utils.ErrInternal("操作失败", err)

// 控制器处理错误
if err != nil {
    return utils.HandleError(c, err)
}
```

**参数验证：**
```go
// 解析路径参数
id, err := utils.ParseUintParam(c, "id")

// 绑定并验证请求体
var user models.User
if err := utils.BindAndValidate(c, &user); err != nil {
    return utils.HandleError(c, err)
}
```

### Swagger 文档

**添加 Swagger 注释：**

```go
// GetUsers 获取用户列表
// @Summary      获取用户列表
// @Description  获取所有用户信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.Response{data=[]models.User}  "成功返回用户列表"
// @Failure      500  {object}  utils.ErrorResponse  "服务器错误"
// @Router       /v1/users [get]
func (uc *UserController) GetUsers(c echo.Context) error {
    // ...
}
```

**生成文档：**

修改 API 后，重新生成 Swagger 文档：

```bash
go run github.com/swaggo/swag/cmd/swag@latest init -g server.go -o docs --parseDependency --parseInternal
```

## 数据库支持

项目支持多种数据库，通过 `DB_TYPE` 环境变量切换：

- `postgres` - PostgreSQL（默认）
- `mysql` - MySQL
- `sqlite` - SQLite

## 架构设计

### 分层架构

- **Controller 层** - 处理 HTTP 请求，调用 Service
- **Service 层** - 业务逻辑处理，接口化设计
- **Model 层** - 数据模型定义
- **Utils 层** - 通用工具函数（响应、错误、验证）

### 设计模式

- **接口化服务** - Service 层使用接口，便于测试和扩展
- **统一响应** - 所有 API 使用统一的响应格式
- **统一错误处理** - 自定义错误类型，集中处理
- **工具函数** - 减少重复代码，提高可维护性

## 依赖

- **Echo v4** - Web 框架
- **GORM** - ORM 框架
- **godotenv** - 环境变量管理
- **swaggo/swag** - Swagger 文档生成
- **validator** - 参数验证

## License

Apache 2.0

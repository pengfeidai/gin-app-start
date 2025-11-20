# Gin App Start 项目使用指南

## 📋 目录

1. [项目概述](#项目概述)
2. [技术栈](#技术栈)
3. [目录结构](#目录结构)
4. [快速开始](#快速开始)
5. [配置说明](#配置说明)
6. [API 文档](#api-文档)
7. [开发指南](#开发指南)
8. [部署指南](#部署指南)
9. [最佳实践](#最佳实践)
10. [常见问题](#常见问题)

---

## 项目概述

Gin App Start 是一个基于 Go 1.24 和 Gin 框架构建的现代化 Web 应用脚手架。项目采用清晰的分层架构设计，遵循 Go 语言最佳实践，支持 PostgreSQL 和 Redis，适合快速开发企业级 Web 应用。

### 核心特性

- ✅ **现代化架构**: 采用标准的 `cmd/internal/pkg` 目录结构
- ✅ **清晰分层**: Controller → Service → Repository 三层架构
- ✅ **数据库支持**: PostgreSQL (GORM) + Redis
- ✅ **结构化日志**: 使用 uber/zap 实现高性能日志
- ✅ **统一错误处理**: 自定义业务错误和错误码管理
- ✅ **统一响应格式**: 标准化的 API 响应结构
- ✅ **中间件支持**: 日志、恢复、CORS、限流等
- ✅ **优雅关闭**: 支持 Graceful Shutdown
- ✅ **环境配置**: 基于 Viper 的多环境配置管理
- ✅ **容器化**: 完整的 Docker 和 Docker Compose 支持
- ✅ **热重载**: 支持 Air 热重载开发

---

## 技术栈

### 核心框架

| 技术 | 版本 | 用途 |
|------|------|------|
| Go | 1.24+ | 编程语言 |
| Gin | v1.10.0 | Web 框架 |
| GORM | v1.25.12 | ORM 框架 |
| PostgreSQL | 17+ | 关系型数据库 |
| Redis | 7+ | 缓存数据库 |

### 主要依赖

- **日志**: go.uber.org/zap v1.27.0
- **配置**: github.com/spf13/viper v1.19.0
- **数据库驱动**: gorm.io/driver/postgres v1.5.11
- **Redis客户端**: github.com/redis/go-redis/v9 v9.7.0
- **CORS**: github.com/gin-contrib/cors v1.7.2
- **Session**: github.com/gin-contrib/sessions v1.0.1
- **UUID**: github.com/google/uuid v1.6.0

---

## 目录结构

```
gin-app-start/
├── cmd/                          # 应用程序入口
│   └── server/
│       └── main.go              # 主入口文件，初始化和启动服务
│
├── internal/                     # 私有应用程序代码（不可被外部导入）
│   ├── config/                  # 配置管理
│   │   └── config.go           # 配置结构定义和加载逻辑
│   │
│   ├── controller/              # 控制器层（HTTP 处理）
│   │   ├── health_controller.go    # 健康检查控制器
│   │   └── user_controller.go      # 用户管理控制器
│   │
│   ├── service/                 # 业务逻辑层
│   │   └── user_service.go     # 用户业务逻辑
│   │
│   ├── repository/              # 数据访问层
│   │   ├── base_repository.go  # 基础仓库（泛型）
│   │   └── user_repository.go  # 用户数据访问
│   │
│   ├── model/                   # 数据模型
│   │   ├── base.go             # 基础模型定义
│   │   └── user.go             # 用户模型
│   │
│   ├── middleware/              # 中间件
│   │   ├── cors.go             # CORS 跨域处理
│   │   ├── logger.go           # 日志记录
│   │   ├── rate_limit.go       # 限流控制
│   │   └── recovery.go         # Panic 恢复
│   │
│   └── router/                  # 路由配置
│       └── router.go           # 路由注册和分组
│
├── pkg/                         # 公共库代码（可被外部导入）
│   ├── database/               # 数据库连接
│   │   ├── postgres.go        # PostgreSQL 连接管理
│   │   └── redis.go           # Redis 连接管理
│   │
│   ├── errors/                 # 错误处理
│   │   └── errors.go          # 业务错误定义
│   │
│   ├── logger/                 # 日志系统
│   │   └── logger.go          # 日志初始化和封装
│   │
│   ├── response/               # 响应格式
│   │   └── response.go        # 统一响应结构
│   │
│   └── utils/                  # 工具函数
│       ├── crypto.go          # 加密相关工具
│       └── utils.go           # 通用工具函数
│
├── configs/                     # 配置文件目录
│   ├── config.local.yaml       # 本地开发配置
│   ├── config.dev.yaml         # 开发环境配置
│   └── config.prod.yaml        # 生产环境配置
│
├── docs/                        # 文档目录
│   └── PROJECT_GUIDE.md        # 本文档
│
├── docker-compose.yml           # Docker Compose 配置
├── Dockerfile                   # Docker 镜像构建文件
├── Makefile                     # 常用命令快捷方式
├── go.mod                       # Go 模块依赖
├── go.sum                       # 依赖校验文件
└── README.md                    # 项目说明
```

### 目录说明

#### cmd/
应用程序的入口点。按照 Go 标准项目布局，每个可执行程序都应该有一个对应的子目录。

#### internal/
私有应用程序代码。这是你不希望其他应用程序或库导入的代码。这个布局模式由 Go 编译器本身强制执行。

#### pkg/
可以被外部应用程序使用的库代码。其他项目会导入这些库，所以在这里放代码之前要三思。

#### configs/
配置文件模板或默认配置。

---

## 快速开始

### 前置要求

- Go 1.24 或更高版本
- PostgreSQL 12+ （推荐 17）
- Redis 6.0+ （推荐 7）
- Docker 和 Docker Compose（可选）

### 方式一：使用 Docker Compose（推荐）

这是最简单快速的启动方式，会自动启动所有依赖服务。

```bash
# 1. 克隆项目
git clone <your-repo-url>
cd gin-app-start

# 2. 启动所有服务（PostgreSQL + Redis + App）
docker-compose up -d

# 3. 查看日志
docker-compose logs -f app

# 4. 停止服务
docker-compose down
```

应用将在 http://localhost:9060 上运行。

### 方式二：本地开发

#### 1. 安装依赖

```bash
# 下载 Go 依赖
go mod download

# 验证依赖
go mod verify
```

#### 2. 启动数据库服务

**使用 Docker 启动数据库：**

```bash
# 启动 PostgreSQL
docker run -d \
  --name gin-app-postgres \
  -e POSTGRES_DB=gin_app \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -p 5432:5432 \
  postgres:17-alpine

# 启动 Redis
docker run -d \
  --name gin-app-redis \
  -p 6379:6379 \
  redis:7-alpine
```

**或者使用本地安装的数据库**，确保服务已启动。

#### 3. 配置环境

编辑 `configs/config.local.yaml`，根据实际情况调整数据库连接信息：

```yaml
database:
  host: localhost
  port: 5432
  user: postgres
  password: postgres
  dbname: gin_app
```

#### 4. 运行应用

```bash
# 使用 Makefile
make run

# 或直接使用 go run
SERVER_ENV=local go run cmd/server/main.go
```

#### 5. 验证运行

```bash
# 健康检查
curl http://localhost:9060/health

# 应该返回：
# {"code":0,"message":"success","data":{"message":"service is running","status":"ok"}}
```

### 方式三：使用热重载开发

安装 Air 工具后，可以实现代码自动重新编译和重启：

```bash
# 安装 Air
go install github.com/cosmtrek/air@latest

# 使用 Air 运行（代码变更自动重启）
air

# 或使用 Makefile
make dev
```

---

## 配置说明

### 配置文件结构

项目使用 YAML 格式的配置文件，支持多环境配置。配置文件位于 `configs/` 目录下。

### 环境选择

通过环境变量 `SERVER_ENV` 指定使用哪个配置文件：

```bash
SERVER_ENV=local    # 使用 configs/config.local.yaml
SERVER_ENV=dev      # 使用 configs/config.dev.yaml
SERVER_ENV=prod     # 使用 configs/config.prod.yaml
```

### 配置项说明

#### 服务器配置 (server)

```yaml
server:
  port: 9060              # 服务监听端口
  mode: debug             # 运行模式: debug/release/test
  read_timeout: 60        # 读取超时（秒）
  write_timeout: 60       # 写入超时（秒）
  limit_num: 100          # 限流：每秒允许的请求数，0 表示不限流
```

#### 数据库配置 (database)

```yaml
database:
  host: localhost         # 数据库主机地址
  port: 5432             # 数据库端口
  user: postgres         # 数据库用户名
  password: postgres     # 数据库密码
  dbname: gin_app        # 数据库名称
  sslmode: disable       # SSL 模式: disable/require/verify-full
  max_idle_conns: 10     # 最大空闲连接数
  max_open_conns: 100    # 最大打开连接数
  max_lifetime: 3600     # 连接最大生命周期（秒）
  log_level: info        # 日志级别: silent/error/warn/info
  auto_migrate: true     # 是否自动迁移数据库表结构
```

#### Redis 配置 (redis)

```yaml
redis:
  addr: localhost:6379   # Redis 服务器地址
  password: ""           # Redis 密码（空表示无密码）
  db: 0                  # Redis 数据库编号 (0-15)
  pool_size: 10          # 连接池大小
  min_idle_conns: 5      # 最小空闲连接数
  max_retries: 3         # 最大重试次数
```

#### 日志配置 (log)

```yaml
log:
  level: debug           # 日志级别: debug/info/warn/error
  file_path: logs/app.log  # 日志文件路径
  max_size: 100          # 单个日志文件最大大小（MB）
  max_age: 7             # 日志文件保留天数
```

#### Session 配置 (session)

```yaml
session:
  key: gin-session       # Session 密钥
  max_age: 604800        # Session 过期时间（秒），604800 = 7天
  path: /                # Cookie 路径
  domain: ""             # Cookie 域名
  http_only: true        # 是否仅 HTTP 访问
  secure: false          # 是否仅 HTTPS（生产环境应设为 true）
```

### 环境变量支持

配置文件支持使用环境变量，格式为 `${变量名}`：

```yaml
database:
  host: ${DB_HOST}
  password: ${DB_PASSWORD}
```

启动时设置环境变量：

```bash
export DB_HOST=192.168.1.100
export DB_PASSWORD=secret_password
SERVER_ENV=prod go run cmd/server/main.go
```

---

## API 文档

### 响应格式

所有 API 返回统一的 JSON 格式：

#### 成功响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    // 实际数据
  }
}
```

#### 错误响应

```json
{
  "code": 10001,
  "message": "Invalid parameters",
  "data": null
}
```

#### 分页响应

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [...],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

### 错误码说明

| 错误码 | 说明 |
|--------|------|
| 0 | 成功 |
| 10001 | 参数错误 |
| 10002 | 用户不存在 |
| 10003 | 未授权访问 |
| 10004 | 用户已存在 |
| 10005 | 数据库错误 |
| 42900 | 请求过于频繁 |
| 50000 | 系统内部错误 |

### API 端点

#### 健康检查

**GET /health**

检查服务是否正常运行。

**请求示例：**

```bash
curl http://localhost:9060/health
```

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "status": "ok",
    "message": "service is running"
  }
}
```

---

#### 用户管理

##### 1. 创建用户

**POST /api/v1/users**

**请求体：**

```json
{
  "username": "testuser",
  "email": "test@example.com",
  "phone": "13800138000",
  "password": "password123"
}
```

**字段说明：**

- `username`: 用户名，必填，3-32个字符
- `email`: 邮箱，选填，需符合邮箱格式
- `phone`: 手机号，选填，11位数字
- `password`: 密码，必填，6-32个字符

**请求示例：**

```bash
curl -X POST http://localhost:9060/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "phone": "13800138000",
    "password": "password123"
  }'
```

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "created_at": "2025-11-20T10:00:00Z",
    "update_at": "2025-11-20T10:00:00Z",
    "username": "testuser",
    "email": "test@example.com",
    "phone": "13800138000",
    "avatar": "",
    "status": 1
  }
}
```

##### 2. 获取用户信息

**GET /api/v1/users/:id**

**路径参数：**

- `id`: 用户ID

**请求示例：**

```bash
curl http://localhost:9060/api/v1/users/1
```

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "phone": "13800138000",
    "avatar": "",
    "status": 1
  }
}
```

##### 3. 更新用户信息

**PUT /api/v1/users/:id**

**请求体：**

```json
{
  "email": "newemail@example.com",
  "phone": "13900139000",
  "avatar": "http://example.com/avatar.jpg",
  "status": 1
}
```

**字段说明：**

- 所有字段都是选填
- `status`: 1表示正常，0表示禁用

**请求示例：**

```bash
curl -X PUT http://localhost:9060/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "phone": "13900139000"
  }'
```

##### 4. 删除用户

**DELETE /api/v1/users/:id**

**请求示例：**

```bash
curl -X DELETE http://localhost:9060/api/v1/users/1
```

**响应示例：**

```json
{
  "code": 0,
  "message": "Deleted successfully",
  "data": null
}
```

##### 5. 用户列表

**GET /api/v1/users**

**查询参数：**

- `page`: 页码，默认 1
- `page_size`: 每页数量，默认 10，最大 100

**请求示例：**

```bash
curl "http://localhost:9060/api/v1/users?page=1&page_size=10"
```

**响应示例：**

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "username": "testuser",
        "email": "test@example.com",
        "status": 1
      }
    ],
    "total": 100,
    "page": 1,
    "page_size": 10
  }
}
```

---

## 开发指南

### 添加新功能

按照以下步骤添加新的业务功能：

#### 1. 定义数据模型

在 `internal/model/` 中创建模型文件，例如 `product.go`：

```go
package model

import (
    "time"
    "gorm.io/gorm"
)

type Product struct {
    ID        uint           `gorm:"primarykey" json:"id"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
    Name      string         `gorm:"size:128;not null" json:"name"`
    Price     float64        `gorm:"not null" json:"price"`
    Stock     int            `gorm:"default:0" json:"stock"`
}

func (Product) TableName() string {
    return "products"
}
```

#### 2. 创建 Repository 层

在 `internal/repository/` 中创建 `product_repository.go`：

```go
package repository

import (
    "context"
    "gin-app-start/internal/model"
    "gorm.io/gorm"
)

type ProductRepository interface {
    Create(ctx context.Context, product *model.Product) error
    GetByID(ctx context.Context, id uint) (*model.Product, error)
    Update(ctx context.Context, product *model.Product) error
    Delete(ctx context.Context, id uint) error
    List(ctx context.Context, offset, limit int) ([]*model.Product, int64, error)
}

type productRepository struct {
    *BaseRepository[model.Product]
}

func NewProductRepository(db *gorm.DB) ProductRepository {
    return &productRepository{
        BaseRepository: NewBaseRepository[model.Product](db),
    }
}

func (r *productRepository) List(ctx context.Context, offset, limit int) ([]*model.Product, int64, error) {
    var products []*model.Product
    var total int64

    if err := r.db.WithContext(ctx).Model(&model.Product{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    err := r.db.WithContext(ctx).Offset(offset).Limit(limit).Find(&products).Error
    return products, total, err
}
```

#### 3. 实现 Service 层

在 `internal/service/` 中创建 `product_service.go`：

```go
package service

import (
    "context"
    "gin-app-start/internal/model"
    "gin-app-start/internal/repository"
    "gin-app-start/pkg/errors"
    "gin-app-start/pkg/logger"
    "go.uber.org/zap"
)

type ProductService interface {
    CreateProduct(ctx context.Context, req *CreateProductRequest) (*model.Product, error)
    GetProduct(ctx context.Context, id uint) (*model.Product, error)
    UpdateProduct(ctx context.Context, id uint, req *UpdateProductRequest) (*model.Product, error)
    DeleteProduct(ctx context.Context, id uint) error
    ListProducts(ctx context.Context, page, pageSize int) ([]*model.Product, int64, error)
}

type productService struct {
    productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
    return &productService{
        productRepo: productRepo,
    }
}

type CreateProductRequest struct {
    Name  string  `json:"name" binding:"required,min=1,max=128"`
    Price float64 `json:"price" binding:"required,gt=0"`
    Stock int     `json:"stock" binding:"gte=0"`
}

type UpdateProductRequest struct {
    Name  string  `json:"name" binding:"omitempty,min=1,max=128"`
    Price float64 `json:"price" binding:"omitempty,gt=0"`
    Stock int     `json:"stock" binding:"omitempty,gte=0"`
}

func (s *productService) CreateProduct(ctx context.Context, req *CreateProductRequest) (*model.Product, error) {
    product := &model.Product{
        Name:  req.Name,
        Price: req.Price,
        Stock: req.Stock,
    }

    if err := s.productRepo.Create(ctx, product); err != nil {
        logger.Error("Failed to create product", zap.Error(err))
        return nil, errors.WrapBusinessError(20001, "Failed to create product", err)
    }

    logger.Info("Product created successfully", zap.Uint("product_id", product.ID))
    return product, nil
}

// ... 实现其他方法
```

#### 4. 创建 Controller 层

在 `internal/controller/` 中创建 `product_controller.go`：

```go
package controller

import (
    "gin-app-start/internal/service"
    "gin-app-start/pkg/response"
    "strconv"
    "github.com/gin-gonic/gin"
)

type ProductController struct {
    productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
    return &ProductController{
        productService: productService,
    }
}

func (ctrl *ProductController) CreateProduct(c *gin.Context) {
    var req service.CreateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, 10001, "Parameter binding failed: "+err.Error())
        return
    }

    product, err := ctrl.productService.CreateProduct(c.Request.Context(), &req)
    if err != nil {
        handleServiceError(c, err)
        return
    }

    response.Success(c, product)
}

// ... 实现其他方法
```

#### 5. 注册路由

在 `internal/router/router.go` 中添加路由：

```go
func SetupRouter(
    healthCtrl *controller.HealthController,
    userCtrl *controller.UserController,
    productCtrl *controller.ProductController,  // 新增
    cfg *config.Config,
) *gin.Engine {
    // ... 现有代码 ...

    apiV1 := router.Group("/api/v1")
    {
        // ... 用户路由 ...

        // 产品路由
        products := apiV1.Group("/products")
        {
            products.POST("", productCtrl.CreateProduct)
            products.GET("/:id", productCtrl.GetProduct)
            products.PUT("/:id", productCtrl.UpdateProduct)
            products.DELETE("/:id", productCtrl.DeleteProduct)
            products.GET("", productCtrl.ListProducts)
        }
    }

    return router
}
```

#### 6. 更新 main.go

在 `cmd/server/main.go` 中初始化新功能：

```go
func main() {
    // ... 现有初始化代码 ...

    // 初始化依赖
    userRepo := repository.NewUserRepository(db)
    userService := service.NewUserService(userRepo)
    userController := controller.NewUserController(userService)

    // 新增产品模块
    productRepo := repository.NewProductRepository(db)
    productService := service.NewProductService(productRepo)
    productController := controller.NewProductController(productService)

    healthController := controller.NewHealthController()

    // 设置路由（传入新的控制器）
    r := router.SetupRouter(healthController, userController, productController, cfg)

    // ... 其余代码 ...
}
```

#### 7. 数据库迁移

在 `cmd/server/main.go` 的自动迁移部分添加新模型：

```go
if cfg.Database.AutoMigrate {
    if err := db.AutoMigrate(
        &model.User{},
        &model.Product{},  // 新增
    ); err != nil {
        logger.Fatal("Database migration failed", zap.Error(err))
    }
    logger.Info("Database migration completed")
}
```

### 代码规范

#### 命名规范

- **文件名**: 使用 snake_case，如 `user_service.go`
- **包名**: 简短、小写、单数形式，如 `service`, `repository`
- **接口**: 以功能命名，如 `UserService`, `ProductRepository`
- **结构体**: 使用 PascalCase，如 `UserController`
- **函数/方法**: 使用 camelCase，如 `getUserByID`, `createProduct`
- **常量**: 使用 PascalCase 或全大写+下划线，如 `MaxRetryCount` 或 `MAX_RETRY_COUNT`

#### 错误处理

```go
// Service 层
func (s *userService) CreateUser(ctx context.Context, req *CreateUserRequest) (*model.User, error) {
    // 检查业务规则
    existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
    if err != nil && err != gorm.ErrRecordNotFound {
        logger.Error("Failed to query user", zap.Error(err))
        return nil, errors.WrapBusinessError(10010, "Failed to query user", err)
    }

    if existingUser != nil {
        return nil, errors.ErrUserExists
    }

    // ... 业务逻辑
}

// Controller 层
func (ctrl *UserController) CreateUser(c *gin.Context) {
    var req service.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, 10001, "Parameter binding failed: "+err.Error())
        return
    }

    user, err := ctrl.userService.CreateUser(c.Request.Context(), &req)
    if err != nil {
        handleServiceError(c, err)
        return
    }

    response.Success(c, user)
}
```

#### 日志记录

```go
import (
    "gin-app-start/pkg/logger"
    "go.uber.org/zap"
)

// 信息日志
logger.Info("User created successfully",
    zap.String("username", user.Username),
    zap.Uint("user_id", user.ID),
)

// 错误日志
logger.Error("Failed to create user",
    zap.Error(err),
    zap.String("username", req.Username),
)

// 警告日志
logger.Warn("User not found in cache, querying database",
    zap.Uint("user_id", id),
)

// 调试日志
logger.Debug("Processing request",
    zap.String("method", c.Request.Method),
    zap.String("path", c.Request.URL.Path),
)
```

### 测试

#### 单元测试示例

创建 `internal/service/user_service_test.go`：

```go
package service

import (
    "context"
    "testing"
    "gin-app-start/internal/model"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

// Mock Repository
type MockUserRepository struct {
    mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *model.User) error {
    args := m.Called(ctx, user)
    return args.Error(0)
}

func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
    args := m.Called(ctx, username)
    if args.Get(0) == nil {
        return nil, args.Error(1)
    }
    return args.Get(0).(*model.User), args.Error(1)
}

// 测试用例
func TestCreateUser(t *testing.T) {
    mockRepo := new(MockUserRepository)
    service := NewUserService(mockRepo)

    ctx := context.Background()
    req := &CreateUserRequest{
        Username: "testuser",
        Email:    "test@example.com",
        Password: "password123",
    }

    // 设置 mock 预期
    mockRepo.On("GetByUsername", ctx, "testuser").Return(nil, gorm.ErrRecordNotFound)
    mockRepo.On("Create", ctx, mock.AnythingOfType("*model.User")).Return(nil)

    // 执行测试
    user, err := service.CreateUser(ctx, req)

    // 断言
    assert.NoError(t, err)
    assert.NotNil(t, user)
    assert.Equal(t, "testuser", user.Username)
    mockRepo.AssertExpectations(t)
}
```

运行测试：

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./internal/service/...

# 查看测试覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

---

## 部署指南

### Docker 部署

#### 1. 构建镜像

```bash
docker build -t gin-app-start:latest .
```

#### 2. 运行容器

```bash
docker run -d \
  -p 9060:9060 \
  -e SERVER_ENV=prod \
  -e DB_HOST=your-postgres-host \
  -e DB_USER=postgres \
  -e DB_PASSWORD=your-password \
  -e DB_NAME=gin_app \
  -e REDIS_ADDR=your-redis-host:6379 \
  -e REDIS_PASSWORD=your-redis-password \
  --name gin-app \
  gin-app-start:latest
```

### Docker Compose 部署

使用提供的 `docker-compose.yml`：

```bash
# 启动所有服务
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止服务
docker-compose down

# 停止并删除数据卷
docker-compose down -v
```

### 二进制部署

#### 1. 编译

```bash
# Linux
GOOS=linux GOARCH=amd64 go build -o bin/server cmd/server/main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o bin/server cmd/server/main.go

# Windows
GOOS=windows GOARCH=amd64 go build -o bin/server.exe cmd/server/main.go
```

#### 2. 部署到服务器

```bash
# 1. 上传文件
scp bin/server user@your-server:/opt/gin-app/
scp -r configs user@your-server:/opt/gin-app/

# 2. SSH 到服务器
ssh user@your-server

# 3. 配置环境
cd /opt/gin-app
export SERVER_ENV=prod

# 4. 运行（使用 systemd 或 supervisor 管理）
./server
```

#### 3. Systemd 服务配置

创建 `/etc/systemd/system/gin-app.service`：

```ini
[Unit]
Description=Gin App Service
After=network.target postgresql.service redis.service

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/gin-app
Environment="SERVER_ENV=prod"
ExecStart=/opt/gin-app/server
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

管理服务：

```bash
# 重新加载配置
sudo systemctl daemon-reload

# 启动服务
sudo systemctl start gin-app

# 设置开机自启
sudo systemctl enable gin-app

# 查看状态
sudo systemctl status gin-app

# 查看日志
sudo journalctl -u gin-app -f
```

### 云平台部署

#### Kubernetes

创建 `k8s/deployment.yaml`：

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gin-app
spec:
  replicas: 3
  selector:
    matchLabels:
      app: gin-app
  template:
    metadata:
      labels:
        app: gin-app
    spec:
      containers:
      - name: gin-app
        image: gin-app-start:latest
        ports:
        - containerPort: 9060
        env:
        - name: SERVER_ENV
          value: "prod"
        - name: DB_HOST
          valueFrom:
            secretKeyRef:
              name: gin-app-secrets
              key: db-host
        resources:
          limits:
            cpu: "1"
            memory: "512Mi"
          requests:
            cpu: "100m"
            memory: "128Mi"
        livenessProbe:
          httpGet:
            path: /health
            port: 9060
          initialDelaySeconds: 30
          periodSeconds: 10
        readinessProbe:
          httpGet:
            path: /health
            port: 9060
          initialDelaySeconds: 5
          periodSeconds: 5
```

---

## 最佳实践

### 1. 数据库连接池配置

根据实际负载调整连接池大小：

```yaml
database:
  max_idle_conns: 10    # 空闲连接数 = 预期QPS / 10
  max_open_conns: 100   # 最大连接数 = max_idle_conns * 10
  max_lifetime: 3600    # 1小时，避免长连接问题
```

### 2. 日志级别设置

不同环境使用不同的日志级别：

- 开发环境: `debug`
- 测试环境: `info`
- 生产环境: `warn` 或 `error`

### 3. 错误处理

- Controller 层只处理 HTTP 相关错误
- Service 层处理业务逻辑错误
- Repository 层返回原始数据库错误
- 使用 `pkg/errors` 定义业务错误码

### 4. 性能优化

#### 数据库查询优化

```go
// 使用索引
type User struct {
    Username string `gorm:"size:64;uniqueIndex;not null"`
    Email    string `gorm:"size:128;uniqueIndex"`
}

// 预加载关联数据
db.Preload("Orders").Find(&users)

// 只查询需要的字段
db.Select("id", "username", "email").Find(&users)

// 批量插入
db.CreateInBatches(users, 100)
```

#### Redis 缓存

```go
import "github.com/redis/go-redis/v9"

// 缓存用户信息
func (s *userService) GetUser(ctx context.Context, id uint) (*model.User, error) {
    // 1. 尝试从缓存获取
    key := fmt.Sprintf("user:%d", id)
    val, err := redis.Get(ctx, key).Result()
    if err == nil {
        var user model.User
        json.Unmarshal([]byte(val), &user)
        return &user, nil
    }

    // 2. 从数据库查询
    user, err := s.userRepo.GetByID(ctx, id)
    if err != nil {
        return nil, err
    }

    // 3. 写入缓存
    data, _ := json.Marshal(user)
    redis.Set(ctx, key, data, 1*time.Hour)

    return user, nil
}
```

### 5. 安全建议

- 使用 HTTPS（生产环境）
- 实现 JWT 或 Session 认证
- 密码使用强加密（bcrypt 替代 MD5）
- 启用 CORS 白名单
- 实施 SQL 注入防护（GORM 已内置）
- 添加请求签名验证
- 实现 API 访问频率限制

---

## 常见问题

### 1. 数据库连接失败

**问题**: `failed to connect to database`

**解决方案**:
- 检查数据库是否启动：`docker ps | grep postgres`
- 验证连接信息是否正确
- 检查防火墙设置
- 确认 PostgreSQL 监听地址（`postgresql.conf` 中的 `listen_addresses`）

### 2. Redis 连接失败

**问题**: `failed to connect to redis`

**解决方案**:
- 检查 Redis 是否启动：`docker ps | grep redis`
- 验证连接地址和密码
- 测试连接：`redis-cli -h localhost -p 6379 ping`

### 3. 端口被占用

**问题**: `bind: address already in use`

**解决方案**:
```bash
# 查看占用端口的进程
lsof -i :9060

# 或使用
netstat -tlnp | grep 9060

# 终止进程
kill -9 <PID>

# 或修改配置使用其他端口
```

### 4. 数据库迁移失败

**问题**: 表结构更新失败

**解决方案**:
```bash
# 方法1: 删除所有表重新创建（仅开发环境）
DROP DATABASE gin_app;
CREATE DATABASE gin_app;

# 方法2: 手动执行迁移 SQL

# 方法3: 使用 GORM Migrator
db.Migrator().DropTable(&model.User{})
db.Migrator().CreateTable(&model.User{})
```

### 5. Go 版本不兼容

**问题**: `go: go.mod requires go >= 1.24`

**解决方案**:
```bash
# 检查当前版本
go version

# 升级 Go（macOS）
brew upgrade go

# 或下载安装包
# https://go.dev/dl/
```

### 6. 依赖下载慢

**问题**: `go mod download` 速度慢

**解决方案**:
```bash
# 使用国内代理
go env -w GOPROXY=https://goproxy.cn,direct

# 或使用其他代理
go env -w GOPROXY=https://goproxy.io,direct
```

### 7. Docker 构建失败

**问题**: Docker 镜像构建失败

**解决方案**:
```bash
# 清理 Docker 缓存
docker system prune -a

# 使用多阶段构建加速
# 已在 Dockerfile 中实现

# 使用构建缓存
docker build --cache-from gin-app-start:latest -t gin-app-start:latest .
```

### 8. 日志文件过大

**问题**: 日志文件占用大量磁盘空间

**解决方案**:
```yaml
# 在配置文件中设置日志轮转
log:
  max_size: 100   # 单文件最大 100MB
  max_age: 7      # 保留 7 天
```

或使用 `logrotate` (Linux)：

```bash
# /etc/logrotate.d/gin-app
/opt/gin-app/logs/*.log {
    daily
    rotate 7
    compress
    delaycompress
    missingok
    notifempty
}
```

### 9. 性能问题

**问题**: API 响应慢

**诊断步骤**:
```bash
# 1. 检查日志中的请求响应时间
# 日志中包含 latency 字段

# 2. 使用 pprof 性能分析
import _ "net/http/pprof"

# 访问 http://localhost:9060/debug/pprof/

# 3. 检查数据库慢查询
# PostgreSQL: 
SELECT * FROM pg_stat_statements ORDER BY mean_time DESC LIMIT 10;

# 4. 检查连接池状态
stats := db.DB().Stats()
fmt.Printf("Open connections: %d\n", stats.OpenConnections)
```

### 10. 内存泄漏

**问题**: 内存占用持续增长

**解决方案**:
```bash
# 1. 使用 pprof 分析内存
go tool pprof http://localhost:9060/debug/pprof/heap

# 2. 检查 goroutine 泄漏
go tool pprof http://localhost:9060/debug/pprof/goroutine

# 3. 设置 GOMAXPROCS
export GOMAXPROCS=4

# 4. 定期重启（临时方案）
```

---

## 附录

### Makefile 命令

```bash
make run           # 运行应用
make build         # 编译应用
make test          # 运行测试
make fmt           # 格式化代码
make lint          # 代码检查
make clean         # 清理编译文件
make deps          # 下载依赖
make dev           # 热重载开发
make docker-build  # 构建 Docker 镜像
make docker-run    # 运行 Docker 容器
```

### 推荐工具

- **开发工具**: GoLand / VSCode
- **API 测试**: Postman / Insomnia / curl
- **数据库管理**: DBeaver / pgAdmin
- **Redis 管理**: RedisInsight / Redis Desktop Manager
- **日志查看**: Kibana / Grafana Loki
- **性能监控**: Prometheus + Grafana
- **代码检查**: golangci-lint

### 学习资源

- [Go 官方文档](https://go.dev/doc/)
- [Gin 框架文档](https://gin-gonic.com/docs/)
- [GORM 文档](https://gorm.io/docs/)
- [Go 标准项目布局](https://github.com/golang-standards/project-layout)
- [Uber Go 代码规范](https://github.com/uber-go/guide/blob/master/style.md)

---

## 贡献指南

欢迎提交 Issue 和 Pull Request！

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

---

## 许可证

MIT License

---

**最后更新**: 2025-11-20  
**版本**: v2.0.0


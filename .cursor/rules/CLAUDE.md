## 项目架构规范
### 目录结构
```text
project-name/
├── cmd/                    # 应用程序入口
│   └── server/
│       └── main.go        # 主入口文件
├── internal/              # 私有应用程序代码
│   ├── config/           # 配置加载
│   ├── controller/       # HTTP 控制器层
│   ├── service/          # 业务逻辑层
│   ├── repository/       # 数据访问层
│   ├── model/            # 数据模型
│   ├── middleware/       # Gin 中间件
│   └── pkg/              # 内部共享包
├── pkg/                  # 公共库代码
│   ├── database/         # 数据库连接
│   ├── logger/           # 日志处理
│   ├── utils/            # 工具函数
│   └── response/         # 统一响应格式
├── api/                  # API 定义文件
│   └── docs/             # OpenAPI 文档
├── configs/              # 配置文件
├── deployments/          # 部署配置
├── scripts/              # 构建脚本
├── go.mod
└── go.sum
```

### 分层架构原则

- Controller层: 只处理 HTTP 请求和响应，不包含业务逻辑
- Service层: 业务逻辑核心，处理业务规则和流程
- Repository层: 数据访问，与数据库直接交互
- Model层: 数据模型定义

### 代码风格规范

#### . 命名约定
文件命名
使用 snake_case 命名文件

测试文件: user_spec.go

主要结构文件: service/user.go

#### 变量和函数命名
```go
// 使用驼峰命名法
var userService UserService
func getUserByID() {}
func calculateTotalPrice() {}

// 避免缩写，除非是公认的
var usrSrv UserService  // 不好
var userService UserService  // 好
```

#### 错误处理规范
##### 错误定义
```go
// pkg/errors/errors.go
package errors

import "fmt"

type BusinessError struct {
    Code    int
    Message string
    Cause   error
}

func (e *BusinessError) Error() string {
    if e.Cause != nil {
        return fmt.Sprintf("code: %d, message: %s, cause: %v", e.Code, e.Message, e.Cause)
    }
    return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func NewBusinessError(code int, message string) *BusinessError {
    return &BusinessError{
        Code:    code,
        Message: message,
    }
}

func WrapBusinessError(code int, message string, cause error) *BusinessError {
    return &BusinessError{
        Code:    code,
        Message: message,
        Cause:   cause,
    }
}

// 预定义业务错误
var (
    ErrInvalidParams = NewBusinessError(10001, "参数错误")
    ErrUserNotFound  = NewBusinessError(10002, "用户不存在")
    ErrUnauthorized  = NewBusinessError(10003, "未授权访问")
)
```

##### 错误处理模式
```go
// Service 层错误处理
func (s *userService) CreateUser(ctx context.Context, req *CreateUserRequest) (*model.User, error) {
    existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
    if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, WrapBusinessError(10010, "查询用户失败", err)
    }
    
    if existingUser != nil {
        return nil, ErrUserExists
    }
    
    // 业务逻辑...
}

// Controller 层错误处理
func (ctrl *UserController) CreateUser(c *gin.Context) {
    var req service.CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, 10001, "参数绑定失败: "+err.Error())
        return
    }
    
    user, err := ctrl.userService.CreateUser(c.Request.Context(), &req)
    if err != nil {
        handleServiceError(c, err)
        return
    }
    
    response.Success(c, user)
}

func handleServiceError(c *gin.Context, err error) {
    var bizErr *errors.BusinessError
    if errors.As(err, &bizErr) {
        response.Error(c, bizErr.Code, bizErr.Message)
    } else {
        // 记录未知错误日志
        logger.Error("未知错误", zap.Error(err))
        response.Error(c, 50000, "系统内部错误")
    }
}
```

### 日志规范
结构化日志
```go
// pkg/logger/logger.go
package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

func Init(env string) error {
    var config zap.Config
    
    if env == "production" {
        config = zap.NewProductionConfig()
    } else {
        config = zap.NewDevelopmentConfig()
        config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
    }
    
    logger, err := config.Build()
    if err != nil {
        return err
    }
    
    globalLogger = logger
    return nil
}

func Info(msg string, fields ...zap.Field) {
    globalLogger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
    globalLogger.Error(msg, fields...)
}

// 使用示例
logger.Info("用户创建成功", 
    zap.String("username", user.Username),
    zap.Uint("user_id", user.ID),
    zap.String("operation", "create_user"),
)
```

### 配置管理
配置结构
```go
// internal/config/config.go
package config

import (
    "github.com/spf13/viper"
    "go.uber.org/zap/zapcore"
)

type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    Redis    RedisConfig    `mapstructure:"redis"`
    JWT      JWTConfig      `mapstructure:"jwt"`
    Log      LogConfig      `mapstructure:"log"`
}

type ServerConfig struct {
    Port         string `mapstructure:"port"`
    Mode         string `mapstructure:"mode"`
    ReadTimeout  int    `mapstructure:"read_timeout"`
    WriteTimeout int    `mapstructure:"write_timeout"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
    SSLMode  string `mapstructure:"sslmode"`
}

type LogConfig struct {
    Level zapcore.Level `mapstructure:"level"`
    File  string        `mapstructure:"file"`
}
```
### API 响应规范
统一响应格式
```go
// pkg/response/response.go
package response

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
    TraceID string      `json:"trace_id,omitempty"`
}

// 成功响应
func Success(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, Response{
        Code:    0,
        Message: "success",
        Data:    data,
    })
}

// 分页响应
type PageResponse struct {
    List     interface{} `json:"list"`
    Total    int64       `json:"total"`
    Page     int         `json:"page"`
    PageSize int         `json:"page_size"`
}

func SuccessWithPage(c *gin.Context, list interface{}, total int64, page, pageSize int) {
    pageResponse := PageResponse{
        List:     list,
        Total:    total,
        Page:     page,
        PageSize: pageSize,
    }
    Success(c, pageResponse)
}

// 错误响应
func Error(c *gin.Context, code int, message string) {
    c.JSON(http.StatusOK, Response{
        Code:    code,
        Message: message,
        Data:    nil,
    })
}
```

### 数据库操作规范
```go
// internal/repository/base_repo.go
package repository

import (
    "context"
    "gorm.io/gorm"
)

type BaseRepository[T any] struct {
    db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
    return &BaseRepository[T]{db: db}
}

func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) error {
    return r.db.WithContext(ctx).Create(entity).Error
}

func (r *BaseRepository[T]) GetByID(ctx context.Context, id uint) (*T, error) {
    var entity T
    err := r.db.WithContext(ctx).First(&entity, id).Error
    if err != nil {
        return nil, err
    }
    return &entity, nil
}

func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
    return r.db.WithContext(ctx).Save(entity).Error
}

func (r *BaseRepository[T]) Delete(ctx context.Context, id uint) error {
    return r.db.WithContext(ctx).Delete(new(T), id).Error
}

// 具体 Repository 实现
type UserRepository interface {
    BaseRepository[model.User]
    GetByUsername(ctx context.Context, username string) (*model.User, error)
    GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepo struct {
    *BaseRepository[model.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepo{
        BaseRepository: NewBaseRepository[model.User](db),
    }
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
    var user model.User
    err := r.db.WithContext(ctx).Where("username = ?", username).First(&user).Error
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```

### 中间件规范
通用中间件
```go
// internal/middleware/logger.go
package middleware

import (
    "time"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        // 处理请求
        c.Next()
        
        // 记录日志
        latency := time.Since(start)
        logger.Info("HTTP请求",
            zap.String("method", c.Request.Method),
            zap.String("path", c.Request.URL.Path),
            zap.Int("status", c.Writer.Status()),
            zap.String("ip", c.ClientIP()),
            zap.Duration("latency", latency),
            zap.String("user_agent", c.Request.UserAgent()),
        )
    }
}

// internal/middleware/recovery.go
func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                // 记录 panic 日志
                logger.Error("HTTP Panic",
                    zap.Any("error", err),
                    zap.String("path", c.Request.URL.Path),
                    zap.String("method", c.Request.Method),
                )
                
                response.Error(c, 50000, "内部服务器错误")
                c.Abort()
            }
        }()
        c.Next()
    }
}
```
### 依赖注入规范
依赖注入容器
```go
// cmd/server/wire.go
//go:build wireinject
// +build wireinject

package main

import (
    "github.com/google/wire"
    
    "your-project/internal/config"
    "your-project/internal/controller"
    "your-project/internal/repository"
    "your-project/internal/service"
    "your-project/pkg/database"
    "your-project/pkg/logger"
)

func InitApp() (*App, error) {
    wire.Build(
        // 配置
        config.Load,
        
        // 基础设施
        database.NewDB,
        logger.Init,
        
        // Repository
        repository.NewUserRepository,
        
        // Service
        service.NewUserService,
        
        // Controller
        controller.NewUserController,
        
        // 应用
        NewApp,
    )
    return &App{}, nil
}
```
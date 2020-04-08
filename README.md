# gin-app-start
基于 [Gin](https://github.com/gin-gonic/gin) 封装的开发框架，目录结构清晰，包含了一些常用中间件，功能强大，使用简单。

持续更新... 


## 目录结构

- [x] 使用 go modules 初始化项目
- [x] 安装 Gin 框架
- [x] 支持优雅地重启或停止
- [x] 规划项目目录
- [x] 常用方法封装
  - [x] 返回数据format
  - [x] 参数验证（validator.v9）
    - [x] 模型绑定和验证
    - [] 自定义验证器
- [x] 路由中间件
    - [x] ip白名单
    - [] 日志记录
    - [] 异常捕获
    - [] 限流
- [ ] 存储
    - [ ] MySQL
    - [ ] Redis
    - [ ] MongoDB

## 快速开始

#### 代码仓库

```go
git clone git@github.com:pengfeidai/gin-app-start.git
```

#### 环境配置

- Go version >= 1.13
- Global environment configure

```go
export GO111MODULE=on
// 修改代理
export GOPROXY=https://goproxy.io
// go env -w GOPROXY=https://goproxy.cn,direct 
```

#### 服务启动

```go
cd gin-app-start

go run main.go

输出如下 `Listening and serving HTTP on :9060`，表示 Http Server 启动成功。
```

#### 健康检查

```
curl -X GET http://127.0.0.1:9060/health_check?name=world
```

## 文档



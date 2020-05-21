# gin-app-start
基于 [Gin](https://github.com/gin-gonic/gin) 封装的开发框架，目录结构清晰，包含了一些常用中间件，功能强大，使用简单。

持续更新... 


## 目录结构

该目录结构是结合自己之前写 `nodejs` 的工作经验来的，其中也是参考了 `egg.js` 的目录结构，自认为结构还是比较清晰，一目了然，也欢迎大家提出好的建议，共同学习。

```go
gin-app-strat
├── app
|    ├── common // constant定义、errorCode等
|    ├── config // 定义config的struct，路由加载
|    |     └── config.go
|    ├── controller
|    ├── database // mysql、mongo、redis等
|    ├── middleware // 中间件
|    ├── model
|    ├── public // 静态资源
|    ├── router // 路由分组、路由聚合加载
|    ├── schema // 定义controller中使用的struct，用于参数绑定、校验
|    ├── service 
|    ├── util  // 常用方法
├── logs
|   ├── access.log
|   └── error.log
├── test  // 测试用例
|   ├── controller
|   ├── service
|   └── model
├── .gitignore
├── config.yaml  // 配置文件
├── Dockerfile
├── go.mod
├── go.sum
├── main.go // 启动文件
└── README.md
```

## 功能点
- [x] 使用 go modules 初始化项目
- [x] 安装 Gin 框架
- [x] 支持优雅地重启或停止
- [x] 规划项目目录
- [x] 路由
    - [x] 分组、聚合
    - [x] 路由中间件
        - [x] 参数校验
          - [x] 模型绑定和验证
          - [] 自定义验证器
        - [x] 返回数据format
        - [x] [日志记录](https://github.com/sirupsen/logrus)
          - [x] 日志按天分割，暂时保存到文件
          - [x] 输出流
        - [x] 异常捕获
        - [x] 限流
- [x] [Session](https://github.com/gin-contrib/sessions)
    - [x] cookie-based
    - [x] Redis
- [x] 存储
    - [x] [MySQL](https://github.com/jinzhu/gorm)
      - [x] AutoMigrate
    - [x] [Redis](https://github.com/go-redis/redis)
    - [x] [MongoDB](https://www.godoc.org/gopkg.in/mgo.v2)
    - [x] [ElasticSearch](https://github.com/olivere/elastic)
- [x] 常用中间件
  - [x] 常用方法common.go，如：uuid
  - [x] 发送邮件
- [x] 线上部署
   - [x] dockerfile
   - [x] 自定义配置文件

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

输出如下 `Listening and serving HTTP on Port: :9060, Pid: 15932`，表示 Http Server 启动成功。
```

#### 健康检查

```
curl -X GET http://127.0.0.1:9060/health_check?name=world
```

## 文档



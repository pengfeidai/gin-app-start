# Build stage
FROM golang:1.24-alpine AS builder

# 安装必要的工具
RUN apk add --no-cache git

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 编译应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# Runtime stage
FROM alpine:latest

# 安装 ca-certificates（用于 HTTPS 请求）
RUN apk --no-cache add ca-certificates

# 设置工作目录
WORKDIR /root/

# 从构建阶段复制编译后的二进制文件
COPY --from=builder /app/main .

# 复制配置文件
COPY --from=builder /app/configs ./configs

# 创建日志目录
RUN mkdir -p /var/log/gin-app

# 暴露端口
EXPOSE 9060

# 运行应用
CMD ["./main"]

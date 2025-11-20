.PHONY: run build test clean docker-build docker-run

# 运行应用
run:
	SERVER_ENV=local go run cmd/server/main.go

# 编译应用
build:
	go build -o bin/server cmd/server/main.go

# 运行测试
test:
	go test -v ./...

# 清理编译文件
clean:
	rm -rf bin/
	go clean

# 安装依赖
deps:
	go mod download
	go mod tidy

# 格式化代码
fmt:
	go fmt ./...

# 代码检查
lint:
	golangci-lint run

# 构建 Docker 镜像
docker-build:
	docker build -t gin-app-start:latest .

# 运行 Docker 容器
docker-run:
	docker run -d \
		-p 9060:9060 \
		-e SERVER_ENV=local \
		--name gin-app-start \
		gin-app-start:latest

# 停止 Docker 容器
docker-stop:
	docker stop gin-app-start
	docker rm gin-app-start

# 数据库迁移
migrate:
	SERVER_ENV=local go run cmd/server/main.go -migrate

# 安装开发工具
install-tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest

# 生成 Swagger 文档
swagger:
	swag init -g cmd/server/main.go -o docs --parseDependency --parseInternal

# 格式化 Swagger 注释
swagger-fmt:
	swag fmt

.DEFAULT_GOAL := run

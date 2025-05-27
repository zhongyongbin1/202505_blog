.PHONY: dev build run clean test

# 默认目标
all: dev

# 开发模式（热更新）
dev:
	air

# 构建应用
build:
	go build -o ./bin/app ./cmd/main.go

# 运行应用
run:
	go run ./cmd/main.go

# 清理临时文件
clean:
	rm -rf ./tmp ./bin

# 运行测试
test:
	go test -v ./...
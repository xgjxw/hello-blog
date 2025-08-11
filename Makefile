# 变量定义
PROJECT_NAME=$(shell basename "$(PWD)")
PROJECT_PATH=$(shell cd "$(dirname "$0" )" &&pwd)
DESTDIR=${PROJECT_PATH}/build
MAIN_DIR=cmd/blog

# 获取当前时间
BUILD_TIME=$(shell date +%Y%m%d_%H%M%S)
# 获取 git commit hash
GIT_COMMIT=$(shell git rev-parse --short HEAD)
# 获取当前版本
VERSION=$(shell git describe --tags | sed 's/\(.*\)-.*/\1/')

.PHONY: run build clean

# 创建build目录（如果不存在）
build-dir:
	@mkdir -p $(DESTDIR)

# 运行应用
run:
	-go run $(MAIN_DIR)/main.go

# 构建应用
build: build-dir
	go build -o "$(DESTDIR)/$(PROJECT_NAME)" "$(MAIN_DIR)/main.go"

# 清理构建文件
clean:
	rm -rf $(DESTDIR)/*

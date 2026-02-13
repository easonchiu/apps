.PHONY: build dev clean install web-dev web-build run push start start-bg stop

# 安装依赖
install:
	cd web && npm install

# 开发模式 - 只运行前端开发服务器
web-dev:
	cd web && npm run dev

# 构建前端
web-build:
	cd web && npm run build

# 开发模式 - 运行 Go 服务器（需要先构建前端）
dev: web-build
	go run main.go

# 生产构建（原有的构建命令，添加前端构建）
build: web-build
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./ysgame -tags=jsoniter -trimpath -ldflags '-s -w'

# 部署到服务器（保留原有命令）
push: build
	ssh ys 'sudo systemctl stop ysgame'
	scp ./ysgame ys:/usr/go
	ssh ys 'sudo systemctl restart ysgame && sudo systemctl status ysgame'
	rm -f ./ysgame

# 本地运行
run:
	./ysgame

# 清理构建产物
clean:
	rm -rf web/node_modules
	rm -rf ui/dist
	rm -rf dist
	rm -rf tmp
	rm -f ./ysgame

# 完整构建流程（从零开始）
all: clean install build

# 快速启动（前台运行，用于开发）
start: web-build
	@echo "启动服务器（前台运行，按 Ctrl+C 停止）..."
	go run main.go

# 后台启动（用于长期运行）
start-bg: web-build
	@echo "启动服务器（后台运行）..."
	@nohup go run main.go > /tmp/ysgame.log 2>&1 & echo $$! > /tmp/ysgame.pid
	@echo "服务器已在后台启动，PID: $$(cat /tmp/ysgame.pid)"
	@echo "日志文件: /tmp/ysgame.log"
	@echo "停止服务器: make stop"

# 停止后台服务器
stop:
	@if [ -f /tmp/ysgame.pid ]; then \
		PID=$$(cat /tmp/ysgame.pid); \
		if ps -p $$PID > /dev/null 2>&1; then \
			kill $$PID && echo "已停止服务器 (PID: $$PID)"; \
		else \
			echo "服务器进程不存在"; \
		fi; \
		rm -f /tmp/ysgame.pid; \
	else \
		echo "未找到 PID 文件，尝试查找并停止所有相关进程..."; \
		pkill -f "go run main.go" || pkill main || echo "未找到运行中的服务器"; \
	fi

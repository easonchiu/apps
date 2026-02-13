# 快速开始指南

## 项目已迁移到 React

本项目已从 Go Template 迁移到 React + Vite 实现。以下是快速开始的步骤。

## 首次设置

### 1. 安装依赖

```bash
make install
```

或者：

```bash
cd web && npm install
```

## 开发模式

### 方式一：前端独立开发（推荐）

适合只修改前端代码的情况：

```bash
make web-dev
```

然后访问 http://localhost:5173

### 方式二：完整开发模式

适合需要测试前后端集成的情况：

```bash
make start
```

然后访问 http://localhost:9090

## 生产构建

### 构建应用

```bash
make build
```

这会：
1. 构建 React 前端（输出到 `ui/dist/`）
2. 编译 Go 应用（包含嵌入的前端资源）

### 运行应用

```bash
./ysgame
```

### 部署到服务器

```bash
make push
```

这会自动：
1. 构建前端和后端
2. 停止服务器上的服务
3. 上传新的二进制文件
4. 重启服务

## 常用命令

| 命令 | 说明 |
|------|------|
| `make install` | 安装前端依赖 |
| `make web-dev` | 启动前端开发服务器 |
| `make web-build` | 只构建前端 |
| `make start` | 构建前端并启动 Go 服务器 |
| `make build` | 完整构建（生产环境） |
| `make push` | 部署到服务器 |
| `make clean` | 清理所有构建产物 |

## 项目结构

```
apps/
├── web/                 # React 前端项目
│   ├── src/            # 源代码
│   ├── public/         # 静态资源
│   └── package.json
├── ui/
│   ├── dist/           # 前端构建产物（Git 忽略）
│   └── assets.go       # Go embed 配置
├── routers/            # Go 路由
├── main.go             # Go 入口
└── makefile            # 构建脚本
```

## 修改内容

### 修改前端代码

1. 编辑 `web/src/` 下的文件
2. 如果在开发模式（`make web-dev`），会自动热重载
3. 如果需要测试生产构建，运行 `make web-build`

### 修改游戏信息

编辑 `web/src/types/game.ts` 文件。

### 修改样式

编辑 `web/src/styles/` 下的 CSS 文件。

### 修改 Go 后端

1. 编辑 Go 代码
2. 运行 `make start` 重启服务器

## 故障排除

### 问题：前端修改没有生效

**解决方案：**
- 如果在开发模式，检查 Vite 开发服务器是否在运行
- 如果在生产模式，需要重新运行 `make build`

### 问题：Go 编译失败，提示找不到 dist 目录

**解决方案：**
```bash
make web-build
```

### 问题：npm install 失败

**解决方案：**
```bash
cd web
rm -rf node_modules package-lock.json
npm install
```

### 问题：端口被占用

**解决方案：**
- 前端开发服务器：修改 `web/vite.config.ts` 中的端口
- Go 服务器：修改 `main.go` 中的端口（默认 9090）

## 更多信息

- 详细迁移说明：查看 `PROJECT_MIGRATION.md`
- 前端开发文档：查看 `web/README.md`
- Go 模块文档：查看各目录下的 README

## 技术栈

- **前端：** React 19 + TypeScript + Vite + React Router
- **后端：** Go + Gin
- **构建：** Vite + Go embed
- **部署：** 单一二进制文件（包含所有前端资源）

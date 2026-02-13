# 项目迁移说明

本项目已从 Go Template 迁移到 React + Vite 实现。

## 迁移内容

### 1. 前端技术栈变更

**之前：**
- Go HTML Templates (`ui/templates/*.html`)
- 服务端渲染

**现在：**
- React 19 + TypeScript
- Vite 构建工具
- 客户端路由 (React Router)
- 单页应用 (SPA)

### 2. 项目结构

```
apps/
├── web/                    # React 前端项目
│   ├── src/
│   │   ├── components/    # React 组件
│   │   ├── pages/         # 页面组件
│   │   ├── styles/        # CSS 样式
│   │   └── types/         # TypeScript 类型
│   ├── public/            # 静态资源
│   └── package.json
├── ui/
│   ├── dist/              # React 构建产物（由 embed 加载）
│   └── assets.go          # Go embed 配置
├── routers/
│   └── routers.go         # 路由配置（简化为 SPA 服务）
├── controllers/           # 已废弃，不再使用
└── main.go
```

### 3. 开发流程

#### 开发模式

1. **前端开发：**
   ```bash
   cd web
   npm install
   npm run dev
   ```
   访问 http://localhost:5173

2. **后端开发：**
   ```bash
   # 先构建前端
   cd web
   npm run build
   
   # 再运行 Go 服务
   cd ..
   go run main.go
   ```
   访问 http://localhost:9090

#### 生产构建

```bash
# 1. 构建前端
cd web
npm run build

# 2. 构建 Go 应用（会自动包含前端构建产物）
cd ..
go build -o dist/app

# 3. 运行
./dist/app
```

### 4. 路由处理

**之前：**
- Go 路由器处理每个页面
- 服务端渲染 HTML

**现在：**
- Go 只处理特殊路由（robots.txt, sitemap.xml）
- 所有其他路由返回 index.html
- React Router 处理客户端路由

### 5. 数据管理

**之前：**
- 数据通过 Go 模板传递
- 每个页面独立渲染

**现在：**
- 游戏数据定义在 `web/src/types/game.ts`
- 客户端 JavaScript 对象
- 可以轻松扩展为 API 调用

### 6. 样式处理

**之前：**
- CSS 文件在 `ui/static/styles/`
- 通过 `<link>` 标签引入

**现在：**
- CSS 文件在 `web/src/styles/`
- 通过 Vite 打包
- 支持 CSS 模块化和代码分割

### 7. 静态资源

**之前：**
- 图片在 `ui/static/images/`
- 通过 `/static/images/` 访问

**现在：**
- 图片在 `web/public/images/`
- 构建后在 `ui/dist/images/`
- 通过 `/images/` 访问

### 8. 已废弃的文件

以下文件/目录不再使用，可以保留作为参考或删除：

- `ui/templates/` - 所有 Go 模板文件
- `ui/static/` - 旧的静态资源目录
- `controllers/controllers.go` - 控制器逻辑（已移至 React）
- `models/` - 如果只用于前端数据

### 9. SEO 考虑

由于现在是 SPA，需要注意：

1. **robots.txt 和 sitemap.xml** 仍由 Go 处理
2. **元标签** 在 `web/index.html` 中设置
3. 如需更好的 SEO，可以考虑：
   - 服务端渲染 (SSR)
   - 预渲染 (Prerendering)
   - 或保持当前方案（适合 App 下载页面）

### 10. 部署注意事项

1. **构建顺序：** 必须先构建前端，再构建 Go 应用
2. **embed 限制：** `ui/dist` 目录必须存在且包含构建产物
3. **单一二进制：** 最终只需部署一个 Go 二进制文件
4. **环境变量：** 可以通过环境变量配置端口等

### 11. 添加新功能

#### 添加新页面

1. 在 `web/src/pages/` 创建新组件
2. 在 `web/src/App.tsx` 添加路由
3. 在 Header/Footer 添加导航链接

#### 添加新游戏

编辑 `web/src/types/game.ts`，添加游戏数据。

#### 修改样式

编辑 `web/src/styles/` 中的 CSS 文件。

### 12. 常见问题

**Q: 修改前端代码后，Go 应用没有更新？**
A: 需要重新运行 `npm run build` 并重新编译 Go 应用。

**Q: 开发时如何同时运行前后端？**
A: 前端使用 `npm run dev`，后端使用 `go run main.go`。前端开发时访问 Vite 开发服务器，后端测试时先构建前端。

**Q: 如何调试生产构建？**
A: 可以在本地运行 `npm run build` 和 `go run main.go`，访问 http://localhost:9090 测试。

**Q: 构建失败怎么办？**
A: 
1. 检查 Node.js 版本（建议 18+）
2. 删除 `node_modules` 和 `package-lock.json`，重新安装
3. 检查 TypeScript 错误
4. 确保 `ui/dist` 目录存在

## 迁移完成清单

- [x] 安装 React Router 和必要依赖
- [x] 创建 React 组件结构（Layout, Header, Footer）
- [x] 创建各页面组件（Home, Games, Privacy, Contact, News, 404）
- [x] 转移 CSS 样式到 React 项目
- [x] 转移静态资源到 React public 目录
- [x] 配置 Vite 构建输出目录
- [x] 更新 Go 代码使用 embed 加载 React 构建产物
- [x] 测试构建流程
- [x] 编写文档

## 后续优化建议

1. **性能优化：**
   - 添加图片懒加载
   - 代码分割优化
   - 添加 Service Worker（PWA）

2. **开发体验：**
   - 添加 ESLint 和 Prettier 配置
   - 添加 Git hooks（pre-commit）
   - 添加 CI/CD 流程

3. **功能增强：**
   - 添加多语言支持（i18n）
   - 添加深色模式
   - 添加分析统计

4. **SEO 优化：**
   - 考虑 SSR 或预渲染
   - 添加结构化数据
   - 优化元标签

5. **测试：**
   - 添加单元测试（Vitest）
   - 添加 E2E 测试（Playwright）

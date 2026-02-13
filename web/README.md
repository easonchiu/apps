# YSGame Studio Web Frontend

这是 YSGame Studio 的前端项目，使用 React + TypeScript + Vite 构建。

## 开发

### 安装依赖

```bash
npm install
```

### 开发模式

```bash
npm run dev
```

访问 http://localhost:5173 查看开发服务器。

### 构建生产版本

```bash
npm run build
```

构建产物会输出到 `../ui/dist` 目录，供 Go 应用的 embed 使用。

## 项目结构

```
src/
├── components/          # 可复用组件
│   ├── Header.tsx      # 页头组件
│   ├── Footer.tsx      # 页脚组件
│   ├── Layout.tsx      # 布局组件
│   └── DownloadButton.tsx  # 下载按钮组件
├── pages/              # 页面组件
│   ├── Home.tsx        # 首页
│   ├── GameDetail.tsx  # 游戏详情页
│   ├── Privacy.tsx     # 隐私政策页
│   ├── Contact.tsx     # 联系我们页
│   ├── News.tsx        # 新闻页
│   └── NotFound.tsx    # 404 页面
├── styles/             # 样式文件
│   ├── common.css      # 通用样式
│   ├── games.css       # 游戏相关样式
│   ├── contact.css     # 联系页样式
│   ├── privacy.css     # 隐私政策样式
│   └── news.css        # 新闻页样式
├── types/              # TypeScript 类型定义
│   └── game.ts         # 游戏数据类型和数据
├── App.tsx             # 应用主组件
├── main.tsx            # 应用入口
└── index.css           # 全局样式重置
```

## 路由

- `/` - 首页（游戏列表）
- `/games/:gameId` - 游戏详情页
  - `/games/sudoku-crown` - Sudoku Crown
  - `/games/block-cuties` - Block Cuties
  - `/games/digit-merge` - Digit Merge
- `/privacy` - 隐私政策
- `/contact` - 联系我们
- `/news` - 新闻（暂未启用）

## 技术栈

- React 19
- TypeScript
- React Router DOM
- Vite
- CSS (原生 CSS 变量)

## 与 Go 后端集成

构建后的静态文件会被 Go 的 `embed` 功能嵌入到二进制文件中。Go 应用会：

1. 使用 `embed.FS` 加载 `ui/dist` 目录
2. 通过 Gin 框架提供静态文件服务
3. 对于所有非 API 路由，返回 `index.html`，由 React Router 处理客户端路由

## 添加新游戏

要添加新游戏，编辑 `src/types/game.ts` 文件，在 `games` 对象中添加新的游戏数据：

```typescript
'new-game-id': {
  id: 'new-game-id',
  name: 'Game Name',
  description: 'Game description',
  icon: 'icon-url',
  images: ['image1-url', 'image2-url'],
  appStoreUrl: 'app-store-url',
  content: ['paragraph1', 'paragraph2']
}
```

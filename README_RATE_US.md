# Rate Us 功能文档

## 功能概述

用户反馈（Rate Us）功能允许用户通过网页表单提交反馈，数据会存储到 MongoDB 数据库中。

## 前端部分

### 页面路径
- URL: `/rate-us`
- 组件位置: `web/src/pages/rate-us/`

### 功能特性
- Email 输入框（必填，带邮箱格式验证）
- 内容输入框（Textarea，必填）
- 表单验证
- 加载状态显示
- 成功/错误消息提示

### 使用示例
访问 `http://localhost:9090/rate-us` 即可看到反馈表单。

## 后端部分

### API 接口

**提交反馈**
```
POST /api/feedback
Content-Type: application/json

Request Body:
{
  "email": "user@example.com",
  "content": "用户反馈内容"
}

Success Response (200):
{
  "message": "Feedback submitted successfully",
  "id": "698f8148d7473168cf3b3923"
}

Error Response (400):
{
  "error": "Invalid request data"
}

Error Response (500):
{
  "error": "Failed to save feedback"
}
```

### 代码结构

```
├── models/
│   └── feedback.go         # 反馈数据模型
├── controllers/
│   └── feedback.go         # 反馈控制器
├── db/
│   └── mongodb.go          # MongoDB 连接配置
└── routers/
    └── routers.go          # 路由注册
```

### 数据库配置

MongoDB 连接可以通过环境变量配置：

```bash
# .env 文件
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=ysgame
```

如果未设置环境变量，将使用默认值：
- URI: `mongodb://localhost:27017`
- Database: `ysgame`

### 数据存储

反馈数据存储在 MongoDB 的 `feedbacks` 集合中，包含以下字段：

```json
{
  "_id": "ObjectId",
  "email": "用户邮箱",
  "content": "反馈内容",
  "created_at": "创建时间（ISO 8601格式）"
}
```

## 测试

### 启动服务

```bash
# 构建前端
cd web && npm run build

# 启动服务器
cd .. && go run main.go
```

或使用 Makefile:

```bash
make start
```

### API 测试

```bash
# 成功案例
curl -X POST http://localhost:9090/api/feedback \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","content":"Great app!"}'

# 预期响应：
# {"message":"Feedback submitted successfully","id":"..."}
```

### 前端测试

1. 访问 http://localhost:9090/rate-us
2. 填写 Email 和反馈内容
3. 点击 "Submit Feedback" 按钮
4. 查看成功或错误提示消息

## 注意事项

1. **MongoDB 连接**: 如果 MongoDB 未运行，应用仍会启动，但会在日志中显示警告，且 `/api/feedback` 接口会返回 500 错误。

2. **数据验证**: 
   - Email 必须是有效的邮箱格式
   - Content 不能为空
   - 两个字段都是必填项

3. **错误处理**: 应用包含完整的错误处理机制，包括：
   - 输入验证错误
   - 数据库连接错误
   - 数据插入错误

4. **无入口设计**: 该页面没有在导航栏中显示，需要直接访问 `/rate-us` URL。

## 未来改进建议

- [ ] 添加反馈管理后台界面
- [ ] 添加反馈评分（星级）功能
- [ ] 添加文件上传（截图）功能
- [ ] 添加邮件通知功能
- [ ] 添加反馈列表查询 API
- [ ] 添加反馈数据导出功能

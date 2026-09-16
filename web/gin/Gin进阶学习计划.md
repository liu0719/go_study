# Gin进阶学习计划

## 学习说明
本文档列出了在掌握Gin基础后需要进阶学习的知识点，每个知识点都包含简要说明和学习资源链接。建议按顺序学习，每掌握一个知识点就创建对应的示例代码。

---

## 1. CORS中间件
**说明**：跨域资源共享是Web开发中必不可少的功能，允许前端应用访问不同域的后端API。

### 学习资源：
- [Gin CORS中间件官方文档](https://github.com/gin-contrib/cors)
- [MDN - CORS详解](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/CORS)

### 关键知识点：
- 配置允许的源、方法、头部
- 预检请求处理
- 凭据（Credentials）支持

### 实践建议：
创建`15.CORS中间件/`目录，实现CORS配置示例

---

## 2. Cookie和Session处理
**说明**：Web应用中常用的用户状态管理方式，用于保持用户登录状态等。

### 学习资源：
- [Go Cookie官方文档](https://pkg.go.dev/net/http#Cookie)
- [Gin Cookie使用教程](https://gin-gonic.com/zh-cn/docs/examples/cookie/)

### 关键知识点：
- Cookie的读写、过期时间、域名、路径设置
- Session中间件的选择和使用
- 安全Cookie设置

### 实践建议：
创建`16.Cookie和Session/`目录，实现用户登录状态管理示例

---

## 3. JWT认证
**说明**：现代Web应用最常用的身份验证方式，基于Token的无状态认证。

### 学习资源：
- [JWT.io官方网站](https://jwt.io/)
- [Golang JWT库](https://github.com/golang-jwt/jwt)
- [JWT认证教程](https://gin-gonic.com/zh-cn/docs/examples/jwt-authentication/)

### 关键知识点：
- JWT结构（Header, Payload, Signature）
- Token生成和验证
- 中间件集成
- 刷新机制

### 实践建议：
创建`17.JWT认证/`目录，实现完整的JWT认证系统

---

## 4. 文件上传（进阶）
**说明**：更完整的文件上传功能，包括文件类型验证、大小限制、文件重命名等。

### 学习资源：
- [Gin文件上传文档](https://gin-gonic.com/zh-cn/docs/examples/multipart-upload/)
- [文件类型验证](https://github.com/gabriel-vasile/mimetype)

### 关键知识点：
- 文件类型验证
- 文件大小限制
- 文件存储策略
- 文件名处理（防止冲突、安全过滤）
- 批量上传

### 实践建议：
创建`18.文件上传进阶/`目录，实现完整的文件上传系统

---

## 5. 错误处理和恢复
**说明**：构建健壮的应用，需要完善的错误处理和panic恢复机制。

### 学习资源：
- [Go错误处理最佳实践](https://go.dev/blog/error-handling-and-go)
- [Gin中间件开发](https://gin-gonic.com/zh-cn/docs/examples/custom-middleware/)

### 关键知识点：
- 自定义错误响应格式
- Panic捕获和恢复
- 错误日志记录
- 错误码标准化

### 实践建议：
创建`19.错误处理/`目录，实现全局错误处理中间件

---

## 6. 日志记录
**说明**：结构化的日志记录对调试和监控至关重要。

### 学习资源：
- [Zap日志库](https://go.uber.org/zap)
- [Logrus日志库](https://github.com/sirupsen/logrus)
- [Gin日志中间件](https://github.com/gin-gonic/gin/tree/master/examples/basic/logger)

### 关键知识点：
- 日志格式配置
- 不同级别的日志（Debug, Info, Warn, Error）
- 请求日志中间件
- 日志输出到文件

### 实践建议：
创建`20.日志记录/`目录，集成结构化日志到Gin应用

---

## 7. 请求限流
**说明**：防止API被滥用，保护服务器资源。

### 学习资源：
- [Gin Rate Limiter中间件](https://github.com/ulule/limiter/v3)
- [限流算法介绍](https://github.com/ulule/limiter#strategies)

### 关键知识点：
- 限流策略（固定窗口、滑动窗口、令牌桶）
- 基于IP或用户的限流
- 限流响应处理

### 实践建议：
创建`21.请求限流/`目录，实现API限流中间件

---

## 8. 配置管理
**说明**：使用配置文件管理应用设置，支持不同环境。

### 学习资源：
- [Viper配置库](https://github.com/spf13/viper)
- [环境变量处理](https://pkg.go.dev/os)

### 关键知识点：
- 配置文件格式（JSON, YAML, TOML）
- 环境变量支持
- 配置热重载
- 不同环境的配置管理

### 实践建议：
创建`22.配置管理/`目录，实现配置管理系统

---

## 9. 数据库集成
**说明**：与数据库的集成，包括ORM和原生SQL操作。

### 学习资源：
- [GORM](https://gorm.io/zh_CN/)
- [Go SQL驱动](https://pkg.go.dev/database/sql)
- [Gin + GORM教程](https://gorm.io/zh_CN/docs/gin.html)

### 关键知识点：
- 数据库连接池配置
- 模型定义和迁移
- CRUD操作
- 事务处理
- 查询优化

### 实践建议：
创建`23.数据库集成/`目录，实现完整的CRUD操作

---

## 10. WebSocket支持
**说明**：实现实时通信功能，如聊天室、实时通知等。

### 学习资源：
- [Gin WebSocket示例](https://github.com/gin-gonic/examples/tree/master/websocket)
- [Go WebSocket库](https://github.com/gorilla/websocket)

### 关键知识点：
- WebSocket连接建立
- 消息发送和接收
- 广播和房间机制
- 连接管理

### 实践建议：
创建`24.WebSocket/`目录，实现实时聊天应用

---

## 11. API文档自动生成
**说明**：使用Swagger/OpenAPI自动生成API文档。

### 学习资源：
- [Swaggo/Swagger](https://github.com/swaggo/swagger)
- [Swaggo Gin中间件](https://github.com/swaggo/gin-swagger)

### 关键知识点：
- 注解方式定义API
- 自动生成文档
- Swagger UI集成

### 实践建议：
创建`25.API文档/`目录，实现自动生成的API文档

---

## 12. 单元测试
**说明**：编写单元测试确保代码质量。

### 学习资源：
- [Go Testing官方文档](https://go.dev/doc/tutorial/testing-and-benchmarking)
- [Gin测试工具](https://github.com/gin-gonic/examples/tree/master/test)

### 关键知识点：
- 单元测试编写
- 表驱动测试
- HTTP测试
- Mock测试

### 实践建议：
创建`26.单元测试/`目录，为之前的示例添加测试用例

---

## 13. 优雅关闭
**说明**：确保应用能够安全关闭，处理未完成的请求。

### 学习资源：
- [Go信号处理](https://pkg.go.dev/signal)
- [Gin优雅关闭示例](https://github.com/gin-gonic/examples/tree/master/graceful-shutdown)

### 关键知识点：
- 信号捕获
- 请求超时设置
- 资源清理
- 优雅等待

### 实践建议：
创建`27.优雅关闭/`目录，实现优雅关闭功能

---

## 14. 监控和健康检查
**说明**：应用健康状态检查和监控指标收集。

### 学习资源：
- [Prometheus Go客户端](https://github.com/prometheus/client_golang)
- [健康检查端点设计](https://cloud.google.com/run/docs/tutorials/health-check)

### 关键知识点：
- /health端点实现
- 指标收集
- 服务状态检查
- 监控数据可视化

### 实践建议：
创建`28.监控健康/`目录，实现健康检查和监控

---

## 学习路线图

### 初级进阶（1-5）
1. CORS中间件
2. Cookie和Session
3. JWT认证
4. 文件上传进阶
5. 错误处理和恢复

### 中级进阶（6-10）
6. 日志记录
7. 请求限流
8. 配置管理
9. 数据库集成
10. WebSocket支持

### 高级进阶（11-14）
11. API文档自动生成
12. 单元测试
13. 优雅关闭
14. 监控和健康检查

## 学习建议

1. **循序渐进**：按照顺序学习，掌握一个再学习下一个
2. **动手实践**：每个知识点都要亲手编写代码
3. **阅读源码**：阅读相关库的源码，理解实现原理
4. **做笔记**：记录关键点和遇到的问题
5. **项目实践**：学完后尝试构建一个完整的项目

## 常用工具库

### 必备工具
```bash
# JWT
go get github.com/golang-jwt/jwt/v5

# CORS
go get github.com/gin-contrib/cors

# 文件类型检测
go get github.com/gabriel-vasile/mimetype

# 日志
go get go.uber.org/zap

# 配置管理
go get github.com/spf13/viper

# ORM
go get gorm.io/gorm

# WebSocket
go get github.com/gorilla/websocket

# API文档
go get github.com/swaggo/swagger
go get github.com/swaggo/gin-swagger
go get github.com/swaggo/files

# 限流
go get github.com/ulule/limiter/v3

# 监控
go get github.com/prometheus/client_golang
```

### 可选工具
```bash
# 缓存
go get github.com/patrickmn/go-cache

# 任务队列
go get github.com/go-redis/redis/v8

# 验证器
go get github.com/go-playground/validator/v10

# 环境变量管理
go get github.com/spf13/pflag
```
💪
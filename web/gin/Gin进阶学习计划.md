# Gin 进阶学习计划

> **目标**：后端实习（研一暑假，2027 年 7–8 月投递窗口）
> **起点**：01–15 基础已过，CORS 已在 `15.跨域问题/` 完成
> **本版调整**：原 14 项压缩为 8 项核心 + 1 项贯穿。原清单是按"库"列的，学完会认识一堆中间件但没有一个能交付的服务。本版改为**从 17 起，所有知识点都长进同一个项目**。

---

## 进度总览

| 阶段  | 内容                     | 目录                   | 预计耗时     |
| --- | ---------------------- | -------------------- | -------- |
| 0   | 工程基建（必须先做，半天）          | —                    | 0.5 天    |
| 1-1 | Cookie 和 Session       | `16.cookie和session/` | 2 天      |
| 1-2 | **JWT 认证**（从此进入项目骨架）   | `17.jwt认证/`          | 2 天      |
| 1-3 | 错误处理和恢复                | 项目内                  | 1 天      |
| 1-4 | 日志记录                   | 项目内                  | 1 天      |
| 1-5 | 请求限流                   | 项目内                  | 0.5 天    |
| 1-6 | 配置管理                   | 项目内                  | 0.5 天    |
| 1-7 | **数据库集成**（含事务）         | `22.数据库集成/`          | 3 天      |
| 1-8 | 优雅关闭                   | 项目内                  | 0.5 天    |
| 贯穿  | 单元测试（每节顺手写，不单开）        | 各处 `_test.go`        | 每节 30 分钟 |
| 2   | MySQL → Redis → Docker | 另起计划                 | 6 周      |
| 3   | 算法（并行，每天一道）            | —                    | 6 个月     |

---

## 阶段 0：工程基建（必须先做，半天）

这一步不做，后面所有东西都验不了。

### 0.1 中文目录名让 `go build ./...` 全部失效
已验证：`go build ./...` 对当前 16 个目录全部报 `malformed import path "github.com/liu0719/go_study/web/gin/01.原生_http库": invalid char '原'`。Go 把目录名当成 import path 的一部分，必须 ASCII。

**后果**：`go vet`、`go test ./...`、`golangci-lint`、CI 全部不可用。现在能跑只是因为一直 `go run 单个文件.go`。

**改法**：目录名改 ASCII，中文标题放进本文件（就是上面的总览表）。例如 `16.cookie和session/`、`17.jwt认证/` → `16-cookie-session/`、`17-jwt/`。已完成的 01–15 可以以后有空再改，不急。

### 0.2 `go mod tidy`
`go.mod` 里所有依赖（包括 gin 本身）都被标成 `// indirect`，还有 `mimetype`、`mongo-driver`、`quic-go` 这些代码里没用的残留。跑一次 `go mod tidy` 清掉。

### 0.3 项目骨架（阶段 1 的第一件事）
从 17 起不再写独立 `main` 脚本，建一个分层骨架：

```
ginapp/
├── cmd/server/main.go      // 只负责启动和优雅关闭
├── internal/
│   ├── router/router.go    // 路由注册，不写业务
│   ├── handler/            // 只解析请求和组装响应，不碰业务
│   ├── service/            // 业务逻辑
│   └── middleware/         // JWT、限流、日志、CORS、错误恢复
├── configs/config.yaml
└── pkg/response/           # 从现有 res/enter.go 挪过来
```

16 还是单文件练手没问题，17 开始往这个骨架里长。

---

## 阶段 1：Gin 进阶（8 项核心）

---

### 1. Cookie 和 Session 处理
**说明**：理解凭据的两条路线——服务端查表（有状态）vs 凭据自带签名（无状态）。这一节是 17 的地基。

### 学习资源：
- [Go Cookie 官方文档](https://pkg.go.dev/net/http#Cookie)
- [Gin Cookie 使用教程](https://gin-gonic.com/zh-cn/docs/examples/cookie/)
- [RFC 6265（Cookie 规范）](https://httpwg.org/specs/rfc6265.html)
- 现有材料：[`../../教学/16.cookie和session处理/Cookie和Session处理.go`](../../教学/16.cookie和session处理/Cookie和Session处理.go)（三种方案对比 + 手写签名 token + 双实例丢登录态复现，可直接用，注意它在 go.mod 外面，`go run` 单文件运行）

### 关键知识点：
- `Set-Cookie` 每个属性都是考点：`Path` / `MaxAge` / `Domain` / `HttpOnly` / `Secure` / `SameSite`
- `MaxAge` 三种取值语义完全不同：`-1` 立即删除、`0` 会话 cookie、正数 持久化
- Session 的本质：cookie 里只放无意义的 sessionId，真数据在服务端
- 三种方案对比：裸 Cookie（不安全）/ Session（能踢人，多实例要共享存储）/ 签名 Token（无状态，踢不掉）

### 验收标准：
能讲清"为什么本地 `Secure=true` 会写了带不回来"、"为什么跨站请求带不上 cookie"、"多实例下 Session 为什么会随机丢登录态"。三个都是能复现的坑，不是背结论。

### 实践建议：
把三种方案做成一个演示服务，用一个 HTML 页面点按钮 + F12 看真实报文。

---

### 2. JWT 认证
**说明**：本节起，所有内容都长进阶段 0.3 的项目骨架里。

### 学习资源：
- [JWT 官方规范 RFC 7519](https://www.rfc-editor.org/rfc/rfc7519)
- [jwt.io（在线解 token）](https://jwt.io/)
- [golang-jwt/jwt v5](https://github.com/golang-jwt/jwt)
- [Gin JWT 示例](https://gin-gonic.com/zh-cn/docs/examples/jwt-authentication/)

### 关键知识点：
- 三段结构：Header / Payload / Signature，前两段是 base64url 编码**不是加密**
- 签名校验：签名是对 payload 算的，改 payload 就验不过
- 密钥泄露 = 攻击者可自行签发任意 token
- **刷新机制**：access token（短）+ refresh token（长，存服务端）
- **撤销问题**：无状态 token 无法主动失效，三种解法——短有效期 + refresh、服务端存 token 版本号、黑名单

### 验收标准：
能回答"JWT 怎么主动踢人下线"，并且能实现其中至少一种方案。这是本节真正的考点，`jwt.New()` 只是手段。

### 实践建议：
实现登录 / 刷新 / 登出 / 权限校验四个接口，登出必须真的能让旧 token 失效。

---

### 3. 错误处理和恢复
**说明**：Gin 的 `gin.Default()` 自带 `Recovery()`，但捕获 panic 不等于错误处理。

### 学习资源：
- [Go 错误处理最佳实践](https://go.dev/blog/error-handling-and-go)
- [Gin 中间件开发](https://gin-gonic.com/zh-cn/docs/examples/custom-middleware/)

### 关键知识点：
- `panic` / `recover` 的边界，为什么不能在 goroutine 里靠主流程 recover
- 业务错误 vs 系统错误的区分，错误码标准化
- 统一错误响应格式，错误日志不要直接返回给前端
- 包装错误：`fmt.Errorf("...: %w", err)` 和 `errors.Is/As`

### 验收标准：
项目里有一个全局错误中间件，handler 不再各自 `c.JSON` 拼错误响应。

### 实践建议：
定义 `errors` 包 + `response` 包（现有 `res/enter.go` 挪过来），handler 只 `return err`。

---

### 4. 日志记录
### 学习资源：
- [Zap](https://go.uber.org/zap)
- [Logrus](https://github.com/sirupsen/logrus)
- [Gin 日志中间件](https://github.com/gin-gonic/gin/tree/master/examples/basic/logger)

### 关键知识点：
- 日志分级：Debug / Info / Warn / Error 各自的语义
- 结构化日志（字段化）vs 字符串拼接，为什么后者在排查时不可用
- 请求日志中间件：traceId 贯穿一次请求
- 日志输出到文件 + 轮转

### 验收标准：
能按 traceId 把一次请求的所有日志串起来。

---

### 5. 请求限流
### 学习资源：
- [ulule/limiter v3](https://github.com/ulule/limiter/v3)
- [令牌桶 / 漏桶算法](https://www.geeksforgeeks.org/sliding-window-algorithm-for-ratelimiting/)

### 关键知识点：
- 固定窗口 / 滑动窗口 / 令牌桶三种策略的差异
- 按 IP 还是按用户限流，单机 map 在多实例下失效
- 限流响应：`429 Too Many Requests` + `Retry-After`

### 验收标准：
能画出令牌桶状态变化，并说明"为什么固定窗口在窗口边界会放过 2 倍流量"。

### 实践建议：
手写一个令牌桶中间件（20 行左右），再对比 limiter 库，理解库替你做了什么。

---

### 6. 配置管理
### 学习资源：
- [Viper](https://github.com/spf13/viper)
- [Go os 环境变量](https://pkg.go.dev/os#Environ)

### 关键知识点：
- YAML / JSON / TOML 配置加载
- 环境变量覆盖配置（十二要素应用里"配置在环境里"）
- 配置分层：本地默认值 < 配置文件 < 环境变量
- 启动时校验配置合法性，缺字段直接 fail fast

### 验收标准：
同一份代码能用不同配置跑在不同端口，且配置错误在启动时就报错而不是运行时报错。

---

### 7. 数据库集成
### 学习资源：
- [GORM 文档](https://gorm.io/zh_CN/)
- [Gin + GORM 示例](https://gorm.io/zh_CN/docs/gin.html)
- [Go database/sql](https://pkg.go.dev/database/sql)

### 关键知识点：
- 连接池配置：`SetMaxOpenConns` / `SetMaxIdleConns` / `SetConnMaxLifetime`，不配会踩连接耗尽
- 模型定义与自动迁移
- CRUD 与查询构造（`Where` / `Joins` / `Preload`）
- **事务**：`Transaction` 回调、`SavePoint`、事务里出错要 `Rollback`
- N+1 查询问题与 `Preload` 解决
- 索引对查询计划的影响（配合 MySQL 学习）

### 验收标准：
写一个必须回滚的场景（转账或库存扣减），并用单元测试覆盖"第二步失败则第一步不生效"。CRUD 本身不是考点，**一致性才是**。

### 实践建议：
这一节值得单独开目录 `22.数据库集成/`，因为它是项目的主干。

---

### 8. 优雅关闭
### 学习资源：
- [Go 信号处理](https://pkg.go.dev/signal)
- [Gin 优雅关闭示例](https://github.com/gin-gonic/examples/tree/master/graceful-shutdown)

### 关键知识点：
- `os/signal` 捕获 `SIGTERM` / `SIGINT`
- `http.Server.Shutdown(ctx)` 的超时语义
- 请求处理超时：`context.WithTimeout`
- 资源清理顺序：停止接新请求 → 等存量请求完成 → 关连接池

### 验收标准：
`Ctrl+C` 时正在处理的请求不会被掐断，日志里能看到"正在等待 N 个请求完成"。

---

### 贯穿项：单元测试
不单开一节，每节写 2–3 个用例即可。
- [Go Testing 官方教程](https://go.dev/doc/tutorial/testing-and-benchmarking)
- [Gin 测试示例（httptest）](https://github.com/gin-gonic/examples/tree/master/test)
- 表驱动测试、`httptest.NewRecorder`、路由用 `r.ServeHTTP` 直接驱动

已有手感参考：[07.原始内容/test_body_read.go](07.原始内容/test_body_read.go)。

---

## 阶段 2：存储与部署（另起计划文档）

按序，不要跳：

1. **MySQL**：索引与 B+ 树、事务隔离级别、锁（行锁 / 间隙锁）、慢查询与 `EXPLAIN`。面试八股重灾区，和 7 的事务练习是同一件事的两面
2. **Redis**：五种数据结构的使用场景、过期策略、缓存三大问题（穿透 / 击穿 / 雪崩）、把 Session 从内存挪到 Redis（正好接 1 的多实例问题）
3. **Docker**：Dockerfile、容器化部署自己的项目、`docker-compose` 起 MySQL + Redis + 服务

---

## 阶段 3：算法（并行，从第 1 天开始）

这是硬门槛，和框架能力无关。

- 每天一道 LeetCode，目标 300 道左右（中等难度为主）
- 顺序：数组 / 双指针 → 哈希 → 字符串 → 链表 → 二叉树 → 回溯 / 动态规划 → 栈与队列
- 同一道题三天后回看，能独立重写才算过

---

## 时间线

| 时间 | 内容 |
|---|---|
| 2026-09-18 ~ 09-20 | 阶段 0 工程基建 + Cookie/Session |
| 2026-09-21 ~ 10-20 | JWT → 错误处理 → 日志 → 限流 → 配置（全部长进项目骨架） |
| 2026-10-21 ~ 11-15 | 数据库（含事务）+ 优雅关闭，Gin 部分收尾 |
| 2026-11-16 ~ 2027-01-31 | MySQL |
| 2027-02-01 ~ 02-28 | Redis + Docker |
| 2027-03-01 ~ 06-30 | 项目打磨 + 简历 + 算法冲刺 |
| 2027-07 ~ 08 | 投递 |

算法从 2026-09-18 起并行，不断档。

---

## 实习水平的验收标准

学完以上内容只是**有资格进面**。实际录用看四样：

1. **一个能讲 30 分钟的项目**——不是"我用 gin 做了个接口"，而是"我在哪里做了取舍、为什么、出问题怎么查的"。现在 `15.跨域问题/跨域问题.go` 里关于 `Origin` / `RawPath` / `Host` 那段注释就是这类素材，多攒
2. **Go 语言本身**：goroutine / channel、GMP 调度、concurrent map 的坑、`panic/recover` 边界、GC 三色标记。框架都会但语言说不清，一面就露
3. **数据库八股**：MySQL 索引与事务隔离、Redis 过期与缓存三大问题
4. **简历 + 投递时机**

选题建议：不要做"用户管理系统"。做研究方向相关的——实验数据管理、模型评测结果采集与对比。面试官看到项目内容和你的课题能对上，是白送的加分项。

---

## 暂缓学习的项

不是没用，是现在没有用例。有明确需求再补。

- **文件上传进阶**（类型校验 / 大小限制 / 存储策略）——做需要传文件的业务时再补
- **WebSocket**——没有实时推送用例时纯消耗
- **API 文档 / Swagger**（[swaggo/swag](https://github.com/swaggo/swag) + [gin-swagger](https://github.com/swaggo/gin-swagger)）——需要把 API 交给别人对接时再补
- **Prometheus 监控 / 健康检查**——没有真实部署就上不了手

---

## 常用工具库

```bash
# 认证
go get github.com/golang-jwt/jwt/v5

# CORS
go get github.com/gin-contrib/cors

# 日志
go get go.uber.org/zap

# 配置
go get github.com/spf13/viper

# 限流
go get github.com/ulule/limiter/v3

# 数据库
go get gorm.io/gorm
go get gorm.io/driver/mysql

# 缓存
go get github.com/redis/go-redis/v9

# 文件类型检测（暂缓项用）
go get github.com/gabriel-vasile/mimetype

# 参数校验
go get github.com/go-playground/validator/v10
```

---

## 学习建议

1. **从 17 起只往一个项目里长**，不要再造独立 `main` 脚本
2. **每节写 2–3 个测试**，不单开测试节
3. **读源码**：gin 的 `Context.Next()` / `Abort()`、validator 的 `FieldLevel`、httputil 的反向代理，这三个读透比多学两个中间件有用
4. **每个知识点都要能复现一个坑**，而不是背结论
5. 完成后把 `res/enter.go` 里的历史问题修掉：`Fail()` 收了 `Code` 参数却固定返回 `0`（失败响应也会报成功码）、`FailWithMsg` 硬编码 `1001`（权限错误）语义不符

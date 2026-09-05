## 说明

这里只放学习笔记和小 demo。
# Gin

框架学习，按章节递进。

## 环境

导入gin包，在项目根目录执行一次：

```bash
go get github.com/gin-gonic/gin
```

依赖记在仓库根的 `go.mod` / `go.sum`，不需要单独建 module。

## 目录

| 章节 | 内容 |
| --- | --- |
| 0 | 引擎、路由组、路由参数、HTTP 方法 |
| 02.请求与响应 | context、参数读取、响应 JSON |
| 03.中间件 | 全局/分组中间件、next()、链式调用 |
| 04.表单与上传 | 表单解析、文件上传 |
| 05.模板与静态资源 | HTML 模板、LoadHTMLGlob、静态文件 |
| 06.配置管理 | 环境变量、YAML 配置、优雅退出 |
| 07.数据库与日志 | GORM 集成、zap 日志 |

每章一个目录，目录名带序号，代码放 `main.go`，可运行。

跟 `basic/` 一样，目录名含中文就无法用 `go run ./gin/01.路由`，进目录跑：

```bash
cd gin/01.路由
go run main.go
```



model-view-controler
数据模型层-视图层-控制层
## 项目目录结构

```
webdemo/
├── main.go                    # 程序入口文件
├── go.mod                    # Go 模块文件
├── go.sum                    # 依赖校验文件
├── config/                   # 配置文件目录
│   ├── database.go           # 数据库配置
│   └── config.go             # 应用配置
├── controllers/              # 控制器层
│   ├── user_controller.go   # 用户控制器
│   └── product_controller.go  # 产品控制器
├── models/                  # 模型层
│   ├── user.go              # 用户模型
│   ├── product.go           # 产品模型
│   └── database.go          # 数据库模型和连接
├── services/                # 业务逻辑层
│   ├── user_service.go      # 用户业务逻辑
│   └── product_service.go   # 产品业务逻辑
├── routes/                  # 路由配置
│   └── routes.go            # 路由定义
├── utils/                   # 工具函数
│   ├── validator.go         # 数据验证工具
│   └── response.go          # 响应工具
├── middleware/              # 中间件
│   ├── auth.go              # 认证中间件
│   └── cors.go              # 跨域中间件
├── views/                   # 视图层（可选）
│   ├── layouts/             # 布局模板
│   └── templates/           # 页面模板
└── static/                  # 静态文件
    ├── css/
    ├── js/
    └── images/
```

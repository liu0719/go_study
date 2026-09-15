# 🚀 Go语言基础学习笔记

## 📚 学习目录

- [b01.var](b01.var/) - 变量和常量
- [b02.io](b02.io/) - 输入输出
- [b03.type](b03.type/) - 数据类型
- [b04.slice_map](b04.slice_map/) - 切片和映射
- [b05.control_statement](b05.control_statement/) - 控制语句
- [b06.function](b06.function/) - 函数
- [b07.struct](b07.struct/) - 结构体
- [b08.custom_type_and_aliases](b08.custom_type_and_aliases/) - 自定义类型和别名
- [b09.interface](b09.interface/) - 接口
- [b10.goroutine](b10.goroutine/) - 协程
- [b11.channel](b11.channel/) - 通道
- [b12.select](b12.select/) - select语句
- [b13.thread_safe](b13.thread_safe/) - 线程安全
- [b14.exception_handle](b14.exception_handle/) - 异常处理
- [b15.generic](b15.generic/) - 泛型
- [b16.file_operation](b16.file_operation/) - 文件操作
- [b17.unit_test](b17.unit_test/) - 单元测试
- [b18.reflect](b18.reflect/) - 反射
- [b19.netcoding](b19.netcoding/) - 网络编程
- [b20.deployment](b20.deployment/) - 部署
- [b21.GMPmodel](b21.GMPmodel/) - GMP模型

## 🛠️ 环境配置

### 安装Go语言
```bash
# 下载地址
https://golang.org/dl/

# 验证安装
go version

# 设置GOPATH，存储第三方包
# linux
export GOPATH=$HOME/go

# windows
# 新建GOPATH,不要在c盘
# 将`%GOPATH%/bin`加到环境变量内
```

### 项目初始化
```bash
# 进入basic目录
cd basic

# 初始化Go模块
go mod init basic

# 运行示例代码
go run b01.var/main.go
```
## 🔗 相关资源

- [Go官方文档](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Go语言圣经](https://gopl-zh.github.io/)

## 📝 笔记记录

建议在学习过程中记录关键点和心得体会，可以放在每个章节的`notes.md`文件中。

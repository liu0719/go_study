# web

## 环境安装

### mysql
1. **安装配置环境变量**
    1. 官网地址 ：[dev.mysql.com/downloads](https://dev.mysql.com/downloads/mysql/)
        > 解压目录说明  

        - bin：存放 MySQL 可执行文件（如 mysql.exe、mysqld.exe）
        - data：默认数据存储目录（后续自动生成）
        - my.ini：MySQL 配置文件（需手动创建）
        - README：官方说明文档
    2. 新建环境变量 ：
        1. 新建`MYSQL_HOME`，其目录指向安装mysql的根目录 eg: `D:\environment\mysql-8.0.46-winx6`
        2. 将`%MYSQL_HOME%\bin`加入到path中
2. 初始化
    1. 在mysql安装目录新建 *my.ini* 文件,将以下内容加入进去，并保存  

        ```sql
            [mysql]
            # 设置mysql客户端默认字符集
            default-character-set=utf8
            [mysqld]
            #设置3306端口
            port = 3306
            # 设置mysql的安装目录
            basedir=D:\environment\mysql-8.0.46-winx64
            # 设置mysql数据库的数据的存放目录
            datadir=D:\environment\mysql-8.0.46-winx64\data
            # 允许最大连接数
            max_connections=200
            # 服务端使用的字符集默认为8比特编码的latin1字符集
            character-set-server=utf8
            # 创建新表时将使用的默认存储引擎
            default-storage-engine=INNODB
        ```
    2. 在有管理员权限的命令行内，运行以下命令，初始化mysql
        ```bash
        mysqld --initialize --console
        ```
        > 成功后，最后一行会有登录密码，复制保存一下
    3. 安装命令
        ```bash
        mysqld install
        ```
        > 会提示successful，然后启动mysql
        ```bash
        net start mysql
        ```
    4. 登录，修改密码
        > 登录命令
        ```bash
        mysql -u root -p
        ```
        > 输入刚才的密码，登进去
        > 修改密码
        ```bash
        ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';
        ```
***完成***
### redis


## go web框架

### gin

```bash

```


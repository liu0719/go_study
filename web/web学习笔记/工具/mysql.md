# MySQL

## 安装
1. 安装并配置环境变量
   - 官网地址: [dev.mysql.com/downloads](https://dev.mysql.com/downloads/mysql/)
   - 解压目录说明：
     - bin：存放 MySQL 可执行文件，如 mysql.exe、mysqld.exe
     - data：默认数据存储目录
     - my.ini：MySQL 配置文件
     - README：官方说明文档
   - 新建环境变量：
     1. 新建 MYSQL_HOME，指向 MySQL 安装根目录，例如 D:\environment\mysql-8.0.46-winx64
     2. 将 %MYSQL_HOME%\bin 加入到 PATH 中

2. 初始化数据库
   - 在 MySQL 安装目录中新建 my.ini 文件，内容如下：

```ini
[mysql]
default-character-set=utf8

[mysqld]
port = 3306
basedir=D:\environment\mysql-8.0.46-winx64
datadir=D:\environment\mysql-8.0.46-winx64\data
max_connections=200
character-set-server=utf8
default-storage-engine=INNODB
```

   - 使用管理员权限打开命令行，执行：

```bash
mysqld --initialize --console
```

   - 成功后，控制台最后一行会输出初始密码，记得保存。

3. 安装并启动 MySQL

```bash
mysqld install
net start mysql
```

4. 登录并修改密码

```bash
mysql -u root -p
```

输入刚才的初始密码后，执行：

```sql
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456';
```

MySQL 安装完成 ✅

---
# Git
git是用来控制版本的
> 下载地址[git下载地址 ](https://git-scm.com/install/)

## git修改提交路径
```bash
git config --global --unset http.proxy    # 关梯子
git config --global http.proxy http://127.0.0.1:10808   # 开梯子后恢复,端口为梯子的端口
```
## git文件忽略跟踪
在跟目录`.git`同级文件夹下新建`.gitignore`在该文件中配置要忽略的文件和文件夹

> eg：
```gitignore
# 文件夹名，也可以用路径名
.vscode
.obsidian
/docs/
```
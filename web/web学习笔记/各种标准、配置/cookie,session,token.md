# 1.cookie


*Cookie不是登录态，它只是浏览器每次请求自动携带的文本容，作用是和session和token配合记录登陆证状态*

所以"用 Cookie 做登录"的完整内容是：Cookie 里放一个**凭据**。这个凭据要么自己能被验证（无状态，JWT），要么服务端拿它去查表（有状态，Session）。

---

| 字段                   | 作用         | 踩坑点                                                    |
| -------------------- | ---------- | ------------------------------------------------------ |
| `name`和`value`       | cookie的键值对 |                                                        |
| `Expires` / `MaxAge` | 过期时间       | **`MaxAge`和`Expires`同时设了以`MaxAge`为准**,为-1立刻删除，为0表示单次有效 |
| `Path`               | 只在哪些路径下回传  | `/api` 表示 `/api/xxx` 会带，但 `/` 不带                       |
| `Domain`             | 哪些域名能收到    | 设 `.example.com` 子域也能拿到                                |
| `Secure`             | 是否只走 HTTPS | **设 true 后本地 `http://` 调试时浏览器根本不发这个 Cookie** ← 新手第一大坑  |
| `HttpOnly`           | JS 能不能读    | **true → `document.cookie` 读不到，防 XSS 偷 Cookie**        |
| `SameSite`           | 跨站请求带不带    | gin 的 `SetCookie` 没有这个参数，默认 `Lax`                      |

---

# 2.session和token

1. session的策略就是：

    用户登录时，服务器会生成一个sessionID，自己在存储库存一份，并将其当作cookie的值返回给请求者。之后请求者每次都带上cookie请求，服务器拿到cookie之后和数据库的sessionID对比就知道是哪个用户。

**session缺点** 现实中服务器并不止一台，处理同一个用户请求的服务器并不一定是同一台服务器

**解决方法** ：
1.session复制，每台服务器都配置一份相同的sessionId，浪费资源
2.session粘连，每个请求有指定的服务器解决，这个单一的服务器挂了就卡了
3.session共享，构建所有服务器共享的集群服务器，需要资金

![session工作流程](../../../static/images/session工作流程.jpg)

2. token的策略：

![token组成部分](../../../static/images/token构成部分.jpg)

可以看到token由三部分构成，其中header和payload是以 `base64`方式加密过后的

`header`：指定了签名算法

`payload`：可以指定用户 id，过期时间等非敏感数据

`Signature`: 签名，服务器 根据 header 知道它该用哪种签名算法，再用密钥根据此签名算法对 head + payload 生成签名，这样一个 token 就生成了。


**token流程:** 当服务器收到浏览器传过来的 token 时，它会首先取出 token 中的 header + payload，根据密钥生成签名，然后再与 token 中的签名比对，如果成功则说明签名是合法的，即 token 是合法的。而且你会发现 payload 中存有我们的 userId，所以拿到 token 后直接在 payload 中就可获取 userid，避免了像 session 那样要从 redis 去取的开销。


这种方式非常微妙，只要服务器保证密钥不泄露，那这个token就是安全的
[gin中cookie，session,token处理](../../gin/16.cookie和session处理/cookie和session.go)


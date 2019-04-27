
### gmx作用：
1. 信息采集，比如cpu、内存、在线玩家数等
2. 程序控制，比如设置最大负载、通知热加载表数据、踢人、关闭服务器（这些功能必须提前存在，它可以作为调用功能的入口）

### 使用方式：
在bean的字段后面加入tag `gmx:"name"`,比如：  
```go
type login struct {
	controller.IController
	loginTokenCheck string         `cfg:"url.login.token.check"`
	serviceLogin    *service.Login `bean:""`
	forbidLogin bool				`gmx:"forbidlogin"`
}
```
然后我们就可以通过http协议使用gmx功能  
http://127.0.0.1:10101/gmx/keys 获取所有的key  
http://127.0.0.1:10101/gmx/get?key=forbidlogin 获取key的值  
http://127.0.0.1:10101/gmx/set?key=forbidlogin&value=true 设置key的值  
http://127.0.0.1:10101/gmx/call 调用函数  

[gmx源码](https://github.com/gosrv/gmx)


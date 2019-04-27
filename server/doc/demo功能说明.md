demo客户端使用unity3d编写，有登陆、基本数据处理、抽卡、充值、购买金币等功能，
主要用于演示基本

登陆界面  
![登陆](https://github.com/gosrv/gdemo/tree/master/server/doc/img/login.png)  
使用匿名方式登陆，流程：
1. 去账号服拿登陆token
2. 将token发往baseapp
3. baseapp用token去账号服验证
4. 返回给客户端结果

主界面  
![登陆](https://github.com/gosrv/gdemo/tree/master/server/doc/img/main.png)  
登陆成功后进入主界面  
我们可以通过在线时间获取收益    
点击挑战可以进入下一关

充值  
![登陆](https://github.com/gosrv/gdemo/tree/master/server/doc/img/money.png)  
点击钻石按键可以免费获取钻石
点击金币可以使用钻石购买金币

英雄  
![登陆](https://github.com/gosrv/gdemo/tree/master/server/doc/img/money.png)    
消耗钻石抽取英雄  
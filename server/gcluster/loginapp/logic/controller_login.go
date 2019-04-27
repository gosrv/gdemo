package logic

import (
	"github.com/globalsign/mgo/bson"
	"github.com/gosrv/gbase/ghttp"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gcluster/gcluster/common/entity"
	"github.com/gosrv/glog"
)

const (
	CodeError = "error"
	CodeOk    = "ok"
)

/**
逻辑消息控制器
*/
type controllerLogin struct {
	// 控制器标记
	ghttp.IHttpController
	log          glog.IFieldLogger `log:"app"`
	serviceLogin *serviceLogin     `bean:""`
}

func NewControllerLogin() *controllerLogin {
	return &controllerLogin{
		IHttpController: ghttp.NewHttpRestController("/login"),
	}
}

// 登陆请求，如果密码为空，认为是匿名登陆，返回一个验证token
func (this *controllerLogin) Login(ctx gnet.ISessionCtx, param struct{ account, password string }) *entity.LoginResult {
	// 先尝试缓存
	tokenCtx := this.serviceLogin.loadTokenByAcc(param.account)
	if tokenCtx == nil {
		// 发起数据库登陆
		id := this.serviceLogin.login(param.account, param.password)
		if id > 0 {
			tokdn := bson.NewObjectId().Hex()
			tokenCtx = entity.NewLoginTokenCtx(param.account, param.password, id, tokdn)
		}
	} else if param.password != tokenCtx.Pw {
		tokenCtx = nil
	}

	var loginResult *entity.LoginResult
	if tokenCtx == nil {
		loginResult = entity.NewLoginResult(CodeError, "")
	} else {
		loginResult = entity.NewLoginResult(CodeOk, tokenCtx.Token)
		this.serviceLogin.saveToken(tokenCtx)
	}

	return loginResult
}

// 注册处理
func (this *controllerLogin) Register(ctx gnet.ISessionCtx, param struct{ account, password string }) *entity.LoginRegisterResult {
	id := this.serviceLogin.register(param.account, param.password)
	if id > 0 {
		return entity.NewLoginRegisterResult(true)
	} else {
		return entity.NewLoginRegisterResult(false)
	}
}

// token检查
func (this *controllerLogin) Check(ctx gnet.ISessionCtx, param struct{ token string }) *entity.LoginCheckResult {
	tokenCtx := this.serviceLogin.loadTokenByTk(param.token)
	if tokenCtx == nil {
		return entity.NewLoginCheckResult("error", 0)
	} else {
		return entity.NewLoginCheckResult("ok", tokenCtx.Id)
	}
}

package controller

import (
	"encoding/json"
	"github.com/gosrv/gbase/controller"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/gcluster/gcluster/proto"
	"io/ioutil"
	"net/http"
)

/**
登陆消息处理控制器
*/
type login struct {
	controller.IController
	loginTokenCheck string         `cfg:"url.login.token.check"`
	serviceLogin    *service.Login `bean:""`
	forbidLogin bool				`gmx:"forbidlogin"`
}

func newLogin() *login {
	return &login{
		IController: controller.NewTypeController(""),
	}
}

type LoginTokenCheck struct {
	Id int64
}

// 登陆消息处理
func (this *login) Login(ctx gnet.ISessionCtx, login *netproto.CS_Login, netChannel gproto.INetChannel) *netproto.SC_Login {
	repLogin := &netproto.SC_Login{
		Code: netproto.E_Code_E_ERROR,
	}
	if this.forbidLogin {
		return repLogin
	}
	// 向登录服查询登陆数据
	repCheck, err := http.Get(this.loginTokenCheck + login.Token)
	if err != nil {
		return repLogin
	}
	body, err := ioutil.ReadAll(repCheck.Body)
	if err != nil {
		return repLogin
	}
	loginCheck := &LoginTokenCheck{}
	json.Unmarshal(body, loginCheck)
	if loginCheck.Id <= 0 {
		return repLogin
	}

	repLogin.Code = this.serviceLogin.ProcessLogin(netChannel, ctx, loginCheck.Id)
	// 登陆服授权成功，进入登陆处理
	return repLogin
}

func init() {
	common.BeansInit = append(common.BeansInit, newLogin())
}

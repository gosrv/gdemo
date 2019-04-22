package controller

import (
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gbase/controller"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/proto"
	"io/ioutil"
	"net/http"
)

/**
登陆消息处理控制器
*/
type ControllerLogin struct {
	controller.IController
	loginTokenCheck string                `cfg:"url.login.token.check"`
	serviceLogin    *service.ServiceLogin `bean:""`
}

func NewControllerLogin() *ControllerLogin {
	return &ControllerLogin{
		IController: controller.NewTypeController(""),
	}
}

type LoginTokenCheck struct {
	Id int64
}

// 登陆消息处理
func (this *ControllerLogin) Login(ctx gnet.ISessionCtx, login *netproto.CS_Login, netChannel gproto.INetChannel) *netproto.SC_Login {
	repLogin := &netproto.SC_Login{
		Id:   proto.Int64(-1),
		Code: netproto.E_Code_E_ERROR.Enum(),
	}
	// 向登录服查询登陆数据
	repCheck, err := http.Get(this.loginTokenCheck + *login.Token)
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

	code := this.serviceLogin.ProcessLogin(netChannel, ctx, loginCheck.Id)
	// 登陆服授权成功，进入登陆处理
	repLogin.Code = code.Enum()
	if code == netproto.E_Code_E_OK {
		repLogin.Id = &loginCheck.Id
	}
	return repLogin
}

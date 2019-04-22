package logic

import (
	"encoding/json"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gbase/tcpnet"
	entity2 "github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/common/entity"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc/util"
	"io/ioutil"
	"net/http"
	"reflect"
)

type serviceLogic struct {
	log        glog.IFieldLogger    `log:"app"`
	serverHost string               `cfg:"server.host"`
	net        *tcpnet.TcpNetServer `bean:""`
	urlLogin   string               `cfg:"url.login.login"`
}

func (this *serviceLogic) BeanInit() {
	eventRoute := this.net.GetEventRoute()
	eventRoute.Connect(gnet.NetEventConnect, func(from interface{}, key interface{}, val interface{}) interface{} {
		this.log.Debug("net connect event")
		ctx := from.(gnet.ISessionCtx)
		playerData := entity2.NewPlayerData()
		playerInfo := entity2.NewPlayerInfo()
		ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(playerData), playerData)
		ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(playerInfo), playerInfo)
		netChannel := ctx.GetByType(gproto.INetChannelType).(gproto.INetChannel)
		this.doLogin(ctx, netChannel)
		return nil
	})
	eventRoute.Connect(gnet.NetEventDisconnect, func(from interface{}, key interface{}, val interface{}) interface{} {
		this.log.Debug("net disconnect event")
		return nil
	})
}

func (this *serviceLogic) BeanUninit() {

}

func NewServiceLogic() *serviceLogic {
	return &serviceLogic{}
}

func (this *serviceLogic) BeanStart() {
	this.net.NetConnect(this.serverHost)
}

func (this *serviceLogic) BeanStop() {

}

func (this *serviceLogic) doLogin(ctx gnet.ISessionCtx, netChannel gproto.INetChannel) {
	rep, err := http.Get(this.urlLogin)
	util.VerifyNoError(err)
	body, err := ioutil.ReadAll(rep.Body)
	util.VerifyNoError(err)
	result := entity.LoginResult{}
	err = json.Unmarshal([]byte(body), &result)
	util.VerifyNoError(err)

	netChannel.Send(&netproto.CS_Login{Token: &result.Token})
}

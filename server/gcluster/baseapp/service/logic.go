package service

import (
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gbase/tcpnet"
	"github.com/gosrv/glog"
)

type Logic struct {
	log glog.IFieldLogger    `log:"app"`
	net *tcpnet.TcpNetServer `bean:""`
}

func (this *Logic) BeanInit() {
	eventRoute := this.net.GetEventRoute()
	eventRoute.Connect(gnet.NetEventConnect, func(from interface{}, key interface{}, data interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		netChannel := ctx.Get(gproto.INetChannelType).(gproto.INetChannel)
		this.log.Debug("net connect event %v->%v", netChannel.RemoteAddr(), netChannel.LocalAddr())
		return nil
	})
	eventRoute.Connect(gnet.NetEventDisconnect, func(from interface{}, key interface{}, data interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		netChannel := ctx.Get(gproto.INetChannelType).(gproto.INetChannel)
		this.log.Debug("net disconnect event %v->%v", netChannel.RemoteAddr(), netChannel.LocalAddr())
		return nil
	})
}

func (this *Logic) BeanUninit() {
}

func NewLogic() *Logic {
	return &Logic{}
}

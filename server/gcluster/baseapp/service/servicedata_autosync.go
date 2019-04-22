package service

import (
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/tcpnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"reflect"
)

type serviceDataAutoSync struct {
	net *tcpnet.TcpNetServer `bean:""`
}

func NewServiceDataAutoSync() *serviceDataAutoSync {
	return &serviceDataAutoSync{}
}

func (this *serviceDataAutoSync) BeanInit() {
	eventRoute := this.net.GetEventRoute()
	// 心跳时进行数据同步处理
	eventRoute.Connect(gnet.NetEventTick, func(from interface{}, key interface{}, value interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		sync := ctx.Get(reflect.TypeOf((*entity.PlayerDataSync)(nil)))
		if sync != nil {
			sync.(*entity.PlayerDataSync).TrySyncDirtyData(false)
		}

		return nil
	})
	// 断开连接是，要强制写数据库
	eventRoute.Connect(gnet.NetEventDisconnect, func(from interface{}, key interface{}, value interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		sync := ctx.Get(reflect.TypeOf((*entity.PlayerDataSync)(nil))).(*entity.PlayerDataSync)
		sync.TrySyncDirtyData(true)
		return nil
	})
}

func (this *serviceDataAutoSync) BeanUninit() {

}

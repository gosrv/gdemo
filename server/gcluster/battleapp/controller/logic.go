package controller

import (
	"github.com/gosrv/gbase/controller"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/glog"
)

/**
逻辑消息控制器
*/
type logic struct {
	log glog.IFieldLogger `log:"app"`
	// 控制器标记
	controller.IController
}

func newLogic() *logic {
	return &logic{
		// 路由收集器，它会收集这样的函数作为路由器：
		// 第一个变量是gnet.ISessionCtx，第二个是消息，可以返回一个一个消息，也可以不返回
		IController: controller.NewTypeController(""),
	}
}

// 心跳消息处理，playerData和playerInfo是自动注入的，如果你不需要也可以不写，比如下面的写法
func (this *logic) Tick(ctx gnet.ISessionCtx, msg *netproto.BCS_Tick) *netproto.BSC_Tick {
	return &netproto.BSC_Tick{}
}

func init() {
	common.BeansInit = append(common.BeansInit, newLogic())
}

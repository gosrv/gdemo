package logic

import (
	"github.com/gosrv/gbase/controller"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/glog"
)

/**
逻辑消息控制器
*/
type ControllerLogic struct {
	// 控制器标记
	controller.IController
	log glog.IFieldLogger `log:"app"`
	// 逻辑处理
	serviceLogic *serviceLogic `bean:""`
}

func NewControllerLogic() *ControllerLogic {
	return &ControllerLogic{
		IController: controller.NewTypeController(""),
	}
}

// 心跳消息处理
func (this *ControllerLogic) Tick(ctx gnet.ISessionCtx, msg *netproto.SC_Tick) {
	this.log.Debug("tick msg %v", msg.String())
}

func (this *ControllerLogic) Logic(ctx gnet.ISessionCtx, msg *netproto.SC_Login) {
	this.log.Debug("login result %v", msg.String())
}

func (this *ControllerLogic) SyncData(ctx gnet.ISessionCtx, msg *netproto.PlayerData, playerData *entity.PlayerData) {
	playerData.FromProto(msg)
	this.log.Debug("login result %v", msg.String())
}

func (this *ControllerLogic) SyncInfo(ctx gnet.ISessionCtx, msg *netproto.PlayerInfo, playerInfo *entity.PlayerInfo) {
	playerInfo.FromProto(msg)
	this.log.Debug("login result %v", msg.String())
}

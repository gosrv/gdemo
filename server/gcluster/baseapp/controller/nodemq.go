package controller

import (
	"github.com/gosrv/gbase/cluster"
	"github.com/gosrv/gbase/controller"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gutil"
	"github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/glog"
	"time"
)

// 这是一个集群的消息处理路由器
type nodeMq struct {
	log glog.IFieldLogger `log:"app"`
	controller.IController
	nodeMgr   *cluster.NodeMgr   `bean:""`
	nodeMq    *cluster.NodeMQ    `bean:""`
	playerMgr *service.PlayerMgr `bean:""`
	playerBaseappInfo *common.PlayerBaseappInfo	`bean:""`
}

func newNodeMq() *nodeMq {
	return &nodeMq{
		IController: controller.NewTypeController(cluster.NodeMQName),
	}
}

func (this *nodeMq) BeanStart() {
	gutil.RecoverGo(func() {
		for {
			for _, node := range this.nodeMgr.GetAllNodesInfo() {
				if node.NodeUuid == this.nodeMgr.GetMyNodeInfo().NodeUuid {
					continue
				}
				_ = this.nodeMq.Push(node.NodeUuid, &netproto.SS_Tick{})
			}
			time.Sleep(time.Second * 5)
		}
	})
}

func (this *nodeMq) BeanStop() {

}

func (this *nodeMq) Tick(ctx gnet.ISessionCtx, msg *netproto.SS_Tick) {
	this.log.Print("cluster tick msg")
}

func (this *nodeMq) KickPlayer(ctx gnet.ISessionCtx, msg *netproto.SS_KickPlayer) *netproto.SS_PlayerBaseappRedirect {
	if this.playerMgr.GetNetchannelByPlayerId(msg.PlayerId) == nil {
		return &netproto.SS_PlayerBaseappRedirect{PlayerId:msg.PlayerId}
	}
	this.playerMgr.KickPlayer(msg.PlayerId)
	return nil
}

func (this *nodeMq) Player2Player(ctx gnet.ISessionCtx, msg *netproto.SS_Player2Player) *netproto.SS_PlayerBaseappRedirect {
	netChannel := this.playerMgr.GetNetchannelByPlayerId(msg.ToPlayerId)
	if netChannel == nil {
		return &netproto.SS_PlayerBaseappRedirect{PlayerId:msg.ToPlayerId}
	}
	netChannel.Send(&netproto.SC_P2P{From:msg.FromPlayerId, To:msg.ToPlayerId, Msg:msg.Data})
	return nil
}

func (this *nodeMq) PlayerBaseappRedirect(ctx gnet.ISessionCtx, msg *netproto.SS_PlayerBaseappRedirect) {
	this.playerBaseappInfo.InvalidCacheUuid(msg.PlayerId)
}

func init() {
	common.BeansInit = append(common.BeansInit, newNodeMq())
}

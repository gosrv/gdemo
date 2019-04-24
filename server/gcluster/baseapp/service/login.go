package service

import (
	"github.com/gosrv/gbase/cluster"
	"github.com/gosrv/gbase/gdb"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gbase/tcpnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/common/meta"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/glog"
)

type Login struct {
	log                   glog.IFieldLogger    `log:"app"`
	net                   *tcpnet.TcpNetServer `bean:""`
	playerMgr             *PlayerMgr           `bean:""`
	nodeUuid              string               `bean:""`
	nodeMgr			      *cluster.NodeMgr  `bean:""`
	nodeMq				  *cluster.NodeMQ    `bean:""`
	servicePlayerMsgQueue *PlayerMsgQueue      `bean:""`
}

func (this *Login) BeanInit() {
	eventRoute := this.net.GetEventRoute()
	eventRoute.Connect(gnet.NetEventConnect, func(from interface{}, key interface{}, data interface{}) interface{} {
		return nil
	})
	// 玩家断开连接时，要清理数据
	eventRoute.Connect(gnet.NetEventDisconnect, func(from interface{}, key interface{}, data interface{}) interface{} {
		ctx := from.(gnet.ISessionCtx)
		playerIdIns := ctx.Get(entity.PlayerIdKey)
		if playerIdIns != nil {
			this.playerMgr.PlayerLogout(playerIdIns.(int64))
			attributeGroup := ctx.Get(gdb.IDBAttributeGroupType).(gdb.IDBAttributeGroup)
			attributeGroup.CasSetAttribute(meta.BaseApp, this.nodeUuid, "")
			this.log.Debug("player %v disconnect", playerIdIns.(int64))
		} else {
			this.log.Debug("player disconnect")
		}
		return nil
	})
}

func (this *Login) BeanUninit() {

}

func NewLogin() *Login {
	return &Login{}
}

func (this *Login) ProcessLogin(netChannel gproto.INetChannel, ctx gnet.ISessionCtx, playerId int64) netproto.E_Code {
	// 重复登陆请求
	if ctx.Get(gdb.IDBAttributeGroupType) != nil {
		this.log.Debug("player %v login failed, relogin in same connection.", playerId)
		return netproto.E_Code_E_INVALID_OPT
	}
	// 同服重登陆处理
	oldChannel := this.playerMgr.GetNetchannelByPlayerId(playerId)
	if oldChannel != nil {
		oldChannel.Close()
		this.log.Debug("player %v relogin in same server", playerId)
		return netproto.E_Code_E_RELOGIN
	}

	// 禁止登陆
	if this.playerMgr.IsForbidLogin(playerId) {
		this.log.Debug("player %v login failed, forbid login.", playerId)
		return netproto.E_Code_E_INVALID_OPT
	}

	// 获取玩家数据访问接口
	dataAccessor := this.playerMgr.GetDBDataAccessor(playerId)
	attributeGroup := dataAccessor.GetAttributeGroup()

	// 跨服登陆处理
	code := this.doClusterLogin(playerId, attributeGroup)
	if code != netproto.E_Code_E_OK {
		this.log.Debug("player %v login failed.", playerId)
		return code
	}
	if !this.playerMgr.PlayerLogin(playerId, netChannel, ctx, dataAccessor) {
		this.log.Debug("player %v login failed. same time login.", playerId)
		return netproto.E_Code_E_RELOGIN
	}
	ctx.Set(entity.PlayerIdKey, playerId)
	this.servicePlayerMsgQueue.GoProcessPlayerMsg(ctx)

	return netproto.E_Code_E_OK
}

func (this *Login) doClusterLogin(playerId int64, attributeGroup gdb.IDBAttributeGroup) netproto.E_Code {
	// 异服重复登陆检查
	if attributeGroup.CasSetAttribute(meta.BaseApp, "", this.nodeUuid) {
		return netproto.E_Code_E_OK
	}

	oldBaseAppUuid, err := attributeGroup.GetAttribute(meta.BaseApp)
	if err != nil {
		this.log.WithFields(glog.LF{"playerId": playerId, "attr": meta.BaseApp}).
			Warn("get player base app attribute error")
		return netproto.E_Code_E_ERROR
	}

	if oldBaseAppUuid == this.nodeUuid {
		// 同服有两个人同时登陆一个账号
		this.log.WithFields(glog.LF{"playerId": playerId, "uuid": this.nodeUuid}).
			Warn("some server relogin process")
		return netproto.E_Code_E_RELOGIN
	}

	if !this.nodeMgr.IsNodeActive(oldBaseAppUuid) {
		// 服务器已死，抢先登陆
		if attributeGroup.CasSetAttribute(meta.BaseApp, oldBaseAppUuid, this.nodeUuid) {
			this.log.WithFields(glog.LF{"playerId": playerId}).
				Warn("player relogin in race server success")
			return netproto.E_Code_E_OK
		} else {
			this.log.WithFields(glog.LF{"playerId": playerId}).
				Warn("player relogin in race server failed")
			return netproto.E_Code_E_RELOGIN
		}
	}

	// 重复登陆处理，通知踢人
	err = this.nodeMq.Push(oldBaseAppUuid, &netproto.SS_KickPlayer{PlayerId:playerId})
	if err != nil {
		this.log.Debug("node message error %v", err)
	}
	this.log.WithFields(glog.LF{"playerId": playerId}).
		Warn("player relogin failed, kick other server player")
	return netproto.E_Code_E_RELOGIN
}

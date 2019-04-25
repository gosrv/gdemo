package service

import (
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gbase/gutil"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/goioc/util"
)

type PlayerMsgQueue struct {
	msgCenter *common.ClusterMsgCenter `bean:""`
}

func newPlayerMsgQueue() *PlayerMsgQueue {
	return &PlayerMsgQueue{}
}

func (this *PlayerMsgQueue) GoProcessPlayerMsg(ctx gnet.ISessionCtx) {
	playerId := ctx.Get(entity.PlayerIdKey).(int64)
	netChannel := ctx.Get(gproto.INetChannelType).(gproto.INetChannel)
	util.Verify(playerId > 0)
	gutil.RecoverGo(func() {
		this.msgCenter.BlockProcessPlayerMsg(playerId, netChannel)
	})
}

func init() {
	common.BeansInit = append(common.BeansInit, newPlayerMsgQueue())
}

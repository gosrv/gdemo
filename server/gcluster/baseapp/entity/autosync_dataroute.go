package entity

import (
	"github.com/gosrv/gbase/gl"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gbase/gutil"
	"reflect"
)

/**
数据自动同步处理路由器
在处理完玩家消息后主动触发一次增量数据更新
*/
type AutoSyncDataRoute struct {
	raw        gproto.IRoute
	playerSync *PlayerDataSync
}

func NewAutoSyncDataRoute() *AutoSyncDataRoute {
	return &AutoSyncDataRoute{}
}

func (this *AutoSyncDataRoute) GetRouteKeys() []interface{} {
	return this.raw.GetRouteKeys()
}

func (this *AutoSyncDataRoute) SetDelegate(raw gproto.IRoute) {
	this.raw = raw
}

func (this *AutoSyncDataRoute) Connect(key interface{}, processor gproto.FProcessor) {
	gl.Panic("not support operation")
}

func (this *AutoSyncDataRoute) GetRoute(key interface{}) []gproto.FProcessor {
	return this.raw.GetRoute(key)
}

func (this *AutoSyncDataRoute) Trigger(from interface{}, key interface{}, value interface{}) interface{} {
	ctx := from.(gnet.ISessionCtx)
	// 先处理消息
	rep := this.raw.Trigger(from, key, value)
	// 在ctx中查找玩家同步模块
	if this.playerSync == nil {
		psync := ctx.Get(reflect.TypeOf((*PlayerDataSync)(nil)))
		if psync != nil {
			this.playerSync = psync.(*PlayerDataSync)
		}
	}
	// 尝试增量更新同步
	if this.playerSync != nil {
		this.playerSync.TrySyncDirtyData(false)
	}
	// 日志记录
	if gutil.IsNilValue(rep) {
		gl.Debug("msg process %v:%v	->	nil", key, value)
	} else {
		gl.Debug("msg process %v:%v	->	%v:%v",
			key, value, reflect.TypeOf(rep), rep)
	}
	return rep
}

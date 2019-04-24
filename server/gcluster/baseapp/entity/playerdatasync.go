package entity

import (
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gbase/datasync"
	"github.com/gosrv/gbase/gdb"
	"github.com/gosrv/gbase/gl"
	"github.com/gosrv/gbase/gproto"
	"reflect"
	"time"
)

const (
	IdxPlayerDataMonitor = 0
	IdxPlayerInfoMonitor = 1
)

/**
数据自动同步，同步目标有两个
1. 客户端，采用增量更新，如果有改变就竟可能快的更新
2. 数据库，采用全量更新，有更新间隔，并不是每次都会更新，有多重数据库，第一重做热数据缓存
*/
type PlayerDataSync struct {
	// 玩家数据，会同步到客户端，会存入数据库
	data *PlayerData
	// 玩家数据，会同步到客户端，不会存入数据库
	info *PlayerInfo
	// 缓存刷新间隔
	cacheFlushGapSec int64
	// 数据库刷新间隔
	dbFlushGapSec int64
	// 网络通道
	netChannel gproto.INetChannel
	// 存储器，会有多级存储
	dataSaver *gdb.TheDataSaverChain
	// 上一次更新每日数据时间
	prefreshDailyDataTime time.Time
	// 脏数据监控器，就是PlayerData和PlayerInfo的监控
	dirtyMonitor datasync.IDirtyContainerMark
	// 上一次flush数据的时间
	preFlushCacheTime int64
	preFlushDbTime    int64
	// 设置缓存超时时间的接口
	dataExpireable gdb.IDataExpireable
}
var PPlayerDataSyncType = reflect.TypeOf((*PlayerDataSync)(nil))

func NewPlayerDataSync(data *PlayerData, info *PlayerInfo, cacheFlushGapSec int64,
	dbFlushGapSec int64, netChannel gproto.INetChannel,
	dataSaver *gdb.TheDataSaverChain, attributeGroup gdb.IDBAttributeGroup) *PlayerDataSync {
	datasync := &PlayerDataSync{
		data:             data,
		info:             info,
		cacheFlushGapSec: cacheFlushGapSec,
		dbFlushGapSec:    dbFlushGapSec,
		netChannel:       netChannel,
		dataSaver:        dataSaver,
		dirtyMonitor:     datasync.NewDirtyContainerVector(),
	}
	// 如果有多级属性存储，我们可以把第一级设为可过期
	// 这里使用了redis和mongo，相当于把redis设为热数据
	if dataExpireable, ok := attributeGroup.(gdb.IDataExpireable); ok {
		datasync.dataExpireable = dataExpireable
	}
	datasync.data.Init(datasync.dirtyMonitor, IdxPlayerDataMonitor)
	datasync.info.Init(datasync.dirtyMonitor, IdxPlayerInfoMonitor)
	return datasync
}

func (this *PlayerDataSync) markDBDataDirty() {
	if this.preFlushCacheTime == 0 {
		this.preFlushCacheTime = time.Now().Unix()
	}
	if this.preFlushDbTime == 0 {
		this.preFlushDbTime = time.Now().Unix()
	}
}

func (this *PlayerDataSync) clearDBDataDirty() {
	this.preFlushCacheTime = 0
	this.preFlushDbTime = 0
}

func (this *PlayerDataSync) TrySyncDirtyData(flushdb bool) {
	// 向客户的同步脏数据
	if this.dirtyMonitor.IsDirty(IdxPlayerInfoMonitor) {
		this.netChannel.Send(this.info.ToProtoDirty())
	}

	if this.dirtyMonitor.IsDirty(IdxPlayerDataMonitor) {
		this.markDBDataDirty()
		this.netChannel.Send(this.data.ToProtoDirty())
	}
	// 清理脏数据
	this.dirtyMonitor.ClearDirty()

	// 是否达到flushdb的时间间隔？
	if this.preFlushDbTime > 0 && (flushdb || time.Now().Unix()-this.preFlushCacheTime > this.cacheFlushGapSec) {
		this.clearDBDataDirty()
		bindata, err := proto.Marshal(this.data.ToProto())
		if err != nil {
			gl.Panic("save player data error %v", err)
		}
		this.dataSaver.SaveDepth(string(bindata), 0)
		if this.dataExpireable != nil {
			err = this.dataExpireable.SetExpireDuration(7 * 24 * time.Hour)
			if err != nil {
				gl.Warn("set cache data expire duration failed")
			}
		}
	}

	// 是否达到flushcache的时间间隔？
	if this.preFlushCacheTime > 0 && time.Now().Unix()-this.preFlushCacheTime > this.cacheFlushGapSec {
		this.preFlushCacheTime = 0
		bindata, err := proto.Marshal(this.data.ToProto())
		if err != nil {
			gl.Panic("save player data error %v", err)
		}
		this.dataSaver.SaveDepth(string(bindata), 1)
	}
}

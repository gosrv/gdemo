package service

import (
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gbase/gdb"
	"github.com/gosrv/gbase/gdb/dbaccessor"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/common/meta"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc/util"
	"reflect"
	"sort"
	"strconv"
	"sync"
)

const (
	cacheFlushGapSec = 30
	dbFlushGapSec    = 300
)

type PlayerMgr struct {
	playerId2Channel sync.Map
	log              glog.IFieldLogger `log:"app"`

	messageQueueFactories   []gdb.IMessageQueueFactory     `bean:""`
	attributeGroupFactories []gdb.IDBAttributeGroupFactory `bean:""`
	dbAccessorFactory       *dbaccessor.DBDataAccessorFactory
}

func (this *PlayerMgr) BeanInit() {
	if len(this.messageQueueFactories) == 0 {
		util.Panic("no message queue found")
	}
	if len(this.attributeGroupFactories) == 0 {
		util.Panic("no attribute group found")
	}
	sort.Slice(this.messageQueueFactories, func(i, j int) bool {
		return this.messageQueueFactories[i].GetPriority() < this.messageQueueFactories[j].GetPriority()
	})
	sort.Slice(this.attributeGroupFactories, func(i, j int) bool {
		return this.attributeGroupFactories[i].GetPriority() < this.attributeGroupFactories[j].GetPriority()
	})
	this.dbAccessorFactory = dbaccessor.NewDBDataAccessorFactory(this.messageQueueFactories[0], this.attributeGroupFactories)
}

func (this *PlayerMgr) BeanUninit() {

}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{}
}

func (this *PlayerMgr) IsForbidLogin(playerId int64) bool {
	return false
}

func (this *PlayerMgr) GetDBDataAccessor(playerId int64) *dbaccessor.DBDataAccessor {
	return this.dbAccessorFactory.GetDataAccessor(meta.PlayerAttribute, strconv.FormatInt(playerId, 10))
}

func (this *PlayerMgr) LoadPlayerData(playerId int64, loader *gdb.TheDataLoaderChain) *entity.PlayerData {
	val, err := loader.Load()
	if err != nil {
		this.log.Debug("load player data error %v", err)
	}

	pd := entity.NewPlayerData()
	if len(val) == 0 {
		this.InitPlayerData(playerId, pd)
	} else {
		npd := &netproto.PlayerData{}
		err = proto.Unmarshal([]byte(val), npd)
		if err != nil {
			this.log.Debug("unmarshal player data error %v", err)
		}
		pd.FromProto(npd)
	}
	return pd
}

func (this *PlayerMgr) InitPlayerData(playerId int64, playerData *entity.PlayerData) {
	baseInfo := playerData.GetBaseInfo()
	baseInfo.SetId(playerId)
	baseInfo.SetLevel(1)
	baseInfo.SetName("name" + strconv.FormatInt(playerId, 10))
}

func (this *PlayerMgr) PlayerLogin(playerId int64, netChannel gproto.INetChannel, ctx gnet.ISessionCtx, dataAccessor *dbaccessor.DBDataAccessor) bool {
	_, loaded := this.playerId2Channel.LoadOrStore(playerId, netChannel)
	if loaded {
		return false
	}
	// 玩家相关数据数据初始化
	// db数据读取
	attributeGroup := dataAccessor.GetAttributeGroup()
	// 加载器
	loader := dataAccessor.GetDataLoader(meta.PlayerBaseData)
	// 存储器
	saver := dataAccessor.GetDataSaver(meta.PlayerBaseData)
	// 加载玩家数据
	playerData := this.LoadPlayerData(playerId, loader)
	playerInfo := entity.NewPlayerInfo()
	onlineData := entity.NewPlayerOnlineData()
	// 数据同步
	syncData := entity.NewPlayerDataSync(playerData, playerInfo,
		cacheFlushGapSec, dbFlushGapSec, netChannel,
		saver, attributeGroup)
	// 将相关属性保存到session上
	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(playerData), playerData)
	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(playerInfo), playerInfo)
	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(onlineData), onlineData)
	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(syncData), syncData)

	ctx.SetAttribute(gnet.ScopeSession, reflect.TypeOf(dataAccessor), dataAccessor)
	ctx.SetAttribute(gnet.ScopeSession, gdb.IDBAttributeGroupType, attributeGroup)
	ctx.SetAttribute(gnet.ScopeSession, gdb.IMessageQueueType,
		this.dbAccessorFactory.GetMessageQueue(meta.PlayerAttribute, strconv.FormatInt(playerId, 10)))

	return true
}

func (this *PlayerMgr) PlayerLogout(playerId int64) {
	this.playerId2Channel.Delete(playerId)
}

func (this *PlayerMgr) GetNetchannelByPlayerId(playerId int64) gproto.INetChannel {
	net, ok := this.playerId2Channel.Load(playerId)
	if !ok {
		return nil
	}
	return net.(gproto.INetChannel)
}

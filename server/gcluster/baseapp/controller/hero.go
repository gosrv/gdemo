package controller

import (
	"github.com/gosrv/gbase/controller"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/gcluster/gcluster/table"
	"github.com/gosrv/glog"
	"math/rand"
)

/**
逻辑消息控制器
*/
type Hero struct {
	log glog.IFieldLogger `log:"app"`
	// 控制器标记
	controller.IController
	serviceHero *service.Hero `bean:""`
	serviceShop *service.Shop `bean:""`
}

func newHero() *Hero {
	return &Hero{
		// 路由收集器，它会收集这样的函数作为路由器：
		// 第一个变量是gnet.ISessionCtx，第二个是消息，可以返回一个一个消息，也可以不返回
		IController: controller.NewTypeController(""),
	}
}

// 抽卡
func (this *Hero) Draw(ctx gnet.ISessionCtx, msg *netproto.CS_HeroDraw, playerData *entity.PlayerData) *netproto.SC_HeroDraw {
	if this.serviceHero.IsPackLimit(playerData) {
		return &netproto.SC_HeroDraw{Code: netproto.E_Code_E_LIMIT_CAPACITY}
	}
	switch msg.Num {
	case 1:
		// 消耗钻石
		if playerData.GetBaseInfo().GetDiamond() < table.TableMgr.TableCommon[7].Val {
			return &netproto.SC_HeroDraw{Code: netproto.E_Code_E_LIMIT_DIAMOND}
		}
		this.serviceShop.ConsumeDiamond(playerData, table.TableMgr.TableCommon[7].Val)
		// 可用英雄id 从 1-100
		this.serviceHero.GiveHero(playerData, rand.Int31()%100+1)
	case 10:
		// 消耗钻石
		if playerData.GetBaseInfo().GetDiamond() < table.TableMgr.TableCommon[8].Val {
			return &netproto.SC_HeroDraw{Code: netproto.E_Code_E_LIMIT_DIAMOND}
		}
		this.serviceShop.ConsumeDiamond(playerData, table.TableMgr.TableCommon[8].Val)
		for i := 0; i < 10; i++ {
			this.serviceHero.GiveHero(playerData, rand.Int31()%100+1)
		}
	default:
		return &netproto.SC_HeroDraw{Code: netproto.E_Code_E_INVALID_OPT}
	}
	return &netproto.SC_HeroDraw{Code: netproto.E_Code_E_OK}
}

// 升级英雄等级
func (this *Hero) UpLevel(ctx gnet.ISessionCtx, msg *netproto.CS_HeroUplevel, playerData *entity.PlayerData) *netproto.SC_HeroUpLevel {
	// 消耗金币
	consumeGold := table.TableMgr.TableCommon[9].Val * playerData.GetBaseInfo().GetLevel() * playerData.GetBaseInfo().GetLevel()
	if playerData.GetBaseInfo().GetGold() < consumeGold {
		return &netproto.SC_HeroUpLevel{Code: netproto.E_Code_E_LIMIT_GOLD}
	}
	this.serviceShop.GiveGold(playerData, consumeGold)
	return &netproto.SC_HeroUpLevel{Code: netproto.E_Code_E_OK}
}

// 英雄装备
func (this *Hero) Equip(ctx gnet.ISessionCtx, msg *netproto.CS_HeroEquip, playerData *entity.PlayerData) *netproto.SC_HeroEquip {
	return &netproto.SC_HeroEquip{Code: netproto.E_Code_E_INVALID_OPT}
}

func init() {
	common.BeansInit = append(common.BeansInit, newHero())
}

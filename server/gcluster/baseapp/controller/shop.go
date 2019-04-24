package controller

import (
	"github.com/gosrv/gbase/controller"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/gcluster/gcluster/table"
	"github.com/gosrv/glog"
)

/**
逻辑消息控制器
*/
type Shop struct {
	log glog.IFieldLogger `log:"app"`
	// 控制器标记
	controller.IController
	serviceShop *service.Shop	`bean:""`
}

func NewShop() *Shop {
	return &Shop{
		// 路由收集器，它会收集这样的函数作为路由器：
		// 第一个变量是gnet.ISessionCtx，第二个是消息，可以返回一个一个消息，也可以不返回
		IController: controller.NewTypeController(""),
	}
}

// 购买钻石
func (this *Shop) AddDiamond(ctx gnet.ISessionCtx, msg *netproto.CS_ShopAddDiamond, playerData *entity.PlayerData) *netproto.SC_ShopAddDiamond {
	switch msg.Idx {
	case 0:
		this.serviceShop.GiveDiamond(playerData, table.TableMgr.TableCommon[1].Val)
	case 1:
		this.serviceShop.GiveDiamond(playerData, table.TableMgr.TableCommon[2].Val)
	case 2:
		this.serviceShop.GiveDiamond(playerData, table.TableMgr.TableCommon[3].Val)
	case 3:
		this.serviceShop.GiveDiamond(playerData, table.TableMgr.TableCommon[4].Val)
	default:
		return &netproto.SC_ShopAddDiamond{Code:netproto.E_Code_E_INVALID_OPT}
	}
	return &netproto.SC_ShopAddDiamond{Code:netproto.E_Code_E_OK}
}

// 购买金币
func (this *Shop) BuyGold(ctx gnet.ISessionCtx, msg *netproto.CS_ShopBuyGold, playerData *entity.PlayerData) *netproto.SC_ShopBuyGold {
	if playerData.GetBaseInfo().GetDiamond() < table.TableMgr.TableCommon[6].Val {
		return &netproto.SC_ShopBuyGold{Code:netproto.E_Code_E_LIMIT_DIAMOND}
	}
	this.serviceShop.ConsumeDiamond(playerData, table.TableMgr.TableCommon[6].Val)
	this.serviceShop.GiveGold(playerData, table.TableMgr.TableCommon[5].Val * table.TableMgr.TableCommon[6].Val)
	return &netproto.SC_ShopBuyGold{Code:netproto.E_Code_E_OK}
}


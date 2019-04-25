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
	"time"
)

/**
逻辑消息控制器
*/
type Chapter struct {
	log glog.IFieldLogger `log:"app"`
	// 控制器标记
	controller.IController
	serviceShop *service.Shop `bean:""`
}

func newChapter() *Chapter {
	return &Chapter{
		// 路由收集器，它会收集这样的函数作为路由器：
		// 第一个变量是gnet.ISessionCtx，第二个是消息，可以返回一个一个消息，也可以不返回
		IController: controller.NewTypeController(""),
	}
}

// 主线关卡挑战
func (this *Chapter) Challenge(ctx gnet.ISessionCtx, msg *netproto.CS_ChapterChallenge, playerData *entity.PlayerData) *netproto.SC_ChapterChallenge {
	if table.TableMgr.TableChapter[playerData.GetChapter().GetLevel()+1] == nil {
		return &netproto.SC_ChapterChallenge{Code: netproto.E_Code_E_INVALID_OPT}
	}
	playerData.GetChapter().SetLevel(playerData.GetChapter().GetLevel() + 1)
	return &netproto.SC_ChapterChallenge{Code: netproto.E_Code_E_OK}
}

// 获取挂机奖励
func (this *Chapter) GetPrize(ctx gnet.ISessionCtx, msg *netproto.CS_ChapterGetPrize, playerData *entity.PlayerData) *netproto.SC_ChapterGetPrize {
	// 收益金币 = 挂机时间 * 每秒收益
	cfgChapter := table.TableMgr.TableChapter[playerData.GetChapter().GetLevel()]
	now := time.Now().Unix()
	delta := now - playerData.GetChapter().GetPrizeCheckTime()
	playerData.GetChapter().SetPrizeCheckTime(now)
	if delta <= 0 {
		return &netproto.SC_ChapterGetPrize{Code: netproto.E_Code_E_OK}
	}
	gold := cfgChapter.GetGold() * int32(delta)
	this.serviceShop.GiveGold(playerData, gold)
	return &netproto.SC_ChapterGetPrize{Code: netproto.E_Code_E_OK}
}

func init() {
	common.BeansInit = append(common.BeansInit, newChapter())
}

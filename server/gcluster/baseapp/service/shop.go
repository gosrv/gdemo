package service

import (
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/glog"
)

type Shop struct {
	log glog.IFieldLogger `log:"app"`
}

func newShop() *Shop {
	return &Shop{}
}

func (this *Shop) GiveDiamond(data *entity.PlayerData, diamond int32) {
	data.GetBaseInfo().SetDiamond(data.GetBaseInfo().GetDiamond() + diamond)
}

func (this *Shop) ConsumeDiamond(data *entity.PlayerData, diamond int32) {
	data.GetBaseInfo().SetDiamond(data.GetBaseInfo().GetDiamond() - diamond)
}

func (this *Shop) GiveGold(data *entity.PlayerData, gold int32) {
	data.GetBaseInfo().SetGold(data.GetBaseInfo().GetGold() + gold)
}

func (this *Shop) ConsumeGold(data *entity.PlayerData, gold int32) {
	data.GetBaseInfo().SetGold(data.GetBaseInfo().GetGold() - gold)
}

func init() {
	common.BeansInit = append(common.BeansInit, newShop())
}

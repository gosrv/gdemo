package service

import (
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/gcluster/gcluster/table"
	"github.com/gosrv/glog"
)

type Hero struct {
	log glog.IFieldLogger `log:"app"`
}

func newHero() *Hero {
	return &Hero{}
}

// 背包是否达到上限
func (this *Hero) IsPackLimit(data *entity.PlayerData) bool {
	return int(data.GetHeroPack().GetLimit()) <= data.GetHeroPack().GetHeros().Size()
}

func (this *Hero) GiveHero(data *entity.PlayerData, id int32) bool {
	if this.IsPackLimit(data) {
		return false
	}

	if table.TableMgr.TableHero[id] == nil {
		this.log.Debug("hero %v config miss", id)
		return false
	}

	hero := data.GetHeroPack().GetHeros().Get(id)
	if hero != nil {
		hero.SetNum(hero.GetNum() + 1)
	} else {
		hero := entity.NewHero()
		hero.SetId(id)
		hero.SetLevel(1)
		hero.SetNum(0)
		data.GetHeroPack().GetHeros().Set(id, hero)
	}

	return true
}

func init() {
	common.BeansInit = append(common.BeansInit, newHero())
}

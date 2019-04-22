package common

import (
	"encoding/json"
	"github.com/gosrv/gbase/app"
	"github.com/gosrv/gmx"
	"github.com/gosrv/goioc/util"
)

type GmxAppStats struct {
	gmxMgr gmx.IMXManager `bean:""`
}

func NewGmxAppStats() *GmxAppStats {
	return &GmxAppStats{}
}

func (this *GmxAppStats) BeanInit() {
	err := this.gmxMgr.AddItemOpt("app.stats.mem", gmx.FuncGetter(func() (string, error) {
		stats := app.AppStats.GetMemStats()
		val, err := json.Marshal(stats)
		return string(val), err
	}), nil, "")
	util.VerifyNoError(err)

	err = this.gmxMgr.AddItemOpt("app.stats.runtime", gmx.FuncGetter(func() (string, error) {
		stats := app.AppStats.GetRuntimeStats()
		val, err := json.Marshal(stats)
		return string(val), err
	}), nil, "")
	util.VerifyNoError(err)
}

func (this *GmxAppStats) BeanUninit() {

}

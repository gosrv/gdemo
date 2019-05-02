package service

import (
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/glog"
)

type Logic struct {
	log glog.IFieldLogger    `log:"app"`
}

func newLogic() *Logic {
	return &Logic{}
}

func init() {
	common.BeansInit = append(common.BeansInit, newLogic())
}

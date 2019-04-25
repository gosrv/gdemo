package service

import (
	"github.com/gosrv/excelreader/tableloader"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/gcluster/gcluster/table"
)

type Table struct {
}

func newTable() *Table {
	return &Table{}
}

func (this *Table) BeanInit() {
	tableloader.DataDir = "gcluster/conf/table"
	this.Reload()
}

func (this *Table) BeanUninit() {

}

func (this *Table) Reload() {
	table.TableMgr = table.Load()
}

func init() {
	common.BeansInit = append(common.BeansInit, newTable())
}

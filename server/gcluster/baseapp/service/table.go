package service

import (
	"github.com/gosrv/excelreader/tableloader"
	"github.com/gosrv/gcluster/gcluster/table"
)

type Table struct {

}

func NewTable() *Table {
	return &Table{}
}

func (this *Table) BeanInit() {
	tableloader.DataDir = "gcluster/conf/table"
	this.Reload()
}

func (this *Table) BeanUninit() {

}

func (this *Table) Reload()  {
	table.TableMgr = table.Load()
}




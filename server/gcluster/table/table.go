// generate by excelreader
package table

import (
    "github.com/golang/protobuf/proto"
    "github.com/gosrv/excelreader/tableloader"
)

var TableMgr *tableMgr

type tableMgr struct {
    TableChapter map[int32]*Chapter
    TableCommon map[int32]*Common
    TableEquip map[int32]*Equip
    TableHero map[int32]*Hero
    
}

func (this *tableMgr)loadChapter() {
    data := tableloader.Load("Chapter")
    itemArray := &ChapterArray{}
    err := proto.Unmarshal(data, itemArray)
    if err != nil {panic(err)}
    this.TableChapter = make(map[int32]*Chapter)
    for i := 0; i < len(itemArray.Keys); i++ {
        this.TableChapter[itemArray.Keys[i]] = itemArray.Items[i]
    }
}
func (this *tableMgr)loadCommon() {
    data := tableloader.Load("Common")
    itemArray := &CommonArray{}
    err := proto.Unmarshal(data, itemArray)
    if err != nil {panic(err)}
    this.TableCommon = make(map[int32]*Common)
    for i := 0; i < len(itemArray.Keys); i++ {
        this.TableCommon[itemArray.Keys[i]] = itemArray.Items[i]
    }
}
func (this *tableMgr)loadEquip() {
    data := tableloader.Load("Equip")
    itemArray := &EquipArray{}
    err := proto.Unmarshal(data, itemArray)
    if err != nil {panic(err)}
    this.TableEquip = make(map[int32]*Equip)
    for i := 0; i < len(itemArray.Keys); i++ {
        this.TableEquip[itemArray.Keys[i]] = itemArray.Items[i]
    }
}
func (this *tableMgr)loadHero() {
    data := tableloader.Load("Hero")
    itemArray := &HeroArray{}
    err := proto.Unmarshal(data, itemArray)
    if err != nil {panic(err)}
    this.TableHero = make(map[int32]*Hero)
    for i := 0; i < len(itemArray.Keys); i++ {
        this.TableHero[itemArray.Keys[i]] = itemArray.Items[i]
    }
}
func Load() *tableMgr {
    mgr := &tableMgr{}
    mgr.loadChapter()
    mgr.loadCommon()
    mgr.loadEquip()
    mgr.loadHero()
    
    return mgr
}
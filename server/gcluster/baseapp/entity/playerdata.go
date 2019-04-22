package entity

import (
	"github.com/golang/protobuf/proto"
    "github.com/gosrv/gbase/datasync"
    "github.com/gosrv/gcluster/gcluster/proto"
)
            
                
type ItemIterPlayerDatapets func(k int32, v *Pet)
type ItemStatusIterPlayerDatapets func(k int32, v *Pet, dirty bool)
    
type ContainerPlayerDatapets struct {
    _impl *datasync.DirtyContainerVector
}
func (this *ContainerPlayerDatapets)Get(key int32) *Pet {
    return this._impl.Get(key).(*Pet)
}
func (this *ContainerPlayerDatapets)Set(key int32, value *Pet) {
    this._impl.Set(key, value)
}
func (this *ContainerPlayerDatapets)Foreach(iter ItemIterPlayerDatapets) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(int32), v.(*Pet))})
}
func (this *ContainerPlayerDatapets)ForeachStatus(iter ItemStatusIterPlayerDatapets) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(int32), v.(*Pet), dirty)})
}
func (this *ContainerPlayerDatapets)ForeachDirty(iter ItemIterPlayerDatapets) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(int32), v.(*Pet))})
}
func (this *ContainerPlayerDatapets)Size() int {
    return this._impl.Size()
}
func (this *ContainerPlayerDatapets)Clear() {
    this._impl.Clear()
}

func NewPlayerDatapetsProto(k int32, v *Pet) *netproto.PlayerDatapets {
        return &netproto.PlayerDatapets{
            Key:&k,
            Val:v.ToProto(),
        }
}

func NewPlayerDatapetsProtoDirty(k int32, v *Pet) *netproto.PlayerDatapets {
        return &netproto.PlayerDatapets{
            Key:&k,
            Val:v.ToProtoDirty(),
        }
}
            
                
type ItemIterPlayerDatapacks func(k int32, v *Pack)
type ItemStatusIterPlayerDatapacks func(k int32, v *Pack, dirty bool)
    
type ContainerPlayerDatapacks struct {
    _impl *datasync.DirtyContainerVector
}
func (this *ContainerPlayerDatapacks)Get(key int32) *Pack {
    return this._impl.Get(key).(*Pack)
}
func (this *ContainerPlayerDatapacks)Set(key int32, value *Pack) {
    this._impl.Set(key, value)
}
func (this *ContainerPlayerDatapacks)Foreach(iter ItemIterPlayerDatapacks) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(int32), v.(*Pack))})
}
func (this *ContainerPlayerDatapacks)ForeachStatus(iter ItemStatusIterPlayerDatapacks) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(int32), v.(*Pack), dirty)})
}
func (this *ContainerPlayerDatapacks)ForeachDirty(iter ItemIterPlayerDatapacks) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(int32), v.(*Pack))})
}
func (this *ContainerPlayerDatapacks)Size() int {
    return this._impl.Size()
}
func (this *ContainerPlayerDatapacks)Clear() {
    this._impl.Clear()
}

func NewPlayerDatapacksProto(k int32, v *Pack) *netproto.PlayerDatapacks {
        return &netproto.PlayerDatapacks{
            Key:&k,
            Val:v.ToProto(),
        }
}

func NewPlayerDatapacksProtoDirty(k int32, v *Pack) *netproto.PlayerDatapacks {
        return &netproto.PlayerDatapacks{
            Key:&k,
            Val:v.ToProtoDirty(),
        }
}
            
                
type ItemIterTestListListPrimitiveInt func(k int32, v int32)
type ItemStatusIterTestListListPrimitiveInt func(k int32, v int32, dirty bool)
    
type ContainerTestListListPrimitiveInt struct {
    _impl *datasync.DirtyContainerVector
}
func (this *ContainerTestListListPrimitiveInt)Get(key int32) int32 {
    return this._impl.Get(key).(int32)
}
func (this *ContainerTestListListPrimitiveInt)Set(key int32, value int32) {
    this._impl.Set(key, value)
}
func (this *ContainerTestListListPrimitiveInt)Foreach(iter ItemIterTestListListPrimitiveInt) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(int32), v.(int32))})
}
func (this *ContainerTestListListPrimitiveInt)ForeachStatus(iter ItemStatusIterTestListListPrimitiveInt) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(int32), v.(int32), dirty)})
}
func (this *ContainerTestListListPrimitiveInt)ForeachDirty(iter ItemIterTestListListPrimitiveInt) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(int32), v.(int32))})
}
func (this *ContainerTestListListPrimitiveInt)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestListListPrimitiveInt)Clear() {
    this._impl.Clear()
}

func NewTestListListPrimitiveIntProto(k int32, v int32) *netproto.TestListListPrimitiveInt {
        return &netproto.TestListListPrimitiveInt{
            Key:&k,
            Val:&v,
        }
}

func NewTestListListPrimitiveIntProtoDirty(k int32, v int32) *netproto.TestListListPrimitiveInt {
        return &netproto.TestListListPrimitiveInt{
            Key:&k,
            Val:&v,
        }
}
            
                
type ItemIterTestListListPrimitiveStr func(k int32, v string)
type ItemStatusIterTestListListPrimitiveStr func(k int32, v string, dirty bool)
    
type ContainerTestListListPrimitiveStr struct {
    _impl *datasync.DirtyContainerVector
}
func (this *ContainerTestListListPrimitiveStr)Get(key int32) string {
    return this._impl.Get(key).(string)
}
func (this *ContainerTestListListPrimitiveStr)Set(key int32, value string) {
    this._impl.Set(key, value)
}
func (this *ContainerTestListListPrimitiveStr)Foreach(iter ItemIterTestListListPrimitiveStr) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(int32), v.(string))})
}
func (this *ContainerTestListListPrimitiveStr)ForeachStatus(iter ItemStatusIterTestListListPrimitiveStr) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(int32), v.(string), dirty)})
}
func (this *ContainerTestListListPrimitiveStr)ForeachDirty(iter ItemIterTestListListPrimitiveStr) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(int32), v.(string))})
}
func (this *ContainerTestListListPrimitiveStr)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestListListPrimitiveStr)Clear() {
    this._impl.Clear()
}

func NewTestListListPrimitiveStrProto(k int32, v string) *netproto.TestListListPrimitiveStr {
        return &netproto.TestListListPrimitiveStr{
            Key:&k,
            Val:&v,
        }
}

func NewTestListListPrimitiveStrProtoDirty(k int32, v string) *netproto.TestListListPrimitiveStr {
        return &netproto.TestListListPrimitiveStr{
            Key:&k,
            Val:&v,
        }
}
            
                
type ItemIterTestListListPrimitiveCom func(k int32, v *PlayerData)
type ItemStatusIterTestListListPrimitiveCom func(k int32, v *PlayerData, dirty bool)
    
type ContainerTestListListPrimitiveCom struct {
    _impl *datasync.DirtyContainerVector
}
func (this *ContainerTestListListPrimitiveCom)Get(key int32) *PlayerData {
    return this._impl.Get(key).(*PlayerData)
}
func (this *ContainerTestListListPrimitiveCom)Set(key int32, value *PlayerData) {
    this._impl.Set(key, value)
}
func (this *ContainerTestListListPrimitiveCom)Foreach(iter ItemIterTestListListPrimitiveCom) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(int32), v.(*PlayerData))})
}
func (this *ContainerTestListListPrimitiveCom)ForeachStatus(iter ItemStatusIterTestListListPrimitiveCom) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(int32), v.(*PlayerData), dirty)})
}
func (this *ContainerTestListListPrimitiveCom)ForeachDirty(iter ItemIterTestListListPrimitiveCom) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(int32), v.(*PlayerData))})
}
func (this *ContainerTestListListPrimitiveCom)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestListListPrimitiveCom)Clear() {
    this._impl.Clear()
}

func NewTestListListPrimitiveComProto(k int32, v *PlayerData) *netproto.TestListListPrimitiveCom {
        return &netproto.TestListListPrimitiveCom{
            Key:&k,
            Val:v.ToProto(),
        }
}

func NewTestListListPrimitiveComProtoDirty(k int32, v *PlayerData) *netproto.TestListListPrimitiveCom {
        return &netproto.TestListListPrimitiveCom{
            Key:&k,
            Val:v.ToProtoDirty(),
        }
}
            
            
type ItemIterTestMapMapPrimitiveIntInt func(k int32, v int32)
type ItemStatusIterTestMapMapPrimitiveIntInt func(k int32, v int32, dirty bool)
    
type ContainerTestMapMapPrimitiveIntInt struct {
    _impl *datasync.DirtyContainerMap
}
func (this *ContainerTestMapMapPrimitiveIntInt)Get(key int32) int32 {
    return this._impl.Get(key).(int32)
}
func (this *ContainerTestMapMapPrimitiveIntInt)Set(key int32, value int32) {
    this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveIntInt)Foreach(iter ItemIterTestMapMapPrimitiveIntInt) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(int32), v.(int32))})
}
func (this *ContainerTestMapMapPrimitiveIntInt)ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveIntInt) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(int32), v.(int32), dirty)})
}
func (this *ContainerTestMapMapPrimitiveIntInt)ForeachDirty(iter ItemIterTestMapMapPrimitiveIntInt) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(int32), v.(int32))})
}
func (this *ContainerTestMapMapPrimitiveIntInt)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveIntInt)Clear() {
    this._impl.Clear()
}

func NewTestMapMapPrimitiveIntIntProto(k int32, v int32) *netproto.TestMapMapPrimitiveIntInt {
        return &netproto.TestMapMapPrimitiveIntInt{
            Key:&k,
            Val:&v,
        }
}

func NewTestMapMapPrimitiveIntIntProtoDirty(k int32, v int32) *netproto.TestMapMapPrimitiveIntInt {
        return &netproto.TestMapMapPrimitiveIntInt{
            Key:&k,
            Val:&v,
        }
}
            
            
type ItemIterTestMapMapPrimitiveIntStr func(k int32, v string)
type ItemStatusIterTestMapMapPrimitiveIntStr func(k int32, v string, dirty bool)
    
type ContainerTestMapMapPrimitiveIntStr struct {
    _impl *datasync.DirtyContainerMap
}
func (this *ContainerTestMapMapPrimitiveIntStr)Get(key int32) string {
    return this._impl.Get(key).(string)
}
func (this *ContainerTestMapMapPrimitiveIntStr)Set(key int32, value string) {
    this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveIntStr)Foreach(iter ItemIterTestMapMapPrimitiveIntStr) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(int32), v.(string))})
}
func (this *ContainerTestMapMapPrimitiveIntStr)ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveIntStr) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(int32), v.(string), dirty)})
}
func (this *ContainerTestMapMapPrimitiveIntStr)ForeachDirty(iter ItemIterTestMapMapPrimitiveIntStr) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(int32), v.(string))})
}
func (this *ContainerTestMapMapPrimitiveIntStr)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveIntStr)Clear() {
    this._impl.Clear()
}

func NewTestMapMapPrimitiveIntStrProto(k int32, v string) *netproto.TestMapMapPrimitiveIntStr {
        return &netproto.TestMapMapPrimitiveIntStr{
            Key:&k,
            Val:&v,
        }
}

func NewTestMapMapPrimitiveIntStrProtoDirty(k int32, v string) *netproto.TestMapMapPrimitiveIntStr {
        return &netproto.TestMapMapPrimitiveIntStr{
            Key:&k,
            Val:&v,
        }
}
            
            
type ItemIterTestMapMapPrimitiveStrInt func(k string, v int32)
type ItemStatusIterTestMapMapPrimitiveStrInt func(k string, v int32, dirty bool)
    
type ContainerTestMapMapPrimitiveStrInt struct {
    _impl *datasync.DirtyContainerMap
}
func (this *ContainerTestMapMapPrimitiveStrInt)Get(key string) int32 {
    return this._impl.Get(key).(int32)
}
func (this *ContainerTestMapMapPrimitiveStrInt)Set(key string, value int32) {
    this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveStrInt)Foreach(iter ItemIterTestMapMapPrimitiveStrInt) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(string), v.(int32))})
}
func (this *ContainerTestMapMapPrimitiveStrInt)ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveStrInt) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(string), v.(int32), dirty)})
}
func (this *ContainerTestMapMapPrimitiveStrInt)ForeachDirty(iter ItemIterTestMapMapPrimitiveStrInt) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(string), v.(int32))})
}
func (this *ContainerTestMapMapPrimitiveStrInt)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveStrInt)Clear() {
    this._impl.Clear()
}

func NewTestMapMapPrimitiveStrIntProto(k string, v int32) *netproto.TestMapMapPrimitiveStrInt {
        return &netproto.TestMapMapPrimitiveStrInt{
            Key:&k,
            Val:&v,
        }
}

func NewTestMapMapPrimitiveStrIntProtoDirty(k string, v int32) *netproto.TestMapMapPrimitiveStrInt {
        return &netproto.TestMapMapPrimitiveStrInt{
            Key:&k,
            Val:&v,
        }
}
            
            
type ItemIterTestMapMapPrimitiveStrStr func(k string, v string)
type ItemStatusIterTestMapMapPrimitiveStrStr func(k string, v string, dirty bool)
    
type ContainerTestMapMapPrimitiveStrStr struct {
    _impl *datasync.DirtyContainerMap
}
func (this *ContainerTestMapMapPrimitiveStrStr)Get(key string) string {
    return this._impl.Get(key).(string)
}
func (this *ContainerTestMapMapPrimitiveStrStr)Set(key string, value string) {
    this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveStrStr)Foreach(iter ItemIterTestMapMapPrimitiveStrStr) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(string), v.(string))})
}
func (this *ContainerTestMapMapPrimitiveStrStr)ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveStrStr) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(string), v.(string), dirty)})
}
func (this *ContainerTestMapMapPrimitiveStrStr)ForeachDirty(iter ItemIterTestMapMapPrimitiveStrStr) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(string), v.(string))})
}
func (this *ContainerTestMapMapPrimitiveStrStr)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveStrStr)Clear() {
    this._impl.Clear()
}

func NewTestMapMapPrimitiveStrStrProto(k string, v string) *netproto.TestMapMapPrimitiveStrStr {
        return &netproto.TestMapMapPrimitiveStrStr{
            Key:&k,
            Val:&v,
        }
}

func NewTestMapMapPrimitiveStrStrProtoDirty(k string, v string) *netproto.TestMapMapPrimitiveStrStr {
        return &netproto.TestMapMapPrimitiveStrStr{
            Key:&k,
            Val:&v,
        }
}
            
            
type ItemIterTestMapMapPrimitiveIntCom func(k int32, v *BaseInfo)
type ItemStatusIterTestMapMapPrimitiveIntCom func(k int32, v *BaseInfo, dirty bool)
    
type ContainerTestMapMapPrimitiveIntCom struct {
    _impl *datasync.DirtyContainerMap
}
func (this *ContainerTestMapMapPrimitiveIntCom)Get(key int32) *BaseInfo {
    return this._impl.Get(key).(*BaseInfo)
}
func (this *ContainerTestMapMapPrimitiveIntCom)Set(key int32, value *BaseInfo) {
    this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveIntCom)Foreach(iter ItemIterTestMapMapPrimitiveIntCom) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(int32), v.(*BaseInfo))})
}
func (this *ContainerTestMapMapPrimitiveIntCom)ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveIntCom) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(int32), v.(*BaseInfo), dirty)})
}
func (this *ContainerTestMapMapPrimitiveIntCom)ForeachDirty(iter ItemIterTestMapMapPrimitiveIntCom) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(int32), v.(*BaseInfo))})
}
func (this *ContainerTestMapMapPrimitiveIntCom)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveIntCom)Clear() {
    this._impl.Clear()
}

func NewTestMapMapPrimitiveIntComProto(k int32, v *BaseInfo) *netproto.TestMapMapPrimitiveIntCom {
        return &netproto.TestMapMapPrimitiveIntCom{
            Key:&k,
            Val:v.ToProto(),
        }
}

func NewTestMapMapPrimitiveIntComProtoDirty(k int32, v *BaseInfo) *netproto.TestMapMapPrimitiveIntCom {
        return &netproto.TestMapMapPrimitiveIntCom{
            Key:&k,
            Val:v.ToProtoDirty(),
        }
}
            
            
type ItemIterTestMapMapPrimitiveStrCom func(k string, v *PlayerData)
type ItemStatusIterTestMapMapPrimitiveStrCom func(k string, v *PlayerData, dirty bool)
    
type ContainerTestMapMapPrimitiveStrCom struct {
    _impl *datasync.DirtyContainerMap
}
func (this *ContainerTestMapMapPrimitiveStrCom)Get(key string) *PlayerData {
    return this._impl.Get(key).(*PlayerData)
}
func (this *ContainerTestMapMapPrimitiveStrCom)Set(key string, value *PlayerData) {
    this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveStrCom)Foreach(iter ItemIterTestMapMapPrimitiveStrCom) {
    this._impl.Foreach(func(k interface{}, v interface{}) {iter(k.(string), v.(*PlayerData))})
}
func (this *ContainerTestMapMapPrimitiveStrCom)ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveStrCom) {
    this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) {iter(k.(string), v.(*PlayerData), dirty)})
}
func (this *ContainerTestMapMapPrimitiveStrCom)ForeachDirty(iter ItemIterTestMapMapPrimitiveStrCom) {
    this._impl.ForeachDirty(func(k interface{}, v interface{}) {iter(k.(string), v.(*PlayerData))})
}
func (this *ContainerTestMapMapPrimitiveStrCom)Size() int {
    return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveStrCom)Clear() {
    this._impl.Clear()
}

func NewTestMapMapPrimitiveStrComProto(k string, v *PlayerData) *netproto.TestMapMapPrimitiveStrCom {
        return &netproto.TestMapMapPrimitiveStrCom{
            Key:&k,
            Val:v.ToProto(),
        }
}

func NewTestMapMapPrimitiveStrComProtoDirty(k string, v *PlayerData) *netproto.TestMapMapPrimitiveStrCom {
        return &netproto.TestMapMapPrimitiveStrCom{
            Key:&k,
            Val:v.ToProtoDirty(),
        }
}


    type PlayerData struct {
        *datasync.DirtyContainerMarkVector
                _baseInfo *BaseInfo
                _pets *ContainerPlayerDatapets
                _daily *Daily
                _packs *ContainerPlayerDatapacks
    }
    // ctor
    func NewPlayerData() *PlayerData {
        ins := &PlayerData {
            DirtyContainerMarkVector:datasync.NewDirtyContainerMarkVector(),
        }
                ins._baseInfo = NewBaseInfo()
                ins._baseInfo.Init(ins, 1)
                ins._pets = &ContainerPlayerDatapets {
                    _impl : datasync.NewDirtyContainerVector(),
                }
                ins._pets._impl.Init(ins, 2)
                ins._daily = NewDaily()
                ins._daily.Init(ins, 3)
                ins._packs = &ContainerPlayerDatapacks {
                    _impl : datasync.NewDirtyContainerVector(),
                }
                ins._packs._impl.Init(ins, 4)
        return ins
    }
// BaseInfo getter and setter
            func (this *PlayerData)GetBaseInfo() *BaseInfo {
                return this._baseInfo
            }
// Pets getter and setter
            func (this *PlayerData)GetPets() *ContainerPlayerDatapets {
                return this._pets
            }
// Daily getter and setter
            func (this *PlayerData)GetDaily() *Daily {
                return this._daily
            }
// Packs getter and setter
            func (this *PlayerData)GetPacks() *ContainerPlayerDatapacks {
                return this._packs
            }

// read from proto    
func (this *PlayerData)FromProto(pdata *netproto.PlayerData) {
        // BaseInfo getter and setter 
            if pdata.BaseInfo != nil {
                this.GetBaseInfo().FromProto(pdata.BaseInfo)
            }
        // Pets getter and setter   
                for _,val := range pdata.Pets {
                    ele := NewPet()
                    ele.FromProto(val.Val)
                    this._pets.Set(*val.Key, ele)
                }
        // Daily getter and setter 
            if pdata.Daily != nil {
                this.GetDaily().FromProto(pdata.Daily)
            }
        // Packs getter and setter   
                for _,val := range pdata.Packs {
                    ele := NewPack()
                    ele.FromProto(val.Val)
                    this._packs.Set(*val.Key, ele)
                }
}

// write to proto    
func (this *PlayerData)ToProto() *netproto.PlayerData {
    pdata := &netproto.PlayerData{}
        pdata.BaseInfo = this._baseInfo.ToProto()
        this._pets.Foreach(func (k int32, v *Pet) {
            pdata.Pets = append(pdata.Pets, NewPlayerDatapetsProto(k,v))
        })
        pdata.Daily = this._daily.ToProto()
        this._packs.Foreach(func (k int32, v *Pack) {
            pdata.Packs = append(pdata.Packs, NewPlayerDatapacksProto(k,v))
        })
    return pdata
}

// write dirty to proto
func (this *PlayerData)ToProtoDirty() *netproto.PlayerData {
    if (this.IsAllDirty()) {
        return this.ToProto()
    }
    pdata := &netproto.PlayerData{}
    if this.IsDirty(1) {
            pdata.BaseInfo = this._baseInfo.ToProtoDirty()
    }
    if this.IsDirty(2) {
            this._pets.ForeachDirty(func (k int32, v *Pet) {
                    pdata.Pets = append(pdata.Pets, NewPlayerDatapetsProtoDirty(k,v))
                })
    }
    if this.IsDirty(3) {
            pdata.Daily = this._daily.ToProtoDirty()
    }
    if this.IsDirty(4) {
            this._packs.ForeachDirty(func (k int32, v *Pack) {
                    pdata.Packs = append(pdata.Packs, NewPlayerDatapacksProtoDirty(k,v))
                })
    }
    return pdata
}
    type PlayerInfo struct {
        *datasync.DirtyContainerMarkVector
                _serverTime int64
                _serverName string
    }
    // ctor
    func NewPlayerInfo() *PlayerInfo {
        ins := &PlayerInfo {
            DirtyContainerMarkVector:datasync.NewDirtyContainerMarkVector(),
        }
        return ins
    }
// ServerTime getter and setter
            func (this *PlayerInfo)GetServerTime() int64 {
                return this._serverTime
            }
            func (this *PlayerInfo)SetServerTime(val int64) {
                this._serverTime = val
                this.MarkDirtyUp(1)
            }
// ServerName getter and setter
            func (this *PlayerInfo)GetServerName() string {
                return this._serverName
            }
            func (this *PlayerInfo)SetServerName(val string) {
                this._serverName = val
                this.MarkDirtyUp(2)
            }

// read from proto    
func (this *PlayerInfo)FromProto(pdata *netproto.PlayerInfo) {
        // ServerTime getter and setter
            if pdata.ServerTime != nil {
                this.SetServerTime(*pdata.ServerTime)
            }
        // ServerName getter and setter
            if pdata.ServerName != nil {
                this.SetServerName(*pdata.ServerName)
            }
}

// write to proto    
func (this *PlayerInfo)ToProto() *netproto.PlayerInfo {
    pdata := &netproto.PlayerInfo{}
        pdata.ServerTime = proto.Int64(this._serverTime)
        pdata.ServerName = proto.String(this._serverName)
    return pdata
}

// write dirty to proto
func (this *PlayerInfo)ToProtoDirty() *netproto.PlayerInfo {
    if (this.IsAllDirty()) {
        return this.ToProto()
    }
    pdata := &netproto.PlayerInfo{}
    if this.IsDirty(1) {
            pdata.ServerTime = proto.Int64(this._serverTime)
    }
    if this.IsDirty(2) {
            pdata.ServerName = proto.String(this._serverName)
    }
    return pdata
}
    type BaseInfo struct {
        *datasync.DirtyContainerMarkVector
                _id int64
                _name string
                _level int32
                _gold int32
                _exp int32
    }
    // ctor
    func NewBaseInfo() *BaseInfo {
        ins := &BaseInfo {
            DirtyContainerMarkVector:datasync.NewDirtyContainerMarkVector(),
        }
        return ins
    }
// Id getter and setter
            func (this *BaseInfo)GetId() int64 {
                return this._id
            }
            func (this *BaseInfo)SetId(val int64) {
                this._id = val
                this.MarkDirtyUp(1)
            }
// Name getter and setter
            func (this *BaseInfo)GetName() string {
                return this._name
            }
            func (this *BaseInfo)SetName(val string) {
                this._name = val
                this.MarkDirtyUp(2)
            }
// Level getter and setter
            func (this *BaseInfo)GetLevel() int32 {
                return this._level
            }
            func (this *BaseInfo)SetLevel(val int32) {
                this._level = val
                this.MarkDirtyUp(3)
            }
// Gold getter and setter
            func (this *BaseInfo)GetGold() int32 {
                return this._gold
            }
            func (this *BaseInfo)SetGold(val int32) {
                this._gold = val
                this.MarkDirtyUp(4)
            }
// Exp getter and setter
            func (this *BaseInfo)GetExp() int32 {
                return this._exp
            }
            func (this *BaseInfo)SetExp(val int32) {
                this._exp = val
                this.MarkDirtyUp(5)
            }

// read from proto    
func (this *BaseInfo)FromProto(pdata *netproto.BaseInfo) {
        // Id getter and setter
            if pdata.Id != nil {
                this.SetId(*pdata.Id)
            }
        // Name getter and setter
            if pdata.Name != nil {
                this.SetName(*pdata.Name)
            }
        // Level getter and setter
            if pdata.Level != nil {
                this.SetLevel(*pdata.Level)
            }
        // Gold getter and setter
            if pdata.Gold != nil {
                this.SetGold(*pdata.Gold)
            }
        // Exp getter and setter
            if pdata.Exp != nil {
                this.SetExp(*pdata.Exp)
            }
}

// write to proto    
func (this *BaseInfo)ToProto() *netproto.BaseInfo {
    pdata := &netproto.BaseInfo{}
        pdata.Id = proto.Int64(this._id)
        pdata.Name = proto.String(this._name)
        pdata.Level = proto.Int32(this._level)
        pdata.Gold = proto.Int32(this._gold)
        pdata.Exp = proto.Int32(this._exp)
    return pdata
}

// write dirty to proto
func (this *BaseInfo)ToProtoDirty() *netproto.BaseInfo {
    if (this.IsAllDirty()) {
        return this.ToProto()
    }
    pdata := &netproto.BaseInfo{}
    if this.IsDirty(1) {
            pdata.Id = proto.Int64(this._id)
    }
    if this.IsDirty(2) {
            pdata.Name = proto.String(this._name)
    }
    if this.IsDirty(3) {
            pdata.Level = proto.Int32(this._level)
    }
    if this.IsDirty(4) {
            pdata.Gold = proto.Int32(this._gold)
    }
    if this.IsDirty(5) {
            pdata.Exp = proto.Int32(this._exp)
    }
    return pdata
}
    type Pet struct {
        *datasync.DirtyContainerMarkVector
                _id int64
                _level int32
                _pack int32
    }
    // ctor
    func NewPet() *Pet {
        ins := &Pet {
            DirtyContainerMarkVector:datasync.NewDirtyContainerMarkVector(),
        }
        return ins
    }
// Id getter and setter
            func (this *Pet)GetId() int64 {
                return this._id
            }
            func (this *Pet)SetId(val int64) {
                this._id = val
                this.MarkDirtyUp(1)
            }
// Level getter and setter
            func (this *Pet)GetLevel() int32 {
                return this._level
            }
            func (this *Pet)SetLevel(val int32) {
                this._level = val
                this.MarkDirtyUp(2)
            }
// Pack getter and setter
            func (this *Pet)GetPack() int32 {
                return this._pack
            }
            func (this *Pet)SetPack(val int32) {
                this._pack = val
                this.MarkDirtyUp(3)
            }

// read from proto    
func (this *Pet)FromProto(pdata *netproto.Pet) {
        // Id getter and setter
            if pdata.Id != nil {
                this.SetId(*pdata.Id)
            }
        // Level getter and setter
            if pdata.Level != nil {
                this.SetLevel(*pdata.Level)
            }
        // Pack getter and setter
            if pdata.Pack != nil {
                this.SetPack(*pdata.Pack)
            }
}

// write to proto    
func (this *Pet)ToProto() *netproto.Pet {
    pdata := &netproto.Pet{}
        pdata.Id = proto.Int64(this._id)
        pdata.Level = proto.Int32(this._level)
        pdata.Pack = proto.Int32(this._pack)
    return pdata
}

// write dirty to proto
func (this *Pet)ToProtoDirty() *netproto.Pet {
    if (this.IsAllDirty()) {
        return this.ToProto()
    }
    pdata := &netproto.Pet{}
    if this.IsDirty(1) {
            pdata.Id = proto.Int64(this._id)
    }
    if this.IsDirty(2) {
            pdata.Level = proto.Int32(this._level)
    }
    if this.IsDirty(3) {
            pdata.Pack = proto.Int32(this._pack)
    }
    return pdata
}
    type Pack struct {
        *datasync.DirtyContainerMarkVector
                _id int64
                _num int64
    }
    // ctor
    func NewPack() *Pack {
        ins := &Pack {
            DirtyContainerMarkVector:datasync.NewDirtyContainerMarkVector(),
        }
        return ins
    }
// Id getter and setter
            func (this *Pack)GetId() int64 {
                return this._id
            }
            func (this *Pack)SetId(val int64) {
                this._id = val
                this.MarkDirtyUp(1)
            }
// Num getter and setter
            func (this *Pack)GetNum() int64 {
                return this._num
            }
            func (this *Pack)SetNum(val int64) {
                this._num = val
                this.MarkDirtyUp(2)
            }

// read from proto    
func (this *Pack)FromProto(pdata *netproto.Pack) {
        // Id getter and setter
            if pdata.Id != nil {
                this.SetId(*pdata.Id)
            }
        // Num getter and setter
            if pdata.Num != nil {
                this.SetNum(*pdata.Num)
            }
}

// write to proto    
func (this *Pack)ToProto() *netproto.Pack {
    pdata := &netproto.Pack{}
        pdata.Id = proto.Int64(this._id)
        pdata.Num = proto.Int64(this._num)
    return pdata
}

// write dirty to proto
func (this *Pack)ToProtoDirty() *netproto.Pack {
    if (this.IsAllDirty()) {
        return this.ToProto()
    }
    pdata := &netproto.Pack{}
    if this.IsDirty(1) {
            pdata.Id = proto.Int64(this._id)
    }
    if this.IsDirty(2) {
            pdata.Num = proto.Int64(this._num)
    }
    return pdata
}
    type Daily struct {
        *datasync.DirtyContainerMarkVector
                _lastUpdateTime int64
                _sign bool
    }
    // ctor
    func NewDaily() *Daily {
        ins := &Daily {
            DirtyContainerMarkVector:datasync.NewDirtyContainerMarkVector(),
        }
        return ins
    }
// LastUpdateTime getter and setter
            func (this *Daily)GetLastUpdateTime() int64 {
                return this._lastUpdateTime
            }
            func (this *Daily)SetLastUpdateTime(val int64) {
                this._lastUpdateTime = val
                this.MarkDirtyUp(1)
            }
// Sign getter and setter
            func (this *Daily)GetSign() bool {
                return this._sign
            }
            func (this *Daily)SetSign(val bool) {
                this._sign = val
                this.MarkDirtyUp(2)
            }

// read from proto    
func (this *Daily)FromProto(pdata *netproto.Daily) {
        // LastUpdateTime getter and setter
            if pdata.LastUpdateTime != nil {
                this.SetLastUpdateTime(*pdata.LastUpdateTime)
            }
        // Sign getter and setter
            if pdata.Sign != nil {
                this.SetSign(*pdata.Sign)
            }
}

// write to proto    
func (this *Daily)ToProto() *netproto.Daily {
    pdata := &netproto.Daily{}
        pdata.LastUpdateTime = proto.Int64(this._lastUpdateTime)
        pdata.Sign = proto.Bool(this._sign)
    return pdata
}

// write dirty to proto
func (this *Daily)ToProtoDirty() *netproto.Daily {
    if (this.IsAllDirty()) {
        return this.ToProto()
    }
    pdata := &netproto.Daily{}
    if this.IsDirty(1) {
            pdata.LastUpdateTime = proto.Int64(this._lastUpdateTime)
    }
    if this.IsDirty(2) {
            pdata.Sign = proto.Bool(this._sign)
    }
    return pdata
}
    type TestList struct {
        *datasync.DirtyContainerMarkVector
                _id int64
                _ListPrimitiveInt *ContainerTestListListPrimitiveInt
                _ListPrimitiveStr *ContainerTestListListPrimitiveStr
                _ListPrimitiveCom *ContainerTestListListPrimitiveCom
    }
    // ctor
    func NewTestList() *TestList {
        ins := &TestList {
            DirtyContainerMarkVector:datasync.NewDirtyContainerMarkVector(),
        }
                ins._ListPrimitiveInt = &ContainerTestListListPrimitiveInt {
                    _impl : datasync.NewDirtyContainerVector(),
                }
                ins._ListPrimitiveInt._impl.Init(ins, 1)
                ins._ListPrimitiveStr = &ContainerTestListListPrimitiveStr {
                    _impl : datasync.NewDirtyContainerVector(),
                }
                ins._ListPrimitiveStr._impl.Init(ins, 2)
                ins._ListPrimitiveCom = &ContainerTestListListPrimitiveCom {
                    _impl : datasync.NewDirtyContainerVector(),
                }
                ins._ListPrimitiveCom._impl.Init(ins, 3)
        return ins
    }
// Id getter and setter
            func (this *TestList)GetId() int64 {
                return this._id
            }
            func (this *TestList)SetId(val int64) {
                this._id = val
                this.MarkDirtyUp(7)
            }
// ListPrimitiveInt getter and setter
            func (this *TestList)GetListPrimitiveInt() *ContainerTestListListPrimitiveInt {
                return this._ListPrimitiveInt
            }
// ListPrimitiveStr getter and setter
            func (this *TestList)GetListPrimitiveStr() *ContainerTestListListPrimitiveStr {
                return this._ListPrimitiveStr
            }
// ListPrimitiveCom getter and setter
            func (this *TestList)GetListPrimitiveCom() *ContainerTestListListPrimitiveCom {
                return this._ListPrimitiveCom
            }

// read from proto    
func (this *TestList)FromProto(pdata *netproto.TestList) {
        // Id getter and setter
            if pdata.Id != nil {
                this.SetId(*pdata.Id)
            }
        // ListPrimitiveInt getter and setter
                for _,val := range pdata.ListPrimitiveInt {
                    this._ListPrimitiveInt.Set(*val.Key, *val.Val)
                }
        // ListPrimitiveStr getter and setter
                for _,val := range pdata.ListPrimitiveStr {
                    this._ListPrimitiveStr.Set(*val.Key, *val.Val)
                }
        // ListPrimitiveCom getter and setter   
                for _,val := range pdata.ListPrimitiveCom {
                    ele := NewPlayerData()
                    ele.FromProto(val.Val)
                    this._ListPrimitiveCom.Set(*val.Key, ele)
                }
}

// write to proto    
func (this *TestList)ToProto() *netproto.TestList {
    pdata := &netproto.TestList{}
        pdata.Id = proto.Int64(this._id)
        this._ListPrimitiveInt.Foreach(func (k int32, v int32) {
            pdata.ListPrimitiveInt = append(pdata.ListPrimitiveInt, NewTestListListPrimitiveIntProto(k,v))
        })
        this._ListPrimitiveStr.Foreach(func (k int32, v string) {
            pdata.ListPrimitiveStr = append(pdata.ListPrimitiveStr, NewTestListListPrimitiveStrProto(k,v))
        })
        this._ListPrimitiveCom.Foreach(func (k int32, v *PlayerData) {
            pdata.ListPrimitiveCom = append(pdata.ListPrimitiveCom, NewTestListListPrimitiveComProto(k,v))
        })
    return pdata
}

// write dirty to proto
func (this *TestList)ToProtoDirty() *netproto.TestList {
    if (this.IsAllDirty()) {
        return this.ToProto()
    }
    pdata := &netproto.TestList{}
    if this.IsDirty(7) {
            pdata.Id = proto.Int64(this._id)
    }
    if this.IsDirty(1) {
            this._ListPrimitiveInt.ForeachDirty(func (k int32, v int32) {
                    pdata.ListPrimitiveInt = append(pdata.ListPrimitiveInt, NewTestListListPrimitiveIntProtoDirty(k,v))
                })
    }
    if this.IsDirty(2) {
            this._ListPrimitiveStr.ForeachDirty(func (k int32, v string) {
                    pdata.ListPrimitiveStr = append(pdata.ListPrimitiveStr, NewTestListListPrimitiveStrProtoDirty(k,v))
                })
    }
    if this.IsDirty(3) {
            this._ListPrimitiveCom.ForeachDirty(func (k int32, v *PlayerData) {
                    pdata.ListPrimitiveCom = append(pdata.ListPrimitiveCom, NewTestListListPrimitiveComProtoDirty(k,v))
                })
    }
    return pdata
}
    type TestMap struct {
        *datasync.DirtyContainerMarkVector
                _id *BaseInfo
                _MapPrimitiveIntInt *ContainerTestMapMapPrimitiveIntInt
                _MapPrimitiveIntStr *ContainerTestMapMapPrimitiveIntStr
                _MapPrimitiveStrInt *ContainerTestMapMapPrimitiveStrInt
                _MapPrimitiveStrStr *ContainerTestMapMapPrimitiveStrStr
                _MapPrimitiveIntCom *ContainerTestMapMapPrimitiveIntCom
                _MapPrimitiveStrCom *ContainerTestMapMapPrimitiveStrCom
    }
    // ctor
    func NewTestMap() *TestMap {
        ins := &TestMap {
            DirtyContainerMarkVector:datasync.NewDirtyContainerMarkVector(),
        }
                ins._id = NewBaseInfo()
                ins._id.Init(ins, 7)
                ins._MapPrimitiveIntInt = &ContainerTestMapMapPrimitiveIntInt {
                    _impl : datasync.NewDirtyContainerMap(),
                }
                ins._MapPrimitiveIntInt._impl.Init(ins, 1)
                ins._MapPrimitiveIntStr = &ContainerTestMapMapPrimitiveIntStr {
                    _impl : datasync.NewDirtyContainerMap(),
                }
                ins._MapPrimitiveIntStr._impl.Init(ins, 2)
                ins._MapPrimitiveStrInt = &ContainerTestMapMapPrimitiveStrInt {
                    _impl : datasync.NewDirtyContainerMap(),
                }
                ins._MapPrimitiveStrInt._impl.Init(ins, 3)
                ins._MapPrimitiveStrStr = &ContainerTestMapMapPrimitiveStrStr {
                    _impl : datasync.NewDirtyContainerMap(),
                }
                ins._MapPrimitiveStrStr._impl.Init(ins, 4)
                ins._MapPrimitiveIntCom = &ContainerTestMapMapPrimitiveIntCom {
                    _impl : datasync.NewDirtyContainerMap(),
                }
                ins._MapPrimitiveIntCom._impl.Init(ins, 5)
                ins._MapPrimitiveStrCom = &ContainerTestMapMapPrimitiveStrCom {
                    _impl : datasync.NewDirtyContainerMap(),
                }
                ins._MapPrimitiveStrCom._impl.Init(ins, 6)
        return ins
    }
// Id getter and setter
            func (this *TestMap)GetId() *BaseInfo {
                return this._id
            }
// MapPrimitiveIntInt getter and setter
            func (this *TestMap)GetMapPrimitiveIntInt() *ContainerTestMapMapPrimitiveIntInt {
                return this._MapPrimitiveIntInt
            }
// MapPrimitiveIntStr getter and setter
            func (this *TestMap)GetMapPrimitiveIntStr() *ContainerTestMapMapPrimitiveIntStr {
                return this._MapPrimitiveIntStr
            }
// MapPrimitiveStrInt getter and setter
            func (this *TestMap)GetMapPrimitiveStrInt() *ContainerTestMapMapPrimitiveStrInt {
                return this._MapPrimitiveStrInt
            }
// MapPrimitiveStrStr getter and setter
            func (this *TestMap)GetMapPrimitiveStrStr() *ContainerTestMapMapPrimitiveStrStr {
                return this._MapPrimitiveStrStr
            }
// MapPrimitiveIntCom getter and setter
            func (this *TestMap)GetMapPrimitiveIntCom() *ContainerTestMapMapPrimitiveIntCom {
                return this._MapPrimitiveIntCom
            }
// MapPrimitiveStrCom getter and setter
            func (this *TestMap)GetMapPrimitiveStrCom() *ContainerTestMapMapPrimitiveStrCom {
                return this._MapPrimitiveStrCom
            }

// read from proto    
func (this *TestMap)FromProto(pdata *netproto.TestMap) {
        // Id getter and setter 
            if pdata.Id != nil {
                this.GetId().FromProto(pdata.Id)
            }
        // MapPrimitiveIntInt getter and setter
                for _,val := range pdata.MapPrimitiveIntInt {
                        this._MapPrimitiveIntInt.Set(*val.Key, *val.Val)
                }
        // MapPrimitiveIntStr getter and setter
                for _,val := range pdata.MapPrimitiveIntStr {
                        this._MapPrimitiveIntStr.Set(*val.Key, *val.Val)
                }
        // MapPrimitiveStrInt getter and setter
                for _,val := range pdata.MapPrimitiveStrInt {
                        this._MapPrimitiveStrInt.Set(*val.Key, *val.Val)
                }
        // MapPrimitiveStrStr getter and setter
                for _,val := range pdata.MapPrimitiveStrStr {
                        this._MapPrimitiveStrStr.Set(*val.Key, *val.Val)
                }
        // MapPrimitiveIntCom getter and setter
                for _,val := range pdata.MapPrimitiveIntCom {
                        ele := NewBaseInfo()
                        ele.FromProto(val.Val)
                        this._MapPrimitiveIntCom.Set(*val.Key, ele)
                }
        // MapPrimitiveStrCom getter and setter
                for _,val := range pdata.MapPrimitiveStrCom {
                        ele := NewPlayerData()
                        ele.FromProto(val.Val)
                        this._MapPrimitiveStrCom.Set(*val.Key, ele)
                }
}

// write to proto    
func (this *TestMap)ToProto() *netproto.TestMap {
    pdata := &netproto.TestMap{}
        pdata.Id = this._id.ToProto()
        this._MapPrimitiveIntInt.Foreach(func (k int32, v int32) {
            pdata.MapPrimitiveIntInt = append(pdata.MapPrimitiveIntInt, NewTestMapMapPrimitiveIntIntProto(k,v))
            })
        this._MapPrimitiveIntStr.Foreach(func (k int32, v string) {
            pdata.MapPrimitiveIntStr = append(pdata.MapPrimitiveIntStr, NewTestMapMapPrimitiveIntStrProto(k,v))
            })
        this._MapPrimitiveStrInt.Foreach(func (k string, v int32) {
            pdata.MapPrimitiveStrInt = append(pdata.MapPrimitiveStrInt, NewTestMapMapPrimitiveStrIntProto(k,v))
            })
        this._MapPrimitiveStrStr.Foreach(func (k string, v string) {
            pdata.MapPrimitiveStrStr = append(pdata.MapPrimitiveStrStr, NewTestMapMapPrimitiveStrStrProto(k,v))
            })
        this._MapPrimitiveIntCom.Foreach(func (k int32, v *BaseInfo) {
            pdata.MapPrimitiveIntCom = append(pdata.MapPrimitiveIntCom, NewTestMapMapPrimitiveIntComProto(k,v))
            })
        this._MapPrimitiveStrCom.Foreach(func (k string, v *PlayerData) {
            pdata.MapPrimitiveStrCom = append(pdata.MapPrimitiveStrCom, NewTestMapMapPrimitiveStrComProto(k,v))
            })
    return pdata
}

// write dirty to proto
func (this *TestMap)ToProtoDirty() *netproto.TestMap {
    if (this.IsAllDirty()) {
        return this.ToProto()
    }
    pdata := &netproto.TestMap{}
    if this.IsDirty(7) {
            pdata.Id = this._id.ToProtoDirty()
    }
    if this.IsDirty(1) {
            this._MapPrimitiveIntInt.ForeachDirty(func (k int32, v int32) {
                    pdata.MapPrimitiveIntInt = append(pdata.MapPrimitiveIntInt, NewTestMapMapPrimitiveIntIntProtoDirty(k,v))
                    })
    }
    if this.IsDirty(2) {
            this._MapPrimitiveIntStr.ForeachDirty(func (k int32, v string) {
                    pdata.MapPrimitiveIntStr = append(pdata.MapPrimitiveIntStr, NewTestMapMapPrimitiveIntStrProtoDirty(k,v))
                    })
    }
    if this.IsDirty(3) {
            this._MapPrimitiveStrInt.ForeachDirty(func (k string, v int32) {
                    pdata.MapPrimitiveStrInt = append(pdata.MapPrimitiveStrInt, NewTestMapMapPrimitiveStrIntProtoDirty(k,v))
                    })
    }
    if this.IsDirty(4) {
            this._MapPrimitiveStrStr.ForeachDirty(func (k string, v string) {
                    pdata.MapPrimitiveStrStr = append(pdata.MapPrimitiveStrStr, NewTestMapMapPrimitiveStrStrProtoDirty(k,v))
                    })
    }
    if this.IsDirty(5) {
            this._MapPrimitiveIntCom.ForeachDirty(func (k int32, v *BaseInfo) {
                    pdata.MapPrimitiveIntCom = append(pdata.MapPrimitiveIntCom, NewTestMapMapPrimitiveIntComProtoDirty(k,v))
                    })
    }
    if this.IsDirty(6) {
            this._MapPrimitiveStrCom.ForeachDirty(func (k string, v *PlayerData) {
                    pdata.MapPrimitiveStrCom = append(pdata.MapPrimitiveStrCom, NewTestMapMapPrimitiveStrComProtoDirty(k,v))
                    })
    }
    return pdata
}
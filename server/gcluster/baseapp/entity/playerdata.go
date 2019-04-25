package entity

import (
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gbase/datasync"
	"github.com/gosrv/gcluster/gcluster/proto"
)

type ItemIterChapterheros func(k int32, v int32)
type ItemStatusIterChapterheros func(k int32, v int32, dirty bool)

type ContainerChapterheros struct {
	_impl *datasync.DirtyContainerVector
}

func (this *ContainerChapterheros) Get(key int32) int32 {
	return this._impl.Get(key).(int32)
}
func (this *ContainerChapterheros) Set(key int32, value int32) {
	this._impl.Set(key, value)
}
func (this *ContainerChapterheros) Foreach(iter ItemIterChapterheros) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(int32)) })
}
func (this *ContainerChapterheros) ForeachStatus(iter ItemStatusIterChapterheros) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(int32), dirty) })
}
func (this *ContainerChapterheros) ForeachDirty(iter ItemIterChapterheros) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(int32)) })
}
func (this *ContainerChapterheros) Size() int {
	return this._impl.Size()
}
func (this *ContainerChapterheros) Clear() {
	this._impl.Clear()
}

func NewChapterherosProto(k int32, v int32) *netproto.Chapterheros {
	return &netproto.Chapterheros{
		Key: &k,
		Val: &v,
	}
}

func NewChapterherosProtoDirty(k int32, v int32) *netproto.Chapterheros {
	return &netproto.Chapterheros{
		Key: &k,
		Val: &v,
	}
}

type ItemIterChapterprizeEquip func(k int32, v *Equip)
type ItemStatusIterChapterprizeEquip func(k int32, v *Equip, dirty bool)

type ContainerChapterprizeEquip struct {
	_impl *datasync.DirtyContainerVector
}

func (this *ContainerChapterprizeEquip) Get(key int32) *Equip {
	val := this._impl.Get(key)
	if val == nil {
		return nil
	} else {
		return val.(*Equip)
	}
}
func (this *ContainerChapterprizeEquip) Set(key int32, value *Equip) {
	this._impl.Set(key, value)
}
func (this *ContainerChapterprizeEquip) Foreach(iter ItemIterChapterprizeEquip) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(*Equip)) })
}
func (this *ContainerChapterprizeEquip) ForeachStatus(iter ItemStatusIterChapterprizeEquip) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(*Equip), dirty) })
}
func (this *ContainerChapterprizeEquip) ForeachDirty(iter ItemIterChapterprizeEquip) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(*Equip)) })
}
func (this *ContainerChapterprizeEquip) Size() int {
	return this._impl.Size()
}
func (this *ContainerChapterprizeEquip) Clear() {
	this._impl.Clear()
}

func NewChapterprizeEquipProto(k int32, v *Equip) *netproto.ChapterprizeEquip {
	return &netproto.ChapterprizeEquip{
		Key: &k,
		Val: v.ToProto(),
	}
}

func NewChapterprizeEquipProtoDirty(k int32, v *Equip) *netproto.ChapterprizeEquip {
	return &netproto.ChapterprizeEquip{
		Key: &k,
		Val: v.ToProtoDirty(),
	}
}

type ItemIterHeroequips func(k int32, v *Equip)
type ItemStatusIterHeroequips func(k int32, v *Equip, dirty bool)

type ContainerHeroequips struct {
	_impl *datasync.DirtyContainerVector
}

func (this *ContainerHeroequips) Get(key int32) *Equip {
	val := this._impl.Get(key)
	if val == nil {
		return nil
	} else {
		return val.(*Equip)
	}
}
func (this *ContainerHeroequips) Set(key int32, value *Equip) {
	this._impl.Set(key, value)
}
func (this *ContainerHeroequips) Foreach(iter ItemIterHeroequips) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(*Equip)) })
}
func (this *ContainerHeroequips) ForeachStatus(iter ItemStatusIterHeroequips) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(*Equip), dirty) })
}
func (this *ContainerHeroequips) ForeachDirty(iter ItemIterHeroequips) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(*Equip)) })
}
func (this *ContainerHeroequips) Size() int {
	return this._impl.Size()
}
func (this *ContainerHeroequips) Clear() {
	this._impl.Clear()
}

func NewHeroequipsProto(k int32, v *Equip) *netproto.Heroequips {
	return &netproto.Heroequips{
		Key: &k,
		Val: v.ToProto(),
	}
}

func NewHeroequipsProtoDirty(k int32, v *Equip) *netproto.Heroequips {
	return &netproto.Heroequips{
		Key: &k,
		Val: v.ToProtoDirty(),
	}
}

type ItemIterHeroPackheros func(k int32, v *Hero)
type ItemStatusIterHeroPackheros func(k int32, v *Hero, dirty bool)

type ContainerHeroPackheros struct {
	_impl *datasync.DirtyContainerMap
}

func (this *ContainerHeroPackheros) Get(key int32) *Hero {
	val := this._impl.Get(key)
	if val == nil {
		return nil
	} else {
		return val.(*Hero)
	}
}
func (this *ContainerHeroPackheros) Set(key int32, value *Hero) {
	this._impl.Set(key, value)
}
func (this *ContainerHeroPackheros) Foreach(iter ItemIterHeroPackheros) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(*Hero)) })
}
func (this *ContainerHeroPackheros) ForeachStatus(iter ItemStatusIterHeroPackheros) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(*Hero), dirty) })
}
func (this *ContainerHeroPackheros) ForeachDirty(iter ItemIterHeroPackheros) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(*Hero)) })
}
func (this *ContainerHeroPackheros) Size() int {
	return this._impl.Size()
}
func (this *ContainerHeroPackheros) Clear() {
	this._impl.Clear()
}

func NewHeroPackherosProto(k int32, v *Hero) *netproto.HeroPackheros {
	return &netproto.HeroPackheros{
		Key: &k,
		Val: v.ToProto(),
	}
}

func NewHeroPackherosProtoDirty(k int32, v *Hero) *netproto.HeroPackheros {
	return &netproto.HeroPackheros{
		Key: &k,
		Val: v.ToProtoDirty(),
	}
}

type ItemIterEquipPackequips func(k int32, v *Equip)
type ItemStatusIterEquipPackequips func(k int32, v *Equip, dirty bool)

type ContainerEquipPackequips struct {
	_impl *datasync.DirtyContainerMap
}

func (this *ContainerEquipPackequips) Get(key int32) *Equip {
	val := this._impl.Get(key)
	if val == nil {
		return nil
	} else {
		return val.(*Equip)
	}
}
func (this *ContainerEquipPackequips) Set(key int32, value *Equip) {
	this._impl.Set(key, value)
}
func (this *ContainerEquipPackequips) Foreach(iter ItemIterEquipPackequips) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(*Equip)) })
}
func (this *ContainerEquipPackequips) ForeachStatus(iter ItemStatusIterEquipPackequips) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(*Equip), dirty) })
}
func (this *ContainerEquipPackequips) ForeachDirty(iter ItemIterEquipPackequips) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(*Equip)) })
}
func (this *ContainerEquipPackequips) Size() int {
	return this._impl.Size()
}
func (this *ContainerEquipPackequips) Clear() {
	this._impl.Clear()
}

func NewEquipPackequipsProto(k int32, v *Equip) *netproto.EquipPackequips {
	return &netproto.EquipPackequips{
		Key: &k,
		Val: v.ToProto(),
	}
}

func NewEquipPackequipsProtoDirty(k int32, v *Equip) *netproto.EquipPackequips {
	return &netproto.EquipPackequips{
		Key: &k,
		Val: v.ToProtoDirty(),
	}
}

type ItemIterTestListListPrimitiveInt func(k int32, v int32)
type ItemStatusIterTestListListPrimitiveInt func(k int32, v int32, dirty bool)

type ContainerTestListListPrimitiveInt struct {
	_impl *datasync.DirtyContainerVector
}

func (this *ContainerTestListListPrimitiveInt) Get(key int32) int32 {
	return this._impl.Get(key).(int32)
}
func (this *ContainerTestListListPrimitiveInt) Set(key int32, value int32) {
	this._impl.Set(key, value)
}
func (this *ContainerTestListListPrimitiveInt) Foreach(iter ItemIterTestListListPrimitiveInt) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(int32)) })
}
func (this *ContainerTestListListPrimitiveInt) ForeachStatus(iter ItemStatusIterTestListListPrimitiveInt) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(int32), dirty) })
}
func (this *ContainerTestListListPrimitiveInt) ForeachDirty(iter ItemIterTestListListPrimitiveInt) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(int32)) })
}
func (this *ContainerTestListListPrimitiveInt) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestListListPrimitiveInt) Clear() {
	this._impl.Clear()
}

func NewTestListListPrimitiveIntProto(k int32, v int32) *netproto.TestListListPrimitiveInt {
	return &netproto.TestListListPrimitiveInt{
		Key: &k,
		Val: &v,
	}
}

func NewTestListListPrimitiveIntProtoDirty(k int32, v int32) *netproto.TestListListPrimitiveInt {
	return &netproto.TestListListPrimitiveInt{
		Key: &k,
		Val: &v,
	}
}

type ItemIterTestListListPrimitiveStr func(k int32, v string)
type ItemStatusIterTestListListPrimitiveStr func(k int32, v string, dirty bool)

type ContainerTestListListPrimitiveStr struct {
	_impl *datasync.DirtyContainerVector
}

func (this *ContainerTestListListPrimitiveStr) Get(key int32) string {
	return this._impl.Get(key).(string)
}
func (this *ContainerTestListListPrimitiveStr) Set(key int32, value string) {
	this._impl.Set(key, value)
}
func (this *ContainerTestListListPrimitiveStr) Foreach(iter ItemIterTestListListPrimitiveStr) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(string)) })
}
func (this *ContainerTestListListPrimitiveStr) ForeachStatus(iter ItemStatusIterTestListListPrimitiveStr) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(string), dirty) })
}
func (this *ContainerTestListListPrimitiveStr) ForeachDirty(iter ItemIterTestListListPrimitiveStr) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(string)) })
}
func (this *ContainerTestListListPrimitiveStr) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestListListPrimitiveStr) Clear() {
	this._impl.Clear()
}

func NewTestListListPrimitiveStrProto(k int32, v string) *netproto.TestListListPrimitiveStr {
	return &netproto.TestListListPrimitiveStr{
		Key: &k,
		Val: &v,
	}
}

func NewTestListListPrimitiveStrProtoDirty(k int32, v string) *netproto.TestListListPrimitiveStr {
	return &netproto.TestListListPrimitiveStr{
		Key: &k,
		Val: &v,
	}
}

type ItemIterTestListListPrimitiveCom func(k int32, v *PlayerData)
type ItemStatusIterTestListListPrimitiveCom func(k int32, v *PlayerData, dirty bool)

type ContainerTestListListPrimitiveCom struct {
	_impl *datasync.DirtyContainerVector
}

func (this *ContainerTestListListPrimitiveCom) Get(key int32) *PlayerData {
	val := this._impl.Get(key)
	if val == nil {
		return nil
	} else {
		return val.(*PlayerData)
	}
}
func (this *ContainerTestListListPrimitiveCom) Set(key int32, value *PlayerData) {
	this._impl.Set(key, value)
}
func (this *ContainerTestListListPrimitiveCom) Foreach(iter ItemIterTestListListPrimitiveCom) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(*PlayerData)) })
}
func (this *ContainerTestListListPrimitiveCom) ForeachStatus(iter ItemStatusIterTestListListPrimitiveCom) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(*PlayerData), dirty) })
}
func (this *ContainerTestListListPrimitiveCom) ForeachDirty(iter ItemIterTestListListPrimitiveCom) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(*PlayerData)) })
}
func (this *ContainerTestListListPrimitiveCom) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestListListPrimitiveCom) Clear() {
	this._impl.Clear()
}

func NewTestListListPrimitiveComProto(k int32, v *PlayerData) *netproto.TestListListPrimitiveCom {
	return &netproto.TestListListPrimitiveCom{
		Key: &k,
		Val: v.ToProto(),
	}
}

func NewTestListListPrimitiveComProtoDirty(k int32, v *PlayerData) *netproto.TestListListPrimitiveCom {
	return &netproto.TestListListPrimitiveCom{
		Key: &k,
		Val: v.ToProtoDirty(),
	}
}

type ItemIterTestMapMapPrimitiveIntInt func(k int32, v int32)
type ItemStatusIterTestMapMapPrimitiveIntInt func(k int32, v int32, dirty bool)

type ContainerTestMapMapPrimitiveIntInt struct {
	_impl *datasync.DirtyContainerMap
}

func (this *ContainerTestMapMapPrimitiveIntInt) Get(key int32) int32 {
	return this._impl.Get(key).(int32)
}
func (this *ContainerTestMapMapPrimitiveIntInt) Set(key int32, value int32) {
	this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveIntInt) Foreach(iter ItemIterTestMapMapPrimitiveIntInt) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(int32)) })
}
func (this *ContainerTestMapMapPrimitiveIntInt) ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveIntInt) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(int32), dirty) })
}
func (this *ContainerTestMapMapPrimitiveIntInt) ForeachDirty(iter ItemIterTestMapMapPrimitiveIntInt) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(int32)) })
}
func (this *ContainerTestMapMapPrimitiveIntInt) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveIntInt) Clear() {
	this._impl.Clear()
}

func NewTestMapMapPrimitiveIntIntProto(k int32, v int32) *netproto.TestMapMapPrimitiveIntInt {
	return &netproto.TestMapMapPrimitiveIntInt{
		Key: &k,
		Val: &v,
	}
}

func NewTestMapMapPrimitiveIntIntProtoDirty(k int32, v int32) *netproto.TestMapMapPrimitiveIntInt {
	return &netproto.TestMapMapPrimitiveIntInt{
		Key: &k,
		Val: &v,
	}
}

type ItemIterTestMapMapPrimitiveIntStr func(k int32, v string)
type ItemStatusIterTestMapMapPrimitiveIntStr func(k int32, v string, dirty bool)

type ContainerTestMapMapPrimitiveIntStr struct {
	_impl *datasync.DirtyContainerMap
}

func (this *ContainerTestMapMapPrimitiveIntStr) Get(key int32) string {
	return this._impl.Get(key).(string)
}
func (this *ContainerTestMapMapPrimitiveIntStr) Set(key int32, value string) {
	this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveIntStr) Foreach(iter ItemIterTestMapMapPrimitiveIntStr) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(string)) })
}
func (this *ContainerTestMapMapPrimitiveIntStr) ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveIntStr) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(string), dirty) })
}
func (this *ContainerTestMapMapPrimitiveIntStr) ForeachDirty(iter ItemIterTestMapMapPrimitiveIntStr) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(string)) })
}
func (this *ContainerTestMapMapPrimitiveIntStr) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveIntStr) Clear() {
	this._impl.Clear()
}

func NewTestMapMapPrimitiveIntStrProto(k int32, v string) *netproto.TestMapMapPrimitiveIntStr {
	return &netproto.TestMapMapPrimitiveIntStr{
		Key: &k,
		Val: &v,
	}
}

func NewTestMapMapPrimitiveIntStrProtoDirty(k int32, v string) *netproto.TestMapMapPrimitiveIntStr {
	return &netproto.TestMapMapPrimitiveIntStr{
		Key: &k,
		Val: &v,
	}
}

type ItemIterTestMapMapPrimitiveStrInt func(k string, v int32)
type ItemStatusIterTestMapMapPrimitiveStrInt func(k string, v int32, dirty bool)

type ContainerTestMapMapPrimitiveStrInt struct {
	_impl *datasync.DirtyContainerMap
}

func (this *ContainerTestMapMapPrimitiveStrInt) Get(key string) int32 {
	return this._impl.Get(key).(int32)
}
func (this *ContainerTestMapMapPrimitiveStrInt) Set(key string, value int32) {
	this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveStrInt) Foreach(iter ItemIterTestMapMapPrimitiveStrInt) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(string), v.(int32)) })
}
func (this *ContainerTestMapMapPrimitiveStrInt) ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveStrInt) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(string), v.(int32), dirty) })
}
func (this *ContainerTestMapMapPrimitiveStrInt) ForeachDirty(iter ItemIterTestMapMapPrimitiveStrInt) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(string), v.(int32)) })
}
func (this *ContainerTestMapMapPrimitiveStrInt) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveStrInt) Clear() {
	this._impl.Clear()
}

func NewTestMapMapPrimitiveStrIntProto(k string, v int32) *netproto.TestMapMapPrimitiveStrInt {
	return &netproto.TestMapMapPrimitiveStrInt{
		Key: &k,
		Val: &v,
	}
}

func NewTestMapMapPrimitiveStrIntProtoDirty(k string, v int32) *netproto.TestMapMapPrimitiveStrInt {
	return &netproto.TestMapMapPrimitiveStrInt{
		Key: &k,
		Val: &v,
	}
}

type ItemIterTestMapMapPrimitiveStrStr func(k string, v string)
type ItemStatusIterTestMapMapPrimitiveStrStr func(k string, v string, dirty bool)

type ContainerTestMapMapPrimitiveStrStr struct {
	_impl *datasync.DirtyContainerMap
}

func (this *ContainerTestMapMapPrimitiveStrStr) Get(key string) string {
	return this._impl.Get(key).(string)
}
func (this *ContainerTestMapMapPrimitiveStrStr) Set(key string, value string) {
	this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveStrStr) Foreach(iter ItemIterTestMapMapPrimitiveStrStr) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(string), v.(string)) })
}
func (this *ContainerTestMapMapPrimitiveStrStr) ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveStrStr) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(string), v.(string), dirty) })
}
func (this *ContainerTestMapMapPrimitiveStrStr) ForeachDirty(iter ItemIterTestMapMapPrimitiveStrStr) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(string), v.(string)) })
}
func (this *ContainerTestMapMapPrimitiveStrStr) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveStrStr) Clear() {
	this._impl.Clear()
}

func NewTestMapMapPrimitiveStrStrProto(k string, v string) *netproto.TestMapMapPrimitiveStrStr {
	return &netproto.TestMapMapPrimitiveStrStr{
		Key: &k,
		Val: &v,
	}
}

func NewTestMapMapPrimitiveStrStrProtoDirty(k string, v string) *netproto.TestMapMapPrimitiveStrStr {
	return &netproto.TestMapMapPrimitiveStrStr{
		Key: &k,
		Val: &v,
	}
}

type ItemIterTestMapMapPrimitiveIntCom func(k int32, v *BaseInfo)
type ItemStatusIterTestMapMapPrimitiveIntCom func(k int32, v *BaseInfo, dirty bool)

type ContainerTestMapMapPrimitiveIntCom struct {
	_impl *datasync.DirtyContainerMap
}

func (this *ContainerTestMapMapPrimitiveIntCom) Get(key int32) *BaseInfo {
	val := this._impl.Get(key)
	if val == nil {
		return nil
	} else {
		return val.(*BaseInfo)
	}
}
func (this *ContainerTestMapMapPrimitiveIntCom) Set(key int32, value *BaseInfo) {
	this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveIntCom) Foreach(iter ItemIterTestMapMapPrimitiveIntCom) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(int32), v.(*BaseInfo)) })
}
func (this *ContainerTestMapMapPrimitiveIntCom) ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveIntCom) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(int32), v.(*BaseInfo), dirty) })
}
func (this *ContainerTestMapMapPrimitiveIntCom) ForeachDirty(iter ItemIterTestMapMapPrimitiveIntCom) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(int32), v.(*BaseInfo)) })
}
func (this *ContainerTestMapMapPrimitiveIntCom) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveIntCom) Clear() {
	this._impl.Clear()
}

func NewTestMapMapPrimitiveIntComProto(k int32, v *BaseInfo) *netproto.TestMapMapPrimitiveIntCom {
	return &netproto.TestMapMapPrimitiveIntCom{
		Key: &k,
		Val: v.ToProto(),
	}
}

func NewTestMapMapPrimitiveIntComProtoDirty(k int32, v *BaseInfo) *netproto.TestMapMapPrimitiveIntCom {
	return &netproto.TestMapMapPrimitiveIntCom{
		Key: &k,
		Val: v.ToProtoDirty(),
	}
}

type ItemIterTestMapMapPrimitiveStrCom func(k string, v *PlayerData)
type ItemStatusIterTestMapMapPrimitiveStrCom func(k string, v *PlayerData, dirty bool)

type ContainerTestMapMapPrimitiveStrCom struct {
	_impl *datasync.DirtyContainerMap
}

func (this *ContainerTestMapMapPrimitiveStrCom) Get(key string) *PlayerData {
	val := this._impl.Get(key)
	if val == nil {
		return nil
	} else {
		return val.(*PlayerData)
	}
}
func (this *ContainerTestMapMapPrimitiveStrCom) Set(key string, value *PlayerData) {
	this._impl.Set(key, value)
}
func (this *ContainerTestMapMapPrimitiveStrCom) Foreach(iter ItemIterTestMapMapPrimitiveStrCom) {
	this._impl.Foreach(func(k interface{}, v interface{}) { iter(k.(string), v.(*PlayerData)) })
}
func (this *ContainerTestMapMapPrimitiveStrCom) ForeachStatus(iter ItemStatusIterTestMapMapPrimitiveStrCom) {
	this._impl.ForeachStatus(func(k interface{}, v interface{}, dirty bool) { iter(k.(string), v.(*PlayerData), dirty) })
}
func (this *ContainerTestMapMapPrimitiveStrCom) ForeachDirty(iter ItemIterTestMapMapPrimitiveStrCom) {
	this._impl.ForeachDirty(func(k interface{}, v interface{}) { iter(k.(string), v.(*PlayerData)) })
}
func (this *ContainerTestMapMapPrimitiveStrCom) Size() int {
	return this._impl.Size()
}
func (this *ContainerTestMapMapPrimitiveStrCom) Clear() {
	this._impl.Clear()
}

func NewTestMapMapPrimitiveStrComProto(k string, v *PlayerData) *netproto.TestMapMapPrimitiveStrCom {
	return &netproto.TestMapMapPrimitiveStrCom{
		Key: &k,
		Val: v.ToProto(),
	}
}

func NewTestMapMapPrimitiveStrComProtoDirty(k string, v *PlayerData) *netproto.TestMapMapPrimitiveStrCom {
	return &netproto.TestMapMapPrimitiveStrCom{
		Key: &k,
		Val: v.ToProtoDirty(),
	}
}

type PlayerData struct {
	*datasync.DirtyContainerMarkVector
	_baseInfo  *BaseInfo
	_vipInfo   *VipInfo
	_heroPack  *HeroPack
	_equipPack *EquipPack
	_guide     *Guide
	_chapter   *Chapter
}

// ctor
func NewPlayerData() *PlayerData {
	ins := &PlayerData{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	ins._baseInfo = NewBaseInfo()
	ins._baseInfo.Init(ins, 1)
	ins._vipInfo = NewVipInfo()
	ins._vipInfo.Init(ins, 2)
	ins._heroPack = NewHeroPack()
	ins._heroPack.Init(ins, 3)
	ins._equipPack = NewEquipPack()
	ins._equipPack.Init(ins, 4)
	ins._guide = NewGuide()
	ins._guide.Init(ins, 5)
	ins._chapter = NewChapter()
	ins._chapter.Init(ins, 6)
	return ins
}

// BaseInfo getter and setter
func (this *PlayerData) GetBaseInfo() *BaseInfo {
	return this._baseInfo
}

// VipInfo getter and setter
func (this *PlayerData) GetVipInfo() *VipInfo {
	return this._vipInfo
}

// HeroPack getter and setter
func (this *PlayerData) GetHeroPack() *HeroPack {
	return this._heroPack
}

// EquipPack getter and setter
func (this *PlayerData) GetEquipPack() *EquipPack {
	return this._equipPack
}

// Guide getter and setter
func (this *PlayerData) GetGuide() *Guide {
	return this._guide
}

// Chapter getter and setter
func (this *PlayerData) GetChapter() *Chapter {
	return this._chapter
}

// read from proto
func (this *PlayerData) FromProto(pdata *netproto.PlayerData) {
	// BaseInfo getter and setter
	if pdata.BaseInfo != nil {
		this.GetBaseInfo().FromProto(pdata.BaseInfo)
	}
	// VipInfo getter and setter
	if pdata.VipInfo != nil {
		this.GetVipInfo().FromProto(pdata.VipInfo)
	}
	// HeroPack getter and setter
	if pdata.HeroPack != nil {
		this.GetHeroPack().FromProto(pdata.HeroPack)
	}
	// EquipPack getter and setter
	if pdata.EquipPack != nil {
		this.GetEquipPack().FromProto(pdata.EquipPack)
	}
	// Guide getter and setter
	if pdata.Guide != nil {
		this.GetGuide().FromProto(pdata.Guide)
	}
	// Chapter getter and setter
	if pdata.Chapter != nil {
		this.GetChapter().FromProto(pdata.Chapter)
	}
}

// write to proto
func (this *PlayerData) ToProto() *netproto.PlayerData {
	pdata := &netproto.PlayerData{}
	pdata.BaseInfo = this._baseInfo.ToProto()
	pdata.VipInfo = this._vipInfo.ToProto()
	pdata.HeroPack = this._heroPack.ToProto()
	pdata.EquipPack = this._equipPack.ToProto()
	pdata.Guide = this._guide.ToProto()
	pdata.Chapter = this._chapter.ToProto()
	return pdata
}

// write dirty to proto
func (this *PlayerData) ToProtoDirty() *netproto.PlayerData {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.PlayerData{}
	if this.IsDirty(1) {
		pdata.BaseInfo = this._baseInfo.ToProtoDirty()
	}
	if this.IsDirty(2) {
		pdata.VipInfo = this._vipInfo.ToProtoDirty()
	}
	if this.IsDirty(3) {
		pdata.HeroPack = this._heroPack.ToProtoDirty()
	}
	if this.IsDirty(4) {
		pdata.EquipPack = this._equipPack.ToProtoDirty()
	}
	if this.IsDirty(5) {
		pdata.Guide = this._guide.ToProtoDirty()
	}
	if this.IsDirty(6) {
		pdata.Chapter = this._chapter.ToProtoDirty()
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
	ins := &PlayerInfo{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	return ins
}

// ServerTime getter and setter
func (this *PlayerInfo) GetServerTime() int64 {
	return this._serverTime
}
func (this *PlayerInfo) SetServerTime(val int64) {
	this._serverTime = val
	this.MarkDirtyUp(1)
}

// ServerName getter and setter
func (this *PlayerInfo) GetServerName() string {
	return this._serverName
}
func (this *PlayerInfo) SetServerName(val string) {
	this._serverName = val
	this.MarkDirtyUp(2)
}

// read from proto
func (this *PlayerInfo) FromProto(pdata *netproto.PlayerInfo) {
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
func (this *PlayerInfo) ToProto() *netproto.PlayerInfo {
	pdata := &netproto.PlayerInfo{}
	pdata.ServerTime = proto.Int64(this._serverTime)
	pdata.ServerName = proto.String(this._serverName)
	return pdata
}

// write dirty to proto
func (this *PlayerInfo) ToProtoDirty() *netproto.PlayerInfo {
	if this.IsAllDirty() {
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
	_id      int64
	_name    string
	_level   int32
	_exp     int32
	_gold    int32
	_head    int32
	_diamond int32
}

// ctor
func NewBaseInfo() *BaseInfo {
	ins := &BaseInfo{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	return ins
}

// Id getter and setter
func (this *BaseInfo) GetId() int64 {
	return this._id
}
func (this *BaseInfo) SetId(val int64) {
	this._id = val
	this.MarkDirtyUp(1)
}

// Name getter and setter
func (this *BaseInfo) GetName() string {
	return this._name
}
func (this *BaseInfo) SetName(val string) {
	this._name = val
	this.MarkDirtyUp(2)
}

// Level getter and setter
func (this *BaseInfo) GetLevel() int32 {
	return this._level
}
func (this *BaseInfo) SetLevel(val int32) {
	this._level = val
	this.MarkDirtyUp(3)
}

// Exp getter and setter
func (this *BaseInfo) GetExp() int32 {
	return this._exp
}
func (this *BaseInfo) SetExp(val int32) {
	this._exp = val
	this.MarkDirtyUp(4)
}

// Gold getter and setter
func (this *BaseInfo) GetGold() int32 {
	return this._gold
}
func (this *BaseInfo) SetGold(val int32) {
	this._gold = val
	this.MarkDirtyUp(5)
}

// Head getter and setter
func (this *BaseInfo) GetHead() int32 {
	return this._head
}
func (this *BaseInfo) SetHead(val int32) {
	this._head = val
	this.MarkDirtyUp(6)
}

// Diamond getter and setter
func (this *BaseInfo) GetDiamond() int32 {
	return this._diamond
}
func (this *BaseInfo) SetDiamond(val int32) {
	this._diamond = val
	this.MarkDirtyUp(7)
}

// read from proto
func (this *BaseInfo) FromProto(pdata *netproto.BaseInfo) {
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
	// Exp getter and setter
	if pdata.Exp != nil {
		this.SetExp(*pdata.Exp)
	}
	// Gold getter and setter
	if pdata.Gold != nil {
		this.SetGold(*pdata.Gold)
	}
	// Head getter and setter
	if pdata.Head != nil {
		this.SetHead(*pdata.Head)
	}
	// Diamond getter and setter
	if pdata.Diamond != nil {
		this.SetDiamond(*pdata.Diamond)
	}
}

// write to proto
func (this *BaseInfo) ToProto() *netproto.BaseInfo {
	pdata := &netproto.BaseInfo{}
	pdata.Id = proto.Int64(this._id)
	pdata.Name = proto.String(this._name)
	pdata.Level = proto.Int32(this._level)
	pdata.Exp = proto.Int32(this._exp)
	pdata.Gold = proto.Int32(this._gold)
	pdata.Head = proto.Int32(this._head)
	pdata.Diamond = proto.Int32(this._diamond)
	return pdata
}

// write dirty to proto
func (this *BaseInfo) ToProtoDirty() *netproto.BaseInfo {
	if this.IsAllDirty() {
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
		pdata.Exp = proto.Int32(this._exp)
	}
	if this.IsDirty(5) {
		pdata.Gold = proto.Int32(this._gold)
	}
	if this.IsDirty(6) {
		pdata.Head = proto.Int32(this._head)
	}
	if this.IsDirty(7) {
		pdata.Diamond = proto.Int32(this._diamond)
	}
	return pdata
}

type Guide struct {
	*datasync.DirtyContainerMarkVector
	_id int64
}

// ctor
func NewGuide() *Guide {
	ins := &Guide{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	return ins
}

// Id getter and setter
func (this *Guide) GetId() int64 {
	return this._id
}
func (this *Guide) SetId(val int64) {
	this._id = val
	this.MarkDirtyUp(1)
}

// read from proto
func (this *Guide) FromProto(pdata *netproto.Guide) {
	// Id getter and setter
	if pdata.Id != nil {
		this.SetId(*pdata.Id)
	}
}

// write to proto
func (this *Guide) ToProto() *netproto.Guide {
	pdata := &netproto.Guide{}
	pdata.Id = proto.Int64(this._id)
	return pdata
}

// write dirty to proto
func (this *Guide) ToProtoDirty() *netproto.Guide {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.Guide{}
	if this.IsDirty(1) {
		pdata.Id = proto.Int64(this._id)
	}
	return pdata
}

type Chapter struct {
	*datasync.DirtyContainerMarkVector
	_heros          *ContainerChapterheros
	_level          int32
	_prizeCheckTime int64
	_prizeEquip     *ContainerChapterprizeEquip
}

// ctor
func NewChapter() *Chapter {
	ins := &Chapter{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	ins._heros = &ContainerChapterheros{
		_impl: datasync.NewDirtyContainerVector(),
	}
	ins._heros._impl.Init(ins, 1)
	ins._prizeEquip = &ContainerChapterprizeEquip{
		_impl: datasync.NewDirtyContainerVector(),
	}
	ins._prizeEquip._impl.Init(ins, 5)
	return ins
}

// Heros getter and setter
func (this *Chapter) GetHeros() *ContainerChapterheros {
	return this._heros
}

// Level getter and setter
func (this *Chapter) GetLevel() int32 {
	return this._level
}
func (this *Chapter) SetLevel(val int32) {
	this._level = val
	this.MarkDirtyUp(2)
}

// PrizeCheckTime getter and setter
func (this *Chapter) GetPrizeCheckTime() int64 {
	return this._prizeCheckTime
}
func (this *Chapter) SetPrizeCheckTime(val int64) {
	this._prizeCheckTime = val
	this.MarkDirtyUp(3)
}

// PrizeEquip getter and setter
func (this *Chapter) GetPrizeEquip() *ContainerChapterprizeEquip {
	return this._prizeEquip
}

// read from proto
func (this *Chapter) FromProto(pdata *netproto.Chapter) {
	// Heros getter and setter
	for _, val := range pdata.Heros {
		this._heros.Set(*val.Key, *val.Val)
	}
	// Level getter and setter
	if pdata.Level != nil {
		this.SetLevel(*pdata.Level)
	}
	// PrizeCheckTime getter and setter
	if pdata.PrizeCheckTime != nil {
		this.SetPrizeCheckTime(*pdata.PrizeCheckTime)
	}
	// PrizeEquip getter and setter
	for _, val := range pdata.PrizeEquip {
		ele := NewEquip()
		ele.FromProto(val.Val)
		this._prizeEquip.Set(*val.Key, ele)
	}
}

// write to proto
func (this *Chapter) ToProto() *netproto.Chapter {
	pdata := &netproto.Chapter{}
	this._heros.Foreach(func(k int32, v int32) {
		pdata.Heros = append(pdata.Heros, NewChapterherosProto(k, v))
	})
	pdata.Level = proto.Int32(this._level)
	pdata.PrizeCheckTime = proto.Int64(this._prizeCheckTime)
	this._prizeEquip.Foreach(func(k int32, v *Equip) {
		pdata.PrizeEquip = append(pdata.PrizeEquip, NewChapterprizeEquipProto(k, v))
	})
	return pdata
}

// write dirty to proto
func (this *Chapter) ToProtoDirty() *netproto.Chapter {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.Chapter{}
	if this.IsDirty(1) {
		this._heros.ForeachDirty(func(k int32, v int32) {
			pdata.Heros = append(pdata.Heros, NewChapterherosProtoDirty(k, v))
		})
	}
	if this.IsDirty(2) {
		pdata.Level = proto.Int32(this._level)
	}
	if this.IsDirty(3) {
		pdata.PrizeCheckTime = proto.Int64(this._prizeCheckTime)
	}
	if this.IsDirty(5) {
		this._prizeEquip.ForeachDirty(func(k int32, v *Equip) {
			pdata.PrizeEquip = append(pdata.PrizeEquip, NewChapterprizeEquipProtoDirty(k, v))
		})
	}
	return pdata
}

type VipInfo struct {
	*datasync.DirtyContainerMarkVector
	_level         int32
	_exp           int32
	_totalRecharge int32
}

// ctor
func NewVipInfo() *VipInfo {
	ins := &VipInfo{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	return ins
}

// Level getter and setter
func (this *VipInfo) GetLevel() int32 {
	return this._level
}
func (this *VipInfo) SetLevel(val int32) {
	this._level = val
	this.MarkDirtyUp(1)
}

// Exp getter and setter
func (this *VipInfo) GetExp() int32 {
	return this._exp
}
func (this *VipInfo) SetExp(val int32) {
	this._exp = val
	this.MarkDirtyUp(2)
}

// TotalRecharge getter and setter
func (this *VipInfo) GetTotalRecharge() int32 {
	return this._totalRecharge
}
func (this *VipInfo) SetTotalRecharge(val int32) {
	this._totalRecharge = val
	this.MarkDirtyUp(3)
}

// read from proto
func (this *VipInfo) FromProto(pdata *netproto.VipInfo) {
	// Level getter and setter
	if pdata.Level != nil {
		this.SetLevel(*pdata.Level)
	}
	// Exp getter and setter
	if pdata.Exp != nil {
		this.SetExp(*pdata.Exp)
	}
	// TotalRecharge getter and setter
	if pdata.TotalRecharge != nil {
		this.SetTotalRecharge(*pdata.TotalRecharge)
	}
}

// write to proto
func (this *VipInfo) ToProto() *netproto.VipInfo {
	pdata := &netproto.VipInfo{}
	pdata.Level = proto.Int32(this._level)
	pdata.Exp = proto.Int32(this._exp)
	pdata.TotalRecharge = proto.Int32(this._totalRecharge)
	return pdata
}

// write dirty to proto
func (this *VipInfo) ToProtoDirty() *netproto.VipInfo {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.VipInfo{}
	if this.IsDirty(1) {
		pdata.Level = proto.Int32(this._level)
	}
	if this.IsDirty(2) {
		pdata.Exp = proto.Int32(this._exp)
	}
	if this.IsDirty(3) {
		pdata.TotalRecharge = proto.Int32(this._totalRecharge)
	}
	return pdata
}

type Hero struct {
	*datasync.DirtyContainerMarkVector
	_id     int32
	_level  int32
	_status int64
	_equips *ContainerHeroequips
	_num    int32
}

// ctor
func NewHero() *Hero {
	ins := &Hero{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	ins._equips = &ContainerHeroequips{
		_impl: datasync.NewDirtyContainerVector(),
	}
	ins._equips._impl.Init(ins, 4)
	return ins
}

// Id getter and setter
func (this *Hero) GetId() int32 {
	return this._id
}
func (this *Hero) SetId(val int32) {
	this._id = val
	this.MarkDirtyUp(1)
}

// Level getter and setter
func (this *Hero) GetLevel() int32 {
	return this._level
}
func (this *Hero) SetLevel(val int32) {
	this._level = val
	this.MarkDirtyUp(2)
}

// Status getter and setter
func (this *Hero) GetStatus() int64 {
	return this._status
}
func (this *Hero) SetStatus(val int64) {
	this._status = val
	this.MarkDirtyUp(3)
}

// Equips getter and setter
func (this *Hero) GetEquips() *ContainerHeroequips {
	return this._equips
}

// Num getter and setter
func (this *Hero) GetNum() int32 {
	return this._num
}
func (this *Hero) SetNum(val int32) {
	this._num = val
	this.MarkDirtyUp(5)
}

// read from proto
func (this *Hero) FromProto(pdata *netproto.Hero) {
	// Id getter and setter
	if pdata.Id != nil {
		this.SetId(*pdata.Id)
	}
	// Level getter and setter
	if pdata.Level != nil {
		this.SetLevel(*pdata.Level)
	}
	// Status getter and setter
	if pdata.Status != nil {
		this.SetStatus(*pdata.Status)
	}
	// Equips getter and setter
	for _, val := range pdata.Equips {
		ele := NewEquip()
		ele.FromProto(val.Val)
		this._equips.Set(*val.Key, ele)
	}
	// Num getter and setter
	if pdata.Num != nil {
		this.SetNum(*pdata.Num)
	}
}

// write to proto
func (this *Hero) ToProto() *netproto.Hero {
	pdata := &netproto.Hero{}
	pdata.Id = proto.Int32(this._id)
	pdata.Level = proto.Int32(this._level)
	pdata.Status = proto.Int64(this._status)
	this._equips.Foreach(func(k int32, v *Equip) {
		pdata.Equips = append(pdata.Equips, NewHeroequipsProto(k, v))
	})
	pdata.Num = proto.Int32(this._num)
	return pdata
}

// write dirty to proto
func (this *Hero) ToProtoDirty() *netproto.Hero {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.Hero{}
	if this.IsDirty(1) {
		pdata.Id = proto.Int32(this._id)
	}
	if this.IsDirty(2) {
		pdata.Level = proto.Int32(this._level)
	}
	if this.IsDirty(3) {
		pdata.Status = proto.Int64(this._status)
	}
	if this.IsDirty(4) {
		this._equips.ForeachDirty(func(k int32, v *Equip) {
			pdata.Equips = append(pdata.Equips, NewHeroequipsProtoDirty(k, v))
		})
	}
	if this.IsDirty(5) {
		pdata.Num = proto.Int32(this._num)
	}
	return pdata
}

type HeroPack struct {
	*datasync.DirtyContainerMarkVector
	_limit int32
	_heros *ContainerHeroPackheros
}

// ctor
func NewHeroPack() *HeroPack {
	ins := &HeroPack{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	ins._heros = &ContainerHeroPackheros{
		_impl: datasync.NewDirtyContainerMap(),
	}
	ins._heros._impl.Init(ins, 2)
	return ins
}

// Limit getter and setter
func (this *HeroPack) GetLimit() int32 {
	return this._limit
}
func (this *HeroPack) SetLimit(val int32) {
	this._limit = val
	this.MarkDirtyUp(1)
}

// Heros getter and setter
func (this *HeroPack) GetHeros() *ContainerHeroPackheros {
	return this._heros
}

// read from proto
func (this *HeroPack) FromProto(pdata *netproto.HeroPack) {
	// Limit getter and setter
	if pdata.Limit != nil {
		this.SetLimit(*pdata.Limit)
	}
	// Heros getter and setter
	for _, val := range pdata.Heros {
		ele := NewHero()
		ele.FromProto(val.Val)
		this._heros.Set(*val.Key, ele)
	}
}

// write to proto
func (this *HeroPack) ToProto() *netproto.HeroPack {
	pdata := &netproto.HeroPack{}
	pdata.Limit = proto.Int32(this._limit)
	this._heros.Foreach(func(k int32, v *Hero) {
		pdata.Heros = append(pdata.Heros, NewHeroPackherosProto(k, v))
	})
	return pdata
}

// write dirty to proto
func (this *HeroPack) ToProtoDirty() *netproto.HeroPack {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.HeroPack{}
	if this.IsDirty(1) {
		pdata.Limit = proto.Int32(this._limit)
	}
	if this.IsDirty(2) {
		this._heros.ForeachDirty(func(k int32, v *Hero) {
			pdata.Heros = append(pdata.Heros, NewHeroPackherosProtoDirty(k, v))
		})
	}
	return pdata
}

type Equip struct {
	*datasync.DirtyContainerMarkVector
	_id     int32
	_level  int32
	_status int64
	_num    int32
}

// ctor
func NewEquip() *Equip {
	ins := &Equip{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	return ins
}

// Id getter and setter
func (this *Equip) GetId() int32 {
	return this._id
}
func (this *Equip) SetId(val int32) {
	this._id = val
	this.MarkDirtyUp(1)
}

// Level getter and setter
func (this *Equip) GetLevel() int32 {
	return this._level
}
func (this *Equip) SetLevel(val int32) {
	this._level = val
	this.MarkDirtyUp(2)
}

// Status getter and setter
func (this *Equip) GetStatus() int64 {
	return this._status
}
func (this *Equip) SetStatus(val int64) {
	this._status = val
	this.MarkDirtyUp(3)
}

// Num getter and setter
func (this *Equip) GetNum() int32 {
	return this._num
}
func (this *Equip) SetNum(val int32) {
	this._num = val
	this.MarkDirtyUp(4)
}

// read from proto
func (this *Equip) FromProto(pdata *netproto.Equip) {
	// Id getter and setter
	if pdata.Id != nil {
		this.SetId(*pdata.Id)
	}
	// Level getter and setter
	if pdata.Level != nil {
		this.SetLevel(*pdata.Level)
	}
	// Status getter and setter
	if pdata.Status != nil {
		this.SetStatus(*pdata.Status)
	}
	// Num getter and setter
	if pdata.Num != nil {
		this.SetNum(*pdata.Num)
	}
}

// write to proto
func (this *Equip) ToProto() *netproto.Equip {
	pdata := &netproto.Equip{}
	pdata.Id = proto.Int32(this._id)
	pdata.Level = proto.Int32(this._level)
	pdata.Status = proto.Int64(this._status)
	pdata.Num = proto.Int32(this._num)
	return pdata
}

// write dirty to proto
func (this *Equip) ToProtoDirty() *netproto.Equip {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.Equip{}
	if this.IsDirty(1) {
		pdata.Id = proto.Int32(this._id)
	}
	if this.IsDirty(2) {
		pdata.Level = proto.Int32(this._level)
	}
	if this.IsDirty(3) {
		pdata.Status = proto.Int64(this._status)
	}
	if this.IsDirty(4) {
		pdata.Num = proto.Int32(this._num)
	}
	return pdata
}

type EquipPack struct {
	*datasync.DirtyContainerMarkVector
	_limit  int32
	_equips *ContainerEquipPackequips
}

// ctor
func NewEquipPack() *EquipPack {
	ins := &EquipPack{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	ins._equips = &ContainerEquipPackequips{
		_impl: datasync.NewDirtyContainerMap(),
	}
	ins._equips._impl.Init(ins, 2)
	return ins
}

// Limit getter and setter
func (this *EquipPack) GetLimit() int32 {
	return this._limit
}
func (this *EquipPack) SetLimit(val int32) {
	this._limit = val
	this.MarkDirtyUp(1)
}

// Equips getter and setter
func (this *EquipPack) GetEquips() *ContainerEquipPackequips {
	return this._equips
}

// read from proto
func (this *EquipPack) FromProto(pdata *netproto.EquipPack) {
	// Limit getter and setter
	if pdata.Limit != nil {
		this.SetLimit(*pdata.Limit)
	}
	// Equips getter and setter
	for _, val := range pdata.Equips {
		ele := NewEquip()
		ele.FromProto(val.Val)
		this._equips.Set(*val.Key, ele)
	}
}

// write to proto
func (this *EquipPack) ToProto() *netproto.EquipPack {
	pdata := &netproto.EquipPack{}
	pdata.Limit = proto.Int32(this._limit)
	this._equips.Foreach(func(k int32, v *Equip) {
		pdata.Equips = append(pdata.Equips, NewEquipPackequipsProto(k, v))
	})
	return pdata
}

// write dirty to proto
func (this *EquipPack) ToProtoDirty() *netproto.EquipPack {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.EquipPack{}
	if this.IsDirty(1) {
		pdata.Limit = proto.Int32(this._limit)
	}
	if this.IsDirty(2) {
		this._equips.ForeachDirty(func(k int32, v *Equip) {
			pdata.Equips = append(pdata.Equips, NewEquipPackequipsProtoDirty(k, v))
		})
	}
	return pdata
}

type Daily struct {
	*datasync.DirtyContainerMarkVector
	_lastUpdateTime int64
	_sign           bool
	_totalSign      int32
}

// ctor
func NewDaily() *Daily {
	ins := &Daily{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	return ins
}

// LastUpdateTime getter and setter
func (this *Daily) GetLastUpdateTime() int64 {
	return this._lastUpdateTime
}
func (this *Daily) SetLastUpdateTime(val int64) {
	this._lastUpdateTime = val
	this.MarkDirtyUp(1)
}

// Sign getter and setter
func (this *Daily) GetSign() bool {
	return this._sign
}
func (this *Daily) SetSign(val bool) {
	this._sign = val
	this.MarkDirtyUp(2)
}

// TotalSign getter and setter
func (this *Daily) GetTotalSign() int32 {
	return this._totalSign
}
func (this *Daily) SetTotalSign(val int32) {
	this._totalSign = val
	this.MarkDirtyUp(3)
}

// read from proto
func (this *Daily) FromProto(pdata *netproto.Daily) {
	// LastUpdateTime getter and setter
	if pdata.LastUpdateTime != nil {
		this.SetLastUpdateTime(*pdata.LastUpdateTime)
	}
	// Sign getter and setter
	if pdata.Sign != nil {
		this.SetSign(*pdata.Sign)
	}
	// TotalSign getter and setter
	if pdata.TotalSign != nil {
		this.SetTotalSign(*pdata.TotalSign)
	}
}

// write to proto
func (this *Daily) ToProto() *netproto.Daily {
	pdata := &netproto.Daily{}
	pdata.LastUpdateTime = proto.Int64(this._lastUpdateTime)
	pdata.Sign = proto.Bool(this._sign)
	pdata.TotalSign = proto.Int32(this._totalSign)
	return pdata
}

// write dirty to proto
func (this *Daily) ToProtoDirty() *netproto.Daily {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.Daily{}
	if this.IsDirty(1) {
		pdata.LastUpdateTime = proto.Int64(this._lastUpdateTime)
	}
	if this.IsDirty(2) {
		pdata.Sign = proto.Bool(this._sign)
	}
	if this.IsDirty(3) {
		pdata.TotalSign = proto.Int32(this._totalSign)
	}
	return pdata
}

type TestList struct {
	*datasync.DirtyContainerMarkVector
	_id               int64
	_ListPrimitiveInt *ContainerTestListListPrimitiveInt
	_ListPrimitiveStr *ContainerTestListListPrimitiveStr
	_ListPrimitiveCom *ContainerTestListListPrimitiveCom
}

// ctor
func NewTestList() *TestList {
	ins := &TestList{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	ins._ListPrimitiveInt = &ContainerTestListListPrimitiveInt{
		_impl: datasync.NewDirtyContainerVector(),
	}
	ins._ListPrimitiveInt._impl.Init(ins, 1)
	ins._ListPrimitiveStr = &ContainerTestListListPrimitiveStr{
		_impl: datasync.NewDirtyContainerVector(),
	}
	ins._ListPrimitiveStr._impl.Init(ins, 2)
	ins._ListPrimitiveCom = &ContainerTestListListPrimitiveCom{
		_impl: datasync.NewDirtyContainerVector(),
	}
	ins._ListPrimitiveCom._impl.Init(ins, 3)
	return ins
}

// Id getter and setter
func (this *TestList) GetId() int64 {
	return this._id
}
func (this *TestList) SetId(val int64) {
	this._id = val
	this.MarkDirtyUp(7)
}

// ListPrimitiveInt getter and setter
func (this *TestList) GetListPrimitiveInt() *ContainerTestListListPrimitiveInt {
	return this._ListPrimitiveInt
}

// ListPrimitiveStr getter and setter
func (this *TestList) GetListPrimitiveStr() *ContainerTestListListPrimitiveStr {
	return this._ListPrimitiveStr
}

// ListPrimitiveCom getter and setter
func (this *TestList) GetListPrimitiveCom() *ContainerTestListListPrimitiveCom {
	return this._ListPrimitiveCom
}

// read from proto
func (this *TestList) FromProto(pdata *netproto.TestList) {
	// Id getter and setter
	if pdata.Id != nil {
		this.SetId(*pdata.Id)
	}
	// ListPrimitiveInt getter and setter
	for _, val := range pdata.ListPrimitiveInt {
		this._ListPrimitiveInt.Set(*val.Key, *val.Val)
	}
	// ListPrimitiveStr getter and setter
	for _, val := range pdata.ListPrimitiveStr {
		this._ListPrimitiveStr.Set(*val.Key, *val.Val)
	}
	// ListPrimitiveCom getter and setter
	for _, val := range pdata.ListPrimitiveCom {
		ele := NewPlayerData()
		ele.FromProto(val.Val)
		this._ListPrimitiveCom.Set(*val.Key, ele)
	}
}

// write to proto
func (this *TestList) ToProto() *netproto.TestList {
	pdata := &netproto.TestList{}
	pdata.Id = proto.Int64(this._id)
	this._ListPrimitiveInt.Foreach(func(k int32, v int32) {
		pdata.ListPrimitiveInt = append(pdata.ListPrimitiveInt, NewTestListListPrimitiveIntProto(k, v))
	})
	this._ListPrimitiveStr.Foreach(func(k int32, v string) {
		pdata.ListPrimitiveStr = append(pdata.ListPrimitiveStr, NewTestListListPrimitiveStrProto(k, v))
	})
	this._ListPrimitiveCom.Foreach(func(k int32, v *PlayerData) {
		pdata.ListPrimitiveCom = append(pdata.ListPrimitiveCom, NewTestListListPrimitiveComProto(k, v))
	})
	return pdata
}

// write dirty to proto
func (this *TestList) ToProtoDirty() *netproto.TestList {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.TestList{}
	if this.IsDirty(7) {
		pdata.Id = proto.Int64(this._id)
	}
	if this.IsDirty(1) {
		this._ListPrimitiveInt.ForeachDirty(func(k int32, v int32) {
			pdata.ListPrimitiveInt = append(pdata.ListPrimitiveInt, NewTestListListPrimitiveIntProtoDirty(k, v))
		})
	}
	if this.IsDirty(2) {
		this._ListPrimitiveStr.ForeachDirty(func(k int32, v string) {
			pdata.ListPrimitiveStr = append(pdata.ListPrimitiveStr, NewTestListListPrimitiveStrProtoDirty(k, v))
		})
	}
	if this.IsDirty(3) {
		this._ListPrimitiveCom.ForeachDirty(func(k int32, v *PlayerData) {
			pdata.ListPrimitiveCom = append(pdata.ListPrimitiveCom, NewTestListListPrimitiveComProtoDirty(k, v))
		})
	}
	return pdata
}

type TestMap struct {
	*datasync.DirtyContainerMarkVector
	_id                 *BaseInfo
	_MapPrimitiveIntInt *ContainerTestMapMapPrimitiveIntInt
	_MapPrimitiveIntStr *ContainerTestMapMapPrimitiveIntStr
	_MapPrimitiveStrInt *ContainerTestMapMapPrimitiveStrInt
	_MapPrimitiveStrStr *ContainerTestMapMapPrimitiveStrStr
	_MapPrimitiveIntCom *ContainerTestMapMapPrimitiveIntCom
	_MapPrimitiveStrCom *ContainerTestMapMapPrimitiveStrCom
}

// ctor
func NewTestMap() *TestMap {
	ins := &TestMap{
		DirtyContainerMarkVector: datasync.NewDirtyContainerMarkVector(),
	}
	ins._id = NewBaseInfo()
	ins._id.Init(ins, 7)
	ins._MapPrimitiveIntInt = &ContainerTestMapMapPrimitiveIntInt{
		_impl: datasync.NewDirtyContainerMap(),
	}
	ins._MapPrimitiveIntInt._impl.Init(ins, 1)
	ins._MapPrimitiveIntStr = &ContainerTestMapMapPrimitiveIntStr{
		_impl: datasync.NewDirtyContainerMap(),
	}
	ins._MapPrimitiveIntStr._impl.Init(ins, 2)
	ins._MapPrimitiveStrInt = &ContainerTestMapMapPrimitiveStrInt{
		_impl: datasync.NewDirtyContainerMap(),
	}
	ins._MapPrimitiveStrInt._impl.Init(ins, 3)
	ins._MapPrimitiveStrStr = &ContainerTestMapMapPrimitiveStrStr{
		_impl: datasync.NewDirtyContainerMap(),
	}
	ins._MapPrimitiveStrStr._impl.Init(ins, 4)
	ins._MapPrimitiveIntCom = &ContainerTestMapMapPrimitiveIntCom{
		_impl: datasync.NewDirtyContainerMap(),
	}
	ins._MapPrimitiveIntCom._impl.Init(ins, 5)
	ins._MapPrimitiveStrCom = &ContainerTestMapMapPrimitiveStrCom{
		_impl: datasync.NewDirtyContainerMap(),
	}
	ins._MapPrimitiveStrCom._impl.Init(ins, 6)
	return ins
}

// Id getter and setter
func (this *TestMap) GetId() *BaseInfo {
	return this._id
}

// MapPrimitiveIntInt getter and setter
func (this *TestMap) GetMapPrimitiveIntInt() *ContainerTestMapMapPrimitiveIntInt {
	return this._MapPrimitiveIntInt
}

// MapPrimitiveIntStr getter and setter
func (this *TestMap) GetMapPrimitiveIntStr() *ContainerTestMapMapPrimitiveIntStr {
	return this._MapPrimitiveIntStr
}

// MapPrimitiveStrInt getter and setter
func (this *TestMap) GetMapPrimitiveStrInt() *ContainerTestMapMapPrimitiveStrInt {
	return this._MapPrimitiveStrInt
}

// MapPrimitiveStrStr getter and setter
func (this *TestMap) GetMapPrimitiveStrStr() *ContainerTestMapMapPrimitiveStrStr {
	return this._MapPrimitiveStrStr
}

// MapPrimitiveIntCom getter and setter
func (this *TestMap) GetMapPrimitiveIntCom() *ContainerTestMapMapPrimitiveIntCom {
	return this._MapPrimitiveIntCom
}

// MapPrimitiveStrCom getter and setter
func (this *TestMap) GetMapPrimitiveStrCom() *ContainerTestMapMapPrimitiveStrCom {
	return this._MapPrimitiveStrCom
}

// read from proto
func (this *TestMap) FromProto(pdata *netproto.TestMap) {
	// Id getter and setter
	if pdata.Id != nil {
		this.GetId().FromProto(pdata.Id)
	}
	// MapPrimitiveIntInt getter and setter
	for _, val := range pdata.MapPrimitiveIntInt {
		this._MapPrimitiveIntInt.Set(*val.Key, *val.Val)
	}
	// MapPrimitiveIntStr getter and setter
	for _, val := range pdata.MapPrimitiveIntStr {
		this._MapPrimitiveIntStr.Set(*val.Key, *val.Val)
	}
	// MapPrimitiveStrInt getter and setter
	for _, val := range pdata.MapPrimitiveStrInt {
		this._MapPrimitiveStrInt.Set(*val.Key, *val.Val)
	}
	// MapPrimitiveStrStr getter and setter
	for _, val := range pdata.MapPrimitiveStrStr {
		this._MapPrimitiveStrStr.Set(*val.Key, *val.Val)
	}
	// MapPrimitiveIntCom getter and setter
	for _, val := range pdata.MapPrimitiveIntCom {
		ele := NewBaseInfo()
		ele.FromProto(val.Val)
		this._MapPrimitiveIntCom.Set(*val.Key, ele)
	}
	// MapPrimitiveStrCom getter and setter
	for _, val := range pdata.MapPrimitiveStrCom {
		ele := NewPlayerData()
		ele.FromProto(val.Val)
		this._MapPrimitiveStrCom.Set(*val.Key, ele)
	}
}

// write to proto
func (this *TestMap) ToProto() *netproto.TestMap {
	pdata := &netproto.TestMap{}
	pdata.Id = this._id.ToProto()
	this._MapPrimitiveIntInt.Foreach(func(k int32, v int32) {
		pdata.MapPrimitiveIntInt = append(pdata.MapPrimitiveIntInt, NewTestMapMapPrimitiveIntIntProto(k, v))
	})
	this._MapPrimitiveIntStr.Foreach(func(k int32, v string) {
		pdata.MapPrimitiveIntStr = append(pdata.MapPrimitiveIntStr, NewTestMapMapPrimitiveIntStrProto(k, v))
	})
	this._MapPrimitiveStrInt.Foreach(func(k string, v int32) {
		pdata.MapPrimitiveStrInt = append(pdata.MapPrimitiveStrInt, NewTestMapMapPrimitiveStrIntProto(k, v))
	})
	this._MapPrimitiveStrStr.Foreach(func(k string, v string) {
		pdata.MapPrimitiveStrStr = append(pdata.MapPrimitiveStrStr, NewTestMapMapPrimitiveStrStrProto(k, v))
	})
	this._MapPrimitiveIntCom.Foreach(func(k int32, v *BaseInfo) {
		pdata.MapPrimitiveIntCom = append(pdata.MapPrimitiveIntCom, NewTestMapMapPrimitiveIntComProto(k, v))
	})
	this._MapPrimitiveStrCom.Foreach(func(k string, v *PlayerData) {
		pdata.MapPrimitiveStrCom = append(pdata.MapPrimitiveStrCom, NewTestMapMapPrimitiveStrComProto(k, v))
	})
	return pdata
}

// write dirty to proto
func (this *TestMap) ToProtoDirty() *netproto.TestMap {
	if this.IsAllDirty() {
		return this.ToProto()
	}
	pdata := &netproto.TestMap{}
	if this.IsDirty(7) {
		pdata.Id = this._id.ToProtoDirty()
	}
	if this.IsDirty(1) {
		this._MapPrimitiveIntInt.ForeachDirty(func(k int32, v int32) {
			pdata.MapPrimitiveIntInt = append(pdata.MapPrimitiveIntInt, NewTestMapMapPrimitiveIntIntProtoDirty(k, v))
		})
	}
	if this.IsDirty(2) {
		this._MapPrimitiveIntStr.ForeachDirty(func(k int32, v string) {
			pdata.MapPrimitiveIntStr = append(pdata.MapPrimitiveIntStr, NewTestMapMapPrimitiveIntStrProtoDirty(k, v))
		})
	}
	if this.IsDirty(3) {
		this._MapPrimitiveStrInt.ForeachDirty(func(k string, v int32) {
			pdata.MapPrimitiveStrInt = append(pdata.MapPrimitiveStrInt, NewTestMapMapPrimitiveStrIntProtoDirty(k, v))
		})
	}
	if this.IsDirty(4) {
		this._MapPrimitiveStrStr.ForeachDirty(func(k string, v string) {
			pdata.MapPrimitiveStrStr = append(pdata.MapPrimitiveStrStr, NewTestMapMapPrimitiveStrStrProtoDirty(k, v))
		})
	}
	if this.IsDirty(5) {
		this._MapPrimitiveIntCom.ForeachDirty(func(k int32, v *BaseInfo) {
			pdata.MapPrimitiveIntCom = append(pdata.MapPrimitiveIntCom, NewTestMapMapPrimitiveIntComProtoDirty(k, v))
		})
	}
	if this.IsDirty(6) {
		this._MapPrimitiveStrCom.ForeachDirty(func(k string, v *PlayerData) {
			pdata.MapPrimitiveStrCom = append(pdata.MapPrimitiveStrCom, NewTestMapMapPrimitiveStrComProtoDirty(k, v))
		})
	}
	return pdata
}

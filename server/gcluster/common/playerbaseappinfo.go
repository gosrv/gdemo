package common

import (
	"github.com/gosrv/gbase/gdb/gredis"
	"github.com/gosrv/gbase/gutil"
	"github.com/gosrv/gcluster/gcluster/common/meta"
	"strconv"
	"sync"
	"time"
)

const (
	MaxCacheSize = 10000
	MaxCacheTimeSeconds = 600
)

type BaseappUuid struct {
	uuid string
	activeTime int64
}

type PlayerBaseappInfo struct {
	plBaseappUuid sync.Map
	keyPrefix string
	hashOpt *gredis.HashOperation	`redis:""`
	maxCacheSize int
	maxCacheTimeSeconds int
}

func (this *PlayerBaseappInfo) BeanStart() {
	gutil.RecoverGo(func() {
		for {
			this.disentangle()
			time.Sleep(10 * time.Second)
		}
	})
}

func (this *PlayerBaseappInfo) BeanStop() {

}

func NewPlayerBaseappInfo(keyPrefix string, maxCacheSize int, maxCacheTimeSeconds int) *PlayerBaseappInfo {
	if maxCacheSize == 0 {
		maxCacheSize = MaxCacheSize
	}
	if maxCacheTimeSeconds == 0{
		maxCacheTimeSeconds = MaxCacheTimeSeconds
	}
	return &PlayerBaseappInfo{keyPrefix: keyPrefix, maxCacheSize: maxCacheSize, maxCacheTimeSeconds: maxCacheTimeSeconds}
}

func (this *PlayerBaseappInfo) getKey(playerId int64) string {
	return this.keyPrefix + strconv.FormatInt(playerId, 10)
}

func (this *PlayerBaseappInfo) GetPlayerBaseappUuidDirect(playerId int64) (string, error) {
	uuid, err := this.hashOpt.HGet(this.getKey(playerId), meta.BaseApp)
	if err != nil {
		return "", err
	}
	luuid, _ := this.plBaseappUuid.LoadOrStore(playerId, &BaseappUuid{uuid:uuid})
	luuid.(*BaseappUuid).activeTime = time.Now().Unix()
	return uuid, nil
}

func (this *PlayerBaseappInfo)GetPlayerBaseappUuid(playerId int64) (string, error) {
	uuid, ok := this.plBaseappUuid.Load(playerId)
	if ok {
		baseUuid := uuid.(*BaseappUuid)
		baseUuid.activeTime = time.Now().Unix()
		return baseUuid.uuid, nil
	}
	return this.GetPlayerBaseappUuidDirect(playerId)
}

func (this *PlayerBaseappInfo)InvalidCacheUuid(playerId int64) {
	this.plBaseappUuid.Delete(playerId)
}

func (this *PlayerBaseappInfo)disentangle() {
	toSize := this.maxCacheSize
	curTime := time.Now().Unix()

	for i := uint(0); i < 8; i++ {
		count := 0
		urgent := this.maxCacheTimeSeconds >> i
		this.plBaseappUuid.Range(func(key, value interface{}) bool {
			uuidData := value.(*BaseappUuid)
			if uuidData.activeTime+int64(urgent) < curTime {
				this.plBaseappUuid.Delete(key)
			} else {
				count++
			}
			return true
		})
		if count <= toSize {
			break
		}
	}
}
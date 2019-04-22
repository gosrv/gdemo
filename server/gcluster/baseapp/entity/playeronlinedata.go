package entity

import "time"

/**
在线玩家数据，只存在于内存中
*/
type PlayerOnlineData struct {
	// 登陆时间
	loginTime time.Time
}

func (this *PlayerOnlineData) LoginTime() time.Time {
	return this.loginTime
}

func NewPlayerOnlineData() *PlayerOnlineData {
	return &PlayerOnlineData{
		loginTime: time.Now(),
	}
}

var PlayerIdKey = struct{}{}

package logic

import (
	"encoding/json"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gosrv/gbase/gdb/gmongo"
	"github.com/gosrv/gbase/gdb/gredis"
	"github.com/gosrv/gbase/gutil"
	"github.com/gosrv/gcluster/gcluster/common/entity"
	"github.com/gosrv/glog"
	"time"
)

const (
	TokenExpireTimeSeconds = 3600
	TokenPrefix            = "tk:"
	AccTokenPrefix         = "acctk:"
	AccIdGenerator         = "_id.generator"
)

type serviceLogin struct {
	log           glog.IFieldLogger      `log:"app"`
	tokenCacheOpt *gredis.ValueOperation `redis:""`
	tabAccount    *mgo.Collection        `mongo.c:"account"`
}

func NewServiceLogin() *serviceLogin {
	return &serviceLogin{}
}

type LoginAccount struct {
	Account  string `bson:"_id"`
	Pw       string
	PlayerId int64
}

func NewLoginAccount(Account string, pw string, playerId int64) *LoginAccount {
	return &LoginAccount{Account: Account, Pw: pw, PlayerId: playerId}
}

func (this *serviceLogin) loadTokenByAcc(account string) *entity.LoginTokenCtx {
	tokenCache, err := this.tokenCacheOpt.Get(AccTokenPrefix + account)
	if err != nil {
		this.log.Debug("load acc token %v error %v", account, err)
		return nil
	}
	tokcnIns := &entity.LoginTokenCtx{}
	err = json.Unmarshal([]byte(tokenCache), tokcnIns)
	if err != nil {
		this.log.Debug("unmarshal acc token %v:%v error %v", account, tokenCache, err)
		return nil
	}
	return tokcnIns
}

func (this *serviceLogin) loadTokenByTk(token string) *entity.LoginTokenCtx {
	tokenCache, err := this.tokenCacheOpt.Get(TokenPrefix + token)
	if err != nil {
		this.log.Debug("load token %v error %v", token, err)
		return nil
	}
	tokcnIns := &entity.LoginTokenCtx{}
	err = json.Unmarshal([]byte(tokenCache), tokcnIns)
	if err != nil {
		this.log.Debug("unmarshal token %v:%v error %v", token, tokenCache, err)
		return nil
	}
	return tokcnIns
}

func (this *serviceLogin) saveToken(token *entity.LoginTokenCtx) {
	tokenCache := gutil.Json(token)
	_, err := this.tokenCacheOpt.SetTimeout(TokenPrefix+token.Token, tokenCache, time.Second*TokenExpireTimeSeconds)
	if err != nil {
		this.log.Debug("save token error %v", err)
	}
	_, err = this.tokenCacheOpt.SetTimeout(AccTokenPrefix+token.Account, tokenCache, time.Second*TokenExpireTimeSeconds)
	if err != nil {
		this.log.Debug("save acc token error %v", err)
	}
}

func (this *serviceLogin) login(account, pw string) int64 {
	findKey := bson.M{gmongo.MONGO_KEY_ID: account}
	loginAccount := &LoginAccount{}
	err := this.tabAccount.Find(findKey).One(loginAccount)
	if err == nil {
		if pw == loginAccount.Pw {
			return loginAccount.PlayerId
		} else {
			return 0
		}
	}
	if len(pw) == 0 {
		return this.register(account, pw)
	}
	return 0
}

func (this *serviceLogin) register(account, pw string) int64 {
	if account == AccIdGenerator {
		return 0
	}
	inc := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"count": 1}},
		ReturnNew: true,
		Upsert:    true,
	}
	res := &struct {
		ID    string `bson:"_id"`
		Count int64  `bson:"count"`
	}{}
	_, err := this.tabAccount.Find(bson.M{gmongo.MONGO_KEY_ID: AccIdGenerator}).Apply(inc, res)
	if err != nil {
		this.log.Debug("get account id error %v", err)
		return 0
	}
	playerId := res.Count
	err = this.tabAccount.Insert(NewLoginAccount(account, pw, playerId))
	if err != nil {
		this.log.Debug("insert account error %v", err)
		return 0
	}
	return playerId
}

package common

import (
	"github.com/golang/protobuf/proto"
	"github.com/gosrv/gbase/gnet"
	"github.com/gosrv/gcluster/gcluster/proto"
	"github.com/gosrv/goioc/util"
	"reflect"
)

/**
网络消息类型和id的映射关系建立
*/
var LogicMsgIds gnet.ITypeID
var BattleMsgIds gnet.ITypeID
var ClusterMsgIds gnet.ITypeID

func init() {
	LogicMsgIds = initLogicMsgIds()
	BattleMsgIds = initBattleMsgIds()
	ClusterMsgIds = initClusterMsgIds()
}

func initLogicMsgIds() gnet.ITypeID {
	msgIds := gnet.NewTypeID()
	// 添加增量更新消息
	err := msgIds.AddIDType(1, reflect.TypeOf((*netproto.PlayerData)(nil)))
	util.VerifyNoError(err)
	err = msgIds.AddIDType(2, reflect.TypeOf((*netproto.PlayerInfo)(nil)))
	util.VerifyNoError(err)
	// 添加其它逻辑消息
	for id, name := range netproto.EMsgIds_name {
		err := msgIds.AddIDType(int(id), proto.MessageType("netproto."+name[1:]))
		util.VerifyNoError(err)
	}

	return msgIds
}

func initBattleMsgIds() gnet.ITypeID {
	msgIds := gnet.NewTypeID()
	// 添加其它逻辑消息
	for id, name := range netproto.EBMsgIds_name {
		err := msgIds.AddIDType(int(id), proto.MessageType("netproto."+name[1:]))
		util.VerifyNoError(err)
	}

	return msgIds
}

func initClusterMsgIds() gnet.ITypeID {
	msgIds := gnet.NewTypeID()
	// 添加其它逻辑消息
	for id, name := range netproto.ECMsgIds_name {
		err := msgIds.AddIDType(int(id), proto.MessageType("netproto."+name[1:]))
		util.VerifyNoError(err)
	}

	return msgIds
}

package entity

import (
	"github.com/gosrv/gbase/gdb/gredis"
	"github.com/gosrv/gbase/gproto"
	"github.com/gosrv/gcluster/gcluster/common/meta"
)

type PlayerMQBundle struct {
	*gredis.RedisMQBundle
	encoder gproto.IEncoder
	decoder gproto.IDecoder
	listOpt *gredis.ListOperation	`redis:""`
}

func NewPlayerMQBundle(encoder gproto.IEncoder, decoder gproto.IDecoder) *PlayerMQBundle {
	return &PlayerMQBundle{encoder: encoder, decoder: decoder}
}

func (this *PlayerMQBundle) BeanInit() {
	this.RedisMQBundle = gredis.NewRedisMQBundle(meta.MessageQueuePlayer, this.encoder, this.decoder, this.listOpt)
}

func (this *PlayerMQBundle) BeanUninit() {

}
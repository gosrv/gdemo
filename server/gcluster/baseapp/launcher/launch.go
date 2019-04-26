package main

import (
	"github.com/gosrv/gbase/app"
	"github.com/gosrv/gbase/cluster"
	"github.com/gosrv/gbase/codec"
	"github.com/gosrv/gbase/gdb/gmongo"
	"github.com/gosrv/gbase/gdb/gredis"
	"github.com/gosrv/gbase/ghttp"
	"github.com/gosrv/gbase/gl"
	"github.com/gosrv/gbase/gmxdriver"
	"github.com/gosrv/gbase/tcpnet"
	_ "github.com/gosrv/gcluster/gcluster/baseapp/controller"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	_ "github.com/gosrv/gcluster/gcluster/baseapp/service"
	"github.com/gosrv/gcluster/gcluster/common"
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
)

// 客户端网络出了
func initBaseNet(builder gioc.IBeanContainerBuilder) {
	// pb消息编解码器，4字节长度 + 2字节proto id + proto
	idtype := common.LogicMsgIds
	encoder := codec.NewNetMsgFixLenProtobufEncoder(idtype)
	decoder := codec.NewNetMsgFixLenProtobufDecoder(idtype)
	// 创建网络模块，这里使用了数据自动同步路由器AutoSyncDataRoute
	net := tcpnet.NewTcpNetServer("pcluster.basenet", "",
		encoder, decoder, nil, entity.NewAutoSyncDataRoute())
	builder.AddBean(net)
}

func initCluster(builder gioc.IBeanContainerBuilder, app *app.Application) {
	idtype := common.ClusterMsgIds
	encoder := codec.NewNetMsgFixLenProtobufEncoder(idtype)
	decoder := codec.NewNetMsgFixLenProtobufDecoder(idtype)
	builder.AddBean(common.NewClusterMsgCenter(encoder, decoder))
	app.InitClusterMqBuilder(builder, "cluster", cluster.NodeMQName, encoder, decoder, nil, nil)
}

func initServices(builder gioc.IBeanContainerBuilder) {
	builder.AddBean(
		// redis 自动配置
		gredis.NewAutoConfigReids("pcluster.redis", ""),
		// mongo 自动配置
		gmongo.NewAutoConfigMongo("pcluster.mongo", ""),
		ghttp.NewHttpServer("pcluster.http", nil),
		gmxdriver.NewGMXDriver("/gmx"),
		common.NewGmxAppStats(),
	)

	builder.AddBean(common.BeansInit...)
}

func initLog(configLoader gioc.IConfigLoader, builder gioc.IBeanContainerBuilder) {
	logroot := &glog.ConfigLogRoot{}
	err := configLoader.Config().Get("pcluster.log").Scan(logroot)
	util.VerifyNoError(err)

	logBuilder := glog.NewLogFactoryBuilder()
	logFactory, err := logBuilder.Build(logroot)
	util.VerifyNoError(err)

	// 重定向系统日志
	err = gl.Redirect(logFactory.GetLogger("engine"))
	util.VerifyNoError(err)
	builder.AddBean(logFactory)
}

func main() {
	application := app.NewApplication()
	configLoader := application.InitCli()
	builder := application.InitBuilder()
	// 初始化日志
	initLog(configLoader, builder)

	application.InitBaseBeanBuilder(builder, configLoader)
	initCluster(builder, application)

	initServices(builder)
	initBaseNet(builder)
	beanContainer := application.Build(builder)
	application.Start(beanContainer)
}

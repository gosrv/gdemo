package main

import (
	"github.com/gosrv/gbase/app"
	"github.com/gosrv/gbase/codec"
	"github.com/gosrv/gbase/gl"
	"github.com/gosrv/gbase/tcpnet"
	"github.com/gosrv/gcluster/gcluster/baseapp/entity"
	"github.com/gosrv/gcluster/gcluster/testclient/logic"
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
)

func initBaseNet(builder gioc.IBeanContainerBuilder) {
	idtype := entity.NewLogicMsgIds()
	encoder := codec.NewNetMsgFixLenProtobufEncoder(idtype)
	decoder := codec.NewNetMsgFixLenProtobufDecoder(idtype)

	net := tcpnet.NewTcpNetServer("pcluster.basenet", "", encoder, decoder, nil, nil)
	builder.AddBean(net)
}

func initServices(beanContainerBuilder gioc.IBeanContainerBuilder) {
	beanContainerBuilder.AddBean(
		logic.NewControllerLogic(),
		logic.NewServiceLogic(),
	)
}

func initLog(configLoader gioc.IConfigLoader, builder gioc.IBeanContainerBuilder) {
	logroot := &glog.ConfigLogRoot{}
	configLoader.Config().Get("pcluster.log").Scan(logroot)

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
	initLog(configLoader, builder)
	application.InitBaseBeanBuilder(builder, configLoader)

	initServices(builder)
	initBaseNet(builder)
	beanContainer := application.Build(builder)

	application.Start(beanContainer)
}

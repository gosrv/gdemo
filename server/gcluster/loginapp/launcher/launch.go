package main

import (
	"github.com/gosrv/gbase/app"
	"github.com/gosrv/gbase/gdb/gmongo"
	"github.com/gosrv/gbase/gdb/gredis"
	"github.com/gosrv/gbase/ghttp"
	"github.com/gosrv/gbase/gl"
	"github.com/gosrv/gcluster/gcluster/loginapp/logic"
	"github.com/gosrv/glog"
	"github.com/gosrv/goioc"
	"github.com/gosrv/goioc/util"
)

func initServices(beanContainerBuilder gioc.IBeanContainerBuilder) {
	beanContainerBuilder.AddBean(
		// redis 自动配置
		gredis.NewAutoConfigReids("pcluster.redis", ""),
		// mongo 自动配置
		gmongo.NewAutoConfigMongo("pcluster.mongo", ""),
		// http 自动配置
		ghttp.NewHttpServer("pcluster.http", nil),
		logic.NewControllerLogin(),
		logic.NewServiceLogin(),
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
	// 初始化日志
	initLog(configLoader, builder)

	application.InitBaseBeanBuilder(builder, configLoader)

	initServices(builder)
	beanContainer := application.Build(builder)

	gl.Debug("application start...")
	application.Start(beanContainer)
	gl.Debug("application finished...")
}

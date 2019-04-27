### 使用步骤
这是自定义容器的过程  
1. 创建Application，它为我们提供了很多默认便捷配置
```$
    application := app.NewApplication()
```

2. 处理命令行输入
```$go
    configLoader := application.InitCli()
```
命令行的格式是app start -c cfgfile.json
主要是用来拿到配置文件

3. 创建一个IBeanContainerBuilder
```$go
    builder := application.InitBuilder()
```
builder是用来组装bean，然后创建bean容器的工厂

4. 加入默认bean，主要是tag解析器，tag处理器
```$go
    application.InitBaseBeanBuilder(builder, configLoader)
```

5. 加入应用专属bean
```$go
    builder.AddBean(otherBeans...)
```

6. 构建bean容器
```$go
    beanContainer := application.Build(builder)
```

7. 开始主线程死循环
```$go
    application.Start(beanContainer)
```

除了第5步外其它都是固定的，可以保持不变
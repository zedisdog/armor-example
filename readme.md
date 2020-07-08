# armor-example

这是一个实际项目的初始代码，还没来得及整理成更清晰的示例。

### 启动

先配置config.yml里面的数据库，然后执行下面命令
```bash
go generate
go run main.go
```

### 路由
路由跟gin的一样，其实就是用的gin。

### 队列
默认没有开启queue，若要开启需在main.go里面这样写：
```go
server := armor.NewArmor(models.AutoMigrate, armor.WithQueue(true))
```
别忘了配置config.yml里面的beanstalkd相关配置

### 优雅退出
支持优雅退出，在启动后按ctrl+c，可以看到相关日志。

### 请求处理流程
相对完整的一个请求处理流程可以看 `web/v1/enterprise.go` 的 `Register` 方法
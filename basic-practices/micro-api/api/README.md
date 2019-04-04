# API

本示例介绍使用**Micro API**（以下简称**API**）有的请求处理类型**api**.

我们专门定义api请求响应的proto文件，[api.Request/Response](https://github.com/micro/go-api/blob/master/proto/api.proto)，

要使用**api**类型的**API**服务，我们得使用这些proto原型。

## 使用方式

运行**API**网关，我们传入api指令运行：

```
micro api --handler=api
```

再运行本api服务

```
go run api.go
```


## 调用服务

通过URL**/example/call**，就会调用**go.micro.api.example**服务的**Example.Call**接口

```
curl "http://localhost:8080/example/call?name=john"
```

我们也加了POST路由 **/example/foo/bar**，可以通过它调用**go.micro.api.example**的**Foo.Bar**

```
curl -H 'Content-Type: application/json' -d '{data:123}' http://localhost:8080/example/foo/bar
```

## 设置命名空间

可以通过`--namespace`指定服务命令空间

```
micro api --handler=api --namespace=com.foobar.api
```

或者通过环境变量的方式

```
MICRO_API_NAMESPACE=com.foobar.api micro api --handler=api
```

切记，如果启动时指定命名空间，则代码中的服务名也要注意同步修改前缀

```
service := micro.NewService(
        micro.Name("com.foobar.api.example"),
)
```   

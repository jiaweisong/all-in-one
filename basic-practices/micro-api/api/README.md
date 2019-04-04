# API

本示例介绍使用**Micro API**（以下简称**API**）有的请求处理类型**api**.

我们专门定义api请求响应的proto文件，[api.Request/Response](https://github.com/micro/go-api/blob/master/proto/api.proto)，

要使用**api**类型的**API**服务，我们得使用这些proto原型。


The micro api request handler gives you full control over the http request and response while still leveraging RPC and 
any transport plugins that use other protocols beyond http in your stack such as grpc, nats, kafka.

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

Make a POST request to /example/foo/bar which will call go.micro.api.example Foo.Bar

```
curl -H 'Content-Type: application/json' -d '{}' http://localhost:8080/example/foo/bar
```

## 设置命名空间

Run the micro API with custom namespace

```
micro api --handler=api --namespace=com.foobar.api
```

or
```
MICRO_API_NAMESPACE=com.foobar.api micro api --handler=api
```

Set service name with the namespace

```
service := micro.NewService(
        micro.Name("com.foobar.api.example"),
)
```   

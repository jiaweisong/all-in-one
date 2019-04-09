# Function

Function是只响应一次请求的服务，我们现在介绍如何使用Function。

## 运行Function

Mac/Linux

```shell
watch -n1 go run main.go
```

Windows（Powershell）

```shell
while (1) {go run main.go; sleep 1}
```

## 调用Function

```shell
micro call greeter Greeter.Hello '{"name": "Micro中国"}'
```

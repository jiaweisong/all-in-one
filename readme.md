# micro 最佳实践

本项目旨在以全面的视角向大家演示micro服务的开发，并尽可能介绍其功能。我们尝试由浅入深，从基础篇开始，每个目录对应一个级别的示例。

本项目不是教程，只侧重如何使用，需要更深的教程，请翻阅[micro-tutorials][micro-tutorials]


## Micro源码运行环境安装

这里省略了Golang环境安装，未安装的同学可参考[Golang][golang-cn]，选择合适自己的安装包搭建Go开发环境。

### 安装micro

```bash
$ go get -u github.com/micro/micro
```

## 目录

- [micro-api](./basic-practices/micro-api) 
  - [rpc-handler](./basic-practices/micro-api/rpc) rpc模式
  - [api-handler](./basic-practices/micro-api/api) api模式
  - [proxy-handler](./basic-practices/micro-api/proxy) proxy（http）模式
  - [web-handler](./basic-practices/micro-api/web) web（websocket）模式
  - [event-handler](./basic-practices/micro-api/event) event模式
  - [meta-handler](./basic-practices/micro-api/meta) meta模式
- [micro-broker](./basic-practices/micro-broker) 
  - [发布与订阅](./basic-practices/micro-broker/basic) 
- [micro-service](./basic-practices/micro-service) 编写Micro服务，包含service和function
  - [function](./basic-practices/micro-service/function)
  - [service](./basic-practices/micro-service/service)
- [micro-cli](./middle-practices/micro-cli) 如何使用命令行接口
  - [flag](./middle-practices/micro-cli/flags) 如何使用flag
- [micro-config](./middle-practices/micro-config) 如何从配置中心读取配置
  - [file](./basic-practices/micro-config) 基于文件配置
  - [center](./middle-practices/micro-config) 使用配置中心 todo

[golang-cn]: https://golang.google.cn/
[micro-tutorials]: https://github.com/micro-in-cn/micro-tutorials

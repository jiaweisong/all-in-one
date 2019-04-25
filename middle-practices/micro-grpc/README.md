# GRPC

本篇给大家演示如何使用gRPC构建服务及使用gRPC gateway，代码请见[go-grpc](https://gitub.com/micro/go-grpc)

gRPC网关[gRPC-gateway](https://github.com/grpc-ecosystem/grpc-gateway)是[protoc](http://github.com/google/protobuf)的一个插件。

它遵循[gRPC](http://github.com/grpc/grpc-common)中对服务的定义，生成反向代理服务，这个代理就会把RESTful风格的JSON API转成gRPC请求。


## 预置条件

我们需要安装protobuf工具以及插件protoc-gen-grpc-gateway和protoc-gen-go。如果已经安装，请忽略以下步骤

```bash
mkdir tmp
cd tmp
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
```

```bash
$ go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
$ go get -u github.com/micro/protobuf/protoc-gen-go
```

## 内容

- [gateway](gateway) - gRPC网关
- [greeter](greeter) - 示例程序

## 开始上手

我们先创建需要的目录

```text
├── gateway    # 网关代码
├── greeter    # 示例程序
│   ├── cli    # 示例程序 服务端
│   └── srv    # 示例程序 客户端
└── proto      # proto 原型文件目录
    └── hello  # hello 原型文件及存根类目录
```

定义hello.proto原型

```proto
syntax = "proto3";

package go.micro.srv.greeter;

service Say {
	rpc Hello(Request) returns (Response) {}
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}
```

### 导入应用

```
import "github.com/micro/go-grpc"
```

### 创建 grpc.Service

使用grpc包创建服务

```
service := grpc.NewService()
```

### 构建二进制文件

```
// For local use
go build -i -o service ./main.go ./plugins.go
```

### Run

Because the default client/server have been replaced we can just run as usual

```
./service
```

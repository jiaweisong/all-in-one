package main

import (
	"github.com/micro-in-cn/all-in-one/senior-practices/micro-config/gRPC/structs"
	"github.com/micro/go-config"
	grpcConfig "github.com/micro/go-config/source/grpc"
	"github.com/micro/go-log"
)

func main() {

	source := grpcConfig.NewSource(
		grpcConfig.WithAddress("127.0.0.1:8600"),
		grpcConfig.WithPath("/micro"),
	)

	conf := config.NewConfig()
	err := conf.Load(source)
	if err != nil {
		log.Fatal(err)
	}

	configs := &structs.Micro{}
	err = conf.Scan(configs)
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Read configs: %s", string(conf.Bytes()))

	watcher, err := conf.Watch()
	if err != nil {
		log.Fatal(err)
	}

	log.Logf("Watch changes ...")

	for {
		v, err := watcher.Next()
		if err != nil {
			log.Fatal(err)
		}

		log.Logf("Watch changes: %v", string(v.Bytes()))
	}
}

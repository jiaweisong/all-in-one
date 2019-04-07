package basic

import (
	"fmt"
	"log"
	"time"

	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
)

var (
	topic = "go.micro.topic.foo"
)

func pub() {
	tick := time.NewTicker(time.Second)
	i := 0
	for range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := broker.Publish(topic, msg); err != nil {
			log.Printf("[pub] 发布消息失败： %v", err)
		} else {
			fmt.Println("[pub] 发布消息：", string(msg.Body))
		}
		i++
	}
}

func sub() {
	_, err := broker.Subscribe(topic, func(p broker.Publication) error {
		fmt.Println("[sub] 订阅收到消息：", string(p.Message().Body), "header", p.Message().Header)
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	cmd.Init()

	if err := broker.Init(); err != nil {
		log.Fatalf("Broker 初始化错误：%v", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatalf("Broker 连接错误：%v", err)
	}

	go pub()
	go sub()

	<-time.After(time.Second * 10)
}

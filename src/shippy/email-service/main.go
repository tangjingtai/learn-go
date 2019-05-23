package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/broker/rabbitmq"
	"log"
	userProto "shippy/user-service/proto/user"
)

const topic = "user.created"

type Subscriber struct{}

func (sub *Subscriber) Process2(ctx context.Context, user *userProto.User) error {
	log.Println("[Picked up a new message]")
	log.Println("[Sending email to]:", user.Name)
	return nil
}

func main() {
	// 通过这种方法来设置RabbitMQ相关参数
	opt := broker.SubscribeOptions{}
	rabbitmq.DurableQueue()(&opt)
	rabbitmq.AckOnSuccess()(&opt)

	service := micro.NewService(micro.Name("go.micro.srv.email"),
		micro.Version("latest"),
		micro.Address("localhost:0"),
		micro.Broker(rabbitmq.NewBroker(
			broker.Addrs("amqp://guest:guest@localhost:5672/"), // 使用RabbitMQ
			rabbitmq.PrefetchCount(5),
			rabbitmq.Exchange("shippy"),
			rabbitmq.DurableExchange(),
		)),
	)
	service.Init()

	err := micro.RegisterSubscriber(topic, service.Server(), func(ctx context.Context, user *userProto.User) error {
		log.Println("[1-Picked up a new message]")
		log.Println("[1-Sending email to]:", user.Name)
		return nil
	},
		server.SubscriberQueue("SendMail"),
		server.SubscriberContext(opt.Context),
	)

	//err = micro.RegisterSubscriber(topic, service.Server(), func(ctx context.Context, user *userProto.User) error {
	//	log.Println("[2-Picked up a new message]")
	//	log.Println("[2-Sending email to]:", user.Name)
	//	return nil
	//})

	//service.Server().Subscribe(server.NewSubscriber(topic, nil))

	if err != nil {
		log.Printf("sub error: %v\n", err)
	}

	if err := service.Run(); err != nil {
		log.Fatalf("srv run error: %v\n", err)
	}
}

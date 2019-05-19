package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-plugins/broker/rabbitmq"
	_ "github.com/micro/go-plugins/broker/rabbitmq"
	userPb "shippy/user-service/proto/user"
)

const topic = "user.created"

func main() {
	// 连接到数据库
	db, err := CreateConnection()

	fmt.Printf("%+v\n", db)
	fmt.Printf("err: %v\n", err)

	defer db.Close()

	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}

	repo := &UserRepository{db}
	token_service := &TokenService{repo: repo}

	// 自动检查 User 结构是否变化
	db.AutoMigrate(&userPb.User{})

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Address("localhost:0"),
		micro.Broker(rabbitmq.NewBroker(
			broker.Addrs("amqp://guest:guest@localhost:5672/"),
			rabbitmq.Exchange("shippy"),
			rabbitmq.DurableExchange(),
		)), // 使用RabbitMQ
	)

	s.Init()
	publisher := micro.NewPublisher(topic, s.Client())
	userPb.RegisterUserServiceHandler(s.Server(), &handler{repo: repo, tokenService: token_service, publisher: publisher})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}

}

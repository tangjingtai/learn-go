package main

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/micro/go-micro"
	userPb "shippy/user-service/proto/user"
)

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

	// 自动检查 User 结构是否变化
	db.AutoMigrate(&userPb.User{})

	s := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
		micro.Address("localhost:0"),
	)

	s.Init()

	userPb.RegisterUserServiceHandler(s.Server(), &handler{repo})

	if err := s.Run(); err != nil {
		log.Fatalf("user service error: %v\n", err)
	}

}

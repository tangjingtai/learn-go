package main

import (
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"golang.org/x/net/context"
	"log"
	pb "shippy/user-service/proto/user"
)

func main() {

	cmd.Init()

	// 创建 user-service 微服务的客户端
	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)
	r, err := client.Create(context.Background(), &pb.User{
		Name:     "tangjingtai",
		Email:    "a253210810@qq.com",
		Password: "123456",
		Company:  "ruanyun",
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %v", r.User.Id)
	//// 设置命令行参数
	//service := micro.NewService(
	//	micro.Flags(
	//		cli.StringFlag{
	//			Name:  "name",
	//			Usage: "tangjingtai",
	//		},
	//		cli.StringFlag{
	//			Name:  "email",
	//			Usage: "a253210810@qq.com",
	//		},
	//		cli.StringFlag{
	//			Name:  "password",
	//			Usage: "123456",
	//		},
	//		cli.StringFlag{
	//			Name:  "company",
	//			Usage: "ruanyun",
	//		},
	//	),
	//)

	//service.Init(
	//	micro.Action(func(c *cli.Context) {
	//		name := c.String("name")
	//		email := c.String("email")
	//		password := c.String("password")
	//		company := c.String("company")
	//
	//		r, err := client.Create(context.TODO(), &pb.User{
	//			Name:     name,
	//			Email:    email,
	//			Password: password,
	//			Company:  company,
	//		})
	//		if err != nil {
	//			log.Fatalf("Could not create: %v", err)
	//		}
	//		log.Printf("Created: %v", r.User.Id)
	//
	//		getAll, err := client.GetAll(context.Background(), &pb.Request{})
	//		if err != nil {
	//			log.Fatalf("Could not list users: %v", err)
	//		}
	//		for _, v := range getAll.Users {
	//			log.Println(v)
	//		}
	//
	//		os.Exit(0)
	//	}),
	//)
	//
	//// 启动客户端
	//if err := service.Run(); err != nil {
	//	log.Println(err)
	//}
}
